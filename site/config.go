// Package site tiene los datos y la lógica de la landing, sin nada de HTTP ni de HTML.
package site

import (
	"net/url"
	"regexp"
	"strings"
)

// Contact son los datos de contacto de la clienta. Cambiá acá y se actualiza toda la página.
type Contact struct {
	// WhatsApp en formato internacional, solo dígitos (para wa.me).
	WhatsApp string
	// Cómo se muestra el número en pantalla.
	WhatsAppLabel string
	// Instagram sin la arroba.
	Instagram string
	// Mensaje por defecto del botón flotante y de los links directos.
	DefaultMessage string
}

var Config = Contact{
	WhatsApp:       "5493496440587",
	WhatsAppLabel:  "+54 9 3496 44-0587",
	Instagram:      "sedh2000.mt",
	DefaultMessage: "Hola! Vi la página del SEDH 2000 y quiero más información sobre las sesiones.",
}

var nonDigits = regexp.MustCompile(`\D`)

// WaLink arma el link de wa.me con el mensaje codificado.
// wa.me necesita los espacios como %20, no como "+".
func WaLink(message string) string {
	number := nonDigits.ReplaceAllString(Config.WhatsApp, "")
	text := strings.ReplaceAll(url.QueryEscape(message), "+", "%20")
	return "https://wa.me/" + number + "?text=" + text
}

// InstagramURL devuelve el link al perfil de Instagram.
func InstagramURL() string {
	return "https://www.instagram.com/" + Config.Instagram + "/"
}
