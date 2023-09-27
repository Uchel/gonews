package handler

import (
	"go_news/database"
	"go_news/model/entity"
	"go_news/model/req"
	"go_news/model/res"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(req.UserReq)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	var newUser = entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreate := database.DB.Debug().Create(&newUser).Error

	if errCreate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   errCreate,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success to create user",
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	userRes := res.UserRes{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userRes,
	})
}

func UserHandlerUpdateById(ctx *fiber.Ctx) error {
	var userReq *req.UserReq

	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "bad request`",
		})
	}

	var user *entity.User
	id := ctx.Params("id")
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Address != "" {
		user.Address = userReq.Address
	}
	if userReq.Phone != "" {
		user.Phone = userReq.Phone
	}

	errorUpdate := database.DB.Debug().Save(&user).Error

	if errorUpdate != nil {
		ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Success to update user",
	})

}
