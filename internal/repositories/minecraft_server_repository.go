package repositories

import (
	"api-mongo-go/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MinecraftServerRepository struct {
	collection *mongo.Collection
}

func NewMinecraftServerRepository(db *mongo.Database) *MinecraftServerRepository {
	return &MinecraftServerRepository{
		collection: db.Collection("minecraft_servers"),
	}
}

func (r *MinecraftServerRepository) Create(server *models.MinecraftServer) error {
	_, err := r.collection.InsertOne(context.Background(), server)
	return err
}

func (r *MinecraftServerRepository) GetByID(id string) (*models.MinecraftServer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var server models.MinecraftServer
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (r *MinecraftServerRepository) Update(id string, server *models.MinecraftServer) (*models.MinecraftServer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"name":          server.Name,
			"address":       server.Address,
			"port":          server.Port,
			"version":       server.Version,
			"playersOnline": server.PlayersOnline,
			"maxPlayers":    server.MaxPlayers,
			"status":        server.Status,
			"motd":          server.MOTD,
			"lastChecked":   server.LastChecked,
			"tags":          server.Tags,
		},
	}

	_, err = r.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *MinecraftServerRepository) GetAll() ([]models.MinecraftServer, error) {
	ctx := context.Background()

	// We only want to retrieve _id and name fields
	projection := bson.M{"_id": 1, "name": 1}

	cursor, err := r.collection.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var servers []models.MinecraftServer
	if err = cursor.All(ctx, &servers); err != nil {
		return nil, err
	}

	return servers, nil
}

func (r *MinecraftServerRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
