# Quick Reference Guide

## 🚀 Adding New Content (The Easy Way)

### New Talk
```bash
npm run new-talk
npm run build
```

### New Blog Post
```bash
npm run new-blog
npm run build
```

### Deploy Changes
```bash
npm run deploy
```

## 📝 Manual Content Creation

### Talk File Format (`content/talks/YYYY-MM-DD-slug.md`)
```markdown
---
title: "Your Talk Title"
date: 2024-12-01
venue: "KubeCon NA 2024"
recording: "https://youtu.be/your-video"
year: 2024
---

Brief description of your talk...
```

### Blog File Format (`content/blogs/YYYY-MM-DD-slug.md`)
```markdown
---
title: "Your Blog Title"
date: 2024-12-01
source: "Kubernetes.io"
url: "https://kubernetes.io/blog/your-post/"
year: 2024
---

Brief description of your blog post...
```

## 🔄 Workflow Summary

1. **Create content** → Add `.md` file to `content/talks/` or `content/blogs/`
2. **Build** → Run `npm run build`
3. **Preview** → Run `npm run serve` (opens on http://localhost:8000)
4. **Deploy** → Commit & push to GitHub

## 🎯 Key Benefits

- ✅ **Super Simple**: Just write markdown files
- ✅ **No Dependencies**: Pure Node.js, no frameworks
- ✅ **Fast**: Instant builds, no compilation
- ✅ **Git-Friendly**: All content is version controlled
- ✅ **GitHub Pages Ready**: Automatic deployment

## 🔧 File Structure

```
your-repo/
├── content/
│   ├── talks/           ← Add .md files here
│   └── blogs/           ← Add .md files here
├── build.js             ← The magic happens here
├── index.template.html  ← HTML template
├── index.html           ← Generated (don't edit)
└── package.json         ← npm scripts
```

## 💡 Pro Tips

- Use descriptive filenames: `2024-12-01-my-awesome-talk.md`
- Always run `npm run build` after adding content
- The `year` field is used for display, `date` for sorting
- Use `featured: true` in blog frontmatter for special styling
- Commit your `.md` files to git for version control

---

That's it! Your website is now as easy to update as writing a markdown file. 🎉
