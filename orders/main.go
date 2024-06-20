package main

import "context"

func main() {

	store := NewStore()
	svc := NewService()

	svc.CreateOrder(context.Background())
}
