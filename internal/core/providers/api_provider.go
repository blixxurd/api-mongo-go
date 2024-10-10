package providers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiProvider struct {
	DB     *mongo.Database
	Router *gin.Engine
}

func NewApiProvider(db *mongo.Database, router *gin.Engine) *ApiProvider {
	return &ApiProvider{
		DB:     db,
		Router: router,
	}
}
