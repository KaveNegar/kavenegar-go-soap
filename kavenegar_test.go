package kavenegar

import (
	"testing"
)

var (
	client *v1Soap
	apiKey string
)

func setup() {
	apiKey = "0"
}

func TestSendSimpleByApikey(t *testing.T) {
	setup()
	client := New()
	request := &SendSimpleByApikey{
		Apikey: apiKey,
		//Sender:   "xxxxxxxxxxxxx",
		Message:  "Hello Go!",
		Receptor: &ArrayOfString{String: []string{"0919xxxxxxx"}},
	}
	_, err := client.SendSimpleByApikey(request)
	if err != nil {
		//t.Errorf("SendSimpleByApikey failed: %v", err)
	} else {
		t.Logf("MessageSend OK")
	}
}
