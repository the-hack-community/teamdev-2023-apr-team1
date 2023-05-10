package handler

import (
	"net/http"
	"stray-cat-api/domain/model"
	"stray-cat-api/usecase/interactor"

	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	UserInfoInteractor interactor.UserInfoInteractor
}

func (h *UserInfoHandler) GetAll(c *gin.Context) {
	users, err := h.UserInfoInteractor.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserInfoHandler) GetByID(c *gin.Context) {
	userId := c.Param("userId")
	user, err := h.UserInfoInteractor.FindByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserInfoHandler) Create(c *gin.Context) {
	userId := c.Param("userId")
	var user model.UserInfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.UserID = userId
	if err := h.UserInfoInteractor.Store(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserInfoHandler) Update(c *gin.Context) {
	userId := c.Param("userId")
	var user model.UserInfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.UserID = userId
	if err := h.UserInfoInteractor.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserInfoHandler) Delete(c *gin.Context) {
	userId := c.Param("userId")
	if err := h.UserInfoInteractor.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
