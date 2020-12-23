package controller

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"github.com/wambugucoder/fiber-todo-api/model"
	"go.mongodb.org/mongo-driver/bson"
)
// GetAllTodos - GEY /api/todos/
func GetAllTodos(ctx *fiber.Ctx) {
	collection := mgm.Coll(&model.Todo{})
	todos := []model.Todo{}

	capturedData := collection.SimpleFind(&todos, bson.D{})

	if capturedData == nil {
		ctx.Status(500).JSON(fiber.Map{
			"error": capturedData.Error(),
		})
		return
	}
	ctx.Status(200).JSON(fiber.Map{
		"todo": todos,
	})

}
// GetTodoByID - GEY /api/todos/:id
func GetTodoByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &model.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
// CreateTodo - POST /api/todos
func CreateTodo(ctx *fiber.Ctx) {
	params := new(struct {
		Title       string
		Description string
	})

	ctx.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Description) == 0 {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title or description not specified.",
		})
		return
	}

	todo := model.CreateTodo(params.Title, params.Description)
	err := mgm.Coll(todo).Create(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
//DeleteTodo deletes a todo by id
func DeleteTodo(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &model.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	err = collection.Delete(todo)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
