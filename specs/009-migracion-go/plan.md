# [009] Plan técnico

**Spec:** ./spec.md

## Enfoque
Un solo `http.Handler` en Go (paquete `server`) arma la página con componentes templ (paquete `views`).
Lo usan dos entradas: `cmd/server` para desarrollo local y `api/index.go` como función serverless de
Vercel. Los estáticos (css, js, imágenes, favicon, og.jpg) viven en `public/` y en Vercel los sirve el
CDN directo, sin pasar por Go. htmx solo maneja el formulario de contacto.

## Estructura
```
api/index.go            → función de Vercel: Handler(w, r) que llama a server.New()
cmd/server/main.go      → servidor local (:8080): sirve public/ + server.New()
server/server.go        → rutas: GET / (página), POST /contacto (formulario)
site/config.go          → datos de contacto + WaLink()            (antes src/config.ts)
site/contact.go         → ContactData, Validate(), Message()      (antes contactForm.ts)
views/*.templ           → layout, secciones y formulario          (antes index.html)
views/*_templ.go        → generado por `templ generate` (se commitea: Vercel no corre templ)
public/css/main.css     → mismo CSS, sin el bloque .form-fallback  (antes src/styles/main.css)
public/js/app.js        → nav + reveal + faq en JS plano          (antes src/modules/*.ts)
public/js/htmx.min.js   → htmx 2.0.4 copiado en el repo (sin CDN)
public/img/*.jpg        → fotos sedh2000-*.jpg                    (antes info/, procesadas por Vite)
vercel.json             → sin framework ni build, output public/, rewrites a /api/index
```

## Cambios por archivo
| Archivo | Cambio | Cubre |
|---|---|---|
| `views/layout.templ` | `<head>`, header, footer, botón flotante; analytics solo si `VERCEL=1` | RF1, RF8 |
| `views/page.templ` | Secciones; pasos, galería y FAQ salen de slices en `views/data.go` | RF1 |
| `views/contact.templ` | `ContactForm` (página + fragmento htmx con errores) | RF3–RF5 |
| `server/server.go` | Ruteo, `MaxBytesReader`, `HX-Redirect` / 303 a WhatsApp / fragmento o 422 con errores | RF3–RF5 |
| `site/*.go` | Config, validación y armado del mensaje, con tests | RF3, RF7 |
| `public/js/app.js` | Port 1:1 de nav.ts, reveal.ts y faq.ts | RF6 |
| `public/css/main.css` | Saca `.form-fallback`; suma el estado `htmx-request` del botón | RF4, RF5 |
| `vercel.json` | Config de deploy | RNF |

## Rutas
| Método | Ruta | Respuesta |
|---|---|---|
| GET/HEAD | `/` | Página completa (`Cache-Control: s-maxage=3600`) |
| POST | `/contacto` + `HX-Request`, válido | `200` + `HX-Redirect: https://wa.me/…` (htmx navega a WhatsApp) |
| POST | `/contacto` + `HX-Request`, con errores | Fragmento del formulario con los errores |
| POST | `/contacto` sin htmx, válido | `303` → `https://wa.me/…` |
| POST | `/contacto` sin htmx, con errores | `422` + página completa con errores |

En Vercel, `/` y `/contacto` se reescriben a `/api/index`, así que el handler acepta también esa ruta.
El ruteo depende del método (GET = página, POST = formulario), y no de cómo llegue la ruta.

## Decisiones y alternativas
- **Vercel serverless en vez de un servidor común** (Fly/Render): misma URL, previews por rama y CDN para estáticos. A cambio, cada request es una función y no hay estado en memoria.
- **JS plano en vez de TypeScript**: son ~90 líneas; así no hace falta Node para nada.
- **htmx copiado en el repo** en vez de CDN: una dependencia externa menos en tiempo de carga.
- **`HX-Redirect` a wa.me al enviar por htmx**: abre WhatsApp en un paso, igual que sin JS. Se descartó abrir una pestaña nueva (`window.open` desde una respuesta asíncrona lo bloquean los navegadores) y un botón "Abrir WhatsApp" intermedio (un toque más). Contra: en desktop la pestaña pasa a wa.me; con "atrás" se vuelve a la página.
- **FAQ con el mismo JS** en vez de pasarlo a `<details name>`: mismo look, sin riesgo. Queda como mejora.
- **Archivos `_templ.go` commiteados**: Vercel compila Go pero no corre `templ generate`.

## Riesgos
- `buildCommand: ""` en `vercel.json` tiene que pisar el preset Vite del proyecto. Si el preview falla, en Vercel → Settings → Build & Development poner Framework Preset = **Other**.
- La versión del CLI de templ tiene que coincidir con la de `go.mod` (v0.3.1001); si no, `templ generate` avisa y el código generado puede no compilar.
