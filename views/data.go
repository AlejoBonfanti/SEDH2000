// Package views tiene las plantillas templ de la landing.
package views

import "github.com/AlejoBonfanti/SEDH2000/site"

// Page es todo lo que necesita la página completa.
type Page struct {
	Year      int
	Analytics bool // script de Vercel Web Analytics (solo en Vercel)
	Form      FormState
}

// FormState es el estado del formulario de contacto.
type FormState struct {
	Data   site.ContactData
	Errors site.Errors
}

// NewForm devuelve el formulario vacío.
func NewForm() FormState {
	return FormState{Data: site.ContactData{Horario: site.Horarios[0]}, Errors: site.Errors{}}
}

func (f FormState) hasError(field string) bool { _, ok := f.Errors[field]; return ok }

// firstError es el primer campo con error, en el orden del formulario (para el autofocus).
func (f FormState) firstError() string {
	for _, k := range []string{"nombre", "telefono"} {
		if f.hasError(k) {
			return k
		}
	}
	return ""
}

type faqItem struct{ Q, A string }

var faqs = []faqItem{
	{"¿Cuánto dura una sesión?", "Entre 5 y 35 minutos, según el peso de cada persona. Sumale unos minutos para la preparación."},
	{"¿Se siente algo durante la sesión?", "El equipo trabaja con bobinas externas: no se colocan electrodos ni agujas. En la charla previa te explico qué vas a notar."},
	{"¿Tengo que preparar algo?", "Venir con los pies limpios. Si tenés callosidades, conviene retirarlas antes. Cualquier otra indicación te la doy al coordinar."},
	{"¿Cuántas sesiones se hacen?", "El protocolo del fabricante indica una sesión diaria durante un ciclo de 5 a 8 sesiones."},
	{"¿Qué puedo esperar?", "Escribime y te explico cómo funciona el equipo y respondo todas tus preguntas antes de que decidas."},
	{"¿Dónde se hacen las sesiones?", "En Esperanza, Santa Fe. Te paso la dirección y los horarios disponibles cuando coordinamos."},
}

type step struct{ Num, Title, Text string }

var steps = []step{
	{"01", "Charla previa", "Te cuento cómo funciona el equipo, respondo tus dudas y vemos juntas si la sesión es para vos."},
	{"02", "Preparación", "Te sentás cómodo, apoyás los pies en la batea con agua y se colocan las bobinas en los tobillos."},
	{"03", "Ajuste del equipo", "Se calibran los tres controles según tu peso, tu edad y tu contextura."},
	{"04", "Sesión", "Entre 5 y 35 minutos, según tu peso. Solo tenés que relajarte."},
}

type photo struct{ Src, Alt string }

var gallery = []photo{
	{"/img/sedh2000-sesion-1.jpg", "Persona sentada en un sillón con los pies en las bateas y las bobinas en los tobillos, junto al equipo"},
	{"/img/sedh2000-sesion-2.jpg", "Vista de frente de una sesión: pies en las dos bateas con las bobinas colocadas"},
	{"/img/sedh2000-sesion-3.jpg", "Sesión vista desde el equipo, con el panel de control encendido en primer plano"},
}

// Atajos para usar en las plantillas.
func waDefault() string  { return site.WaLink(site.Config.DefaultMessage) }
func igURL() string      { return site.InstagramURL() }
func igHandle() string   { return "@" + site.Config.Instagram }
func waLabel() string    { return site.Config.WhatsAppLabel }
func horarios() []string { return site.Horarios }
