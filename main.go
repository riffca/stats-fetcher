// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


func main() {
	
	
	createConnection()
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error  {
		return c.SendString("Hello, World 👋!")
	})
	
	type LeadResponse struct {
		Ip string `json:"ip" xml:"ip" form:"ip"`
	}
	
	app.Post("/info", func(c *fiber.Ctx) error {
		r := new(LeadResponse)
		
    if err := c.BodyParser(r); err != nil {
      return err
    }
		
		//fetchLeadInfo(r.Ip)
		
	  return c.JSON(r.Ip)
	})

	
	app.Post("/proceed", func(c *fiber.Ctx) error {
		
		r := new(LeadResponse)
		
    if err := c.BodyParser(r); err != nil {
      return err
    }
		
		saveLeadStats(r.Ip)
		
		//info := getServerInfo(r.Ip)
		//saveServerInfo(r.Ip, info)
		
	  return c.JSON(r.Ip)
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}


