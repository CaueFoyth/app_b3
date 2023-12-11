package entity

//Ordem das listas de compras e vendas
type OrderQeue struct {
	Orders []*Order
}

//Less Valor i < Valor j
func (oq *OrderQeue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

//Swap Valor i < Valor j
func (oq *OrderQeue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

//Len tamanho dos dados
func (oq *OrderQeue) Len() int {
	return len(oq.Orders)
}

//Push Adiciona novos (Append)
func (oq *OrderQeue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

//POP Remove
func (oq *OrderQeue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQeue() *OrderQeue {
	return &OrderQeue{}
}
