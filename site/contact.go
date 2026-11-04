package site

import (
	"strings"
	"unicode/utf8"
)

// ContactData son los datos que junta el formulario de contacto.
type ContactData struct {
	Nombre    string
	Telefono  string
	Localidad string
	Horario   string
	Mensaje   string
}

// Errors guarda un mensaje por campo con error (la clave es el name del input).
type Errors map[string]string

// Horarios son las opciones válidas del select. La primera es la que va por defecto.
var Horarios = []string{"Indistinto", "Mañana", "Tarde"}

// Límites para que nadie mande un texto gigante.
const (
	maxCorto   = 120
	maxMensaje = 1000
)

// NewContactData limpia lo que llega del formulario: saca espacios, recorta largos
// y fuerza un horario válido.
func NewContactData(get func(string) string) ContactData {
	d := ContactData{
		Nombre:    clip(get("nombre"), maxCorto),
		Telefono:  clip(get("telefono"), maxCorto),
		Localidad: clip(get("localidad"), maxCorto),
		Horario:   clip(get("horario"), maxCorto),
		Mensaje:   clip(get("mensaje"), maxMensaje),
	}
	if !validHorario(d.Horario) {
		d.Horario = Horarios[0]
	}
	return d
}

// Validate aplica las mismas reglas que tenía contactForm.ts.
func (d ContactData) Validate() Errors {
	errs := Errors{}
	if utf8.RuneCountInString(d.Nombre) < 2 {
		errs["nombre"] = "Contame tu nombre."
	}
	if len(nonDigits.ReplaceAllString(d.Telefono, "")) < 8 {
		errs["telefono"] = "Ingresá un teléfono válido."
	}
	return errs
}

// Message arma el texto que se manda por WhatsApp.
func (d ContactData) Message() string {
	lines := []string{
		"Hola! Quiero coordinar una sesión con el SEDH 2000.",
		"Nombre: " + d.Nombre,
		"Teléfono: " + d.Telefono,
	}
	if d.Localidad != "" {
		lines = append(lines, "Localidad: "+d.Localidad)
	}
	lines = append(lines, "Horario preferido: "+d.Horario)
	if d.Mensaje != "" {
		lines = append(lines, "Consulta: "+d.Mensaje)
	}
	return strings.Join(lines, "\n")
}

func validHorario(h string) bool {
	for _, v := range Horarios {
		if v == h {
			return true
		}
	}
	return false
}

func clip(s string, max int) string {
	s = strings.TrimSpace(s)
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	return strings.TrimSpace(string([]rune(s)[:max]))
}
