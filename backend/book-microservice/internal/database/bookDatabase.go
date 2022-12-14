package database

import (
	"book-microservice/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func CreateDatabase(books []model.Book) error {
	var err error

	dsn := "host=localhost user=postgres password=admin dbname=crypto-bazaar-book-db port=5432 sslmode=disable"
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	database.Migrator().DropTable("Books")

	database.AutoMigrate(&books)

	for indx := range books {
		database.Create(&books[indx])
	}

	return nil
}

func GetAll(pagination Pagination) *Pagination {

	var rows []*model.Book

	database.Scopes(paginate(rows, &pagination, database)).Find(&rows)
	pagination.Rows = rows

	return &pagination
}

func GetById(id uint) *model.Book {

	var book *model.Book
	database.First(&book, id)

	if book.ID != id {
		return nil
	}
	return book
}

func Create() {

}

func Update() {

}

func Delete() {

}
