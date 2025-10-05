package main

import (
    //"fmt"
    "github.com/gofiber/fiber/v2"
    "time"
)

type jsonTest struct {
    Message string `json:"message"`
    Timestamp string `json:"timestamp"`
}

// https://gobyexample.com/structs
func newJsonTest(message string) *jsonTest {
    j := jsonTest{Message: message}
    // https://stackoverflow.com/questions/32015364/timestamps-in-golang
    j.Timestamp = time.Now().Format(time.RFC3339)
    return &j
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(newJsonTest("My name is Stephen"))
	})

	app.Listen(":3000")
}


