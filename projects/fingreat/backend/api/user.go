package api

import (
	"context"
	db "github/ronnie5562/fingreat_backend/db/sqlc"
	"github/ronnie5562/fingreat_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type User struct {
	server *Server
}

func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users")

	serverGroup.GET("", u.listUsers)
	serverGroup.POST("", u.createUser)
}

type UserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) createUser(c *gin.Context) {
	user := new(UserParams)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.GenerateHashedPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateUserParams{
		Email:          user.Email,
		HashedPassword: hashedPassword,
	}

	newUser, err := u.server.queries.CreateUser(
		context.Background(),
		arg,
	)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(
					http.StatusBadRequest,
					gin.H{"error": "user already exists"},
				)
				return
			}
		}

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
}

func (u *User) listUsers(c *gin.Context) {
	arg := db.ListUsersParams{
		Limit:  10,
		Offset: 0,
	}
	users, err := u.server.queries.ListUsers(
		context.Background(),
		arg,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
