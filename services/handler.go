package services

import (
	"github.com/Taeho0927/goLogger/logger"
	"github.com/gin-gonic/gin"
	"github/Taeho0927/Standard_API/core/models"
	"github/Taeho0927/Standard_API/core/pgLayer"
	"github/Taeho0927/Standard_API/middlewares"
	"net/http"
)

type Handler struct {
	pg pgLayer.PGLayer
}

func NewHandler() (*Handler, error) {
	pg, err := pgLayer.PGConn()
	if err != nil {
		logger.Error(middlewares.SetLogger("NewHandler()", err.Error()))
		return nil, err
	}
	return &Handler{
		pg: pg,
	}, nil
}

func (h *Handler) CheckDBConnection(c *gin.Context) {
	if h.pg == nil {
		logger.Error(middlewares.SetLogger("CheckDBConnection", "Postgres Server Error"))
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "PostgreSQL Server Error"})
		return
	}
}
func (h *Handler) BindingParam(c *gin.Context) {
	var params models.Params
	if err := c.ShouldBindUri(&params); err != nil {
		logger.Error(middlewares.SetLogger("BindingParam", err.Error()))
		c.JSON(http.StatusAccepted, gin.H{"httpCode": http.StatusBadRequest, "Error": err.Error()})
		return
	}
	c.Set("Params", params)
	c.Next()
}
func (h *Handler) BindingFileParam(c *gin.Context) {
	var params models.FileParams
	if err := c.ShouldBindUri(&params); err != nil {
		logger.Error(middlewares.SetLogger("BindingFileParam", err.Error()))
		c.JSON(http.StatusAccepted, gin.H{"httpCode": http.StatusBadRequest, "Error": err.Error()})
		return
	}
	c.Set("Params", params)
	c.Next()
}
