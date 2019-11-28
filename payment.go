package main

import (
	"errors"
	"flag"
	"fmt"
)

type paymentSummary struct {
	amount int
	fee float32
}

func main() {
	amount := flag.Int("amount", 100, "amount to pay")
	interestRate := flag.Float64("rate", 0.01, "interest rate to apply")
	flag.Parse()

	ps, err := pay(*amount, float32(*interestRate))
	if err == nil{
		fmt.Printf("You paid a fee of %v on %v", ps.fee * float32(ps.amount), ps.amount)

	} else{
		fmt.Printf("%v", err)
	}

}

func pay(amount int, fee float32) (paymentSummary, error){
	if fee > 0.09 {
		return paymentSummary{}, errors.New("fee is too high")
	}
	ps := paymentSummary{amount: amount, fee: fee}


	return ps, nil
}
