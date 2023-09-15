## Saga in Golang 


#### Step to implement saga.
- Take package from the www.github.com/gondsuryaprakash/saga
```go
go get www.github.com/gondsuryaprakash/sagapattern
```

- Example 
```go 
package main 

import (
	"errors"
	"fmt"
	"log"

	"sagapattern/saga"
)
type Order struct {
	OrderId  int
	Product  string
	Amount   int
	Qauntity int
}
func (o *Order) AddOrder(product string, amount int, qantity int) error {
	fmt.Println("Order Successfull")
	return nil
}
func (o *Order) CancelOrder(orderId int) error {
	fmt.Println("Oreder Cancle Successfully")
	return nil
}
type Payment struct {
	PaymentId int
	Amount    int
	OrderId   int
}

// Add Payment
func (p *Payment) AddPayment(orderId int, amount int) error {
	// fmt.Println("Payment Successfull")
	return errors.New("Payment failed !!")
}

// Refund done
func (p *Payment) Refund(paymentId int, amount int) error {
	fmt.Println("Refund Done ")
	return nil
}

func main() {
	saga := saga.NewSaga()
	orderService := Order{}
	paymentService := Payment{}
	// Add step for Place order
	saga.AddStep(
		func() error { return orderService.AddOrder("Product1", 250, 1) },
		func() error { return orderService.CancelOrder(1) })

	saga.AddStep(
		func() error { return paymentService.AddPayment(1, 250) },
		func() error { return paymentService.Refund(1, 250) })
	if err := saga.Exec(); err != nil {
		log.Fatalf("Failed To Execute saga %v", err)
	}
	fmt.Println("Saga run completely !!")
}
```