package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rizgust/go-finance/providers/db/sqlc"
)

type createMBankRequest struct {
	OwnerID   int32          `json:"owner_id" binding:"required"`
	Code      string         `json:"code" binding:"required"`
	Name      string         `json:"name" binding:"required"`
	Alias     string		 `json:"alias"`
	UserId 	  int32          `json:"user_id" binding:"required"`
}

func (server *Server) createMBank(ctx *gin.Context) {
	var req createMBankRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var alias sql.NullString
	alias.Scan(req.Alias)

	arg := db.CreateMBankParams{
		OwnerID: req.OwnerID,
		Code: req.Code,
		Name: req.Name,
		Alias: alias,
		CreatedBy: req.UserId,
		UpdatedBy: req.UserId,
	}

	mBank, err := server.store.CreateMBank(ctx, arg)
	if err != nil {
		// if pqErr, ok := err.(*pq.Error); ok {
		// 	switch pqErr.Code.Name() {
		// 	case "foreign_key_violation", "unique_violation":
		// 		ctx.JSON(http.StatusForbidden, errorResponse(err))
		// 		return
		// 	}
		// }
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, mBank)
}


