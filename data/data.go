package data

import (
	"github.com/LeoLapo/ramen-go.git/models"
)

var Caldos = []string{"Miso", "Shoyu", "Tonkotsu"}

var Proteinas = []string{"Frango", "Carne de Porco", "Carne de Vaca"}

func GenerateBrothsMap() map[int]string {
	brothsMap := make(map[int]string)
	for i, broth := range Caldos {
		brothsMap[i] = broth
	}
	return brothsMap
}

func GenerateProteinsMap() map[int]string {
	proteinsMap := make(map[int]string)
	for i, protein := range Proteinas {
		proteinsMap[i] = protein
	}
	return proteinsMap
}

var BrothsMap = GenerateBrothsMap()
var ProteinsMap = GenerateProteinsMap()

var Pedidos []models.Pedido
