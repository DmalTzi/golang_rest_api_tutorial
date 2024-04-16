package main

import (
    "gorm.io/gorm"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type Book struct {
    gorm.Model
    Name            string  `json:"name"`
    Author          string  `json:"author"`
    Descriptions    string  `json:"descriptions"`
    Price           uint    `json:"price"`
}

func CreateBook(db *gorm.DB, c *fiber.Ctx) error {
    book := new(Book)

    if err := c.BodyParser(book); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    result := db.Create(book)
    if result.Error != nil {
        return result.Error
    }

    return c.JSON(book)
}

func GetBook(db *gorm.DB, c *fiber.Ctx) error {
    var book Book
    bookId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }
    result := db.First(&book, bookId)
    if result.Error != nil {
        return result.Error
    }

    return c.JSON(book)
}

func GetBooks(db *gorm.DB, c *fiber.Ctx) error {
    var books []Book

    result := db.Find(&books)
    if result.Error != nil {
        return result.Error 
    }

    return c.JSON(books)
}

func UpdateBook(db *gorm.DB, c *fiber.Ctx) error {
    bookId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    book := new(Book)

    result := db.First(&book, bookId)
    if result.Error != nil {
        return result.Error
    }

    if err := c.BodyParser(book); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    result = db.Save(book)
    if result.Error != nil {
        return result.Error
    }

    return c.JSON(book)
}

func DeleteBook(db *gorm.DB, c *fiber.Ctx) error {
    bookId, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }
    result := db.Delete(&Book{}, bookId)

    if result.Error != nil {
        return result.Error
    }

    return c.SendString("DeleteSuccessfully")
}
