// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Print current process
	if fiber.IsChild() {
		fmt.Printf("[%d] Child", os.Getppid())
	} else {
		fmt.Printf("[%d] Master", os.Getppid())
	}

	// Fiber instance
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))

	// Run the following command to see all processes sharing port 3000:
	// sudo lsof -i -P -n | grep LISTEN
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
