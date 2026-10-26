import { defineConfig, type Plugin } from "vite";
import { CONTACT, waLink } from "./src/config";

// Escribe los datos de contacto de src/config.ts directo en el HTML (en dev y en el build),
// para que los links funcionen aunque el JavaScript no llegue a cargar.
function contactHtml(): Plugin {
  const values: Record<string, string> = {
    WA_URL: waLink(CONTACT.defaultMessage),
    WA_LABEL: CONTACT.whatsappLabel,
    MAIL_URL: `mailto:${CONTACT.email}?subject=${encodeURIComponent("Consulta SEDH 2000")}`,
    EMAIL: CONTACT.email,
    YEAR: String(new Date().getFullYear()),
  };
  const escape = (s: string): string =>
    s.replace(/&/g, "&amp;").replace(/"/g, "&quot;").replace(/</g, "&lt;");

  return {
    name: "sedh-contact-html",
    transformIndexHtml: {
      order: "pre",
      handler: (html) =>
        html.replace(/\{\{(\w+)\}\}/g, (match, key: string) =>
          key in values ? escape(values[key]) : match,
        ),
    },
  };
}

export default defineConfig({
  plugins: [contactHtml()],
});
