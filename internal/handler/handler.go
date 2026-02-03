package handler

import (
	"net/http"
	"time"

	"github.com/G0tem/go-service-entity/internal/config"
	grpcClient "github.com/G0tem/go-service-entity/internal/grpc"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB, cfg *config.Config) *Handler {
	// Инициализируем gRPC клиент
	authClient, err := grpcClient.NewAuthClient(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create auth gRPC client")
		// Продолжаем работу без gRPC клиента, но логируем ошибку
	}

	return &Handler{
		db:         db,
		cfg:        cfg,
		authClient: authClient,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type Handler struct {
	db         *gorm.DB
	cfg        *config.Config
	authClient *grpcClient.AuthClient
	client     *http.Client
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
}
