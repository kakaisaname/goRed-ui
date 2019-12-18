package main

import (
	"encoding/json"
	"fmt"
	acservices "github.com/kakaisaname/account/services"
)

func main() {

	d, e := json.Marshal(&acservices.AccountTransferDTO{})
	fmt.Println(e)
	fmt.Println(string(d))
}
