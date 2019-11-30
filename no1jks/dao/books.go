package dao

import "no1jks/no1jks/models"

type BookSet struct {
	DataSet
	Books []*models.Book
}

func (d *Dao)GetHomepageBooks(limit uint8) *BookSet {
	var Books BookSet

	books := d.mysql.Find(&Books.Books)
	if err := books.Error; err != nil {
		panic(err)
	}
	return &Books
}