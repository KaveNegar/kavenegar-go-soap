package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go-soap"
)

func main() {
	client := kavenegar.New()
	request := &kavenegar.SendSimpleByApikey{
		Apikey: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		//Sender:   "xxxxxxxxxxxxx",
		Message:  "Hello Go!",
		Receptor: &kavenegar.ArrayOfString{String: []string{"0919xxxxxxx", "0936xxxxxxx"}},
	}
	res, err := client.SendSimpleByApikey(request)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, messageid := range res.SendSimpleByApikeyResult.Long {
			fmt.Println(messageid)
		}
	}
}
