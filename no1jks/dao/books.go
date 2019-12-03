package dao

import (
	"github.com/jinzhu/gorm"
	"math"
	"no1jks/no1jks/models"
)

const numPerLine int = 4

type BookSet struct {
	DataSet
	Books [][]models.Book
}

func BooksBaseFilter(db *gorm.DB) *gorm.DB {
	return db.Where("book.is_deleted = ?", models.False)
}

func BooksHomepageFilter(db *gorm.DB) *gorm.DB {
	return db.Where("book.display_homepage = ?", models.True)
}

func assembleBookSet(rawBooks *[]models.Book) *BookSet {
	var bootSet BookSet
	length := len(*rawBooks)
	if length == 0 {
		return &bootSet
	}
	lines := int(math.Ceil(float64(length / numPerLine)))
	for i := 0; i < lines; i++ {
		start, end := i*numPerLine, i*numPerLine+numPerLine
		if i == lines-1 {
			slice := (*rawBooks)[start:]
			bootSet.Books = append(bootSet.Books, slice)
		} else {
			slice := (*rawBooks)[start:end]
			bootSet.Books = append(bootSet.Books, slice)
		}
	}
	return &bootSet
}

// Book Homepage
func (d *Dao) GetBooksHomepage(page int, onlyCount bool, filters *map[string]interface{}) interface{} {
	var rawBooks []models.Book
	var totalCount int
	db := d.mysql.
		Table("book").
		Scopes(BooksBaseFilter)
	db.Count(&totalCount)
	if onlyCount {
		return totalCount
	}
	err := db.Order("book.display_homepage asc, book.create_at desc").
		Offset(getOffset(page)).
		Limit(models.Limit).
		Find(&rawBooks).Error
	if err != nil {
		panic(err)
	}
	bookSet := assembleBookSet(&rawBooks)
	(*bookSet).Page = page
	(*bookSet).TotalCount = totalCount
	return bookSet
}

// Homepage recommendation
func (d *Dao) GetHomepageBooks(limit uint8) *BookSet {
	var rawBooks []models.Book
	books := d.mysql.
		Scopes(BooksBaseFilter, BooksHomepageFilter).
		Find(&rawBooks)
	if err := books.Error; err != nil {
		panic(err)
	}
	return assembleBookSet(&rawBooks)
}
