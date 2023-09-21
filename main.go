package main

import (
	"log"
	"os"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Blog struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Comments string `json:"comments"`
	Summary  string `json:"summary"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBlog(context *fiber.Ctx) error {
	blog := Blog{}
	if err := context.BodyParser(&blog); err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Invalid request body"})
		return err
	}

	if err := r.DB.Create(&blog).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Could not create a blog"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Blog created successfully",
		"data":    blog, // Return the created blog data.
	})
	return nil
}

func (r *Repository) GetBlogs(context *fiber.Ctx) error {
	blogModels := &[]Blog{}
	if err := r.DB.Find(blogModels).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Could not find blogs"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Blogs retrieved successfully",
		"data":    blogModels, // Return the retrieved blog data.
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	// Blog routes
	app.Get("/api/blogs", r.GetBlogs)
	app.Post("/api/blogs", r.CreateBlog)
	// Add routes for updating and deleting blogs as needed.

	// Add routes for user and comment operations if required.
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":3000")
}
