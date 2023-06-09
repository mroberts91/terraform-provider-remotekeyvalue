// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package exampleserver

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// KeyValuePair defines model for KeyValuePair.
type KeyValuePair struct {
	// Id Id of the Key Value Pair.
	Id int64 `json:"id"`

	// IsSensitive Flag to indicate that the value of the Key Value Pair is a sensitive value.
	IsSensitive bool `json:"isSensitive"`

	// Key Key of the Key Value Pair.
	Key string `json:"key"`

	// Value Value of the Key Value Pair.
	Value string `json:"value"`
}
