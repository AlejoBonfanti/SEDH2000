package site

import (
	"strings"
	"testing"
)

func form(values map[string]string) func(string) string {
	return func(k string) string { return values[k] }
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name     string
		values   map[string]string
		wantErrs []string
	}{
		{"vacío", map[string]string{}, []string{"nombre", "telefono"}},
		{"nombre de 1 letra", map[string]string{"nombre": "A", "telefono": "3496551234"}, []string{"nombre"}},
		{"teléfono corto", map[string]string{"nombre": "Ana", "telefono": "123-45"}, []string{"telefono"}},
		{"válido con espacios", map[string]string{"nombre": "  Ñu ", "telefono": "3496 55 1234"}, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			errs := NewContactData(form(c.values)).Validate()
			if len(errs) != len(c.wantErrs) {
				t.Fatalf("errores = %v, quería %v", errs, c.wantErrs)
			}
			for _, k := range c.wantErrs {
				if _, ok := errs[k]; !ok {
					t.Errorf("falta error en %q", k)
				}
			}
		})
	}
}

func TestHorarioInvalidoVuelveAlDefault(t *testing.T) {
	d := NewContactData(form(map[string]string{"horario": "Medianoche"}))
	if d.Horario != "Indistinto" {
		t.Fatalf("horario = %q", d.Horario)
	}
}

func TestMessageYWaLink(t *testing.T) {
	d := NewContactData(form(map[string]string{
		"nombre": "Ana", "telefono": "3496551234", "horario": "Tarde", "mensaje": "¿Hay turnos & precios?",
	}))
	msg := d.Message()
	if strings.Contains(msg, "Localidad") {
		t.Error("no debería incluir localidad vacía")
	}
	link := WaLink(msg)
	if !strings.HasPrefix(link, "https://wa.me/5493496440587?text=") {
		t.Fatalf("link = %s", link)
	}
	if strings.Contains(link, "+") || strings.Contains(link, " ") || strings.Contains(link, "&precios") {
		t.Errorf("link mal codificado: %s", link)
	}
}
