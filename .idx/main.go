package main

import (
	"fmt"
	"net/http"

	"example.com/m/bankex"
)

func sum(i, j int) int {
	return i + j
}

func main() {
	fmt.Println("go ")
	fmt.Println(sum(1, 2))

	resp, err := http.Get("https://api.production.wealthsimple.com/v1/oauth/v2/token")
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
