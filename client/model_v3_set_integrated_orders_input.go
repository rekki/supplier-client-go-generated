/*
 * Rekki.com Supply API
 *
 * The base URL for all API endpoints is https://api.rekki.com  Api key value contains of word Bearer together with api key that you can get from integrations@rekki.com 
 *
 * API version: 
 * Contact: integrations@rekki.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
// V3SetIntegratedOrdersInput struct for V3SetIntegratedOrdersInput
type V3SetIntegratedOrdersInput struct {
	// Array of References of the orders to mark as integrated, required. Order refs are discoverable when listing orders.
	Orders []string `json:"orders,omitempty"`
}
