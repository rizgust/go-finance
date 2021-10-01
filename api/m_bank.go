package api

import (
	"net/http"
	// "reflect"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/rizgust/go-finance/providers/db/sqlc"
) 

func (server *Server) createMBank(ctx *gin.Context ) {
	var req db.CreateMBankParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	mBank, err := server.store.CreateMBank(ctx, arg)
	// method := reflect.ValueOf(server.store).MethodByName("CreateMBank")
	// var inputs []reflect.Value
	// inputs = append(inputs, reflect.ValueOf(ctx))
	// inputs = append(inputs, reflect.ValueOf(arg))
	// outputs :=  method.Call(inputs)
	// objectTable := reflect.ValueOf(outputs[0])
	// err := reflect.ValueOf(outputs[1])
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, mBank)
}
