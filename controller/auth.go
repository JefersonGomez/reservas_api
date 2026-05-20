package controller

import (
	"api_reservas/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "datos invalidos"})
		return
	}

	var Usuario models.Usuario

	DB.Where("email = ?", body.Email).First(&Usuario)

	if Usuario.ID == 0 {
		c.JSON(404, gin.H{"error": "id no encontrado"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(Usuario.Password), []byte(body.Password))

	if err != nil {
		c.JSON(404, gin.H{"error": " contraseñas no coinciden"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  Usuario.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"error": "error generando token"})
		return
	}
	c.JSON(200, gin.H{"token": tokenString})

}

func Registro(c *gin.Context) {

	var Usuario models.Usuario

	if err := c.BindJSON(&Usuario); err != nil {
		c.JSON(400, gin.H{"error": "datos incorrectos"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(Usuario.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{"error": "error al encriptar la contraseña"})
		return
	}

	Usuario.Password = string(hash)
	DB.Create(&Usuario)
	c.JSON(200, gin.H{"messaje": "usuario creado correctamente"})

}
