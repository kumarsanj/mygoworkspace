package main

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/m/bankex"
)

func sum(i, j int) int {
	return i + j
}

func main() {
	fmt.Println("go ")
	fmt.Println(sum(1, 2))

	payload := strings.NewReader(`{"grant_type":"password","username":"kumar.sanjay@gmail.com","password":"","skip_provision":true,"otp_claim":null,"scope":"invest.read invest.write trade.read trade.write tax.read tax.write","client_id":"4da53ac2b03225bed1550eba8e4611e086c7b905a3855e6ed12ea08c246758fa"}`)
	//payload := strings.NewReader(`{"grant_type":"password","username":"kumar.sanjay@gmail.com","password":"","skip_provision":true,"otp_claim":null,"scope":"invest.read invest.write trade.read trade.write tax.read tax.write","client_id":"4da53ac2b03225bed1550eba8e4611e086c7b905a3855e6ed12ea08c246758fa"}`)
	resp, err := http.Post("https://api.production.wealthsimple.com/v1/oauth/v2/token",
		"application/json",
		payload,
	)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status: ", resp.Status)
	fmt.Println("balance: ", bankex.Balance())
	fmt.Println("withdraw: ", bankex.Withdraw(10))
	fmt.Println("balance: ", bankex.Balance())
	fmt.Println("go end")
}
