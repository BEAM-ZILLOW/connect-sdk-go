// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/tokens"
)

func deleteTokenExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper functions can be found in file Helper.go

	var query tokens.DeleteParams
	query.MandateCancelDate = newString("20150102")

	err := client.Merchant("merchantId").Tokens().Delete("tokenId", query, nil)

	fmt.Println(err)
}