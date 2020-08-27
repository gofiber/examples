package routes

import (
	"github.com/amalshaji/stoyle/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
)

// ResolveURL ...
func ResolveURL(c *fiber.Ctx) {
	// get the short from the url
	url := c.Params("url")
	// query the db to find the original URL, if a match is found
	// increment the redirect counter and redirect to the original URL
	// else return error message
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		_ = c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found on database",
		})
		return
	} else if err != nil {
		_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
		return
	}
	// increment the counter
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")
	// redirect to original URL
	c.Redirect(value, 301)
}
