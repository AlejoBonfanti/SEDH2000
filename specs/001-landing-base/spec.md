# [001] Landing base

**Estado:** Hecha (documentada de forma retroactiva) | **Fecha:** 2026-10-07 | **Pedido por:** Cliente (Ing. Juan José Cinalli — Patagonia Instrumental)

## Problema / objetivo
El cliente no tiene presencia web para el SEDH 2000. Hoy las consultas llegan por WhatsApp sin un lugar
donde la persona pueda ver qué es el equipo, cómo es una sesión y quién lo atiende.
**Objetivo:** una página única que presente el equipo y la sesión, genere confianza y lleve a la persona
a coordinar una sesión por WhatsApp — sin hacer promesas de salud.

## Usuarios
- **Interesado/a desde el celular** — llega por un link compartido (WhatsApp, Instagram, boca en boca). Quiere entender rápido qué es y cómo contactar.
- **Familiar o acompañante** — evalúa si es seguro/serio antes de que otra persona reserve.
- **El cliente** — necesita recibir consultas ordenadas (nombre, teléfono, horario) sin manejar un sistema.

## Historias de usuario
- **H1** — Como interesado/a, quiero ver en la primera pantalla qué es el SEDH 2000 y un botón para contactar, para no tener que buscar.
- **H2** — Como interesado/a, quiero conocer las partes del equipo, para entender qué me van a colocar.
- **H3** — Como interesado/a, quiero saber paso a paso cómo es una sesión y cuánto dura, para saber qué esperar.
- **H4** — Como acompañante, quiero saber quién desarrolló el equipo y dónde está, para evaluar si es confiable.
- **H5** — Como interesado/a, quiero saber en qué casos **no** se usa el equipo, para no reservar si no corresponde.
- **H6** — Como interesado/a, quiero respuestas a las dudas más comunes sin tener que escribir.
- **H7** — Como interesado/a, quiero dejar mis datos y que se abra WhatsApp con el mensaje armado, para coordinar en un paso.
- **H8** — Como cliente, quiero cambiar mi teléfono o email en un solo lugar.

## Requisitos funcionales
- **RF1 — Header fijo** con logo, navegación a las secciones y botón "Coordinar sesión". En mobile, menú hamburguesa que se cierra al tocar un link o con Escape.
- **RF2 — Hero (`#inicio`)**: nombre del equipo, origen (Industria Argentina · Rosario), bajada y dos CTA: coordinar sesión / ver cómo es una sesión.
- **RF3 — El equipo (`#equipo`)**: foto del equipo + 4 tarjetas (módulo de modulación, dos bobinas, panel con tres controles, batea) + bloque "Una sesión a tu medida" con los 3 controles (tiempo, edad, contextura).
- **RF4 — La sesión (`#sesion`)**: 4 pasos numerados (charla previa, preparación, ajuste, sesión), galería de 3 fotos y nota con el protocolo del fabricante (1 diaria, ciclo de 5 a 8).
- **RF5 — Quiénes somos (`#nosotros`)**: Patagonia Instrumental, slogan "Tecnología para la vida", Ing. Cinalli como desarrollador y 3 datos destacados.
- **RF6 — Antes de reservar (`#importante`)**: contraindicaciones del fabricante, aviso de consultar (marcapasos, embarazo, dudas) y disclaimer de que no es tratamiento médico.
- **RF7 — Preguntas frecuentes (`#preguntas`)**: acordeón de 6 preguntas, una abierta a la vez.
- **RF8 — Contacto (`#contacto`)**: WhatsApp, email y ubicación + formulario (nombre*, teléfono*, localidad, horario preferido, consulta). Al enviar, valida y abre WhatsApp con el mensaje armado.
- **RF9 — Botón flotante de WhatsApp** visible en toda la página con mensaje por defecto.
- **RF10 — Footer** con marca, disclaimer médico y año actual automático.
- **RF11 — Datos de contacto centralizados** en `src/config.ts`.

## Requisitos no funcionales
- Responsive (mobile first). Animaciones de aparición que respetan `prefers-reduced-motion`.
- Accesibilidad: `alt` descriptivos, `aria-expanded` en menú y FAQ, foco visible, errores del formulario con `aria-invalid`.
- Sin backend: todo el contacto pasa por WhatsApp / email.
- Textos en español rioplatense, sin promesas de salud (Constitución 3.1).

## Criterios de aceptación
- [x] **CA1** — Dado un celular de 360 px, cuando abro la página, entonces veo el nombre del equipo y un botón para coordinar sin hacer scroll.
- [x] **CA2** — Cuando toco un link del menú mobile, entonces se cierra el menú y la página baja a la sección.
- [x] **CA3** — Cuando toco una pregunta del FAQ, entonces se abre su respuesta y se cierra la que estaba abierta.
- [x] **CA4** — Dado el formulario con nombre de 1 letra o teléfono con menos de 8 dígitos, cuando envío, entonces veo el error en el campo y el foco va al primer campo con error.
- [x] **CA5** — Dado el formulario válido, cuando envío, entonces se abre WhatsApp al número del cliente con nombre, teléfono, horario y (si hay) localidad y consulta.
- [x] **CA6** — Cuando cambio el número en `src/config.ts`, entonces se actualizan el link de contacto, el botón flotante y el destino del formulario.
- [x] **CA7** — La sección "Antes de reservar" y el disclaimer del footer están presentes.
- [x] **CA8** — `tsc` compila sin errores en modo `strict`.
- [ ] **CA9** — Con JavaScript desactivado, todo el contenido sigue siendo legible. → **No se cumple** (ver spec 008).
- [ ] **CA10** — Verificado en dispositivo real (Android + iPhone). → Pendiente.

## Fuera de alcance
- Reservas con calendario / turnos online, pagos, precios.
- Blog, testimonios, multi-idioma.
- Panel de administración.
- Publicación (deploy) y dominio → spec 007.

## Dudas abiertas
- [DUDA] Dirección exacta y horarios de atención → *responsable:* cliente | *bloqueante:* no (hoy se informan por WhatsApp) → spec 002
- [DUDA] Permiso de las personas que aparecen en las fotos de sesiones → *responsable:* cliente | *bloqueante:* **sí para publicar** → spec 003
- [DUDA] ¿Se muestra precio o se informa por WhatsApp? → *responsable:* cliente | *bloqueante:* no
- [DUDA] ¿Tiene Instagram/Facebook para enlazar? → *responsable:* cliente | *bloqueante:* no
