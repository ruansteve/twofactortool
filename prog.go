package main

import (
	//"crypto"
	"encoding/base32"
	"flag"
	"fmt"
	"github.com/gwwfps/onetime"

	//"github.com/gwwfps/onetime"
	//"encoding/base64"
	//"github.com/sec51/twofactor"
)

var secret  string

func parseCommandOptions() {
	flag.StringVar(&secret, "secret", "", "secret to generate code")
}

func main() {
	parseCommandOptions()
	var otp,_  = onetime.Simple(6)
	data, err := base32.StdEncoding.DecodeString(secret)
	if (err == nil){
		var code = otp.TOTP(data)
		fmt.Println(code)
	}else{
		fmt.Println("error");
	}
}
