package main

import (
	"fmt"

	"github.com/kavenegar/kavenegar-go-soap"
)

func main() {
	client := kavenegar.New()

	req := &kavenegar.SendArrayByApikey{
		Apikey:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Sender:   &kavenegar.ArrayOfString{String: []string{"xxxxxxxxxxxxx", "xxxxxxxxxxxxx"}},
		Message:  &kavenegar.ArrayOfString{String: []string{"Hello Go!", "Hello Go!"}},
		Receptor: &kavenegar.ArrayOfString{String: []string{"0919xxxxxxx", "0936xxxxxxx"}},
		Msgmode:  &kavenegar.ArrayOfInt{Int: []int32{1, 1}},
	}
	res, err := client.SendArrayByApikey(req)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, messageid := range res.SendArrayByApikeyResult.Long {
			fmt.Println(messageid)
		}
	}
}
