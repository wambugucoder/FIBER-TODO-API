package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm"
	"github.com/gofiber/fiber"
	"github.com/wambugucoder/fiber-todo-api/controller"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	app := fiber.New()

	app.Get("/api/todos", controller.GetAllTodos)
	app.Get("/api/todos/:id", controller.GetTodoByID)
	app.Post("/api/todos", controller.CreateTodo)
	app.Delete("/api/todos/:id", controller.DeleteTodo)

	app.Listen(3000)
}
