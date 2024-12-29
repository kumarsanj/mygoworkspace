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
		chromedp.Navigate(`https://api.production.wealthsimple.com/v1/oauth/v2/token`),
		chromedp.FullScreenshot("#content", &outer, chromedp.ByQuery),
	)

	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(outer)
	fmt.Println("go end")
}
