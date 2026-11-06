package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// source finds a KEP's directory under keps/<sig>/ and reads its kep.yaml.
type source interface {
	kepYAML(ctx context.Context, sig string, kep int) (dir string, data []byte, err error)
}

// pickDir returns the one directory name for a KEP, named "<kep>-<slug>".
func pickDir(names []string, sig string, kep int) (string, error) {
	prefix := strconv.Itoa(kep) + "-"
	var found []string
	for _, n := range names {
		if strings.HasPrefix(n, prefix) {
			found = append(found, n)
		}
	}
	switch len(found) {
	case 0:
		return "", fmt.Errorf("KEP %d: no keps/%s/%s* directory in kubernetes/enhancements", kep, sig, prefix)
	case 1:
		return found[0], nil
	default:
		return "", fmt.Errorf("KEP %d: more than one keps/%s/%s* directory: %s", kep, sig, prefix, strings.Join(found, ", "))
	}
}

// localSource reads a kubernetes/enhancements clone.
type localSource struct {
	root string
}

func (l localSource) kepYAML(_ context.Context, sig string, kep int) (string, []byte, error) {
	entries, err := os.ReadDir(filepath.Join(l.root, "keps", sig))
	if err != nil {
		return "", nil, fmt.Errorf("KEP %d: %w", kep, err)
	}
	var names []string
	for _, e := range entries {
		if e.IsDir() {
			names = append(names, e.Name())
		}
	}
	dir, err := pickDir(names, sig, kep)
	if err != nil {
		return "", nil, err
	}
	b, err := os.ReadFile(filepath.Join(l.root, "keps", sig, dir, "kep.yaml"))
	if err != nil {
		return "", nil, fmt.Errorf("KEP %d: %w", kep, err)
	}
	return dir, b, nil
}

// remoteSource reads kubernetes/enhancements on GitHub: the contents API lists a SIG
// directory once, and kep.yaml comes from raw.githubusercontent.com.
type remoteSource struct {
	gh  *gitHub
	ref string

	mu   sync.Mutex
	dirs map[string][]string
}

func (r *remoteSource) kepYAML(ctx context.Context, sig string, kep int) (string, []byte, error) {
	names, err := r.list(ctx, sig)
	if err != nil {
		return "", nil, fmt.Errorf("KEP %d: %w", kep, err)
	}
	dir, err := pickDir(names, sig, kep)
	if err != nil {
		return "", nil, err
	}
	b, err := r.gh.get(ctx, fmt.Sprintf("%s/kubernetes/enhancements/%s/keps/%s/%s/kep.yaml", r.gh.raw, r.ref, sig, dir), false)
	if err != nil {
		return "", nil, fmt.Errorf("KEP %d: %w", kep, err)
	}
	return dir, b, nil
}

func (r *remoteSource) list(ctx context.Context, sig string) ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if names, ok := r.dirs[sig]; ok {
		return names, nil
	}
	b, err := r.gh.get(ctx, fmt.Sprintf("%s/repos/kubernetes/enhancements/contents/keps/%s?ref=%s", r.gh.api, sig, r.ref), true)
	if err != nil {
		return nil, err
	}
	var entries []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &entries); err != nil {
		return nil, fmt.Errorf("listing keps/%s: %w", sig, err)
	}
	var names []string
	for _, e := range entries {
		if e.Type == "dir" {
			names = append(names, e.Name)
		}
	}
	if r.dirs == nil {
		r.dirs = map[string][]string{}
	}
	r.dirs[sig] = names
	return names, nil
}

// gitHub is a minimal read-only GitHub client. The token, when set, goes only to the API.
type gitHub struct {
	client   *http.Client
	api, raw string
	token    string
}

func newGitHub(token string) *gitHub {
	return &gitHub{
		client: http.DefaultClient,
		api:    "https://api.github.com",
		raw:    "https://raw.githubusercontent.com",
		token:  token,
	}
}

func (g *gitHub) get(ctx context.Context, url string, api bool) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if api {
		req.Header.Set("Accept", "application/vnd.github+json")
		if g.token != "" {
			req.Header.Set("Authorization", "Bearer "+g.token)
		}
	}
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("GET %s: %s", url, resp.Status)
		var msg struct {
			Message string `json:"message"`
		}
		if json.Unmarshal(b, &msg) == nil && msg.Message != "" {
			err = fmt.Errorf("%w: %s", err, msg.Message)
		}
		if api && g.token == "" && (resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusTooManyRequests) {
			err = fmt.Errorf("%w (set GITHUB_TOKEN, or pass -enhancements with a local clone)", err)
		}
		return nil, err
	}
	return b, nil
}

