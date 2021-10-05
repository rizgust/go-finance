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

func (server *Server) createAccountReceiveable(ctx *gin.Context ) {
	var req db.CreateAccountReceiveableParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.CreateAccountReceiveable(ctx, arg)
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

func (server *Server) getAccountReceiveable(ctx *gin.Context) {
	_id, err := strconv.ParseInt(ctx.Param("id"), 10, 32) 
	if err == nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id := int32(_id)

	data, err := server.store.GetAccountReceiveable(ctx, id)
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

func (server *Server) getListAccountReceiveable(ctx *gin.Context) {
	var req db.ListAccountReceiveablesParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.ListAccountReceiveables(ctx, arg)
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

func (server *Server) updateAccountReceiveable(ctx *gin.Context ) {
	var req db.UpdateAccountReceiveableParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateAccountReceiveable(ctx, arg)
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

func (server *Server) deleteAccountReceiveable(ctx *gin.Context ) {
	var req db.UpdateAccountReceiveableParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateAccountReceiveable(ctx, arg)
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