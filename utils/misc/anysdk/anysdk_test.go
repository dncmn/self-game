package anysdk

import (
	"testing"
)

func TestCheckPaySignFunc(t *testing.T) {
	param := map[string]string{
		"order_id":       "PB71562015060111462144722",
		"product_count":  "1",
		"amount":         "1.0",
		"pay_status":     "1",
		"pay_time":       "2015-06-01 11:46:21",
		"user_id":        "85282",
		"order_type":     "88",
		"game_user_id":   "80973",
		"server_id":      "4",
		"product_name":   "金币",
		"product_id":     "6",
		"private_data":   "buy100gold",
		"channel_number": "000023",
		"enhanced_sign":  "5172b67c5e743bad6885ccc9c3f65f97",
		"sign":           "04e96a1752751d5a05451d0255b2bf47",
	}
	privateKey := "B742866570AA1856BD24124CEE856203"
	enhancedKey := "ZGYzMDIzMzMwODViZWIyNDJkZjY"
	checkSign := CheckPaySignFunc(privateKey, enhancedKey)
	if checkSign == nil {
		t.Fatal("Failed to create checkSign func")
	}
	sign, enhancedSign := checkSign(param)
	t.Logf("Sign: %v, EnhancedSign: %v", sign, enhancedSign)
	if !sign || !enhancedSign {
		t.Fail()
	}
}
