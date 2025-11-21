package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Errors       string      `json:"errors,omitempty"`
}

type ResponsePaginated struct {
	Status       string         `json:"status"`
	ResponseCode int            `json:"response_code"`
	Message      string         `json:"message,omitempty"`
	Data         interface{}    `json:"data,omitempty"`
	Errors       string         `json:"errors,omitempty"`
	Meta         PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page       int    `json:"page" example:"1"`
	Limit      int    `json:"limit" example:"10"`
	Total      int    `json:"total" example:"100"`
	TotalPages int    `json:"total_pages" example:"10"`
	Filter     string `json:"filter" example:"nama=rizky"`
	Sort       string `json:"sort" example:"-id"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
	})
}

func SuccessPagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(ResponsePaginated{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func NotFoundPagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusNotFound).JSON(ResponsePaginated{
		Status:       "Not Found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:       "success",
		ResponseCode: fiber.StatusCreated,
		Message:      message,
		Data:         data,
	})
}

func BadRequest(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:       "Error bad request",
		ResponseCode: fiber.StatusBadRequest,
		Message:      message,
		Errors:       err,
	})
}

func NotFound(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:       "Error not found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Errors:       err,
	})
}

func UnAuthorized(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:       "Unauthorized",
		ResponseCode: fiber.StatusUnauthorized,
		Message:      message,
		Errors:       err,
	})
}