// issue is the part of a KEP issue that says whether a release tracks the KEP.
type issue struct {
	Milestone string
	Labels    []string
}

func (g *gitHub) issue(ctx context.Context, n int) (issue, error) {
	b, err := g.get(ctx, fmt.Sprintf("%s/repos/kubernetes/enhancements/issues/%d", g.api, n), true)
	if err != nil {
		return issue{}, err
	}
	var raw struct {
		Milestone *struct {
			Title string `json:"title"`
		} `json:"milestone"`
		Labels []struct {
			Name string `json:"name"`
		} `json:"labels"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return issue{}, fmt.Errorf("KEP issue %d: %w", n, err)
	}
	var is issue
	if raw.Milestone != nil {
		is.Milestone = raw.Milestone.Title
	}
	for _, l := range raw.Labels {
		is.Labels = append(is.Labels, l.Name)
	}
	return is, nil
}

var releaseTagRE = regexp.MustCompile(`^refs/tags/v1\.(\d+)\.0$`)

// latestRelease returns the minor of the newest v1.N.0 tag in kubernetes/kubernetes.
func latestRelease(ctx context.Context) (int, error) {
	var out bytes.Buffer
	cmd := exec.CommandContext(ctx, "git", "ls-remote", "--tags", "--refs", "https://github.com/kubernetes/kubernetes", "v1.*")
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("listing kubernetes/kubernetes tags: %w", err)
	}
	return parseLatestRelease(out.String())
}

func parseLatestRelease(lsRemote string) (int, error) {
	latest := 0
	for _, line := range strings.Split(lsRemote, "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		m := releaseTagRE.FindStringSubmatch(fields[1])
		if m == nil {
			continue
		}
		if n, _ := strconv.Atoi(m[1]); n > latest {
			latest = n
		}
	}
	if latest == 0 {
		return 0, fmt.Errorf("no v1.N.0 tag in kubernetes/kubernetes")
	}
	return latest, nil
}

// releaseDater gives the date of each Kubernetes release: the commit date of its v1.N.0 tag.
type releaseDater interface {
	releaseDates(ctx context.Context, releases []int) (map[int]time.Time, error)
}

// gitDater reads the v1.N.0 tag dates from a kubernetes/kubernetes remote. It fetches
// only the tagged commits, without trees or blobs, into a temporary bare repository, and
// reads each commit date as `git log -1 --format=%cs` gives it. Tag dates never change,
// so the dates in the output change only when a release gains its first stage.
type gitDater struct {
	remote string
}

const kubernetesRemote = "https://github.com/kubernetes/kubernetes"

func (d gitDater) releaseDates(ctx context.Context, releases []int) (map[int]time.Time, error) {
	dates := make(map[int]time.Time, len(releases))
	if len(releases) == 0 {
		return dates, nil
	}
	dir, err := os.MkdirTemp("", "ledger-tags-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	git := func(args ...string) (string, error) {
		var stderr bytes.Buffer
		cmd := exec.CommandContext(ctx, "git", append([]string{"-C", dir}, args...)...)
		cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
		cmd.Stderr = &stderr
		out, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(stderr.String()))
		}
		return strings.TrimSpace(string(out)), nil
	}

	if _, err := git("init", "--quiet", "--bare"); err != nil {
		return nil, err
	}
	fetch := []string{"fetch", "--quiet", "--no-tags", "--depth=1", "--filter=tree:0", d.remote}
	for _, r := range releases {
		fetch = append(fetch, fmt.Sprintf("refs/tags/v1.%d.0:refs/tags/v1.%d.0", r, r))
	}
	if _, err := git(fetch...); err != nil {
		return nil, fmt.Errorf("fetching the v1.N.0 tags from %s: %w", d.remote, err)
	}
	for _, r := range releases {
		out, err := git("log", "-1", "--format=%cs", fmt.Sprintf("refs/tags/v1.%d.0", r))
		if err != nil {
			return nil, err
		}
		t, err := time.Parse(time.DateOnly, out)
		if err != nil {
			return nil, fmt.Errorf("v1.%d.0: commit date %q: %w", r, out, err)
		}
		dates[r] = t
	}
	return dates, nil
}
