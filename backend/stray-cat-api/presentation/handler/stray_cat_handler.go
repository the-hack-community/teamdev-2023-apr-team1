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
	catIdStr := c.Param("catId")
	// 空文字列のチェック
	if catIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catId must not be empty"})
		return
	}

	// 整数として解釈できるかのチェック
	catId, err := strconv.Atoi(catIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catId must be a number"})
		return
	}
	cat, err := h.StrayCatInteractor.FindByID(catId)
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
		UserID:          "123",
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
	catIdStr := c.Param("catId")
	// 空文字列のチェック
	if catIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catId must not be empty"})
		return
	}

	// 整数として解釈できるかのチェック
	catId, err := strconv.Atoi(catIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catId must be a number"})
		return
	}

	var cat model.StrayCat
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cat.CatID = catId
	if err := h.StrayCatInteractor.Update(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *StrayCatHandler) Delete(c *gin.Context) {
	catIdStr := c.Param("catId")
	// 空文字列のチェック
	if catIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catId must not be empty"})
		return
	}

	// 整数として解釈できるかのチェック
	catId, err := strconv.Atoi(catIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "catId must be a number"})
		return
	}

	if err := h.StrayCatInteractor.Delete(catId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
