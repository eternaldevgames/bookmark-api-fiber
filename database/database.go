package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Bookmark{})

	return nil
}

func GetAllBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return bookmarks, err
	}

	db.Find(&bookmarks)

	return bookmarks, nil
}

func CreateBookmark(name string, url string) (Bookmark, error) {
	var newBookmark = Bookmark{Name: name, Url: url}

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return newBookmark, err
	}
	db.Create(&Bookmark{Name: name, Url: url})

	return newBookmark, nil
}
