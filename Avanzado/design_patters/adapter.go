package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BanckPayment struct{}

type BanckPaymentAdapter struct {
	BanckPayment *BanckPayment
	bankAccount  int
}

func (bpa *BanckPaymentAdapter) Pay() {
	bpa.BanckPayment.Pay(bpa.bankAccount)
}
func (BanckPayment) Pay(bankAccount int) {
	fmt.Printf("Pating using Banckaccount %d\n", bankAccount)
}

func maiin() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	/* banck := &BanckPayment{}
	ProcessPayment(banck) */
	bpa := &BanckPaymentAdapter{
		bankAccount:  5,
		BanckPayment: &BanckPayment{},
	}
	ProcessPayment(bpa)
}
