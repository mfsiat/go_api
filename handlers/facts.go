package handlers
import (
	"github.com/gofiber/fiber/v2"
	"github.com/mfsiat/go_api/models"
	"github.com/mfsiat/go_api/database"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Siat")
}

func ListFacts(c *fiber.Ctx) error {
	// return c.SendString("Hello, Siat")
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// persist the db
	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}