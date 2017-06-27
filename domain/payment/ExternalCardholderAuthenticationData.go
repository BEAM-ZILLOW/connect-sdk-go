// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// ExternalCardholderAuthenticationData represents class ExternalCardholderAuthenticationData
type ExternalCardholderAuthenticationData struct {
	Cavv             *string `json:"cavv,omitempty"`
	CavvAlgorithm    *string `json:"cavvAlgorithm,omitempty"`
	Eci              *int32  `json:"eci,omitempty"`
	ValidationResult *string `json:"validationResult,omitempty"`
	Xid              *string `json:"xid,omitempty"`
}

// NewExternalCardholderAuthenticationData constructs a new ExternalCardholderAuthenticationData
func NewExternalCardholderAuthenticationData() *ExternalCardholderAuthenticationData {
	return &ExternalCardholderAuthenticationData{}
}
