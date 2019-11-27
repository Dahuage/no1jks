package service

func (s *Service) GetHomeContent(IsLogin bool) *map[string]interface{} {
	HomeData := make(map[string]interface{})
	HomeData["News"] = s.dao.GetNewsByID()
	return &HomeData
}
