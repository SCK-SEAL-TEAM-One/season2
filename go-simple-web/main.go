package main

import (
	"fmt"
	vm "./vendingmachine"
	"log"
	"net/http"
)

func main() {
	v := vm.NewVendingMachine()

	http.HandleFunc("/insertcoins", func(w http.ResponseWriter, r *http.Request) {
		// insertcoin
		coin := r.FormValue("coin")
		v.InsertCoins(coin)
		fmt.Fprintf(w, "%v", v.ShowTotalBalance())
	})

	http.HandleFunc("/buydrink", func(w http.ResponseWriter, r *http.Request) {
		// buy
		drink := r.FormValue("drink")
		fmt.Fprintf(w, "%v", v.BuyDrink(drink))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
