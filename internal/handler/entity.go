package handler

import (
	"github.com/G0tem/go-service-entity/internal/model"
	"github.com/G0tem/go-service-entity/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// GetEntity godoc
// @Summary GetEntity info
// @Description Test endpoint crud
// @Tags Entity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} types.SuccessResponseData
// @Failure 400 {object} types.FailureResponse
// @Failure 500 {object} types.FailureErrorResponse
// @Router /entity/get [get]
func (h *Handler) GetEntity(c *fiber.Ctx) error {
	log.Info().Msg("Start GetEntity")

	claims := c.Locals("claims").(*JwtClaims)
	log.Debug().
		Str("email", claims.Email).
		Str("exp", claims.Exp.Format("3:04PM 2006-01-02")).
		Msg("Attempting to get user")

	var result []model.Entity
	if err := h.db.Where("user_id = ?", claims.UserID).Find(&result).Error; err != nil {
		log.Error().Msg("Failed to query entities from database")
		return c.Status(fiber.StatusInternalServerError).JSON(types.FailureErrorResponse{
			Status:  "error",
			Message: "Failed to query entities",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.SuccessResponseData{
		Status:  "Success",
		Message: "Entities retrieved successfully",
		Data:    result,
	})
}

// CreateEntity godoc
// @Summary CreateEntity info
// @Description Test endpoint crud
// @Tags Entity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param company body types.EntityRequest true "Entity data"
// @Success 200 {array} types.SuccessResponseData
// @Failure 400 {object} types.FailureResponse
// @Failure 500 {object} types.FailureErrorResponse
// @Router /entity/create [post]
func (h *Handler) CreateEntity(c *fiber.Ctx) error {
	log.Info().Msg("Start CreateEntity")

	claims := c.Locals("claims").(*JwtClaims)
	log.Debug().
		Str("email", claims.Email).
		Str("exp", claims.Exp.Format("3:04PM 2006-01-02")).
		Msg("Attempting to get user")

	var req types.EntityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.FailureResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	err := h.db.Create(&model.Entity{
		UserID:      uuid.MustParse(claims.UserID),
		Description: &req.Description,
	}).Error
	if err != nil {
		log.Error().Err(err).Msg("Failed to create entity in database")
		return c.Status(fiber.StatusInternalServerError).JSON(types.FailureErrorResponse{
			Status:  "error",
			Message: "Failed to create entity",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.SuccessResponseData{
		Status:  "Success",
		Message: "Entity created successfully",
	})
}

// UpdateEntity godoc
// @Summary UpdateEntity info
// @Description Test endpoint crud
// @Tags Entity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Entity ID"
// @Param company body types.EntityRequest true "Entity data"
// @Success 200 {array} types.SuccessResponseData
// @Failure 400 {object} types.FailureResponse
// @Failure 500 {object} types.FailureErrorResponse
// @Router /entity/update/{id} [patch]
func (h *Handler) UpdateEntity(c *fiber.Ctx) error {
	log.Info().Msg("Start UpdateEntity")

	claims := c.Locals("claims").(*JwtClaims)
	log.Debug().
		Str("email", claims.Email).
		Str("exp", claims.Exp.Format("3:04PM 2006-01-02")).
		Msg("Attempting to get user")

	entityIDStr := c.Params("id")
	entityID, err := uuid.Parse(entityIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.FailureResponse{
			Status:  "error",
			Message: "Invalid entity ID",
		})
	}

	var req types.EntityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.FailureResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	err = h.db.Model(&model.Entity{}).
		Where("user_id = ? AND id = ?", claims.UserID, entityID).
		Updates(model.Entity{
			Description: &req.Description,
		}).Error
	if err != nil {
		log.Error().Err(err).Msg("Failed to update entity in database")
		return c.Status(fiber.StatusInternalServerError).JSON(types.FailureErrorResponse{
			Status:  "error",
			Message: "Failed to update entity",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.SuccessResponseData{
		Status:  "Success",
		Message: "Entities updated successfully",
	})
}

// DeleteEntity godoc
// @Summary DeleteEntity info
// @Description Test endpoint crud
// @Tags Entity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Entity ID"
// @Success 200 {array} types.SuccessResponseData
// @Failure 400 {object} types.FailureResponse
// @Failure 500 {object} types.FailureErrorResponse
// @Router /entity/delete/{id} [delete]
func (h *Handler) DeleteEntity(c *fiber.Ctx) error {
	log.Info().Msg("Start DeleteEntity")

	claims := c.Locals("claims").(*JwtClaims)
	log.Debug().
		Str("email", claims.Email).
		Str("exp", claims.Exp.Format("3:04PM 2006-01-02")).
		Msg("Attempting to get user")

	entityIDStr := c.Params("id")
	entityID, err := uuid.Parse(entityIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.FailureResponse{
			Status:  "error",
			Message: "Invalid entity ID",
		})
	}

	err = h.db.Where("user_id = ? AND id = ?", claims.UserID, entityID).Delete(&model.Entity{}).Error
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete entities from database")
		return c.Status(fiber.StatusInternalServerError).JSON(types.FailureErrorResponse{
			Status:  "error",
			Message: "Failed to delete entities",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.SuccessResponseData{
		Status:  "Success",
		Message: "Entities deleted successfully",
	})
}
