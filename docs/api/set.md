# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/set/set.proto](#atomix/set/set.proto)
    - [AddRequest](#atomix.set.AddRequest)
    - [AddResponse](#atomix.set.AddResponse)
    - [ClearRequest](#atomix.set.ClearRequest)
    - [ClearResponse](#atomix.set.ClearResponse)
    - [CloseRequest](#atomix.set.CloseRequest)
    - [CloseResponse](#atomix.set.CloseResponse)
    - [ContainsRequest](#atomix.set.ContainsRequest)
    - [ContainsResponse](#atomix.set.ContainsResponse)
    - [CreateRequest](#atomix.set.CreateRequest)
    - [CreateResponse](#atomix.set.CreateResponse)
    - [EventRequest](#atomix.set.EventRequest)
    - [EventResponse](#atomix.set.EventResponse)
    - [IterateRequest](#atomix.set.IterateRequest)
    - [IterateResponse](#atomix.set.IterateResponse)
    - [RemoveRequest](#atomix.set.RemoveRequest)
    - [RemoveResponse](#atomix.set.RemoveResponse)
    - [SizeRequest](#atomix.set.SizeRequest)
    - [SizeResponse](#atomix.set.SizeResponse)
  
    - [EventResponse.Type](#atomix.set.EventResponse.Type)
    - [ResponseStatus](#atomix.set.ResponseStatus)
  
  
    - [SetService](#atomix.set.SetService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/set/set.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/set/set.proto



<a name="atomix.set.AddRequest"></a>

### AddRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.set.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.set.ResponseStatus) |  |  |
| added | [bool](#bool) |  |  |






<a name="atomix.set.ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.set.ClearResponse"></a>

### ClearResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.set.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.set.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.set.ContainsRequest"></a>

### ContainsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.set.ContainsResponse"></a>

### ContainsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| contains | [bool](#bool) |  |  |






<a name="atomix.set.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.set.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.set.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| replay | [bool](#bool) |  |  |






<a name="atomix.set.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.set.EventResponse.Type) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.set.IterateRequest"></a>

### IterateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.set.IterateResponse"></a>

### IterateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.set.RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.set.RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.set.ResponseStatus) |  |  |
| removed | [bool](#bool) |  |  |






<a name="atomix.set.SizeRequest"></a>

### SizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.set.SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| size | [int32](#int32) |  |  |





 


<a name="atomix.set.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| REMOVED | 2 |  |



<a name="atomix.set.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| NOOP | 1 |  |
| WRITE_LOCK | 2 |  |


 

 


<a name="atomix.set.SetService"></a>

### SetService
Set service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.set.CreateRequest) | [CreateResponse](#atomix.set.CreateResponse) | Create creates a set session |
| Close | [CloseRequest](#atomix.set.CloseRequest) | [CloseResponse](#atomix.set.CloseResponse) | Close closes a set |
| Size | [SizeRequest](#atomix.set.SizeRequest) | [SizeResponse](#atomix.set.SizeResponse) | Size gets the number of elements in the set |
| Contains | [ContainsRequest](#atomix.set.ContainsRequest) | [ContainsResponse](#atomix.set.ContainsResponse) | Contains returns whether the set contains a value |
| Add | [AddRequest](#atomix.set.AddRequest) | [AddResponse](#atomix.set.AddResponse) | Add adds a value to the set |
| Remove | [RemoveRequest](#atomix.set.RemoveRequest) | [RemoveResponse](#atomix.set.RemoveResponse) | Remove removes a value from the set |
| Clear | [ClearRequest](#atomix.set.ClearRequest) | [ClearResponse](#atomix.set.ClearResponse) | Clear removes all values from the set |
| Events | [EventRequest](#atomix.set.EventRequest) | [EventResponse](#atomix.set.EventResponse) stream | Events listens for set change events |
| Iterate | [IterateRequest](#atomix.set.IterateRequest) | [IterateResponse](#atomix.set.IterateResponse) stream | Iterate iterates through all values in the set |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

