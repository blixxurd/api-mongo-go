package controllers

import (
	"api-mongo-go/internal/models"
	"api-mongo-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MinecraftServerController struct {
	service *services.MinecraftServerService
}

func NewMinecraftServerController(db *mongo.Database) *MinecraftServerController {
	return &MinecraftServerController{
		service: services.NewMinecraftServerService(db),
	}
}

func (c *MinecraftServerController) CreateServer(ctx *gin.Context) {
	var server models.MinecraftServer
	if err := ctx.ShouldBindJSON(&server); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateServer(&server); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, server)
}

func (c *MinecraftServerController) GetServer(ctx *gin.Context) {
	id := ctx.Param("id")
	server, err := c.service.GetServer(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}
	ctx.JSON(http.StatusOK, server)
}

func (c *MinecraftServerController) GetAllServers(ctx *gin.Context) {
	servers, err := c.service.GetAllServers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, servers)
}

func (c *MinecraftServerController) UpdateServer(ctx *gin.Context) {
	id := ctx.Param("id")
	var server models.MinecraftServer
	if err := ctx.ShouldBindJSON(&server); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedServer, err := c.service.UpdateServer(id, &server)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedServer)
}

func (c *MinecraftServerController) DeleteServer(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteServer(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Server deleted successfully"})
}
