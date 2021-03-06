package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/azf-hello-go/internal/adapter/rest/handler"
)

func main() {
	app := fiber.New()

	app.Get("/api/hello", handler.GetHello)

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(app.Listen(listenAddr))
}
