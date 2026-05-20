package routes

import (
	"api_reservas/controller"
	"api_reservas/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/registro", controller.Registro)
	r.POST("/login", controller.Login)

	protegidas := r.Group("/")
	protegidas.Use(middlewares.VerificarToken)
	{
		protegidas.GET("/reservas", controller.ObtenerData)
		protegidas.GET("/reservas/:id", controller.ObtenerDataById)
		protegidas.POST("/reservas", controller.CrearDato)
		protegidas.DELETE("/reservas/:id", controller.EliminarDato)
		protegidas.PUT("/reservas/:id", controller.EditarDato)
		protegidas.PATCH("/reservas/:id/confirmar", controller.ConfirmarReserva)
		protegidas.PATCH("/reservas/:id/cancelar", controller.CancelarReserva)

	}

}
