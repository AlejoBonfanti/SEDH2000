// Package server tiene el handler HTTP de la landing. Lo usan tanto el servidor
// local (cmd/server) como la función serverless de Vercel (api/index.go).
package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"

	"github.com/AlejoBonfanti/SEDH2000/site"
	"github.com/AlejoBonfanti/SEDH2000/views"
)

// En Vercel todo llega por la función /api/index (ver los rewrites de vercel.json),
// así que se aceptan las dos rutas.
const vercelFn = "/api/index"

// New devuelve el handler de la página.
func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet, http.MethodHead:
			if r.URL.Path != "/" && r.URL.Path != vercelFn {
				http.NotFound(w, r)
				return
			}
			// La página es igual para todos: Vercel la puede cachear en su CDN.
			w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=3600, stale-while-revalidate=86400")
			render(w, r, http.StatusOK, views.Home(page(views.NewForm())))

		case http.MethodPost:
			if r.URL.Path != "/contacto" && r.URL.Path != vercelFn {
				http.NotFound(w, r)
				return
			}
			contact(w, r)

		default:
			w.Header().Set("Allow", "GET, HEAD, POST")
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}

func contact(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulario inválido", http.StatusBadRequest)
		return
	}

	data := site.NewContactData(r.PostForm.Get)
	form := views.FormState{Data: data, Errors: data.Validate()}

	w.Header().Set("Cache-Control", "no-store")
	isHTMX := r.Header.Get("HX-Request") == "true"

	switch {
	case len(form.Errors) == 0 && isHTMX:
		// htmx y datos válidos: HX-Redirect hace que htmx navegue directo a WhatsApp.
		w.Header().Set("HX-Redirect", site.WaLink(data.Message()))
		w.WriteHeader(http.StatusOK)
	case len(form.Errors) == 0:
		// Sin JS y datos válidos: directo a WhatsApp.
		http.Redirect(w, r, site.WaLink(data.Message()), http.StatusSeeOther)
	case isHTMX:
		// htmx con errores: solo el fragmento del formulario con los errores marcados.
		render(w, r, http.StatusOK, views.ContactForm(form))
	default:
		// Sin JS y con errores: la página completa con los errores marcados.
		render(w, r, http.StatusUnprocessableEntity, views.Home(page(form)))
	}
}

func page(f views.FormState) views.Page {
	return views.Page{
		Year:      time.Now().Year(),
		Analytics: os.Getenv("VERCEL") == "1",
		Form:      f,
	}
}

func render(w http.ResponseWriter, r *http.Request, status int, c templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if r.Method == http.MethodHead {
		return
	}
	if err := c.Render(r.Context(), w); err != nil {
		log.Printf("render: %v", err)
	}
}
