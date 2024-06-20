package main

import "context"

type OrderService interface {
	CreateOrder(context.Context)
}

type OrdersStore interface {
}
