package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"zero/internal/auth/app"
	"zero/internal/auth/app/service"
	"zero/internal/auth/domain/common"
)

func RegisterUser(app *app.Application) gin.HandlerFunc {
	type Body struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	type Response struct {
		ID        int       `json:"id"`
		UID       string    `json:"uid"`
		Email     string    `json:"email"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// Validate parameters
		var body Body
		err := c.ShouldBind(&body)
		if err != nil {
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		user, err := app.AuthService.RegisterUser(ctx, service.RegisterUserParam{
			Email:    body.Email,
			Name:     body.Name,
			Password: body.Password,
		})
		if err != nil {
			respondWithError(c, err)
			return
		}

		resp := Response{
			ID:        user.ID,
			UID:       user.UID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		}
		respondWithJSON(c, http.StatusCreated, resp)
	}
}

func LoginUser(app *app.Application) gin.HandlerFunc {
	type Body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	type Response struct {
		UserID int    `json:"user_id"`
		Token  string `json:"token"`
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// Validate parameters
		var body Body
		err := c.ShouldBind(&body)
		if err != nil {
			respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("invalid parameter")))
			return
		}

		// Invoke service
		user, err := app.AuthService.LoginUser(ctx, service.LoginUserParam{
			Email:    body.Email,
			Password: body.Password,
		})
		if err != nil {
			respondWithError(c, err)
			return
		}
		token, err := app.AuthService.GenerateUserToken(ctx, *user)
		if err != nil {
			respondWithError(c, err)
			return
		}

		resp := Response{
			UserID: user.ID,
			Token:  token,
		}
		respondWithJSON(c, http.StatusOK, resp)
	}
}
