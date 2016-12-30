package app

import (
	"encoding/json"
	"fmt"
)

func TestMaptoJSONStr() {

	callBackParams := map[string]interface{}{
		"uid":       36247,
		"productid": 32,
		"orderid":   34715624375624,
		"phone":     "646644142",
	}
	jsonStr, _ := json.Marshal(callBackParams)

	fmt.Println(string(jsonStr))

}
