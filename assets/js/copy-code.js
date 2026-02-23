// Simple syntax highlighting setup
document.addEventListener('DOMContentLoaded', () => {
  // Add copy button to code blocks
  const codeBlocks = document.querySelectorAll('pre code');
  
  codeBlocks.forEach(block => {
    const pre = block.parentElement;
    const button = document.createElement('button');
    button.textContent = 'Copy';
    button.className = 'copy-button';
    button.addEventListener('click', () => {
      navigator.clipboard.writeText(block.textContent);
      button.textContent = 'Copied!';
      setTimeout(() => {
        button.textContent = 'Copy';
      }, 2000);
    });
    
    pre.style.position = 'relative';
    pre.appendChild(button);
  });
});
