// handler/reaction_handler.go
package handler

import (
	"net/http"
	"stray-cat-api/domain/model"
	"stray-cat-api/usecase/interactor"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReactionHandler struct {
	ReactionInteractor interactor.ReactionInteractor
}

func (h *ReactionHandler) GetByID(c *gin.Context) {
	reactionIdStr := c.Param("reactionId")
	// 空文字列のチェック
	if reactionIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reactionId must not be empty"})
		return
	}

	// 整数として解釈できるかのチェック
	reactionId, err := strconv.Atoi(reactionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reactionId must be a number"})
		return
	}

	reaction, err := h.ReactionInteractor.FindByID(reactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reaction)
}

func (h *ReactionHandler) Create(c *gin.Context) {
	var reaction model.Reaction
	if err := c.BindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.ReactionInteractor.Store(&reaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reaction)
}

func (h *ReactionHandler) Update(c *gin.Context) {
	reactionIdStr := c.Param("reactionId")
	// 空文字列のチェック
	if reactionIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reactionId must not be empty"})
		return
	}

	// 整数として解釈できるかのチェック
	reactionId, err := strconv.Atoi(reactionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reactionId must be a number"})
		return
	}

	var reaction model.Reaction
	if err := c.BindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reaction.ReactionID = reactionId
	if err := h.ReactionInteractor.Update(&reaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reaction)
}

func (h *ReactionHandler) Delete(c *gin.Context) {
	reactionIdStr := c.Param("reactionId")
	// 空文字列のチェック
	if reactionIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reactionId must not be empty"})
		return
	}

	// 整数として解釈できるかのチェック
	reactionId, err := strconv.Atoi(reactionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reactionId must be a number"})
		return
	}

	if err := h.ReactionInteractor.Delete(reactionId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
