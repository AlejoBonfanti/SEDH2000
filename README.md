# SEDH 2000 — Landing

Landing page para coordinar sesiones con el equipo **SEDH 2000** en Esperanza, Santa Fe.
Página única: presenta el equipo, cómo es una sesión y lleva a coordinar por WhatsApp.

🔗 **[sedh-2000.vercel.app](https://sedh-2000.vercel.app)**

## Stack

| | |
|---|---|
| **Go** | Servidor y lógica del formulario |
| **[templ](https://templ.guide)** | Plantillas HTML tipadas, compiladas a Go |
| **[htmx](https://htmx.org)** | Envío del formulario sin recargar la página |
| **CSS + JS plano** | Estilos y detalles de interfaz (menú, animaciones, FAQ). Sin Node ni build |
| **Vercel** | Deploy como función serverless de Go; estáticos servidos por el CDN |

Algunas decisiones:

- **Funciona sin JavaScript.** El contenido se ve completo y el formulario hace un POST común: el servidor valida y redirige a WhatsApp. htmx solo mejora la experiencia (errores sin recargar).
- **Sin base de datos.** El contacto se resuelve armando un link de `wa.me` con el mensaje; no se guarda nada de lo que escribe la persona.
- **Un mismo handler** sirve en local (`cmd/server`) y en Vercel (`api/index.go`).

## Correrlo local

Requiere Go 1.24+.

```bash
go install github.com/a-h/templ/cmd/templ@v0.3.1001
templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./cmd/server"
```

Abrí **http://localhost:7331** (recarga sola al guardar un `.templ`). Tests: `go test ./...`

## Estructura

```
site/      datos de contacto, validación y armado del mensaje
views/     plantillas templ (*_templ.go se generan y se commitean)
server/    rutas: GET / y POST /contacto
cmd/       servidor local
api/       función serverless de Vercel
public/    css, js, imágenes y archivos estáticos
specs/     especificaciones del proyecto
```

## Metodología

El proyecto se trabaja con **Spec-Driven Development**: cada cambio arranca con una spec
(qué y por qué), un plan (cómo) y una lista de tareas. Las reglas y el índice están en
[`SDD.md`](SDD.md); por ejemplo, la migración de TypeScript + Vite a Go está en
[`specs/009-migracion-go`](specs/009-migracion-go/spec.md).

---

Desarrollado por [Alejo Bonfanti](https://github.com/AlejoBonfanti).
