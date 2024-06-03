package models

type Pedido struct {
	PedidoID  string `json:"pedido_id"`
	BrothID   int    `json:"broth_id"`
	ProteinID int    `json:"protein_id"`
	Broth     string `json:"broth"`
	Protein   string `json:"protein"`
}

type Order struct {
	OrderID   int    `json:"order_id"`
	BrothID   int    `json:"broth_id"`
	ProteinID int    `json:"protein_id"`
	Quantity  int    `json:"quantity"`
	Notes     string `json:"notes"`
}
