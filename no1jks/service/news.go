package service

import (
	"no1jks/no1jks/dao"
)

func (s *Service) GetNewsHomepage(IsLogin bool, page int, filters *map[string]interface{}) (news *dao.NewsHomepageSet) {
	news = s.dao.GetNewsHomepage(page, false, nil).(*dao.NewsHomepageSet)
	return news
}
