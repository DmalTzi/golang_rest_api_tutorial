package main

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "gorm.io/gorm/logger"
    "fmt"
    "log"
    "os"
    "time"
)

const (
    host    =   "localhost"
    port    =   54329
    user    =   "myuser"
    password=   "mypassword"
    dbname  =   "mydatabase"
)

func main() {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second,
            LogLevel: logger.Info,
            Colorful: true,
        },
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(Book{})

    app := fiber.New()

    app.Get("/books", func(c *fiber.Ctx) error {
        return GetBooks(db, c)
    })

    app.Get("/book/:id", func(c *fiber.Ctx) error {
        return GetBook(db, c)
    })

    app.Post("/book", func(c *fiber.Ctx) error {
        return CreateBook(db, c)
    })

    app.Put("/book/:id", func(c *fiber.Ctx) error {
        return UpdateBook(db, c)
    })

    app.Delete("/book/:id", func(c *fiber.Ctx) error {
        return DeleteBook(db, c)
    })

    app.Listen(":8080")
}
