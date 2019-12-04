package service

import (
	"no1jks/no1jks/models"
)

func (s *Service) GetHomeContent(IsLogin bool) *map[string]interface{} {
	HomeData := make(map[string]interface{})
	HomeData["News"] = s.Dao.GetHomepageNews(models.HomepageLimit)
	HomeData["Questions"] = s.Dao.GetHomepageQuestions(models.HomepageLimit)
	HomeData["Books"] = s.Dao.GetHomepageBooks(4)
	HomeData["Blog"] = s.Dao.GetHomepageBlog(models.HomepageLimit)
	return &HomeData
}
