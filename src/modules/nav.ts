// Menú móvil + sombra del header al hacer scroll.
export function initNav(): void {
  const header = document.querySelector<HTMLElement>("#header");
  const toggle = document.querySelector<HTMLButtonElement>("#navToggle");
  const nav = document.querySelector<HTMLElement>("#nav");
  if (!header || !toggle || !nav) return;

  const setOpen = (open: boolean): void => {
    nav.classList.toggle("is-open", open);
    toggle.classList.toggle("is-open", open);
    toggle.setAttribute("aria-expanded", String(open));
    toggle.setAttribute("aria-label", open ? "Cerrar menú" : "Abrir menú");
  };

  toggle.addEventListener("click", () => setOpen(!nav.classList.contains("is-open")));

  // Cerrar al tocar un link
  nav.querySelectorAll<HTMLAnchorElement>("a").forEach((a) =>
    a.addEventListener("click", () => setOpen(false)),
  );

  // Cerrar con Escape
  document.addEventListener("keydown", (e: KeyboardEvent) => {
    if (e.key === "Escape") setOpen(false);
  });

  const onScroll = (): void => {
    header.classList.toggle("is-scrolled", window.scrollY > 10);
  };
  window.addEventListener("scroll", onScroll, { passive: true });
  onScroll();
}
