# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/list/list.proto](#atomix/list/list.proto)
    - [AppendRequest](#atomix.list.AppendRequest)
    - [AppendResponse](#atomix.list.AppendResponse)
    - [ClearRequest](#atomix.list.ClearRequest)
    - [ClearResponse](#atomix.list.ClearResponse)
    - [CloseRequest](#atomix.list.CloseRequest)
    - [CloseResponse](#atomix.list.CloseResponse)
    - [ContainsRequest](#atomix.list.ContainsRequest)
    - [ContainsResponse](#atomix.list.ContainsResponse)
    - [CreateRequest](#atomix.list.CreateRequest)
    - [CreateResponse](#atomix.list.CreateResponse)
    - [EventRequest](#atomix.list.EventRequest)
    - [EventResponse](#atomix.list.EventResponse)
    - [GetRequest](#atomix.list.GetRequest)
    - [GetResponse](#atomix.list.GetResponse)
    - [InsertRequest](#atomix.list.InsertRequest)
    - [InsertResponse](#atomix.list.InsertResponse)
    - [IterateRequest](#atomix.list.IterateRequest)
    - [IterateResponse](#atomix.list.IterateResponse)
    - [KeepAliveRequest](#atomix.list.KeepAliveRequest)
    - [KeepAliveResponse](#atomix.list.KeepAliveResponse)
    - [RemoveRequest](#atomix.list.RemoveRequest)
    - [RemoveResponse](#atomix.list.RemoveResponse)
    - [SetRequest](#atomix.list.SetRequest)
    - [SetResponse](#atomix.list.SetResponse)
    - [SizeRequest](#atomix.list.SizeRequest)
    - [SizeResponse](#atomix.list.SizeResponse)
  
    - [EventResponse.Type](#atomix.list.EventResponse.Type)
    - [ResponseStatus](#atomix.list.ResponseStatus)
  
  
    - [ListService](#atomix.list.ListService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/list/list.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/list/list.proto



<a name="atomix.list.AppendRequest"></a>

### AppendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.AppendResponse"></a>

### AppendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.list.ResponseStatus) |  |  |






<a name="atomix.list.ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.list.ClearResponse"></a>

### ClearResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.list.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.list.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.list.ContainsRequest"></a>

### ContainsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.ContainsResponse"></a>

### ContainsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| contains | [bool](#bool) |  |  |






<a name="atomix.list.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.list.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.list.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| replay | [bool](#bool) |  |  |






<a name="atomix.list.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.list.EventResponse.Type) |  |  |
| index | [uint32](#uint32) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [uint32](#uint32) |  |  |






<a name="atomix.list.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.list.ResponseStatus) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.InsertRequest"></a>

### InsertRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [uint32](#uint32) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.InsertResponse"></a>

### InsertResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.list.ResponseStatus) |  |  |






<a name="atomix.list.IterateRequest"></a>

### IterateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.list.IterateResponse"></a>

### IterateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.KeepAliveRequest"></a>

### KeepAliveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.list.KeepAliveResponse"></a>

### KeepAliveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.list.RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [uint32](#uint32) |  |  |






<a name="atomix.list.RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.list.ResponseStatus) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.SetRequest"></a>

### SetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [uint32](#uint32) |  |  |
| value | [string](#string) |  |  |






<a name="atomix.list.SetResponse"></a>

### SetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.list.ResponseStatus) |  |  |






<a name="atomix.list.SizeRequest"></a>

### SizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.list.SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| size | [int32](#int32) |  |  |





 


<a name="atomix.list.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| ADDED | 1 |  |
| REMOVED | 2 |  |



<a name="atomix.list.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| NOOP | 1 |  |
| WRITE_LOCK | 2 |  |
| OUT_OF_BOUNDS | 3 |  |


 

 


<a name="atomix.list.ListService"></a>

### ListService
List service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.list.CreateRequest) | [CreateResponse](#atomix.list.CreateResponse) |  |
| KeepAlive | [KeepAliveRequest](#atomix.list.KeepAliveRequest) | [KeepAliveResponse](#atomix.list.KeepAliveResponse) |  |
| Close | [CloseRequest](#atomix.list.CloseRequest) | [CloseResponse](#atomix.list.CloseResponse) |  |
| Size | [SizeRequest](#atomix.list.SizeRequest) | [SizeResponse](#atomix.list.SizeResponse) |  |
| Contains | [ContainsRequest](#atomix.list.ContainsRequest) | [ContainsResponse](#atomix.list.ContainsResponse) |  |
| Append | [AppendRequest](#atomix.list.AppendRequest) | [AppendResponse](#atomix.list.AppendResponse) |  |
| Insert | [InsertRequest](#atomix.list.InsertRequest) | [InsertResponse](#atomix.list.InsertResponse) |  |
| Get | [GetRequest](#atomix.list.GetRequest) | [GetResponse](#atomix.list.GetResponse) |  |
| Set | [SetRequest](#atomix.list.SetRequest) | [SetResponse](#atomix.list.SetResponse) |  |
| Remove | [RemoveRequest](#atomix.list.RemoveRequest) | [RemoveResponse](#atomix.list.RemoveResponse) |  |
| Clear | [ClearRequest](#atomix.list.ClearRequest) | [ClearResponse](#atomix.list.ClearResponse) |  |
| Events | [EventRequest](#atomix.list.EventRequest) | [EventResponse](#atomix.list.EventResponse) stream |  |
| Iterate | [IterateRequest](#atomix.list.IterateRequest) | [IterateResponse](#atomix.list.IterateResponse) stream |  |

 



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

