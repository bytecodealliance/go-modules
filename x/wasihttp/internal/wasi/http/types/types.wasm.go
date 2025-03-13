// Code generated by wit-bindgen-go. DO NOT EDIT.

package types

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:http@0.2.0".

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]fields
//go:noescape
func wasmimport_FieldsResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [constructor]fields
//go:noescape
func wasmimport_NewFields() (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [static]fields.from-list
//go:noescape
func wasmimport_FieldsFromList(entries0 *cm.Tuple[FieldKey, FieldValue], entries1 uint32, result *cm.Result[Fields, Fields, HeaderError])

//go:wasmimport wasi:http/types@0.2.0 [method]fields.append
//go:noescape
func wasmimport_FieldsAppend(self0 uint32, name0 *uint8, name1 uint32, value0 *uint8, value1 uint32, result *cm.Result[HeaderError, struct{}, HeaderError])

//go:wasmimport wasi:http/types@0.2.0 [method]fields.clone
//go:noescape
func wasmimport_FieldsClone(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]fields.delete
//go:noescape
func wasmimport_FieldsDelete(self0 uint32, name0 *uint8, name1 uint32, result *cm.Result[HeaderError, struct{}, HeaderError])

//go:wasmimport wasi:http/types@0.2.0 [method]fields.entries
//go:noescape
func wasmimport_FieldsEntries(self0 uint32, result *cm.List[cm.Tuple[FieldKey, FieldValue]])

//go:wasmimport wasi:http/types@0.2.0 [method]fields.get
//go:noescape
func wasmimport_FieldsGet(self0 uint32, name0 *uint8, name1 uint32, result *cm.List[FieldValue])

//go:wasmimport wasi:http/types@0.2.0 [method]fields.has
//go:noescape
func wasmimport_FieldsHas(self0 uint32, name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]fields.set
//go:noescape
func wasmimport_FieldsSet(self0 uint32, name0 *uint8, name1 uint32, value0 *FieldValue, value1 uint32, result *cm.Result[HeaderError, struct{}, HeaderError])

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]incoming-request
//go:noescape
func wasmimport_IncomingRequestResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-request.authority
//go:noescape
func wasmimport_IncomingRequestAuthority(self0 uint32, result *cm.Option[string])

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-request.consume
//go:noescape
func wasmimport_IncomingRequestConsume(self0 uint32, result *cm.Result[IncomingBody, IncomingBody, struct{}])

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-request.headers
//go:noescape
func wasmimport_IncomingRequestHeaders(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-request.method
//go:noescape
func wasmimport_IncomingRequestMethod(self0 uint32, result *Method)

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-request.path-with-query
//go:noescape
func wasmimport_IncomingRequestPathWithQuery(self0 uint32, result *cm.Option[string])

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-request.scheme
//go:noescape
func wasmimport_IncomingRequestScheme(self0 uint32, result *cm.Option[Scheme])

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]outgoing-request
//go:noescape
func wasmimport_OutgoingRequestResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [constructor]outgoing-request
//go:noescape
func wasmimport_NewOutgoingRequest(headers0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.authority
//go:noescape
func wasmimport_OutgoingRequestAuthority(self0 uint32, result *cm.Option[string])

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.body
//go:noescape
func wasmimport_OutgoingRequestBody(self0 uint32, result *cm.Result[OutgoingBody, OutgoingBody, struct{}])

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.headers
//go:noescape
func wasmimport_OutgoingRequestHeaders(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.method
//go:noescape
func wasmimport_OutgoingRequestMethod(self0 uint32, result *Method)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.path-with-query
//go:noescape
func wasmimport_OutgoingRequestPathWithQuery(self0 uint32, result *cm.Option[string])

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.scheme
//go:noescape
func wasmimport_OutgoingRequestScheme(self0 uint32, result *cm.Option[Scheme])

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.set-authority
//go:noescape
func wasmimport_OutgoingRequestSetAuthority(self0 uint32, authority0 uint32, authority1 *uint8, authority2 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.set-method
//go:noescape
func wasmimport_OutgoingRequestSetMethod(self0 uint32, method0 uint32, method1 *uint8, method2 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.set-path-with-query
//go:noescape
func wasmimport_OutgoingRequestSetPathWithQuery(self0 uint32, pathWithQuery0 uint32, pathWithQuery1 *uint8, pathWithQuery2 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-request.set-scheme
//go:noescape
func wasmimport_OutgoingRequestSetScheme(self0 uint32, scheme0 uint32, scheme1 uint32, scheme2 *uint8, scheme3 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]request-options
//go:noescape
func wasmimport_RequestOptionsResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [constructor]request-options
//go:noescape
func wasmimport_NewRequestOptions() (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]request-options.between-bytes-timeout
//go:noescape
func wasmimport_RequestOptionsBetweenBytesTimeout(self0 uint32, result *cm.Option[Duration])

//go:wasmimport wasi:http/types@0.2.0 [method]request-options.connect-timeout
//go:noescape
func wasmimport_RequestOptionsConnectTimeout(self0 uint32, result *cm.Option[Duration])

//go:wasmimport wasi:http/types@0.2.0 [method]request-options.first-byte-timeout
//go:noescape
func wasmimport_RequestOptionsFirstByteTimeout(self0 uint32, result *cm.Option[Duration])

//go:wasmimport wasi:http/types@0.2.0 [method]request-options.set-between-bytes-timeout
//go:noescape
func wasmimport_RequestOptionsSetBetweenBytesTimeout(self0 uint32, duration0 uint32, duration1 uint64) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]request-options.set-connect-timeout
//go:noescape
func wasmimport_RequestOptionsSetConnectTimeout(self0 uint32, duration0 uint32, duration1 uint64) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]request-options.set-first-byte-timeout
//go:noescape
func wasmimport_RequestOptionsSetFirstByteTimeout(self0 uint32, duration0 uint32, duration1 uint64) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]response-outparam
//go:noescape
func wasmimport_ResponseOutparamResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [static]response-outparam.set
//go:noescape
func wasmimport_ResponseOutparamSet(param0 uint32, response0 uint32, response1 uint32, response2 uint32, response3 uint64, response4 uint32, response5 uint32, response6 uint32, response7 uint32)

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]incoming-response
//go:noescape
func wasmimport_IncomingResponseResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-response.consume
//go:noescape
func wasmimport_IncomingResponseConsume(self0 uint32, result *cm.Result[IncomingBody, IncomingBody, struct{}])

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-response.headers
//go:noescape
func wasmimport_IncomingResponseHeaders(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-response.status
//go:noescape
func wasmimport_IncomingResponseStatus(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]incoming-body
//go:noescape
func wasmimport_IncomingBodyResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [static]incoming-body.finish
//go:noescape
func wasmimport_IncomingBodyFinish(this0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]incoming-body.stream
//go:noescape
func wasmimport_IncomingBodyStream(self0 uint32, result *cm.Result[InputStream, InputStream, struct{}])

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]future-trailers
//go:noescape
func wasmimport_FutureTrailersResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]future-trailers.get
//go:noescape
func wasmimport_FutureTrailersGet(self0 uint32, result *cm.Option[cm.Result[cm.Result[ErrorCodeShape, cm.Option[Trailers], ErrorCode], cm.Result[ErrorCodeShape, cm.Option[Trailers], ErrorCode], struct{}]])

