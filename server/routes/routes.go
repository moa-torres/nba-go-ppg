package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moacirtorres/nba-go-ppg/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		cadastrar := main.Group("cadastrar")
		{
			cadastrar.POST("/", controllers.Cadastrar)
		}
	}
	{
		listar := main.Group("listar")
		{
			listar.GET("/", controllers.Listar)
			listar.GET("/:jogador", controllers.ListarUm)
		}
	}
	{
		atualizar := main.Group("atualizar")
		{
			atualizar.PATCH("/:jogador", controllers.Atualizar)
		}
	}
	{
		deletar := main.Group("deletar")
		{
			deletar.DELETE("/:jogador", controllers.Deletar)
		}
	}

	return router
}
