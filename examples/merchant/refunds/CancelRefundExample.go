// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import "fmt"

func cancelRefundExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	err := client.Merchant("merchantId").Refunds().Cancel("refundId", nil)

	fmt.Println(err)
}