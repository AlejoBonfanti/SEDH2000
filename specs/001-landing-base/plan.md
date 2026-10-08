# [001] Plan técnico (retroactivo)

**Spec:** ./spec.md

## Enfoque
Página única estática con HTML semántico, un CSS propio y TypeScript compilado por Vite.
Cada comportamiento interactivo es un módulo independiente que se inicializa en `main.ts`.
El contacto no tiene backend: se arma un link `wa.me` con el mensaje codificado.

## Cambios por archivo
| Archivo | Contenido | Cubre |
|---|---|---|
| `index.html` | Estructura y textos de todas las secciones, header, footer, botón flotante | RF1–RF10 |
| `src/config.ts` | `ContactConfig` (whatsapp, label, email, mensaje por defecto) + `waLink()` | RF11, RF8, RF9 |
| `src/main.ts` | Rellena links de WhatsApp/email y año; inicializa módulos en `DOMContentLoaded` | RF8–RF10 |
| `src/modules/nav.ts` | Menú mobile (toggle, Escape, cerrar al navegar) + sombra del header al scrollear | RF1 |
| `src/modules/reveal.ts` | `IntersectionObserver` para `.reveal`; muestra todo si hay reduced-motion | RNF |
| `src/modules/faq.ts` | Acordeón con `aria-expanded` y `max-height` | RF7 |
| `src/modules/contactForm.ts` | `readForm` → `validate` → `showErrors` → `buildMessage` → `window.open(waLink)` | RF8 |
| `src/styles/main.css` | Variables de color, layout, componentes, breakpoints 960/760 px, reduced-motion | RNF |
| `info/sedh2000-*.jpg` | Fotos usadas en la página (Vite las incluye en el build) | RF3, RF4 |

## Datos / configuración
```ts
interface ContactConfig { whatsapp: string; whatsappLabel: string; email: string; defaultMessage: string; }
interface ContactData   { nombre: string; telefono: string; localidad: string; horario: string; mensaje: string; }
```
Validación: nombre ≥ 2 caracteres; teléfono ≥ 8 dígitos.

## Decisiones y alternativas
- **WhatsApp en vez de backend/email**: es el canal que ya usa el cliente; cero mantenimiento. Contra: no queda registro de consultas fuera de WhatsApp.
- **Vite sin framework**: landing chica, objetivo de aprender TypeScript.
- **Animación hero en CSS** (ondas) en vez de foto: carga rápida y no depende de fotos con personas.

## Riesgos conocidos
- `.reveal` arranca con `opacity: 0` y el FAQ con `max-height: 0`: **sin JS el contenido no se ve** (rompe Constitución 3.3.12).
- Email y número de WhatsApp se inyectan por JS: sin JS los links quedan vacíos.
- Google Fonts externo: dependencia de terceros y algo de demora en la carga.
