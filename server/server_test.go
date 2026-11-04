package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func do(method, path string, form url.Values, htmx bool) *httptest.ResponseRecorder {
	var body *strings.Reader
	if form != nil {
		body = strings.NewReader(form.Encode())
	} else {
		body = strings.NewReader("")
	}
	req := httptest.NewRequest(method, path, body)
	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if htmx {
		req.Header.Set("HX-Request", "true")
	}
	rec := httptest.NewRecorder()
	New().ServeHTTP(rec, req)
	return rec
}

func TestHome(t *testing.T) {
	for _, path := range []string{"/", "/api/index"} {
		rec := do("GET", path, nil, false)
		if rec.Code != 200 {
			t.Fatalf("%s: status %d", path, rec.Code)
		}
		html := rec.Body.String()
		for _, want := range []string{
			`<html lang="es-AR">`, `id="contactForm"`, `hx-post="/contacto"`,
			`https://wa.me/5493496440587?text=`, "Antes de reservar", "© <span id=\"year\">",
		} {
			if !strings.Contains(html, want) {
				t.Errorf("%s: falta %q", path, want)
			}
		}
	}
	if rec := do("GET", "/otra", nil, false); rec.Code != 404 {
		t.Errorf("/otra: status %d", rec.Code)
	}
}

func TestContactoHTMX(t *testing.T) {
	rec := do("POST", "/contacto", url.Values{"nombre": {"A"}, "telefono": {"12"}}, true)
	html := rec.Body.String()
	if rec.Code != 200 || strings.Contains(html, "<html") {
		t.Fatalf("debería devolver solo el fragmento: %d", rec.Code)
	}
	if !strings.Contains(html, "Contame tu nombre.") || !strings.Contains(html, "Ingresá un teléfono válido.") {
		t.Errorf("faltan errores: %s", html)
	}
	if !strings.Contains(html, `id="nombre"`) || !strings.Contains(html, "autofocus") {
		t.Error("el primer campo con error debería tener autofocus")
	}

	rec = do("POST", "/api/index", url.Values{"nombre": {"Ana"}, "telefono": {"3496 55 1234"}, "horario": {"Tarde"}}, true)
	if loc := rec.Header().Get("HX-Redirect"); !strings.HasPrefix(loc, "https://wa.me/5493496440587?text=") {
		t.Errorf("debería redirigir a WhatsApp con HX-Redirect, llegó %q", loc)
	}
}

func TestContactoSinJS(t *testing.T) {
	rec := do("POST", "/contacto", url.Values{"nombre": {"Ana"}, "telefono": {"3496551234"}}, false)
	if rec.Code != http.StatusSeeOther || !strings.HasPrefix(rec.Header().Get("Location"), "https://wa.me/") {
		t.Fatalf("debería redirigir a WhatsApp: %d %s", rec.Code, rec.Header().Get("Location"))
	}

	rec = do("POST", "/contacto", url.Values{"nombre": {"Ana"}}, false)
	if rec.Code != http.StatusUnprocessableEntity || !strings.Contains(rec.Body.String(), "<html") {
		t.Fatalf("con errores debería volver la página completa: %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `value="Ana"`) {
		t.Error("debería conservar lo que escribió")
	}
}
