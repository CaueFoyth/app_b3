package entity

import (
	"container/heap"
	"sync"
)

// Livro com todos os detalhes de entradas e saídas
type Book struct {
	Order         []*Order
	Transaction   []*Transaction
	OrdersChan    chan *Order //Input
	OrdersChanOut chan *Order
	Wg            *sync.WaitGroup
}

// Inserindo ao livro
func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:         []*Order{},
		Transaction:   []*Transaction{},
		OrdersChan:    orderChan,
		OrdersChanOut: orderChanOut,
		Wg:            wg,
	}
}

// Match de ordens
func (b *Book) Trade() {
	buyOrders := NewOrderQeue()
	sellOrders := NewOrderQeue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range b.OrdersChan {
		if order.OrderType == "BUY" {
			buyOrders.Push(order)
			if sellOrders.Len() > 0 && sellOrders.Orders[0].Price <= order.Price {
				sellOrder := sellOrders.Pop().(*Order)
				if sellOrder.PendingShares > 0 {
					trasaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					b.AddTransaction(trasaction, b.Wg)
					sellOrder.Transactions = append(sellOrder.Transactions, trasaction)
					order.Transactions = append(order.Transactions, trasaction)
					b.OrdersChanOut <- sellOrder
					b.OrdersChanOut <- order
					if sellOrder.PendingShares > 0 {
						sellOrders.Push(sellOrder)
					}
				}
			}
		} else if order.OrderType == "SELL" {
			sellOrders.Push(order)
			if buyOrders.Len() > 0 && buyOrders.Orders[0].Price >= order.Price {
				buyOrder := buyOrders.Pop().(*Order)
				if buyOrder.PendingShares > 0 {
					transaction := NewTransaction(order, buyOrder, order.Shares, order.Price)
					b.AddTransaction(transaction, b.Wg)
					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrdersChanOut <- buyOrder
					b.OrdersChanOut <- order
					if buyOrder.PendingShares > 0 {
						buyOrders.Push(buyOrder)
					}
				}
			}
		}
	}
}

func (b *Book) AddTransaction(trasaction *Transaction, wg *sync.WaitGroup) {

	defer wg.Done()
	sellingShares := trasaction.SellingOrder.PendingShares
	buyingShares := trasaction.SellingOrder.PendingShares

	minShares := sellingShares
	if buyingShares < minShares {
		minShares = buyingShares
	}

	trasaction.SellingOrder.Investor.UpdateAssetPosition(trasaction.SellingOrder.Asset.ID, -minShares)
	trasaction.SellingOrder.PendingShares -= minShares
	trasaction.BuyingOrder.Investor.UpdateAssetPosition(trasaction.BuyingOrder.Asset.ID, minShares)
	trasaction.BuyingOrder.PendingShares -= minShares

	trasaction.Total = float64(trasaction.Shares) * trasaction.BuyingOrder.Price

	if trasaction.BuyingOrder.PendingShares == 0 {
		trasaction.BuyingOrder.Status = "CLOSED"
	}
	if trasaction.SellingOrder.PendingShares == 0 {
		trasaction.SellingOrder.Status = "CLOSED"
	}
	b.Transaction = append(b.Transaction, trasaction)
}
