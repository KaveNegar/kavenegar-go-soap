package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go-soap"
)

func main() {
	client := kavenegar.New()

	req := &kavenegar.RemainCreditByApiKey{
		Apikey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	res, err := client.RemainCreditByApiKey(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.RemainCreditByApiKeyResult)
	}
}
