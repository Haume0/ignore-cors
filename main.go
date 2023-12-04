package main

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		return c.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(`
		"Welcome to my project 'Ignore CORS'!"
		I'm Haume!
		My github: https://github.com/haume0
		Project github: https://github.com/haume0/ignore-cors
		---------------------------------------
		In cases where you are not authorized to access the api that gives CORS error,
		your local development will be able to continue without interruption.
		---------------------------------------
		GET: http://localhost:8080/get?url=https://www.google.com
		`)
	})
	app.Get("/get", func(c *fiber.Ctx) error {
		url := c.Query("url")
		if url == "" {
			return c.JSON(fiber.Map{
				"message": "URL is required",
			})
		}
		res, err := http.Get(url)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "FETCH ERROR",
			})
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": "BODY READ ERROR",
			})
		}
		return c.SendString(string(body))
	})
	//LISTEN ON PORT 8080
	app.Listen(":8080")
}
