// Datos de contacto del cliente. Cambiá acá y se actualiza toda la página.
export interface ContactConfig {
  /** Número en formato internacional, solo dígitos (para wa.me). */
  whatsapp: string;
  /** Cómo se muestra el número en pantalla. */
  whatsappLabel: string;
  //email: string;
  /** Mensaje por defecto del botón flotante de WhatsApp. */
  defaultMessage: string;
}

export const CONTACT: ContactConfig = {
  whatsapp: "+54 9 3496 44-0587",
  whatsappLabel: "+54 9 3496 44-0587",
  //email: "cinallijj@yahoo.com",
  defaultMessage: "Hola! Vi la página del SEDH 2000 y quiero más información sobre las sesiones.",
};

export const waLink = (message: string): string =>
  `https://wa.me/${CONTACT.whatsapp}?text=${encodeURIComponent(message)}`;
