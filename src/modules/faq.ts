// Acordeón de preguntas frecuentes: una abierta a la vez.
export function initFaq(): void {
  const faq = document.querySelector<HTMLElement>("#faq");
  if (!faq) return;

  const buttons = Array.from(faq.querySelectorAll<HTMLButtonElement>(".faq__q"));

  const close = (btn: HTMLButtonElement): void => {
    btn.setAttribute("aria-expanded", "false");
    const answer = btn.nextElementSibling as HTMLElement | null;
    if (answer) answer.style.maxHeight = "0px";
  };

  const open = (btn: HTMLButtonElement): void => {
    btn.setAttribute("aria-expanded", "true");
    const answer = btn.nextElementSibling as HTMLElement | null;
    if (answer) answer.style.maxHeight = `${answer.scrollHeight}px`;
  };

  buttons.forEach((btn) => {
    btn.addEventListener("click", () => {
      const isOpen = btn.getAttribute("aria-expanded") === "true";
      buttons.forEach(close);
      if (!isOpen) open(btn);
    });
  });
}
