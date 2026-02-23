// Theme toggle functionality
(function () {
  "use strict";

  function getEffectiveTheme() {
    var saved = localStorage.getItem("theme");
    if (saved) return saved;
    return window.matchMedia("(prefers-color-scheme: light)").matches
      ? "light"
      : "dark";
  }

  function applyTheme(theme) {
    document.documentElement.setAttribute("data-theme", theme);
    localStorage.setItem("theme", theme);
  }

  document.addEventListener("DOMContentLoaded", function () {
    var toggle = document.querySelector(".theme-toggle");
    if (!toggle) return;

    toggle.addEventListener("click", function () {
      var current = getEffectiveTheme();
      applyTheme(current === "dark" ? "light" : "dark");
    });

    // Listen for OS theme changes — only apply if user hasn't manually chosen
    window
      .matchMedia("(prefers-color-scheme: light)")
      .addEventListener("change", function () {
        if (!localStorage.getItem("theme")) {
          // No manual override, OS pref already handled via CSS
          // Remove data-theme so CSS media query takes over
          document.documentElement.removeAttribute("data-theme");
        }
      });
  });
})();
