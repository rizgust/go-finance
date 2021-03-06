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

func (server *Server) createMVirtualAccount(ctx *gin.Context ) {
	var req db.CreateMVirtualAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.CreateMVirtualAccount(ctx, arg)
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

func (server *Server) getMVirtualAccount(ctx *gin.Context) {
	_id, err := strconv.ParseInt(ctx.Param("id"), 10, 32) 
	if err == nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	id := int32(_id)

	data, err := server.store.GetMVirtualAccount(ctx, id)
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

func (server *Server) getListMVirtualAccount(ctx *gin.Context) {
	var req db.ListMVirtualAccountsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.ListMVirtualAccounts(ctx, arg)
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

func (server *Server) updateMVirtualAccount(ctx *gin.Context ) {
	var req db.UpdateMVirtualAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateMVirtualAccount(ctx, arg)
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

func (server *Server) deleteMVirtualAccount(ctx *gin.Context ) {
	var req db.UpdateMVirtualAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req
	data, err := server.store.UpdateMVirtualAccount(ctx, arg)
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