package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/rizgust/go-finance/providers/db/sqlc"
)

type Server struct {
    store  *db.Store
    router *gin.Engine
}

func NewServer(store *db.Store) *Server {
    server := &Server{store: store}

    server.setupRouter()
    return server
}

func (server *Server) setupRouter() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		master := v1.Group("/master")
		{
			master.POST("/bank", server.createMBank)
			master.GET("/bank/:id", server.getMBank)
			master.GET("/banks", server.getListMBank)
			master.PUT("/bank", server.updateMBank)

		}		
	}
	server.router = router
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
