package main

import (
	"gofr.dev/pkg/gofr"
)

const defaultStaticFilePath = `./static`

func main() {
	app := gofr.New()

	staticFilePath := app.Config.GetOrDefault("STATIC_DIR_PATH", defaultStaticFilePath)

	app.AddStaticFiles("/", staticFilePath)

	app.Run()
}
