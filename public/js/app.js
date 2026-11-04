// Comportamiento del lado del navegador: menú móvil, animaciones y acordeón.
// El formulario de contacto lo maneja htmx + el servidor Go (no hace falta JS acá).
(function () {
  "use strict";

  // Avisa al script del <head> que el JS cargó (si no, vuelve a la versión sin JS).
  var root = document.documentElement;
  root.classList.add("js");
  root.dataset.ready = "1";

  // ---------- Menú móvil + sombra del header al hacer scroll ----------
  function initNav() {
    var header = document.getElementById("header");
    var toggle = document.getElementById("navToggle");
    var nav = document.getElementById("nav");
    if (!header || !toggle || !nav) return;

    function setOpen(open) {
      nav.classList.toggle("is-open", open);
      toggle.classList.toggle("is-open", open);
      toggle.setAttribute("aria-expanded", String(open));
      toggle.setAttribute("aria-label", open ? "Cerrar menú" : "Abrir menú");
    }

    toggle.addEventListener("click", function () {
      setOpen(!nav.classList.contains("is-open"));
    });
    nav.querySelectorAll("a").forEach(function (a) {
      a.addEventListener("click", function () { setOpen(false); });
    });
    document.addEventListener("keydown", function (e) {
      if (e.key === "Escape") setOpen(false);
    });

    function onScroll() {
      header.classList.toggle("is-scrolled", window.scrollY > 10);
    }
    window.addEventListener("scroll", onScroll, { passive: true });
    onScroll();
  }

  // ---------- Aparición suave de los .reveal ----------
  function initReveal() {
    var items = document.querySelectorAll(".reveal");
    var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;

    if (reduceMotion || !("IntersectionObserver" in window)) {
      items.forEach(function (el) { el.classList.add("is-visible"); });
      return;
    }

    var observer = new IntersectionObserver(function (entries, obs) {
      entries.forEach(function (entry) {
        if (entry.isIntersecting) {
          entry.target.classList.add("is-visible");
          obs.unobserve(entry.target);
        }
      });
    }, { threshold: 0.15, rootMargin: "0px 0px -40px 0px" });

    items.forEach(function (el) { observer.observe(el); });
  }

  // ---------- Preguntas frecuentes: una abierta a la vez ----------
  function initFaq() {
    var faq = document.getElementById("faq");
    if (!faq) return;
    var buttons = Array.prototype.slice.call(faq.querySelectorAll(".faq__q"));

    function setOpen(btn, open) {
      btn.setAttribute("aria-expanded", String(open));
      var answer = btn.nextElementSibling;
      if (answer) answer.style.maxHeight = open ? answer.scrollHeight + "px" : "0px";
    }

    buttons.forEach(function (btn) {
      btn.addEventListener("click", function () {
        var isOpen = btn.getAttribute("aria-expanded") === "true";
        buttons.forEach(function (b) { setOpen(b, false); });
        if (!isOpen) setOpen(btn, true);
      });
    });
  }

  initNav();
  initReveal();
  initFaq();
})();
