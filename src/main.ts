import { initNav } from "./modules/nav";
import { initReveal } from "./modules/reveal";
import { initFaq } from "./modules/faq";
import { initContactForm } from "./modules/contactForm";
import { inject } from "@vercel/analytics";

// Vercel Web Analytics (el proyecto es Vite sin framework, por eso se usa inject() y no el componente de Next).
inject();

// Avisa al script del <head> que el JS cargó (si no, vuelve a la versión sin JS).
document.documentElement.classList.add("js");
document.documentElement.dataset.ready = "1";

// Los links de WhatsApp/email y el año ya vienen escritos en el HTML desde src/config.ts
// (ver vite.config.ts). Acá solo se actualiza el año por si el build es de un año anterior.
function updateYear(): void {
  const year = document.querySelector<HTMLSpanElement>("#year");
  if (year) year.textContent = String(new Date().getFullYear());
}

document.addEventListener("DOMContentLoaded", () => {
  updateYear();
  initNav();
  initReveal();
  initFaq();
  initContactForm();
});
