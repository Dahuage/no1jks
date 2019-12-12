package service

import (
	"no1jks/no1jks/models"
	"no1jks/no1jks/utils"
)

func (s *Service) AnswerCreate(user *models.User, questionId int, conclusion, content string) (bool, *utils.ServiceErr) {
	ok := s.Dao.AnswerCreate(user.ID, questionId, conclusion, content)
	if ok {
		return ok, utils.Errs["UNKNOWN_ERROR"]
	}
	return ok, nil
}