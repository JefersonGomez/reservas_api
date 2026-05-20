package controller

import (
	"api_reservas/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//operaciones crud

var DB *gorm.DB

func ObtenerData(c *gin.Context) {
	var reservas []models.Reserva
	DB.Find(&reservas)
	c.JSON(200, reservas)

}

func ObtenerDataById(c *gin.Context) {
	var reserva models.Reserva
	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": "id no encotrado"})
		return
	}

	resultado := DB.First(&reserva, id)
	if resultado.Error != nil {
		c.JSON(404, gin.H{"error": "nose encotro la reserva"})
	}
	c.JSON(200, reserva)

}

func CrearDato(c *gin.Context) {

	var nuevaReserva models.Reserva

	if err := c.BindJSON(&nuevaReserva); err != nil {
		c.JSON(400, gin.H{"error": "datos incorrectos"})
		return
	}

	DB.Create(&nuevaReserva)
	c.JSON(200, nuevaReserva)

}

func EliminarDato(c *gin.Context) {

	var reserva models.Reserva

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": "id incorrecto"})
		return
	}

	resultado := DB.First(&reserva, id)

	if resultado.Error != nil {
		c.JSON(404, gin.H{"error": "reserva no encontrada"})
		return
	}

	DB.Delete(&reserva, id)
	c.JSON(200, gin.H{"message": "reserva eliminada correctamente"})

}

func EditarDato(c *gin.Context) {

	var reserva models.Reserva

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": "id incorrecto"})
		return
	}

	resultado := DB.First(&reserva, id)

	if resultado.Error != nil {
		c.JSON(404, gin.H{"error": "reserva no encontrada"})
		return
	}

	if err := c.BindJSON(&reserva); err != nil {
		c.JSON(400, gin.H{"error": "datos incorrectos"})
		return
	}

	DB.Save(&reserva)
	c.JSON(200, reserva)

}

func ConfirmarReserva(c *gin.Context) {

	var reserva models.Reserva

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": "id incorrecto"})
		return
	}

	resultado := DB.First(&reserva, id)

	if resultado.Error != nil {
		c.JSON(404, gin.H{"error": "reserva no encontrada"})
		return
	}

	DB.Model(&reserva).Update("estado", "confirmada")
	c.JSON(200, gin.H{"message": "reserva confirmada"})
}

func CancelarReserva(c *gin.Context) {
	var reserva models.Reserva

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": "id incorrecto"})
		return
	}

	resultado := DB.First(&reserva, id)

	if resultado.Error != nil {
		c.JSON(404, gin.H{"error": "reserva no encontrada"})
		return
	}

	DB.Model(&reserva).Update("estado", "cancelada")
	c.JSON(200, gin.H{"message": "reserva confirmada"})

}
