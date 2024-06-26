package services

import (
	"github.com/sekibuuun/go_api/apperrors"
	"github.com/sekibuuun/go_api/models"
	"github.com/sekibuuun/go_api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)

	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
