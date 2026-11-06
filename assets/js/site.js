// Theme toggle, copy buttons and list filters. The site works without this file: the toggle,
// copy buttons and filters stay hidden, the OS colour scheme applies, the address is plain
// selectable text and every timeline entry shows.
(function () {
  "use strict";
  var root = document.documentElement;
  var dark = window.matchMedia("(prefers-color-scheme: dark)");

  function effective() {
    return root.getAttribute("data-theme") || (dark.matches ? "dark" : "light");
  }

  // The page background (--bg in site.css), so the mobile browser bar follows a manual choice.
  var bars = { light: "#fcfdff", dark: "#0c0f12" };

  document.querySelectorAll("[data-theme-toggle]").forEach(function (btn) {
    btn.hidden = false;
    var sync = function () {
      var theme = effective();
      btn.setAttribute("aria-pressed", theme === "dark" ? "true" : "false");
      if (!root.hasAttribute("data-theme")) return;
      document.querySelectorAll('meta[name="theme-color"]').forEach(function (meta) {
        meta.setAttribute("content", bars[theme]);
      });
    };
    sync();
    btn.addEventListener("click", function () {
      var next = effective() === "dark" ? "light" : "dark";
      root.setAttribute("data-theme", next);
      try { localStorage.setItem("theme", next); } catch (e) { /* private mode: applies to this page only */ }
      sync();
    });
    dark.addEventListener("change", sync);
  });

  // Copy the address. Clipboard API first; then execCommand on a selection; if both fail,
  // leave the address selected and tell the visitor to press the copy shortcut.
  function selectText(el) {
    var range = document.createRange();
    range.selectNodeContents(el);
    var sel = window.getSelection();
    sel.removeAllRanges();
    sel.addRange(range);
  }

  document.querySelectorAll("[data-copy]").forEach(function (btn) {
    var target = document.getElementById(btn.getAttribute("data-copy"));
    var status = document.getElementById(btn.getAttribute("data-status"));
    var label = btn.querySelector("[data-label]");
    var idle = label ? label.textContent : "";
    var timer;
    if (!target) return;
    btn.hidden = false;

    function report(ok) {
      var key = /Mac|iPhone|iPad/.test(navigator.platform) ? "⌘C" : "Ctrl+C";
      var msg = ok ? "Copied " + target.textContent.trim() : "Selected. Press " + key + " to copy.";
      if (status) status.textContent = msg;
      if (label) label.textContent = ok ? "Copied" : "Selected";
      btn.setAttribute("data-state", ok ? "copied" : "selected");
      clearTimeout(timer);
      timer = setTimeout(function () {
        if (label) label.textContent = idle;
        btn.removeAttribute("data-state");
        if (status) status.textContent = "";
      }, 2600);
    }

    function fallback() {
      selectText(target);
      var ok = false;
      try { ok = document.execCommand("copy"); } catch (e) { ok = false; }
      report(ok);
    }

    btn.addEventListener("click", function () {
      var text = target.textContent.trim();
      if (navigator.clipboard && window.isSecureContext) {
        navigator.clipboard.writeText(text).then(function () { report(true); }, fallback);
      } else {
        fallback();
      }
    });
  });

  // Filters on a list: buttons with data-filter toggle [data-kind] items. Without JS every item shows.
  document.querySelectorAll("[data-filters]").forEach(function (group) {
    var list = document.getElementById(group.getAttribute("data-filters"));
    var count = document.getElementById(group.getAttribute("data-count"));
    if (!list) return;
    group.hidden = false;
    var buttons = group.querySelectorAll("[data-filter]");
    buttons.forEach(function (b) {
      b.addEventListener("click", function () {
        var kind = b.getAttribute("data-filter");
        buttons.forEach(function (o) { o.setAttribute("aria-pressed", o === b ? "true" : "false"); });
        var shown = 0;
        list.querySelectorAll("[data-kind]").forEach(function (item) {
          var on = kind === "all" || item.getAttribute("data-kind") === kind;
          item.hidden = !on;
          if (on) shown++;
        });
        // Hide a group whose items are all filtered out.
        list.querySelectorAll("[data-group]").forEach(function (g) {
          g.hidden = !g.querySelector("[data-kind]:not([hidden])");
        });
        if (count) count.textContent = shown + (shown === 1 ? " entry" : " entries");
      });
    });
  });
})();
