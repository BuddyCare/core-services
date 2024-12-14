package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Crear una nueva instancia de Fiber
	app := fiber.New()

	// Ruta raíz con un handler que responde "Hola Mundo"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("hello world v2..., %s", os.Getenv("JWT_KEY")))
	})

	// Escuchar en el puerto 4000
	log.Println("Servidor corriendo en http://localhost:4000")
	log.Fatal(app.Listen(":4000"))
}
