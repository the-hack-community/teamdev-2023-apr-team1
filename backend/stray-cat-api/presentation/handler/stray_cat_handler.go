// handler/stray_cat_handler.go
package handler

import (
	"net/http"
	"stray-cat-api/domain/model"
	"stray-cat-api/usecase/interactor"

	"encoding/base64"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type StrayCatHandler struct {
	StrayCatInteractor interactor.StrayCatInteractor
}

func (h *StrayCatHandler) GetAll(c *gin.Context) {
	cats, err := h.StrayCatInteractor.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

func (h *StrayCatHandler) GetByID(c *gin.Context) {
	catID := c.Param("catID")
	cat, err := h.StrayCatInteractor.FindByID(catID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *StrayCatHandler) Create(c *gin.Context) {
	// マルチパートフォームデータの解析
	err := c.Request.ParseMultipartForm(10 << 20) // Max 10MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// form dataから値を取得
	name := c.Request.FormValue("name")
	features := c.Request.FormValue("features")
	condition := c.Request.FormValue("condition")
	captureDateTimeStr := c.Request.FormValue("captureDateTime")
	latStr := c.Request.FormValue("lat")
	longStr := c.Request.FormValue("long")

	// captureDateTimeを解析する
	captureDateTime, err := time.Parse(time.RFC3339, captureDateTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// locationを解析する

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location := &model.Location{
		Lat:  lat,
		Long: long,
	}

	// ファイルを取得
	file, _, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// ファイルデータを読み込む
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ファイルデータをbase64にエンコードする
	photoData := base64.StdEncoding.EncodeToString(fileData)

	// StrayCatオブジェクトを作成
	strayCat := &model.StrayCat{
		Name:            name,
		Features:        features,
		Condition:       condition,
		CaptureDateTime: captureDateTime,
		Location:        *location,
		PhotoData:       photoData,
	}

	// 保存処理
	if err := h.StrayCatInteractor.Store(strayCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, strayCat)
}

func (h *StrayCatHandler) Update(c *gin.Context) {
	catID := c.Param("catID")
	var cat model.StrayCat
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat.CatID = catID
	if err := h.StrayCatInteractor.Update(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *StrayCatHandler) Delete(c *gin.Context) {
	catID := c.Param("catID")
	if err := h.StrayCatInteractor.Delete(catID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
