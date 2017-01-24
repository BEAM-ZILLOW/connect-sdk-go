// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import "fmt"

func getHostedCheckoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	response, err := client.Merchant("merchantId").Hostedcheckouts().Get("hostedCheckoutId", nil)

	fmt.Println(response, err)
}