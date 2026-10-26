import { waLink } from "../config";

// Datos que junta el formulario
interface ContactData {
  nombre: string;
  telefono: string;
  localidad: string;
  horario: string;
  mensaje: string;
}

type Errors = Partial<Record<keyof ContactData, string>>;

function readForm(form: HTMLFormElement): ContactData {
  const fd = new FormData(form);
  const get = (k: keyof ContactData): string => String(fd.get(k) ?? "").trim();
  return {
    nombre: get("nombre"),
    telefono: get("telefono"),
    localidad: get("localidad"),
    horario: get("horario"),
    mensaje: get("mensaje"),
  };
}

function validate(data: ContactData): Errors {
  const errors: Errors = {};
  if (data.nombre.length < 2) errors.nombre = "Contame tu nombre.";

  const digits = data.telefono.replace(/\D/g, "");
  if (digits.length < 8) errors.telefono = "Ingresá un teléfono válido.";

  return errors;
}

function buildMessage(d: ContactData): string {
  const lines = [
    "Hola! Quiero coordinar una sesión con el SEDH 2000.",
    `Nombre: ${d.nombre}`,
    `Teléfono: ${d.telefono}`,
  ];
  if (d.localidad) lines.push(`Localidad: ${d.localidad}`);
  lines.push(`Horario preferido: ${d.horario}`);
  if (d.mensaje) lines.push(`Consulta: ${d.mensaje}`);
  return lines.join("\n");
}

function showErrors(form: HTMLFormElement, errors: Errors): void {
  form.querySelectorAll<HTMLElement>(".form__row").forEach((row) => {
    const field = row.querySelector<HTMLInputElement>("input, textarea, select");
    const slot = row.querySelector<HTMLElement>(".form__error");
    if (!field || !slot) return;

    const msg = errors[field.name as keyof ContactData];
    slot.textContent = msg ?? "";
    row.classList.toggle("has-error", Boolean(msg));
    field.setAttribute("aria-invalid", String(Boolean(msg)));
  });
}

export function initContactForm(): void {
  const form = document.querySelector<HTMLFormElement>("#contactForm");
  if (!form) return;

  form.addEventListener("submit", (e: SubmitEvent) => {
    e.preventDefault();
    const data = readForm(form);
    const errors = validate(data);
    showErrors(form, errors);

    if (Object.keys(errors).length > 0) {
      form.querySelector<HTMLElement>(".has-error input")?.focus();
      return;
    }

    window.open(waLink(buildMessage(data)), "_blank", "noopener");
  });
}
