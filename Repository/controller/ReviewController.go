package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReview())
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	err = repo.Books.CheckBook(review)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	reviewID := repo.Reviews.CreateNewReview(review)
	book, err := repo.Books.FindBookById(int64(review.BookId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	result := repo.Reviews.AverageRating()
	book.Rating = float32(result[int64(review.BookId)])
	// newRate := repo.Books.GetRateBook(review)
	return c.SendString(fmt.Sprintf("New reivew is created successfully with id = %d", reviewID))
}

func AverageRating(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	_, err := repo.Books.FindBookById(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Not found book for this id",
		})
	}
	result := repo.Reviews.AverageRating()
	return c.JSON(result[int64(id)])

}

func DelReviewByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	} else {
		book, err := repo.Books.FindBookById(int64(id))
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		result := repo.Reviews.AverageRating()
		book.Rating = float32(result[int64(id)])
		return c.SendString("delete successfully")
	}

}
