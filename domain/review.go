package domain

import "github.com/google/uuid"

type Review struct {
	id      string
	user    User
	product Product
	comment string
	rating  int
}

func NewReview(user User, product Product, comment string, rating int) *Review {
	return &Review{
		id:      uuid.New().String(),
		user:    user,
		product: product,
		comment: comment,
		rating:  rating,
	}
}

func (r *Review) AddComment(comment string) string {
	r.comment = comment
	return r.comment
}

func (r *Review) AddRating(rating int) int {
	r.rating = rating
	return r.rating
}

func (r *Review) GetComment() string {
	return r.comment
}

func (r *Review) GetRating() int {
	return r.rating
}