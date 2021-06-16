package routes

import (
	"github.com/TechMaster/golang/08Fiber/Repository/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigBookRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllBook)

	(*router).Get("/:id", controller.GetBookById)

	(*router).Delete("/:id", controller.DeleteBookById)

	(*router).Post("", controller.CreateBook)

	(*router).Put("", controller.UpdateBook)
}
func Setup(app *fiber.App) {
	app.Get("/review", controller.GetAllReview)
	app.Post("/review/add", controller.CreateReview)
	//lấy rate trung bình của sách = id
	app.Get("/average/:id", controller.AverageRating)
	app.Post("/review/del/:id", controller.DelReviewByID)
}
