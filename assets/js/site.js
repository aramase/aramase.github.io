// Theme toggle. The site works without this file: the toggle stays hidden and the OS scheme applies.
(function () {
  "use strict";

  var root = document.documentElement;
  var darkQuery = window.matchMedia("(prefers-color-scheme: dark)");

  function effectiveTheme() {
    return root.getAttribute("data-theme") || (darkQuery.matches ? "dark" : "light");
  }

  // The header strip's colour (--ink in site.css), so the mobile browser bar follows a manual choice.
  var strips = { light: "#111419", dark: "#edeef1" };

  function syncToggle(toggle) {
    var theme = effectiveTheme();
    toggle.setAttribute("aria-pressed", theme === "dark" ? "true" : "false");
    if (!root.hasAttribute("data-theme")) return;
    document.querySelectorAll('meta[name="theme-color"]').forEach(function (meta) {
      meta.setAttribute("content", strips[theme]);
    });
  }

  // Apply a theme change with colour transitions off, so text and paper switch together.
  function switchTheme(apply) {
    root.classList.add("theme-switching");
    if (apply) apply();
    // Reading layout flushes styles, so the new colours are computed with transitions off.
    void document.body.offsetHeight;
    window.requestAnimationFrame(function () {
      root.classList.remove("theme-switching");
    });
  }

  function setupTheme() {
    var toggle = document.querySelector(".toggle");
    if (!toggle) return;
    toggle.hidden = false;
    syncToggle(toggle);

    toggle.addEventListener("click", function () {
      var next = effectiveTheme() === "dark" ? "light" : "dark";
      switchTheme(function () {
        root.setAttribute("data-theme", next);
      });
      try {
        localStorage.setItem("theme", next);
      } catch (e) {
        // Storage can be unavailable (private mode); the choice still applies to this page.
      }
      syncToggle(toggle);
    });

    // The OS scheme can change while the page is open.
    darkQuery.addEventListener("change", function () {
      if (!root.hasAttribute("data-theme")) switchTheme(null);
      syncToggle(toggle);
    });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", setupTheme);
  } else {
    setupTheme();
  }
})();
