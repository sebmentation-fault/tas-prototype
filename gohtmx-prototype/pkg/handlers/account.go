package handlers

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/accounts"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/components/headers"
)

// TODO: Implement the accountHandler function
//
// Handle the request to get account information
func accountButtonHandler(c *fiber.Ctx) error {
	// for now random account (some celebs and some not)

	// get random int either 0 and 1
	isCeleb := rand.Intn(2) == 0

	if isCeleb {
		acc := accounts.NewAccount("1", "John Doe", false, true, "2021-01-01", "2021-01-01")
		return RenderHTML(c, headers.AccountButton(acc))
	} else {
		acc := accounts.NewAccount("0", "Alice Bobina", false, false, "2021-01-01", "2021-01-01")
		return RenderHTML(c, headers.AccountButton(acc))
	}
}
