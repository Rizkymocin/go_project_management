package controllers

import (
	"math"
	"strconv"

	"github.com/Rizkymocin/project-management/models"
	"github.com/Rizkymocin/project-management/services"
	"github.com/Rizkymocin/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Gagal registrasi user", err.Error())
	}

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)

	return utils.Success(ctx, "Register Success", userResp)

}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid Request", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.UnAuthorized(ctx, "Login Failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, user)
	return utils.Success(ctx, "Login successful", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          userResp,
	})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.service.GetByPublicID(id)

	if err != nil {
		return utils.NotFound(ctx, "Data Not Found", err.Error())
	}

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)

	if err != nil {
		return utils.BadRequest(ctx, "Internal Server Error", err.Error())
	}

	return utils.Success(ctx, "User Found", userResp)

}

func (c *UserController) GetUserPagination(ctx *fiber.Ctx) error {
	// Misal ada alamat /user/page?page=1&limit=10&sorting=-id&filter=rizky
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	users, total, err := c.service.GetAllPagination(filter, sort, limit, offset)

	if err != nil {
		return utils.BadRequest(ctx, "Gagal mengambil data", err.Error())
	}

	var userResp []models.UserResponse
	_ = copier.Copy(&userResp, &users)

	meta := utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		Filter:     filter,
		Sort:       sort,
	}

	if total == 0 {
		utils.NotFoundPagination(ctx, "Data pengguna tidak ditemukan", userResp, meta)
	}

	return utils.SuccessPagination(ctx, "Data ditemukan", userResp, meta)

}
