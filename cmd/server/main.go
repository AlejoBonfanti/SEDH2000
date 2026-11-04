// Servidor local para desarrollo: sirve public/ y la página.
//
//	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./cmd/server"
//
// En producción no se usa: en Vercel corre api/index.go y los estáticos los sirve el CDN.
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AlejoBonfanti/SEDH2000/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	static := http.FileServer(http.Dir("public"))
	app := server.New()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Si existe el archivo en public/, se sirve tal cual (css, js, img, favicon…).
		if r.URL.Path != "/" {
			if info, err := os.Stat(filepath.Join("public", filepath.FromSlash(r.URL.Path))); err == nil && !info.IsDir() {
				static.ServeHTTP(w, r)
				return
			}
		}
		app.ServeHTTP(w, r)
	})

	log.Printf("SEDH 2000 en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
