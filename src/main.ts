import { CONTACT, waLink } from "./config";
import { initNav } from "./modules/nav";
import { initReveal } from "./modules/reveal";
import { initFaq } from "./modules/faq";
import { initContactForm } from "./modules/contactForm";

function initContactLinks(): void {
  const waDirect = document.querySelector<HTMLAnchorElement>("#waDirect");
  const waFloat = document.querySelector<HTMLAnchorElement>("#waFloat");
  const mail = document.querySelector<HTMLAnchorElement>("#mailDirect");

  if (waDirect) {
    waDirect.href = waLink(CONTACT.defaultMessage);
    waDirect.textContent = CONTACT.whatsappLabel;
  }
  if (waFloat) waFloat.href = waLink(CONTACT.defaultMessage);
  if (mail) {
    mail.href = `mailto:${CONTACT.email}?subject=${encodeURIComponent("Consulta SEDH 2000")}`;
    mail.textContent = CONTACT.email;
  }

  const year = document.querySelector<HTMLSpanElement>("#year");
  if (year) year.textContent = String(new Date().getFullYear());
}

document.addEventListener("DOMContentLoaded", () => {
  initContactLinks();
  initNav();
  initReveal();
  initFaq();
  initContactForm();
});
