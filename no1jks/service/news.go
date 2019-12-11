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
	// TODO get ad
	// set autocommit false
	// 高并发下一定会出岔子 for update
	// 更好的解决方案是异步+1
	news := s.Dao.GetNewsByID(newsId)
	if news != nil {
		s.Dao.Mysql.Model(news).Update("view_count", news.ViewCount + 1)
	}
	comment := dao.NewsCommentSet{}
	ret.News = news
	ret.NewsComment = &comment
	return &ret
}
