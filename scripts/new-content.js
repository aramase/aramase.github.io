#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function ask(question) {
    return new Promise((resolve) => {
        rl.question(question, resolve);
    });
}

function slugify(text) {
    return text
        .toLowerCase()
        .replace(/[^\w\s-]/g, '')
        .replace(/[\s_-]+/g, '-')
        .replace(/^-+|-+$/g, '');
}

async function createTalk() {
    console.log('📢 Creating a new talk...\n');
    
    const title = await ask('Title: ');
    const venue = await ask('Venue (e.g., KubeCon NA 2024): ');
    const date = await ask('Date (YYYY-MM-DD): ');
    const recording = await ask('Recording URL (optional): ');
    const description = await ask('Brief description: ');
    
    const year = new Date(date).getFullYear();
    const slug = slugify(title);
    const filename = `${date}-${slug}.md`;
    const filePath = path.join('./content/talks', filename);
    
    const content = `---
title: "${title}"
date: ${date}
venue: "${venue}"
${recording ? `recording: "${recording}"` : ''}
year: ${year}
---

${description}`;

    fs.writeFileSync(filePath, content);
    console.log(`✅ Created talk: ${filePath}`);
}

async function createBlog() {
    console.log('📝 Creating a new blog post...\n');
    
    const title = await ask('Title: ');
    const source = await ask('Source (e.g., Kubernetes.io, Microsoft Open Source): ');
    const url = await ask('URL: ');
    const date = await ask('Date (YYYY-MM-DD): ');
    const description = await ask('Brief description: ');
    
    const year = new Date(date).getFullYear();
    const slug = slugify(title);
    const filename = `${date}-${slug}.md`;
    const filePath = path.join('./content/blogs', filename);
    
    const content = `---
title: "${title}"
date: ${date}
source: "${source}"
url: "${url}"
year: ${year}
---

${description}`;

    fs.writeFileSync(filePath, content);
    console.log(`✅ Created blog post: ${filePath}`);
}

async function main() {
    const type = process.argv[2];
    
    if (type === 'talk') {
        await createTalk();
    } else if (type === 'blog') {
        await createBlog();
    } else {
        console.log('Usage: node scripts/new-content.js [talk|blog]');
        process.exit(1);
    }
    
    rl.close();
    
    console.log('\n🔨 Run "npm run build" to regenerate the website!');
}

main().catch(console.error);
