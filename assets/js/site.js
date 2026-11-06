// Theme toggle and the ledger cross-highlight. The site works without this file.
(function () {
  "use strict";

  var root = document.documentElement;
  var darkQuery = window.matchMedia("(prefers-color-scheme: dark)");

  function effectiveTheme() {
    return root.getAttribute("data-theme") || (darkQuery.matches ? "dark" : "light");
  }

  // Matches --ground in site.css; keeps the mobile browser bar in step with a manual choice.
  var grounds = { light: "#f2f9f6", dark: "#041b1d" };

  function syncToggle(toggle) {
    var theme = effectiveTheme();
    toggle.setAttribute("aria-pressed", theme === "dark" ? "true" : "false");
    if (!root.hasAttribute("data-theme")) return;
    document.querySelectorAll('meta[name="theme-color"]').forEach(function (meta) {
      meta.setAttribute("content", grounds[theme]);
    });
  }

  // Apply a theme change with colour transitions off, so text and ground switch
  // together instead of the text fading in against a ground that already changed.
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
    var toggle = document.querySelector(".theme-toggle");
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

    // The OS scheme can change while the page is open; the media query has already
    // flipped, and this event runs before the next style update.
    darkQuery.addEventListener("change", function () {
      if (!root.hasAttribute("data-theme")) switchTheme(null);
      syncToggle(toggle);
    });
  }

  // Focusing or hovering a ledger feature, or an area entry with a KEP, quiets the
  // other rows and the unrelated entries in the record, so its posts and talks step forward.
  function setupLedger() {
    if (!document.querySelector("[data-ledger]")) return;
    var scope = document.querySelector("[data-ledger-scope]") || document.body;
    var features = scope.querySelectorAll(".feature[data-feature], .area-entry[data-feature]");
    var entries = scope.querySelectorAll(".entry[data-features]");
    var active = null;

    function focusFeature(feature) {
      if (feature === active) return;
      active = feature;
      var id = feature ? feature.getAttribute("data-feature") : null;
      if (id) {
        scope.setAttribute("data-focus", id);
      } else {
        scope.removeAttribute("data-focus");
      }
      features.forEach(function (f) {
        f.classList.toggle("is-focused", f === feature);
      });
      var counts = { post: 0, talk: 0 };
      entries.forEach(function (e) {
        var ids = (e.getAttribute("data-features") || "").split(" ");
        var related = !!id && ids.indexOf(id) !== -1;
        e.classList.toggle("is-related", related);
        if (related) counts[e.getAttribute("data-kind")] += 1;
      });
      scope.querySelectorAll(".feature__cue").forEach(function (cue) {
        var mine = !!feature && feature.contains(cue);
        var parts = [];
        if (counts.post) parts.push(counts.post + (counts.post === 1 ? " post" : " posts"));
        if (counts.talk) parts.push(counts.talk + (counts.talk === 1 ? " talk" : " talks"));
        cue.hidden = !(mine && parts.length);
        if (mine) cue.textContent = parts.join(", ") + " in the record below";
      });
    }

    features.forEach(function (feature) {
      feature.addEventListener("pointerenter", function (ev) {
        if (ev.pointerType === "mouse") focusFeature(feature);
      });
      feature.addEventListener("pointerleave", function (ev) {
        if (ev.pointerType === "mouse" && !feature.contains(document.activeElement)) focusFeature(null);
      });
      feature.addEventListener("focusin", function () {
        focusFeature(feature);
      });
      feature.addEventListener("focusout", function (ev) {
        if (!feature.contains(ev.relatedTarget)) focusFeature(null);
      });
    });

    document.addEventListener("keydown", function (ev) {
      if (ev.key === "Escape") focusFeature(null);
    });
  }

  function init() {
    setupTheme();
    setupLedger();
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();
