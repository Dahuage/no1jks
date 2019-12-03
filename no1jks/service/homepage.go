package service

import (
	"no1jks/no1jks/models"
)

func (s *Service) GetHomeContent(IsLogin bool) *map[string]interface{} {
	HomeData := make(map[string]interface{})
	HomeData["News"] = s.dao.GetHomepageNews(models.HomepageLimit)
	HomeData["Questions"] = s.dao.GetHomepageQuestions(models.HomepageLimit)
	HomeData["Books"] = s.dao.GetHomepageBooks(4)
	HomeData["Blog"] = s.dao.GetHomepageBlog(models.HomepageLimit)
	return &HomeData
}
