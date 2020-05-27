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
// V3Seasonality struct for V3Seasonality
type V3Seasonality struct {
	// The end date when the item is in season. In ISO 8601 calendar date format YYYY-MM-DD.
	EndDate string `json:"end_date,omitempty"`
	// The start date when the item is in season. In ISO 8601 calendar date format YYYY-MM-DD.
	StartDate string `json:"start_date,omitempty"`
}
