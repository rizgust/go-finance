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

func (server *Server) createMBankAccount(ctx *gin.Context ) {
	var req db.CreateMBankAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.CreateMBankAccount(ctx, arg)
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

func (server *Server) getMBankAccount(ctx *gin.Context) {
	_id, err := strconv.ParseInt(ctx.Param("id"), 10, 32) 
	if err == nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id := int32(_id)

	data, err := server.store.GetMBankAccount(ctx, id)
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

func (server *Server) getListMBankAccount(ctx *gin.Context) {
	var req db.ListMBankAccountsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.ListMBankAccounts(ctx, arg)
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

func (server *Server) updateMBankAccount(ctx *gin.Context ) {
	var req db.UpdateMBankAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateMBankAccount(ctx, arg)
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

func (server *Server) deleteMBankAccount(ctx *gin.Context ) {
	var req db.UpdateMBankAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateMBankAccount(ctx, arg)
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