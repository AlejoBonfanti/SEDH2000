# [009] Migración a Go + templ + htmx

**Estado:** En desarrollo | **Fecha:** 2026-11-03 | **Pedido por:** interno

## Problema / objetivo
La landing está hecha en HTML + CSS + TypeScript con Vite. Queremos pasarla a **Go** con plantillas
**templ** y **htmx**, para aprender el stack y dejar la base lista para cosas que necesitan servidor
(por ejemplo, reservas). La persona que entra **no tiene que notar ninguna diferencia**, salvo que el
formulario de contacto ahora funciona también sin JavaScript.

## Usuarios
- Los mismos de [001](../001-landing-base/spec.md): la persona que entra desde el celular, el familiar o acompañante, y Mariana.
- Alejo, que mantiene el código.

## Historias de usuario
- **H1** — Como interesado/a, quiero ver la misma página que hoy, con los mismos textos, fotos y botones.
- **H2** — Como interesado/a, quiero que el formulario me avise si me equivoqué sin perder lo que escribí.
- **H3** — Como interesado/a con JS bloqueado o un celular viejo, quiero poder mandar el formulario igual.
- **H4** — Como Alejo, quiero cambiar el WhatsApp o el Instagram en un solo archivo Go.

## Requisitos funcionales
- **RF1** — Todo lo de 001 (RF1–RF11) se mantiene, con el mismo HTML y los mismos estilos.
- **RF2** — La página la arma el servidor Go con templ. No hay paso de build con Node.
- **RF3** — El formulario se valida **en el servidor**, con las mismas reglas: nombre ≥ 2 caracteres, teléfono ≥ 8 dígitos.
- **RF4** — Con JS, el formulario se envía con htmx y se reemplaza solo el formulario:
  si hay errores, se ven marcados y el foco va al primer campo con error; si está bien, se abre WhatsApp directo (header `HX-Redirect`).
- **RF5** — Sin JS, el formulario hace un POST común: si hay errores vuelve la página con los errores y los datos cargados; si está bien, redirige directo a WhatsApp.
- **RF6** — Menú móvil, animaciones al hacer scroll y acordeón del FAQ siguen funcionando, con JS plano (sin TypeScript ni build).
- **RF7** — Datos de contacto centralizados en `site/config.go`.
- **RF8** — Vercel Web Analytics sigue midiendo visitas.

## Requisitos no funcionales
- Se sigue publicando en Vercel, en la misma URL (sedh-2000.vercel.app), con previews por rama.
- Nada de lo que escribe la persona se guarda en el servidor: solo se usa para armar el link de WhatsApp.
- Límite de tamaño en lo que se manda (texto corto ≤ 120, consulta ≤ 1000 caracteres).
- Constitución 3.1 (contenido/legal) y 3.3 (calidad) sin cambios.

## Criterios de aceptación
- [x] **CA1** — Con JS desactivado, la página en 360 px y en 1280 px se ve igual que la versión Vite, hasta la sección de contacto (comparado con captura).
- [x] **CA2** — Dado nombre "A" y teléfono vacío, cuando envío con JS, entonces veo los dos errores sin que se recargue la página y el foco va a "Nombre".
- [x] **CA3** — Dado el formulario válido, cuando envío con JS, entonces el navegador va directo a wa.me con el número de `site/config.go` y el mensaje armado.
- [x] **CA4** — Dado el formulario válido sin JS, cuando envío, entonces el navegador va directo a WhatsApp.
- [x] **CA5** — Dado el formulario con errores sin JS, cuando envío, entonces vuelve la página con los errores y lo que había escrito.
- [x] **CA6** — Menú móvil (abrir, cerrar al tocar un link, Escape) y FAQ (una pregunta abierta a la vez) funcionan.
- [x] **CA7** — `go vet ./...` y `go test ./...` pasan.
- [x] **CA8** — El preview de Vercel de la rama `dev` responde en `/`, `/contacto` (POST), `/css/main.css`, `/og.jpg` y `/favicon.ico`. (probado por Alejo 2026-11-04)
- [ ] **CA9** — Vercel Analytics registra visitas en el preview.
- [ ] **CA10** — Probado en Android e iPhone reales (heredado de 001).

## Fuera de alcance
- Cambiar textos, diseño o secciones.
- Reservas, turnos o guardar consultas: eso va en su propia spec, ahora que hay servidor.
- Optimizar imágenes (spec 004).

## Dudas abiertas
- ~~¿Botón "Abrir WhatsApp" o redirección directa?~~ → Resuelto 2026-11-03: redirección directa con `HX-Redirect`.
