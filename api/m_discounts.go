package api

import (
	"net/http"
	"database/sql"
	"strconv"
	// "reflect"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/rizgust/go-finance/providers/db/sqlc"
) 

func (server *Server) createMDiscount(ctx *gin.Context ) {
	var req db.CreateMDiscountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.CreateMDiscount(ctx, arg)
	// method := reflect.ValueOf(server.store).MethodByName("CreateMDiscount")
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
	
	ctx.JSON(http.StatusOK, data)
}

func (server *Server) getMDiscount(ctx *gin.Context) {
	_id, err := strconv.ParseInt(ctx.Param("id"), 10, 32) 
	if err == nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id := int32(_id)

	data, err := server.store.GetMDiscount(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (server *Server) getListMDiscount(ctx *gin.Context) {
	var req db.ListMDiscountsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.ListMDiscounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (server *Server) updateMDiscount(ctx *gin.Context ) {
	var req db.UpdateMDiscountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateMDiscount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if(pqErr.Code.Name() != ""){
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, data)
}

func (server *Server) deleteMDiscount(ctx *gin.Context ) {
	var req db.UpdateMDiscountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateMDiscount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if(pqErr.Code.Name() != ""){
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, data)
}