package middleware

import (
	"net/http"
	"siakad/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Require auth cek keberadaan dan keabsahan token JWT
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succes":  false,
				"message": "Akses ditolak atau token tidak ditemukan",
			})
			c.Abort()
			return
		}

		//Token Formatnya "Bearer xxxxx.yyyyy.zzzzz"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succes":  false,
				"message": "Format Token tidak Valid",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := utils.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"succes":  false,
				"message": "Token tidak valid atau Sudah Kadaluarsa.",
			})
			c.Abort()
			return
		}

		//Ambil Payload/claims dari token dan simpan di context request
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])
		}
		c.Next()
	}
}

// RequireROle untuk membatasi akses hanya untuk role tertentu
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exist := c.Get("role")
		if !exist {
			c.JSON(http.StatusForbidden, gin.H{
				"succes":  false,
				"message": "akses ditolak",
			})
			c.Abort()
			return
		}

		roleValid := false
		for _, role := range roles {
			if userRole == role {
				roleValid = true
				break
			}
		}
		if !roleValid {
			c.JSON(http.StatusForbidden, gin.H{
				"succes":  false,
				"message": "Anda tidak memiliki izin(Role) untuk mengakses resource ini",
			})
			c.Abort()
			return
		}
	}
}
