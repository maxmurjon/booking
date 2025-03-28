package handler

import (
	"booking/pkg/helper/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization headerni olish
		authHeader := c.GetHeader("Authorization")
		token, err := helper.ExtractToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// Tokenni tekshirish va ma'lumotlarni olish
		info, err := helper.ParseClaims(token, h.cfg.SekretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}

		// Token ichidan rolni olish
		role := info.Role // `info.ClientType` noto'g'ri edi, to'g'ri chaqirilishi kerak

		// Faqat to‘g‘ri ro‘llarga ruxsat berish
		if role != "admin" && role != "customer" && role != "doctor" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid role"})
			return
		}

		// Kontekstga user va rolni qo‘shish
		c.Set("Auth", info)
		c.Set("Role", role)
		c.Next()
	}
}



func (h *Handler) RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("Role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			return
		}

		// Role ni string formatiga o‘girish
		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid role format"})
			return
		}

		// Ruxsat berilgan ro‘llarni tekshirish
		for _, allowedRole := range allowedRoles {
			if roleStr == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied for role: " + roleStr})
	}
}
