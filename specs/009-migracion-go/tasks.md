# [009] Tareas

- [x] **T1** — Spec, plan y tareas + registro de decisiones y constitución en SDD.md
- [x] **T2** — `go.mod`, paquete `site` (config, validación, mensaje) con tests · *cubre:* RF3, RF7
- [x] **T3** — Plantillas templ: layout, secciones, formulario · *cubre:* RF1, RF4, RF5
- [x] **T4** — `server` (rutas GET / y POST /contacto) con tests · *cubre:* RF3–RF5
- [x] **T5** — `public/`: css, `app.js` (nav, reveal, faq), htmx, fotos · *cubre:* RF6
- [x] **T6** — `cmd/server` (local) y `api/index.go` + `vercel.json` (Vercel)
- [x] **T7** — Verificación local: CA1–CA7
- [ ] **T8** — Instalar Go y templ en Windows y correr el proyecto local · *listo cuando:* `http://localhost:8080` muestra la página
- [ ] **T9** — Push de `dev` y revisar el preview de Vercel · *cubre:* CA8, CA9
- [x] **T10** — Borrar el stack viejo: `index.html`, `src/`, `vite.config.ts`, `tsconfig.json`, `package*.json`, `node_modules/`, `dist/` (hecho 2026-11-03, antes de T9; si hay que volver, está en `main`) · *.gitignore* actualizado para Go
- [ ] **T11** — Probar en Android e iPhone reales · *cubre:* CA10
- [ ] **T12** — Merge a `main` + estado → Hecha en SDD.md · *depende de:* T9–T11
