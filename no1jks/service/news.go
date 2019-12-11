package service

import (
	"no1jks/no1jks/dao"
	"no1jks/no1jks/models"
)

type NewsDetail struct {
	News          *models.News
	NewsComment   *dao.NewsCommentSet
	NewsRecommend []models.News
	AD            *string
}

func (s *Service) GetNewsHomepage(IsLogin bool, page int, filters *map[string]interface{}) (news *dao.NewsHomepageSet, pager *Page) {
	news = s.Dao.GetNewsHomepage(page, false, nil).(*dao.NewsHomepageSet)
	pager = MakePager(page, (*news).TotalCount, "/news?page=")
	return news, pager
}

func (s *Service) GetNewsDetail(newsId int, other *map[string]interface{}) *NewsDetail {
	var ret NewsDetail
	// TODO set view + 1
	// TODO log user pict
	// TODO get ad
	news := s.Dao.GetNewsByID(newsId)
	comment := dao.NewsCommentSet{}
	ret.News = news
	ret.NewsComment = &comment
	return &ret
}
