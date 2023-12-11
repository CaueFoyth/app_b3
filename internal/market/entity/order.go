package entity

//Classe da ordem
type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string //Buy or sell
	Status        string
	Transactions  []*Transaction
}

//Adicionar a ordem
func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		ID:           orderID,
		Investor:     investor,
		Asset:        asset,
		Shares:       shares,
		Price:        price,
		OrderType:    orderType,
		Status:       "OPEN",
		Transactions: []*Transaction{},
	}
}
