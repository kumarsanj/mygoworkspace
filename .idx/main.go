package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func sum(i, j int) int {
	return i + j
}

func main() {
	fmt.Println("go")
	fmt.Println(sum(1, 2))

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.google.com`),
		//		chromedp.WaitVisible(`body > div.L3eUgb > div.o3j9`),
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("go end")
}
