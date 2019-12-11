package service

import (
	"fmt"
	"math"
	"no1jks/no1jks/models"
)

type Page struct {
	CurrentPage int
	TotalPage   int
	IsFirst     string
	IsLast      string
	PrePage     string
	NextPage    string
}

func getTotalPage(totalCount int) int {
	return int(math.Ceil(float64(totalCount) / float64(models.Limit)))
}

func MakePager(page int, totalCount int, pagePrefix string) *Page {
	var pager Page
	totalPage := getTotalPage(totalCount)
	pager.TotalPage = totalPage
	pager.CurrentPage = page

	if page == 0 {
		pager.IsFirst = "disabled"
		pager.PrePage = ""
	}else{
		pager.IsFirst = ""
		pager.PrePage = pagePrefix + fmt.Sprintf("%d", page - 1)
	}
	if page == totalPage - 1 {
		pager.IsLast = "disabled"
		pager.NextPage = ""
	}else{
		pager.IsLast = ""
		pager.NextPage = pagePrefix + fmt.Sprintf("%d", page + 1)
	}
	return &pager
}
