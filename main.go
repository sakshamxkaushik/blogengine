package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Blog struct {
	Author   string `json:"author"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Comments string `json:"comments"`
	Summary  string `json:"summary"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBlog(context *fibre.Ctx) error {  
blog := Blog{}

  context.BodyParser(&blog)

if err := nil {
	context.Status (http.StatusBadRequest).JSON(
		&fibre.Map{"message": "Request Failed"}
	)


}

func (r *Repository) SetupRoutes(app *fiber.App) {
	// User routes
	app.Get("/api/users", r.GetUsers)
	app.Post("/api/users", r.CreateUser)
	app.Get("/api/users/:id", r.GetUser)
	app.Put("/api/users/:id", r.UpdateUser)
	app.Delete("/api/users/:id", r.DeleteUser)

	// Blog routes
	app.Get("/api/posts", r.GetPosts)
	app.Post("/api/posts", r.CreatePost)
	app.Get("/api/posts/:id", r.GetPost)
	app.Put("/api/posts/:id", r.UpdatePost)
	app.Delete("/api/posts/:id", r.DeletePost)

	// Comment routes
	app.Get("/api/posts/:id/comments", r.GetComments)
	app.Post("/api/posts/:id/comments", r.CreateComment)
	app.Put("/api/posts/:id/comments/:commentId", r.UpdateComment)
	app.Delete("/api/posts/:id/comments/:commentId", r.DeleteComment)
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
