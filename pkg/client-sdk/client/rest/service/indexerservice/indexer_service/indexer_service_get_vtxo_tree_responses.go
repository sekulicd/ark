// Code generated by go-swagger; DO NOT EDIT.

package indexer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ark-network/ark/pkg/client-sdk/client/rest/service/models"
)

// IndexerServiceGetVtxoTreeReader is a Reader for the IndexerServiceGetVtxoTree structure.
type IndexerServiceGetVtxoTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexerServiceGetVtxoTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexerServiceGetVtxoTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIndexerServiceGetVtxoTreeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIndexerServiceGetVtxoTreeOK creates a IndexerServiceGetVtxoTreeOK with default headers values
func NewIndexerServiceGetVtxoTreeOK() *IndexerServiceGetVtxoTreeOK {
	return &IndexerServiceGetVtxoTreeOK{}
}

/*
IndexerServiceGetVtxoTreeOK describes a response with status code 200, with default header values.

A successful response.
*/
type IndexerServiceGetVtxoTreeOK struct {
	Payload *models.V1GetVtxoTreeResponse
}

// IsSuccess returns true when this indexer service get vtxo tree o k response has a 2xx status code
func (o *IndexerServiceGetVtxoTreeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this indexer service get vtxo tree o k response has a 3xx status code
func (o *IndexerServiceGetVtxoTreeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indexer service get vtxo tree o k response has a 4xx status code
func (o *IndexerServiceGetVtxoTreeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this indexer service get vtxo tree o k response has a 5xx status code
func (o *IndexerServiceGetVtxoTreeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this indexer service get vtxo tree o k response a status code equal to that given
func (o *IndexerServiceGetVtxoTreeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the indexer service get vtxo tree o k response
func (o *IndexerServiceGetVtxoTreeOK) Code() int {
	return 200
}

func (o *IndexerServiceGetVtxoTreeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/batch/{batchOutpoint.txid}/{batchOutpoint.vout}/tree][%d] indexerServiceGetVtxoTreeOK %s", 200, payload)
}

func (o *IndexerServiceGetVtxoTreeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/batch/{batchOutpoint.txid}/{batchOutpoint.vout}/tree][%d] indexerServiceGetVtxoTreeOK %s", 200, payload)
}

func (o *IndexerServiceGetVtxoTreeOK) GetPayload() *models.V1GetVtxoTreeResponse {
	return o.Payload
}

func (o *IndexerServiceGetVtxoTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetVtxoTreeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIndexerServiceGetVtxoTreeDefault creates a IndexerServiceGetVtxoTreeDefault with default headers values
func NewIndexerServiceGetVtxoTreeDefault(code int) *IndexerServiceGetVtxoTreeDefault {
	return &IndexerServiceGetVtxoTreeDefault{
		_statusCode: code,
	}
}

/*
IndexerServiceGetVtxoTreeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IndexerServiceGetVtxoTreeDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this indexer service get vtxo tree default response has a 2xx status code
func (o *IndexerServiceGetVtxoTreeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this indexer service get vtxo tree default response has a 3xx status code
func (o *IndexerServiceGetVtxoTreeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this indexer service get vtxo tree default response has a 4xx status code
func (o *IndexerServiceGetVtxoTreeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this indexer service get vtxo tree default response has a 5xx status code
func (o *IndexerServiceGetVtxoTreeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this indexer service get vtxo tree default response a status code equal to that given
func (o *IndexerServiceGetVtxoTreeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the indexer service get vtxo tree default response
func (o *IndexerServiceGetVtxoTreeDefault) Code() int {
	return o._statusCode
}

func (o *IndexerServiceGetVtxoTreeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/batch/{batchOutpoint.txid}/{batchOutpoint.vout}/tree][%d] IndexerService_GetVtxoTree default %s", o._statusCode, payload)
}

func (o *IndexerServiceGetVtxoTreeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/batch/{batchOutpoint.txid}/{batchOutpoint.vout}/tree][%d] IndexerService_GetVtxoTree default %s", o._statusCode, payload)
}

func (o *IndexerServiceGetVtxoTreeDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *IndexerServiceGetVtxoTreeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
