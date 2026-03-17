package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/shutterscripter/kafka_chat/model"
	"github.com/shutterscripter/kafka_chat/utils"
)

func CreateComment(c fiber.Ctx) error {
	cmt := new(model.Comment)
	if err := c.Bind().Body(cmt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Invalid request body", "data": nil})
	}
	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong, Try again!",
			"data":    nil,
		})
	}

	utils.PushCommentToQueue("comments", cmtInBytes)
	err = c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Comment pushed succesfully",
		"data":    cmt,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  true,
			"message": "Error creating product",
			"data":    nil,
		})
	}
	return err
}

func HealthCheck(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Healthy!", "data": nil})
}
