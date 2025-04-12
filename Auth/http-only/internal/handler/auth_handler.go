package handler

import (
	"net/http"
	"time"

	"github.com/Rai-Sahil/backend/internal/domain"
	"github.com/Rai-Sahil/backend/internal/middleware"
	"github.com/Rai-Sahil/backend/internal/service"
	"github.com/Rai-Sahil/backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	// Setup up auth router groups
	authRouter := router.Group("/auth")
	
	// Public routes
	authRouter.POST("/register", h.Register)
	authRouter.POST("/login", h.Login)
	
	// Setup middleware -- Below are protected routes
	authRouter.Use(middleware.AuthMiddleware())
	authRouter.POST("/test", h.Test)
}

func (h *AuthHandler) Register (c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "user registered"})
}

func (h *AuthHandler) Login (c *gin.Context) {
	var body struct {
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrorWithFileInfo(err)})
		return
	}

	user, err := h.Service.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": utils.ErrorWithFileInfo(err)})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": utils.ErrorWithFileInfo(err)})
		return
	}
	
	http.SetCookie(c.Writer, &http.Cookie{
		Name: 		"auth_token",
		Value:		token,
		Path: 		"/",
		HttpOnly: 	true,
		Secure: 	false, // True if in prod
		SameSite: 	http.SameSiteStrictMode,
		Expires: 	time.Now().Add(24 * time.Hour),
	})

	c.JSON(http.StatusAccepted, gin.H{"message": "login succesful", "user": user})
}

func (h *AuthHandler) Test (c *gin.Context) {
	userIdVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userId, ok := userIdVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userId})
}
