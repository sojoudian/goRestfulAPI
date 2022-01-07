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

//GetComment - retrives comments by their ID from database
func (s *Service) GetComment(Id uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//GetCommentBySlug - retrives all comments by slug (path ~ /article/name)
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

//Post comment - adds new comments to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//UpdateComment -  updates a comment by ID with new comment information
func (s *Comment) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).UpdateComment(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

//DeleteComment - Deletes a comment from the database by ID
func (s *Service) DeleteComment(ID uint) error {
	if result := sDB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

//GetAllComments - retrives all comments from the database
func (s *Service) GetAllComments() ([]*Comment, error) {
	var comments []Comment
	if result := sDB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
