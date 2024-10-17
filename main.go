package main

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/hello", func(c *gofr.Context) (interface{}, error) {

		return fmt.Sprintf("Hello! Welcome to %s", c.GetAppName()), nil
	})

	app.GET("/logs", func(c *gofr.Context) (interface{}, error) {
		c.Logger.Debug("sample debug log")
		c.Logger.Info("sample info log")
		c.Logger.Warn("sample warn log")
		c.Logger.Errorf("sample error log")

		return "success", nil
	})

	app.Run()
}
