// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundCardMethodSpecificOutput represents class RefundCardMethodSpecificOutput
type RefundCardMethodSpecificOutput struct {
	TotalAmountPaid     *int64 `json:"totalAmountPaid,omitempty"`
	TotalAmountRefunded *int64 `json:"totalAmountRefunded,omitempty"`
}

// NewRefundCardMethodSpecificOutput constructs a new RefundCardMethodSpecificOutput
func NewRefundCardMethodSpecificOutput() *RefundCardMethodSpecificOutput {
	return &RefundCardMethodSpecificOutput{}
}
