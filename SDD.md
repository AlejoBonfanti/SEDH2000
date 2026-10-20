# SEDH 2000 — Metodología SDD (Spec-Driven Development)

> **Regla de oro:** no se escribe código sin una spec aprobada.
> La spec dice **qué** y **por qué**; el plan dice **cómo**; las tareas dicen **en qué orden**.

Este archivo es la guía de trabajo del proyecto. Sirve tanto para vos como para cualquier asistente de IA
(Claude, Copilot, etc.): antes de pedir código, se le pasa la spec correspondiente.

---

## 1. El flujo

```
 0. Constitución   →  reglas fijas del proyecto (se escribe una vez, se toca poco)
 1. Spec           →  QUÉ se construye y POR QUÉ (sin hablar de código)
 2. Clarificar     →  preguntas abiertas resueltas con el cliente
 3. Plan           →  CÓMO: archivos, módulos, datos, decisiones técnicas
 4. Tareas         →  lista chica, ordenada y verificable
 5. Implementar    →  una tarea a la vez, commit por tarea
 6. Verificar      →  contra los criterios de aceptación de la spec
 7. Cerrar         →  actualizar estado + registro de cambios
```

| Fase | Pregunta que responde | Quién decide | Sale cuando… |
|---|---|---|---|
| Spec | ¿Qué necesita el cliente/usuario? | Cliente + vos | Tiene criterios de aceptación medibles |
| Clarificar | ¿Qué falta saber? | Cliente | No quedan `[DUDA]` bloqueantes |
| Plan | ¿Cómo lo construyo? | Vos | Cada requisito tiene dónde vivir en el código |
| Tareas | ¿En qué orden? | Vos | Cada tarea entra en ≤ 2 h y tiene un "listo cuando" |
| Implementar | — | Vos / IA | Todas las tareas tildadas |
| Verificar | ¿Cumple la spec? | Vos (+ cliente en demo) | Todos los criterios ✅ |

**Estados de una spec:** `Borrador` → `En revisión` → `Aprobada` → `En desarrollo` → `Hecha` (o `Descartada`).

---

## 2. Organización de carpetas

```
SEDH2000/
├── SDD.md                     ← este archivo (método + constitución + índice)
├── specs/
│   ├── 001-landing-base/
│   │   ├── spec.md            ← qué y por qué
│   │   ├── plan.md            ← cómo
│   │   └── tasks.md           ← checklist
│   ├── 002-<nombre-corto>/
│   └── ...
├── info/                      ← material del cliente (fuente, no se publica salvo fotos sedh2000-*)
├── src/ …                     ← código
└── README.md                  ← cómo correr el proyecto
```

- Numeración correlativa de 3 dígitos + nombre corto en kebab-case.
- Una spec = una funcionalidad entregable (no "toda la web").
- Rama de git sugerida: `spec/002-nombre-corto`. Commits: `[002] T3: descripción`.

---

## 3. Constitución del proyecto (reglas que ninguna spec puede romper)

### 3.1 Contenido y legal
1. **Sin promesas de salud.** Nada de "cura", "elimina", "desintoxica", "trata" ni beneficios clínicos.
   El PDF técnico del cliente es material interno: **no se copia texto de ahí a la web.**
   Se describe la experiencia de la sesión, no resultados médicos.
2. **Sección "Antes de reservar" siempre visible** con el aviso de que no reemplaza una consulta médica.
3. **Fotos de personas** solo con permiso confirmado del cliente.
4. Datos de contacto, dirección y horarios: **solo los confirmados por el cliente**, centralizados en `src/config.ts`.

### 3.2 Técnica
5. Stack: **HTML + CSS + TypeScript + Vite**. Sin frameworks salvo que una spec lo justifique en su plan.
6. TypeScript en modo `strict`; `npm run build` tiene que pasar sin errores antes de cerrar una spec.
7. Un módulo por comportamiento en `src/modules/` (nav, faq, reveal, contactForm…).
8. Textos en español rioplatense, tuteo/voseo consistente con lo existente.

### 3.3 Calidad mínima (aplica a toda spec)
9. Responsive: se prueba en 360 px, 768 px y 1280 px.
10. Accesibilidad: contraste AA, `alt` en imágenes, navegable con teclado, `prefers-reduced-motion` respetado.
11. Performance: imágenes optimizadas (WebP/AVIF o JPG comprimido, `loading="lazy"` debajo del hero).
12. Funciona sin JS para leer el contenido (JS mejora, no es requisito para ver la info).

> Cambiar la constitución requiere anotarlo en el **Registro de decisiones** (sección 7).

---

## 4. Plantillas

### 4.1 `spec.md`

