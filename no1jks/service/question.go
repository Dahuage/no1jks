package service

import "no1jks/no1jks/dao"

func (s *Service) GetQuestionHomepage(page int, isLogin bool,
	filters *map[string]interface{}) *dao.QuestionHomepageDataSet {
	res := s.Dao.GetNewsHomepageNewsList(page, false, filters).(*dao.QuestionHomepageDataSet)
	return res
}
