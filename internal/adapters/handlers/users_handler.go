package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nazrawigedion123/user-service/internal/application"
	"github.com/nazrawigedion123/user-service/internal/domain"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserService *application.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService *application.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		UserService: userService,
		logger:      logger,
	}
}

// @Summary Create a new user
// @Description Creates a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User object to be created"
// @Success 201 {object} domain.User
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.CreateUser(c.Request.Context(), &user); err != nil {
		h.logger.Error("error creating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)

}

// @Summary Get user by ID
// @Description Retrieves a user by their unique ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H "Invalid ID format"
// @Failure 500 {object} gin.H "User not found or internal server error"

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	id_uuid, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error("error parsing id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := h.UserService.GetUserByID(c.Request.Context(), id_uuid)
	if err != nil {
		h.logger.Error("error getting user	by id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

// @Summary Get user by email
// @Description Retrieves a user by their email address
// @Tags users
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H "User not found or invalid email"
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.UserService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {

		h.logger.Error("error getting user by email", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

// @Summary Get user by username
// @Description Retrieves a user by their username
// @Tags users
// @Produce json
// @Param username path string true "User Username"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H "User not found or invalid username"
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.UserService.GetUserByUsername(c.Request.Context(), username)
	if err != nil {

		h.logger.Error("error getting user by username", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

// @Summary List all users
// @Description Retrieves a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} gin.H "Internal server error"
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.UserService.ListUsers(c.Request.Context())
	if err != nil {
		h.logger.Error("error listing users", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

// @Summary Update an existing user
// @Description Updates an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User object with updated fields"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.UpdateUser(c.Request.Context(), &user); err != nil {
		h.logger.Error("error updating user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

// @Summary Delete a user by ID
// @Description Deletes a user by their unique ID (soft delete)
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} gin.H "User deleted successfully"
// @Failure 400 {object} gin.H "Invalid ID format"
// @Failure 500 {object} gin.H "Internal server error"
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	id_uuid, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error("error parsing id", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.UserService.DeleteUser(c.Request.Context(), id_uuid); err != nil {
		h.logger.Error("error deleting user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}
