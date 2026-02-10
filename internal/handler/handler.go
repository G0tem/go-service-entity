package handler

import (
	"net/http"
	"time"

	"github.com/G0tem/go-service-entity/internal/config"
	grpcClient "github.com/G0tem/go-service-entity/internal/grpc"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"

	"gorm.io/gorm"
)

type Handler struct {
	db         *gorm.DB
	mongo      *mongo.Client
	rds        *redis.Client
	cfg        *config.Config
	authClient *grpcClient.AuthClient
	client     *http.Client
}

func NewHandler(db *gorm.DB, mongo *mongo.Client, rds *redis.Client, cfg *config.Config) *Handler {
	// Инициализируем gRPC клиент
	authClient, err := grpcClient.NewAuthClient(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create auth gRPC client")
		// Продолжаем работу без gRPC клиента, но логируем ошибку
	}

	return &Handler{
		db:         db,
		mongo:      mongo,
		rds:        rds,
		cfg:        cfg,
		authClient: authClient,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *Handler) SetupRoutes(app *fiber.App) {
	cfg := config.LoadConfig()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	entity := v1.Group("/entity")

	entity.Use(JWTMiddleware(cfg.SecretKey))

	entity.Get("test-grpc", h.TestGrpc)
	entity.Get("test-grpc-user-info", h.TestGetUserInfo)

	entity.Get("get", h.GetEntity)
	entity.Post("create", h.CreateEntity)
	entity.Patch("update/:id", h.UpdateEntity)
	entity.Delete("delete/:id", h.DeleteEntity)

	entity.Get("check", h.UserInfo)
}
