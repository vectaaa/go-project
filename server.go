package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	log.Println("Listening on localhost:4242...")
    var err error = http.ListenAndServe("invalid", nil) //This is responsible for making the app run continuously
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Something happened")
	}
    

	// var itemReturned string = returnsValue("Hello, i returned this to you.")
	// fmt.Println(itemReturned) 
}

func handleCreatePaymentIntent (res http.ResponseWriter, r *http.Request ) {
	fmt.Println("ENDPOINT CALLED!")
}

// func returnsValue(somethingR string) string {
// 	return somethingR
// }