package service

import "no1jks/no1jks/models"

func (s *Service) GetHomeContent(IsLogin bool) *map[string]interface{} {
	HomeData := make(map[string]interface{})
	HomeData["News"] = s.dao.GetHomePageNews(models.HomepageLimit)
	return &HomeData
}
