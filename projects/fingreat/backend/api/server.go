package api

import (
	"database/sql"
	"fmt"
	db "github/ronnie5562/fingreat_backend/db/sqlc"
	"github/ronnie5562/fingreat_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(envPath string) *Server {
	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Could not load env config: %v", err))
	}

	conn, err := sql.Open(config.DBdriver, config.DB_source_live)

	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	q := db.New(conn)

	g := gin.Default()

	return &Server{
		queries: q,
		router:  g,
	}
}

func (s *Server) Start(port int) {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to fingreat_backend",
		})
	})

	User{}.router(s)

	s.router.Run(fmt.Sprintf(":%v", port))
}
