// Función serverless de Vercel. Vercel busca una función exportada en api/*.go
// con la firma de http.HandlerFunc. Todo lo demás vive en el paquete server.
package handler

import (
	"net/http"

	"github.com/AlejoBonfanti/SEDH2000/server"
)

var app = server.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