```markdown
# [NNN] Nombre de la funcionalidad

**Estado:** Borrador | **Fecha:** AAAA-MM-DD | **Pedido por:** Cliente / interno

## Problema / objetivo
Qué pasa hoy y qué queremos lograr. 2–4 líneas. Sin hablar de código.

## Usuarios
- Quién lo usa y en qué situación (ej: "persona que llega desde Instagram en el celular").

## Historias de usuario
- **H1** — Como <usuario>, quiero <acción>, para <beneficio>.
- **H2** — …

## Requisitos funcionales
- **RF1** — El sistema debe…
- **RF2** — …

## Requisitos no funcionales
- Rendimiento / accesibilidad / SEO / legal (además de la constitución).

## Criterios de aceptación (verificables)
- [ ] **CA1** — Dado <contexto>, cuando <acción>, entonces <resultado>.
- [ ] **CA2** — …

## Fuera de alcance
- Lo que explícitamente NO entra en esta spec.

## Dudas abiertas
- [DUDA] … → *responsable:* cliente | *bloqueante:* sí/no
```

### 4.2 `plan.md`

```markdown
# [NNN] Plan técnico

**Spec:** ./spec.md

## Enfoque
Resumen de la solución en 3–5 líneas.

## Cambios por archivo
| Archivo | Cambio | Cubre |
|---|---|---|
| index.html | Nueva sección `#…` | RF1 |
| src/modules/x.ts | Nuevo módulo | RF2 |
| src/styles/main.css | Estilos de … | RNF |

## Datos / configuración
Tipos, interfaces, valores en `config.ts`.

## Decisiones y alternativas
- Elegí X en vez de Y porque…

## Riesgos
- …
```

### 4.3 `tasks.md`

```markdown
# [NNN] Tareas

- [ ] **T1** — Descripción corta · *listo cuando:* … · *cubre:* RF1
- [ ] **T2** — … · *depende de:* T1
- [ ] **T3** — Verificación: recorrer CA1…CAn en mobile y desktop
- [ ] **T4** — `npm run build` sin errores + actualizar índice en SDD.md
```

---

## 5. Cómo usar SDD con una IA

1. **Escribir la spec vos** (o pedirle a la IA un borrador y corregirlo). La IA no decide alcance.
2. Pedir el plan pasando: `SDD.md` (constitución) + `spec.md`. Revisarlo antes de seguir.
3. Pedir las tareas a partir del plan.
4. Implementar **de a una tarea**: *"Implementá T2 de specs/002-…/tasks.md. No toques nada fuera de lo que dice el plan."*
5. Si la IA propone algo que no está en la spec → se agrega a la spec primero, o se descarta.
6. Al terminar, pedir una revisión: *"Verificá la implementación contra los criterios de aceptación de spec.md y la constitución de SDD.md."*

---

## 6. Índice de specs

| N° | Nombre | Estado | Notas |
|---|---|---|---|
| 001 | [Landing base](specs/001-landing-base/spec.md) | Hecha (retroactiva) | Falta CA9 (sin JS) y CA10 (prueba en celulares reales). |
| 002 | Datos de atención confirmados | Borrador | Dirección y horarios. Espera respuesta del cliente. |
| 003 | Permisos de imágenes | Borrador | **Bloquea la publicación.** Si no hay permiso, recortar/reemplazar fotos. |
| 004 | Optimización de imágenes | Borrador | `public/img/modulo.jpg` ya no existe ✅. Falta: WebP, `width`/`height`, tamaños responsive. ||
| 005 | SEO y metadatos | En desarrollo | Hecho: favicon, Open Graph + `og.jpg`, canonical, robots/sitemap. Falta: datos estructurados y verificar la vista previa en WhatsApp. |
| 006 | Deploy | Borrador | Hosting + dominio + HTTPS. |
| 007 | Robustez sin JavaScript | Borrador | `.reveal` y FAQ ocultan contenido sin JS; links de contacto vacíos sin JS. |
| 008 | Medición de consultas | Idea | Contar clics a WhatsApp / envíos del formulario. |

> Ideas sueltas que todavía no son spec van en el **Backlog** (sección 8).

---

## 7. Registro de decisiones (ADR corto)

| Fecha | Decisión | Motivo |
|---|---|---|
| 2026-09 | HTML + CSS + TS con Vite, sin framework | Landing simple; aprender TypeScript; build liviano |
| 2026-09 | Contacto vía WhatsApp (wa.me) en lugar de backend | Canal que ya usa el cliente; sin servidor que mantener |
| 2026-09 | Datos de contacto centralizados en `src/config.ts` | Un solo lugar para editar |
| 2026-10 | Adopción de SDD | Ordenar pedidos del cliente y el trabajo con IA |
| 2026-10 | Spec 001 documentada de forma retroactiva | La landing se construyó antes de adoptar SDD |

---

## 8. Backlog (ideas sin spec todavía)

- [ ] Galería de fotos de sesiones con lightbox
- [ ] Testimonios (solo con autorización escrita, sin claims de salud)
- [ ] Analítica simple (Plausible / GA4) para medir clics a WhatsApp
- [ ] Versión en inglés

---

## 9. Checklist de cierre de una spec

- [ ] Todos los criterios de aceptación verificados
- [ ] Constitución respetada (en especial 3.1 contenido/legal)
- [ ] Probado en 360 / 768 / 1280 px
- [ ] `npm run build` OK
- [ ] Estado de la spec → `Hecha` y actualizado el índice (sección 6)
- [ ] Decisiones nuevas anotadas en la sección 7
- [ ] Si cambió algo visible: avisar/mostrar al cliente
