package main

import (
	"gofr.dev/pkg/gofr"
	"net/http"
	"path/filepath"
)

const defaultStaticFilePath = `./static`

func main() {
	app := gofr.New()

	staticFilePath := app.Config.GetOrDefault("STATIC_DIR_PATH", defaultStaticFilePath)

	app.UseMiddleware(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)

			filePath := filepath.Join(staticFilePath, "404.html")

			http.ServeFile(w, r, filePath)

			return
		})
	})

	app.AddStaticFiles("/", staticFilePath)

	app.Run()
}
