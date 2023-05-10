// handler/reaction_handler.go
package handler

import (
	"net/http"
	"stray-cat-api/domain/model"
	"stray-cat-api/usecase/interactor"

	"github.com/gin-gonic/gin"
)

type ReactionHandler struct {
	ReactionInteractor interactor.ReactionInteractor
}

func (h *ReactionHandler) GetByID(c *gin.Context) {
	reactionID := c.Param("reactionID")
	reaction, err := h.ReactionInteractor.FindByID(reactionID)
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
	reactionID := c.Param("reactionID")
	var reaction model.Reaction
	if err := c.BindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reaction.ReactionID = reactionID
	if err := h.ReactionInteractor.Update(&reaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reaction)
}

func (h *ReactionHandler) Delete(c *gin.Context) {
	reactionID := c.Param("reactionID")
	if err := h.ReactionInteractor.Delete(reactionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
