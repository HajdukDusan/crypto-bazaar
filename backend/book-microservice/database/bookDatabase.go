package database

import (
	"book-microservice/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_Instance *gorm.DB

func ConnectDb(books []model.Book) {

	//"host=localhost or db
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	db.Migrator().DropTable("Books")

	// create database tables
	db.AutoMigrate(&books)

	for indx := range books {
		db.Create(&books[indx])
	}

	DB_Instance = db
}

func GetAll(pagination Pagination) *Pagination {

	var rows []*model.Book

	DB_Instance.Scopes(paginate(rows, &pagination, DB_Instance)).Find(&rows)
	pagination.Rows = rows

	return &pagination
}

func GetById(id uint) *model.Book {

	var book *model.Book
	DB_Instance.First(&book, id)

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
