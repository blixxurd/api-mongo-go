package services

import (
	"api-mongo-go/internal/models"
	"api-mongo-go/internal/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type MinecraftServerService struct {
	repo *repositories.MinecraftServerRepository
}

func NewMinecraftServerService(db *mongo.Database) *MinecraftServerService {
	return &MinecraftServerService{
		repo: repositories.NewMinecraftServerRepository(db),
	}
}

func (s *MinecraftServerService) CreateServer(server *models.MinecraftServer) error {
	return s.repo.Create(server)
}

func (s *MinecraftServerService) GetServer(id string) (*models.MinecraftServer, error) {
	return s.repo.GetByID(id)
}

func (s *MinecraftServerService) GetAllServers() ([]models.MinecraftServer, error) {
	return s.repo.GetAll()
}

func (s *MinecraftServerService) UpdateServer(id string, server *models.MinecraftServer) (*models.MinecraftServer, error) {
	return s.repo.Update(id, server)
}

func (s *MinecraftServerService) DeleteServer(id string) error {
	return s.repo.Delete(id)
}
