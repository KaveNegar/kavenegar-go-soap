# kavenegar-go-soap
[![Build Status](https://travis-ci.org/KaveNegar/kavenegar-go-soap.svg?branch=master)](https://travis-ci.org/KaveNegar/kavenegar-go-soap)

## Installation
```
go get github.com/kavenegar/kavenegar-go-soap
```
## Usage

### Send
```golang
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

```
## Contribution
Bug fixes, docs, and enhancements welcome! Please let us know support@kavenegar.com



<hr>

<div dir='rtl'>
	
## راهنما

### معرفی سرویس کاوه نگار

کاوه نگار یک وب سرویس ارسال و دریافت پیامک و تماس صوتی است که به راحتی میتوانید از آن استفاده نمایید.

### ساخت حساب کاربری

اگر در وب سرویس کاوه نگار عضو نیستید میتوانید از [لینک عضویت](http://panel.kavenegar.com/client/membership/register) ثبت نام  و اکانت آزمایشی برای تست API دریافت نمایید.

### مستندات

برای مشاهده اطلاعات کامل مستندات [وب سرویس پیامک](http://kavenegar.com/وب-سرویس-پیامک.html)  به صفحه [مستندات وب سرویس](http://kavenegar.com/rest.html) مراجعه نمایید.

### راهنمای فارسی

در صورتی که مایل هستید راهنمای فارسی کیت توسعه کاوه نگار را مطالعه کنید به صفحه [کد ارسال پیامک](http://kavenegar.com/sdk.html) مراجعه نمایید.

### اطالاعات بیشتر
برای مطالعه بیشتر به صفحه معرفی
[وب سرویس اس ام اس ](http://kavenegar.com)
کاوه نگار
مراجعه نمایید .

 اگر در استفاده از کیت های سرویس کاوه نگار مشکلی یا پیشنهادی  داشتید ما را با یک Pull Request  یا  ارسال ایمیل به support@kavenegar.com  خوشحال کنید.
 
##
![http://kavenegar.com](http://kavenegar.com/public/images/logo.png)		

[http://kavenegar.com](http://kavenegar.com)	

</div>


