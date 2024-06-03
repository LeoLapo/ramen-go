package main

import (
	"log"

	"github.com/LeoLapo/ramen-go.git/middlewares"
	"github.com/LeoLapo/ramen-go.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	apiKey := "wgWNLxz68CrPFwWLLHhL37Ymdq1mibdU" // chave de API aqui

	router := gin.Default()

	// Adicionando middleware de CORS e autenticação
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.AuthMiddleware(apiKey)) // Aqui está sendo passada a chave de API

	// Configurando as rotas
	router.GET("/broths", routes.GetCaldos)
	router.OPTIONS("/broths", routes.OptionsHandler)
	router.GET("/proteins", routes.GetProteinas)
	router.OPTIONS("/proteins", routes.OptionsHandler)
	router.POST("/pedido", routes.PlacePedido)
	router.OPTIONS("/pedido", routes.OptionsHandler)

	log.Println("API server is running on port :8080")
	router.Run(":8080")
}
