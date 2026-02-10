package factory

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/G0tem/go-service-entity/internal/config"
	"github.com/G0tem/go-service-entity/internal/database"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

var (
	// Для GORM
	databaseServiceInstance *gorm.DB
	databaseOnce            sync.Once
	databaseErr             error

	// Для MongoDB
	mongoClientInstance *mongo.Client
	mongoOnce           sync.Once
	mongoErr            error

	// Для Redis
	redisClientInstance *redis.Client
	redisOnce           sync.Once
	redisErr            error
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	databaseOnce.Do(func() {
		databaseServiceInstance, databaseErr = database.Connect(cfg)
	})
	return databaseServiceInstance, databaseErr
}

func NewMongo(cfg *config.Config) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mongoClientInstance, mongoErr = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MangoURL))
		if mongoErr != nil {
			log.Printf("Error connecting to MongoDB: %v", mongoErr)
			return
		}

		if mongoErr = mongoClientInstance.Ping(ctx, nil); mongoErr != nil {
			log.Printf("Error pinging MongoDB: %v", mongoErr)
			return
		}

		log.Println("Successfully connected to MongoDB!")
	})
	return mongoClientInstance, mongoErr
}

func NewRedis(cfg *config.Config) (*redis.Client, error) {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     cfg.RedisUrl,
			Password: "",
			DB:       0,
		})

		ctx := context.Background()
		if err := client.Ping(ctx).Err(); err != nil {
			redisErr = fmt.Errorf("failed to connect to Redis: %w", err)
			return
		}

		redisClientInstance = client
		log.Println("[InitRedis] create redisConnect")
	})

	if redisErr != nil {
		return nil, redisErr
	}
	return redisClientInstance, nil
}
