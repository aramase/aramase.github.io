# Anish Ramasekar - Personal Website

A modern, responsive personal website showcasing conference talks and blog posts about Kubernetes, cloud security, and open-source technologies.

## 🌟 Features

- **Markdown-Based Content**: Just add `.md` files to add new talks or blogs
- **Automatic Generation**: Simple build script converts markdown to HTML
- **Responsive Design**: Works perfectly on desktop, tablet, and mobile devices
- **Modern UI**: Clean, professional design with smooth animations
- **Performance Optimized**: Fast loading with optimized images and assets

## 🚀 Quick Start

### Adding New Content

**Add a new talk:**
```bash
npm run new-talk
# Follow the prompts, then:
npm run build
```

**Add a new blog post:**
```bash
npm run new-blog
# Follow the prompts, then:
npm run build
```

**Manual approach:**
1. Create a new `.md` file in `content/talks/` or `content/blogs/`
2. Add frontmatter with required fields
3. Run `npm run build`

### Example Talk (`content/talks/2024-my-awesome-talk.md`):
```markdown
---
title: "My Awesome Kubernetes Talk"
date: 2024-12-01
venue: "KubeCon NA 2024"
recording: "https://youtu.be/example"
year: 2024
---

This talk covers advanced Kubernetes security patterns...
```

### Example Blog (`content/blogs/2024-my-blog-post.md`):
```markdown
---
title: "Advanced Kubernetes Security"
date: 2024-12-01
source: "Kubernetes.io"
url: "https://kubernetes.io/blog/2024/12/01/my-post/"
year: 2024
---

In this post, we explore advanced security patterns...
```

## �️ Available Commands

```bash
npm run build          # Generate website from markdown files
npm run serve          # Serve website locally on port 8000
npm run deploy         # Build, commit, and push to GitHub
npm run new-talk       # Interactive talk creation
npm run new-blog       # Interactive blog post creation
```

## 📁 Project Structure

```
├── content/
│   ├── talks/         # Markdown files for talks
│   └── blogs/         # Markdown files for blog posts
├── scripts/
│   └── new-content.js # Helper script for creating content
├── build.js           # Build script (generates index.html)
├── index.template.html # HTML template
├── index.html         # Generated website (don't edit directly)
├── styles.css         # CSS styles
└── scripts.js         # JavaScript functionality
```

## � Workflow

1. **Add content**: Create new `.md` files or use `npm run new-talk`/`npm run new-blog`
2. **Build**: Run `npm run build` to generate the website
3. **Preview**: Run `npm run serve` to preview locally
4. **Deploy**: Commit and push to GitHub (or use `npm run deploy`)

## 🎯 Why This Approach?

- **Simple**: Just write markdown, no complex build tools
- **Fast**: Single Node.js script, no dependencies
- **Maintainable**: Clear separation of content and presentation
- **Version Controlled**: All content is in git
- **GitHub Pages Ready**: Works perfectly with GitHub Pages

## � Content Format

### Talk Frontmatter Fields:
- `title`: Talk title
- `date`: Date in YYYY-MM-DD format
- `venue`: Conference or event name
- `recording`: YouTube or video URL (optional)
- `year`: Year for display

### Blog Frontmatter Fields:
- `title`: Blog post title
- `date`: Date in YYYY-MM-DD format
- `source`: Publication source
- `url`: Link to the blog post
- `year`: Year for display
- `featured`: Set to `true` for special styling (optional)

## 🚀 Live Site

Visit the live website at: [https://aramase.github.io](https://aramase.github.io)

## 🎨 Customization

The website uses a template-based approach:
- `index.template.html`: Main HTML template with `{{TALKS_CONTENT}}` and `{{BLOGS_CONTENT}}` placeholders
- `build.js`: Generates HTML by replacing placeholders with content from markdown files
- `styles.css`: All styling (modify colors, fonts, layout here)

## 📈 Future Enhancements

- [ ] Tags/categories for filtering
- [ ] Search functionality
- [ ] Dark mode toggle
- [ ] RSS feed generation
- [ ] Image optimization
- [ ] SEO improvements

---

Built with ❤️ for the Kubernetes and open-source community.
