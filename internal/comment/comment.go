package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - define the struct for our comment
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

//CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(Slug string) ([]Comment, error)
	PostComment(comment, Comment) (Comment, error)
	UpdateComment(ID uint, newComment comment) (comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	returns & Service{
		DB: db,
	}
}
