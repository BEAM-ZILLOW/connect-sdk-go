// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payments

import (
	"net/http"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/token"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
)

// Client represents a payments client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// Create represents the resource /{merchantId}/payments
// Create payment
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// DeclinedPaymentError if the GlobalCollect platform declined / rejected the payment. The payment result will be available from the exception.
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Create(body payment.CreateRequest, context communication.CallContext) (payment.CreateResponse, error) {
	var resultObject payment.CreateResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments", nil)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			case http.StatusBadRequest:
				{
					errorObject = &payment.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			case http.StatusPaymentRequired:
				{
					errorObject = &payment.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			case http.StatusForbidden:
				{
					errorObject = &payment.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			case http.StatusBadGateway:
				{
					errorObject = &payment.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			case http.StatusServiceUnavailable:
				{
					errorObject = &payment.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Get represents the resource /{merchantId}/payments/{paymentId}
// Get payment
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Get(paymentID string, context communication.CallContext) (payment.Response, error) {
	var resultObject payment.Response

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, nil, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, getErr
	}

	return resultObject, nil
}

// Approve represents the resource /{merchantId}/payments/{paymentId}/approve
// Capture payment
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Approve(paymentID string, body payment.ApproveRequest, context communication.CallContext) (payment.ApprovalResponse, error) {
	var resultObject payment.ApprovalResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}/approve", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Cancel represents the resource /{merchantId}/payments/{paymentId}/cancel
// Cancel payment
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Cancel(paymentID string, context communication.CallContext) (payment.CancelResponse, error) {
	var resultObject payment.CancelResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}/cancel", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Cancelapproval represents the resource /{merchantId}/payments/{paymentId}/cancelapproval
// Undo capture payment request
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Cancelapproval(paymentID string, context communication.CallContext) (payment.CancelApprovalPaymentResponse, error) {
	var resultObject payment.CancelApprovalPaymentResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}/cancelapproval", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Processchallenged represents the resource /{merchantId}/payments/{paymentId}/processchallenged
// Approves challenged payment
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Processchallenged(paymentID string, context communication.CallContext) (payment.Response, error) {
	var resultObject payment.Response

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}/processchallenged", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Refund represents the resource /{merchantId}/payments/{paymentId}/refund
// Create refund
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// DeclinedRefundError if the GlobalCollect platform declined / rejected the refund. The refund result will be available from the exception.
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Refund(paymentID string, body refund.Request, context communication.CallContext) (refund.Response, error) {
	var resultObject refund.Response

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}/refund", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			case http.StatusBadRequest:
				{
					errorObject = &refund.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			case http.StatusNotFound:
				{
					errorObject = &refund.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Tokenize represents the resource /{merchantId}/payments/{paymentId}/tokenize
// Create a token from payment
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Tokenize(paymentID string, body payment.TokenizeRequest, context communication.CallContext) (token.CreateResponse, error) {
	var resultObject token.CreateResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payments/{paymentId}/tokenize", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// NewClient constructs a Payments Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Payments Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}