package service

import (
	"no1jks/no1jks/models"
)

func (s *Service) GetHomeContent(IsLogin bool) *map[string]interface{} {
	HomeData := make(map[string]interface{})
	HomeData["News"] = s.dao.GetHomepageNews(models.HomepageLimit)
	HomeData["Questions"] = s.dao.GetHomepageQuestions(models.HomepageLimit)
	HomeData["Books"] = s.dao.GetHomepageBooks(models.HomepageLimit)
	HomeData["Blog"] = s.dao.GetHomepageBlog(models.HomepageLimit)
	// Oh, trying print a value deep in the data, hell the map.
	// logs.Info(*((*((*(HomeData["Questions"].(*[]*map[string]interface{})))[0]))
	// 			 ["Answers"].([]*map[string]interface{})[0]))
	return &HomeData
}
