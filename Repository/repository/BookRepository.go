package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type BookRepo struct {
	books  map[int64]*model.Book
	autoID int64
}

var Books BookRepo

func init() {
	Books = BookRepo{autoID: 0}
	Books.books = make(map[int64]*model.Book)
	Books.InitData("sql:45312")
}

func (r *BookRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *BookRepo) CreateNewBook(book *model.Book) int64 {
	nextID := r.getAutoID()
	book.Id = nextID
	r.books[nextID] = book
	return nextID
}

func (r *BookRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewBook(&model.Book{
		Title: "Dế Mèn Phiêu Lưu Ký",
		Authors: []model.Author{
			{FullName: "Tô Hoài", Country: "Vietnam"},
			{FullName: "Hames", Country: "Turkey"},
		},
		Rating: 0})

	r.CreateNewBook(&model.Book{
		Title: "100 năm cô đơn",
		Authors: []model.Author{
			{FullName: "Gabriel Garcia Marquez", Country: "Columbia"},
			{FullName: "Ivan", Country: "Russia"},
		},
		Rating: 0})
}

func (r *BookRepo) GetAllBooks() map[int64]*model.Book {
	return r.books
}

func (r *BookRepo) FindBookById(Id int64) (*model.Book, error) {
	if book, ok := r.books[Id]; ok {
		return book, nil
	} else {
		return nil, errors.New("book not found")
	}
}

func (r *BookRepo) DeleteBookById(Id int64) error {
	if _, ok := r.books[Id]; ok {
		delete(r.books, Id)
		return nil
	} else {
		return errors.New("book not found")
	}
}

func (r *BookRepo) UpdateBook(book *model.Book) error {
	if _, ok := r.books[book.Id]; ok {
		r.books[book.Id] = book
		return nil //tìm được
	} else {
		return errors.New("book not found")
	}
}

//review
type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
}
func (r *ReviewRepo) GetAllReview() map[int64]*model.Review {
	return r.reviews
}
func (r *ReviewRepo) getAutoIDRe() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoIDRe()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

func (r *BookRepo) CheckBook(review *model.Review) error {
	if _, ok := r.books[review.BookId]; ok {
		return nil //tìm được
	} else {
		return errors.New("book not found")
	}
}

func (r *ReviewRepo) AverageRating() (result map[int64]float32) {
	sum := make(map[int64]int)
	number := make(map[int64]int)
	result = make(map[int64]float32)

	for _, value := range r.reviews {
		number[value.BookId]++
		sum[value.BookId] += value.Rating
	}
	for key := range number {
		result[key] = float32(sum[key]) / float32(number[key])
	}
	return result
}
func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}
