const STORAGE_KEY = 'theme';
const DARK_CLASS = 'dark-theme';
const LIGHT_CLASS = 'light-theme';

function getTheme() {
  const theme = localStorage.getItem(STORAGE_KEY);
  if (theme) {
    return theme;
  }
  
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}

function setTheme(theme) {
  const body = document.body;
  
  // Remove both classes first, then add the correct one
  // This prevents any intermediate state issues
  body.classList.remove(DARK_CLASS, LIGHT_CLASS);
  
  if (theme === 'dark') {
    body.classList.add(DARK_CLASS);
  } else {
    body.classList.add(LIGHT_CLASS);
  }
  
  localStorage.setItem(STORAGE_KEY, theme);
}

function toggleTheme() {
  const currentTheme = getTheme();
  const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
  setTheme(newTheme);
}

// Initialize theme toggle button after DOM loads
document.addEventListener('DOMContentLoaded', () => {
  const themeToggle = document.querySelector('.theme-toggle');
  if (themeToggle) {
    themeToggle.addEventListener('click', toggleTheme);
  }
});