//go:wasmimport wasi:http/types@0.2.0 [method]future-trailers.subscribe
//go:noescape
func wasmimport_FutureTrailersSubscribe(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]outgoing-response
//go:noescape
func wasmimport_OutgoingResponseResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [constructor]outgoing-response
//go:noescape
func wasmimport_NewOutgoingResponse(headers0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-response.body
//go:noescape
func wasmimport_OutgoingResponseBody(self0 uint32, result *cm.Result[OutgoingBody, OutgoingBody, struct{}])

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-response.headers
//go:noescape
func wasmimport_OutgoingResponseHeaders(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-response.set-status-code
//go:noescape
func wasmimport_OutgoingResponseSetStatusCode(self0 uint32, statusCode0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-response.status-code
//go:noescape
func wasmimport_OutgoingResponseStatusCode(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]outgoing-body
//go:noescape
func wasmimport_OutgoingBodyResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [static]outgoing-body.finish
//go:noescape
func wasmimport_OutgoingBodyFinish(this0 uint32, trailers0 uint32, trailers1 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

//go:wasmimport wasi:http/types@0.2.0 [method]outgoing-body.write
//go:noescape
func wasmimport_OutgoingBodyWrite(self0 uint32, result *cm.Result[OutputStream, OutputStream, struct{}])

//go:wasmimport wasi:http/types@0.2.0 [resource-drop]future-incoming-response
//go:noescape
func wasmimport_FutureIncomingResponseResourceDrop(self0 uint32)

//go:wasmimport wasi:http/types@0.2.0 [method]future-incoming-response.get
//go:noescape
func wasmimport_FutureIncomingResponseGet(self0 uint32, result *cm.Option[cm.Result[cm.Result[ErrorCodeShape, IncomingResponse, ErrorCode], cm.Result[ErrorCodeShape, IncomingResponse, ErrorCode], struct{}]])

//go:wasmimport wasi:http/types@0.2.0 [method]future-incoming-response.subscribe
//go:noescape
func wasmimport_FutureIncomingResponseSubscribe(self0 uint32) (result0 uint32)

//go:wasmimport wasi:http/types@0.2.0 http-error-code
//go:noescape
func wasmimport_HTTPErrorCode(err0 uint32, result *cm.Option[ErrorCode])
