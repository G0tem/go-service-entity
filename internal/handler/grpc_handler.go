package handler

import (
	"context"
	"time"

	"github.com/G0tem/go-service-entity/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// TestGrpc godoc
// @Summary Test gRPC connection to auth service
// @Description Test endpoint for gRPC communication with auth service
// @Tags grpc
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} types.SuccessResponseData
// @Failure 400 {object} types.FailureResponse
// @Failure 500 {object} types.FailureErrorResponse
// @Router /entity/test_grpc [get]
func (h *Handler) TestGrpc(c *fiber.Ctx) error {
	log.Info().Msg("Start TestGrpc endpoint")

	if h.authClient == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.FailureResponse{
			Status:  "error",
			Message: "gRPC client is not initialized",
		})
	}

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызываем gRPC метод GetTestData
	testMessage := "Hello from entity service!"
	response, err := h.authClient.GetTestData(ctx, testMessage)
	if err != nil {
		log.Error().Err(err).Msg("Failed to call auth gRPC service")
		return c.Status(fiber.StatusInternalServerError).JSON(types.FailureResponse{
			Status:  "error",
			Message: "Failed to call auth service: " + err.Error(),
		})
	}

	log.Info().
		Str("message", response.Message).
		Int32("status", response.Status).
		Str("timestamp", response.Timestamp).
		Msg("Successfully received response from auth gRPC service")

	return c.Status(fiber.StatusOK).JSON(types.SuccessResponseData{
		Status:  "success",
		Message: "gRPC call successful",
		Data: map[string]interface{}{
			"message":   response.Message,
			"status":    response.Status,
			"timestamp": response.Timestamp,
		},
	})
}
