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
			
			master.POST("/bank-acc", server.createMBankAccount)
			master.GET("/bank-acc/:id", server.getMBankAccount)
			master.GET("/bank-accs", server.getListMBankAccount)
			master.PUT("/bank-acc", server.updateMBankAccount)

			master.POST("/book-period", server.createMBookPeriod)
			master.GET("/book-period/:id", server.getMBookPeriod)
			master.GET("/book-periods", server.getListMBookPeriod)
			master.PUT("/book-period", server.updateMBookPeriod)

			master.POST("/coa", server.createMChartOfAccount)
			master.GET("/coa/:id", server.getMChartOfAccount)
			master.GET("/coas", server.getListMChartOfAccount)
			master.PUT("/coa", server.updateMChartOfAccount)

			master.POST("/va", server.createMVirtualAccount)
			master.GET("/va/:id", server.getMVirtualAccount)
			master.GET("/vas", server.getListMVirtualAccount)
			master.PUT("/va", server.updateMVirtualAccount)

			master.POST("/discount", server.createMDiscount)
			master.GET("/discount/:id", server.getMDiscount)
			master.GET("/discounts", server.getListMDiscount)
			master.PUT("/discount", server.updateMDiscount)
		}	
		
		master.POST("/ap", server.createAccountPayable)
		master.GET("/ap/:id", server.getAccountPayable)
		master.GET("/ap", server.getListAccountPayable)
		master.PUT("/ap", server.updateAccountPayable)
	}
	server.router = router
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
