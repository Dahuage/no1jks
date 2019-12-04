package service

import "no1jks/no1jks/dao"

func (s *Service) GetBooksHomepage(page int, isLogin bool, filters *map[string]interface{}) *dao.BookSet {
	books := s.Dao.GetBooksHomepage(page, false, nil)
	return books.(*dao.BookSet)
}