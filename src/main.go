// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"

	"github.com/Naoji1984/go-rana8/config"
	"github.com/Naoji1984/go-rana8/listener"
	"github.com/labstack/echo"
)

func main() {
	config.LoadConfig()
	e := echo.New()
	listener.Init(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Config.AppPort)))
}
