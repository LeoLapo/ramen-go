package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/LeoLapo/ramen-go.git/data"
	"github.com/LeoLapo/ramen-go.git/models"
	"github.com/LeoLapo/ramen-go.git/utils"
	"github.com/gin-gonic/gin"
)

// PedidoResponse representa a resposta desejada para a rota PlacePedido
type PedidoResponse struct {
	PedidoID string `json:"pedido_id"`
	Broth    string `json:"broth"`
	Protein  string `json:"protein"`
}

func GetCaldos(c *gin.Context) {
	log.Println("Handler GetCaldos chamado")
	c.JSON(http.StatusOK, data.Caldos)
}

func GetProteinas(c *gin.Context) {
	log.Println("Handler GetProteinas chamado")
	c.JSON(http.StatusOK, data.Proteinas)
}

func PlacePedido(c *gin.Context) {
	log.Println("Handler PlacePedido chamado")
	var pedido models.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := utils.GenerateOrderID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate order ID"})
		return
	}

	pedido.PedidoID = strconv.Itoa(orderID)
	brothName, ok := data.BrothsMap[pedido.BrothID]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BrothID"})
		return
	}
	proteinName, ok := data.ProteinsMap[pedido.ProteinID]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ProteinID"})
		return
	}

	pedido.Broth = brothName
	pedido.Protein = proteinName
	data.Pedidos = append(data.Pedidos, pedido)

	response := PedidoResponse{
		PedidoID: pedido.PedidoID,
		Broth:    pedido.Broth,
		Protein:  pedido.Protein,
	}

	c.JSON(http.StatusOK, response)
}

func OptionsHandler(c *gin.Context) {
	log.Println("Handler OptionsHandler chamado")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, X-API-Key, Authorization")
	c.JSON(http.StatusOK, struct{}{})
}
