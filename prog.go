package main

import (
	"fmt"
	"github.com/gwwfps/onetime"
	"encoding/base32"
	"io/ioutil"
	"os"
)

func main() {
	var otp,_  = onetime.Simple(6)
	str := readStdin()
	data, err := base32.StdEncoding.DecodeString(str)
	if (err == nil){
		var code = otp.TOTP(data)
		fmt.Println(code)
	}else{
		fmt.Println("error");
	}
}

func readStdin() string{
    bytes, _ := ioutil.ReadAll(os.Stdin)
	return string(bytes)
}
