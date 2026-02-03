package handler

import (
	"github.com/G0tem/go-service-entity/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// UserInfo godoc
// @Summary UserInfo info
// @Description Test endpoint UserInfo
// @Tags Entity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} types.SuccessResponseMe
// @Failure 400 {object} types.FailureResponse
// @Failure 500 {object} types.FailureErrorResponse
// @Router /entity/check [get]
func (h *Handler) UserInfo(c *fiber.Ctx) error {
	log.Info().Msg("Start GetEntity")

	claims := c.Locals("claims").(*JwtClaims)
	log.Debug().
		Str("email", claims.Email).
		Str("exp", claims.Exp.Format("3:04PM 2006-01-02")).
		Msg("Attempting to get user")

	return c.Status(fiber.StatusOK).JSON(types.SuccessResponseMe{
		UserID:      claims.UserID,
		Username:    claims.Username,
		Email:       claims.Email,
		Role:        claims.Role,
		Permissions: claims.Permissions,
		Exp:         claims.Exp,
	})
}
