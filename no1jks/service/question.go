package service

import (
	"github.com/astaxie/beego/logs"
	"no1jks/no1jks/dao"
	"no1jks/no1jks/models"
)

func (s *Service) GetQuestionHomepage(page int, isLogin bool,
	filters *map[string]interface{}) (*dao.QuestionHomepageDataSet, *Page) {
	res := s.Dao.GetNewsHomepageNewsList(page, false, filters).(*dao.QuestionHomepageDataSet)
	pager := MakePager(page, (*res).TotalCount, "/question?page=")
	return res, pager
}

func (s *Service) GetQuestionDetail(questionId int, other *map[string]interface{}) *dao.QuestionSet {
	question := s.Dao.GetNewsDetail(questionId, other)
	if question != nil {
		var question models.Question
		db := s.Dao.Mysql.First(&question, questionId)
		if db.Error != nil {
			logs.Info("query err", db.Error)
		} else {
			db = s.Dao.Mysql.Model(&question).Update("view_count", question.ViewCount + 1)
		}
	}
	return question
}