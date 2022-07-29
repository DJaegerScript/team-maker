package team

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler struct {
	Service  Service
	Validate *validator.Validate
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service:  service,
		Validate: validator.New(),
	}
}

func (h *Handler) GetAllTeams(ctx *fiber.Ctx) error {
	err, status, res := h.Service.GetAllTeams()
	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	return ctx.Status(status).JSON(fiber.Map{
		"success": true,
		"content": res,
		"message": "",
	})
}

func (h *Handler) GetTeam(ctx *fiber.Ctx) error {
	rawId := ctx.Params("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	err, status, res := h.Service.GetTeam(id)
	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	return ctx.Status(status).JSON(fiber.Map{
		"success": true,
		"content": res,
		"message": "",
	})
}

func (h *Handler) CreateTeam(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": "Name can't be empty",
		})
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	err, status, res := h.Service.CreateTeam(name, file)
	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"content": res,
		"message": "",
	})
}

func (h *Handler) UpdateTeam(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name" validate:"required"`
	}

	var r request

	if err := ctx.BodyParser(&r); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	if err := h.Validate.Struct(&r); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	rawId := ctx.Params("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	err, status, res := h.Service.UpdateTeam(id, r.Name)
	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	return ctx.Status(status).JSON(fiber.Map{
		"success": true,
		"content": res,
		"message": "",
	})
}

func (h *Handler) DeleteTeam(ctx *fiber.Ctx) error {
	rawId := ctx.Params("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	err, status, res := h.Service.DeleteTeam(id)
	if err != nil {
		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": err.Error(),
		})
	}

	return ctx.Status(status).JSON(fiber.Map{
		"success": true,
		"content": res,
		"message": "",
	})
}
