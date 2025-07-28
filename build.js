#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Simple YAML frontmatter parser
function parseFrontmatter(content) {
    const lines = content.split('\n');
    if (lines[0] !== '---') return { frontmatter: {}, content };
    
    let i = 1;
    const frontmatterLines = [];
    while (i < lines.length && lines[i] !== '---') {
        frontmatterLines.push(lines[i]);
        i++;
    }
    
    const frontmatter = {};
    frontmatterLines.forEach(line => {
        const match = line.match(/^(\w+):\s*["']?(.+?)["']?$/);
        if (match) {
            frontmatter[match[1]] = match[2];
        }
    });
    
    const bodyContent = lines.slice(i + 1).join('\n').trim();
    return { frontmatter, content: bodyContent };
}

// Read all markdown files from a directory
function readMarkdownFiles(dirPath) {
    if (!fs.existsSync(dirPath)) return [];
    
    const files = fs.readdirSync(dirPath)
        .filter(file => file.endsWith('.md'))
        .map(file => {
            const filePath = path.join(dirPath, file);
            const content = fs.readFileSync(filePath, 'utf8');
            const { frontmatter, content: body } = parseFrontmatter(content);
            return { ...frontmatter, body, filename: file };
        });
    
    return files.sort((a, b) => new Date(b.date) - new Date(a.date));
}

// Generate talk cards HTML
function generateTalkCards(talks) {
    return talks.map(talk => `
        <div class="talk-card">
            <div class="talk-meta">
                <span class="talk-event">${talk.venue || 'Conference'}</span>
                <span class="talk-date">${talk.year || new Date(talk.date).getFullYear()}</span>
            </div>
            <h3>${talk.title}</h3>
            ${talk.recording ? `
            <a href="${talk.recording}" target="_blank" class="talk-link">
                <i class="fab fa-youtube"></i> Watch on YouTube
            </a>
            ` : ''}
        </div>
    `).join('');
}

// Generate blog cards HTML
function generateBlogCards(blogs) {
    return blogs.map(blog => `
        <div class="blog-card${blog.featured === 'true' || blog.source === 'Microsoft Open Source' ? ' featured' : ''}">
            <div class="blog-meta">
                <span class="blog-source">${blog.source || 'Blog'}</span>
                <span class="blog-date">${blog.year || new Date(blog.date).getFullYear()}</span>
            </div>
            <h3>${blog.title}</h3>
            <p>${blog.body}</p>
            ${blog.url ? `
            <a href="${blog.url}" target="_blank" class="blog-link">
                ${blog.source === 'Microsoft Open Source' ? 'View All Posts' : 'Read More'} <i class="fas fa-external-link-alt"></i>
            </a>
            ` : ''}
        </div>
    `).join('');
}

// Read content
const talks = readMarkdownFiles('./content/talks');
const blogs = readMarkdownFiles('./content/blogs');

console.log(`Found ${talks.length} talks and ${blogs.length} blogs`);

// Read the template HTML
const templatePath = './index.template.html';
if (!fs.existsSync(templatePath)) {
    console.error('Template file index.template.html not found. Please create it first.');
    process.exit(1);
}

const template = fs.readFileSync(templatePath, 'utf8');

// Replace placeholders
const html = template
    .replace('{{TALKS_CONTENT}}', generateTalkCards(talks))
    .replace('{{BLOGS_CONTENT}}', generateBlogCards(blogs));

// Write the output
fs.writeFileSync('./index.html', html);

console.log('✅ Generated index.html successfully!');
console.log('📁 To add new content:');
console.log('  - Add talks to content/talks/filename.md');
console.log('  - Add blogs to content/blogs/filename.md');
console.log('  - Run: node build.js');
