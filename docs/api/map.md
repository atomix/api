# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/map/map.proto](#atomix/map/map.proto)
    - [ClearRequest](#atomix.map.ClearRequest)
    - [ClearResponse](#atomix.map.ClearResponse)
    - [CloseRequest](#atomix.map.CloseRequest)
    - [CloseResponse](#atomix.map.CloseResponse)
    - [CreateRequest](#atomix.map.CreateRequest)
    - [CreateResponse](#atomix.map.CreateResponse)
    - [EntriesRequest](#atomix.map.EntriesRequest)
    - [EntriesResponse](#atomix.map.EntriesResponse)
    - [EventRequest](#atomix.map.EventRequest)
    - [EventResponse](#atomix.map.EventResponse)
    - [ExistsRequest](#atomix.map.ExistsRequest)
    - [ExistsResponse](#atomix.map.ExistsResponse)
    - [GetRequest](#atomix.map.GetRequest)
    - [GetResponse](#atomix.map.GetResponse)
    - [KeepAliveRequest](#atomix.map.KeepAliveRequest)
    - [KeepAliveResponse](#atomix.map.KeepAliveResponse)
    - [PutRequest](#atomix.map.PutRequest)
    - [PutResponse](#atomix.map.PutResponse)
    - [RemoveRequest](#atomix.map.RemoveRequest)
    - [RemoveResponse](#atomix.map.RemoveResponse)
    - [ReplaceRequest](#atomix.map.ReplaceRequest)
    - [ReplaceResponse](#atomix.map.ReplaceResponse)
    - [SizeRequest](#atomix.map.SizeRequest)
    - [SizeResponse](#atomix.map.SizeResponse)
  
    - [EventResponse.Type](#atomix.map.EventResponse.Type)
    - [ResponseStatus](#atomix.map.ResponseStatus)
  
  
    - [MapService](#atomix.map.MapService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/map/map.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/map/map.proto



<a name="atomix.map.ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.map.ClearResponse"></a>

### ClearResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.map.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.map.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.map.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.map.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.map.EntriesRequest"></a>

### EntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.map.EntriesResponse"></a>

### EntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.map.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| replay | [bool](#bool) |  |  |
| key | [string](#string) |  |  |






<a name="atomix.map.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.map.EventResponse.Type) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.map.ExistsRequest"></a>

### ExistsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| key | [string](#string) |  |  |






<a name="atomix.map.ExistsResponse"></a>

### ExistsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| contains_key | [bool](#bool) |  |  |






<a name="atomix.map.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| key | [string](#string) |  |  |






<a name="atomix.map.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.map.KeepAliveRequest"></a>

### KeepAliveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.map.KeepAliveResponse"></a>

### KeepAliveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.map.PutRequest"></a>

### PutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| ttl | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.map.PutResponse"></a>

### PutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.map.ResponseStatus) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |






<a name="atomix.map.RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.map.RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.map.ResponseStatus) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |






<a name="atomix.map.ReplaceRequest"></a>

### ReplaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| key | [string](#string) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |
| new_value | [bytes](#bytes) |  |  |
| ttl | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.map.ReplaceResponse"></a>

### ReplaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.map.ResponseStatus) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |






<a name="atomix.map.SizeRequest"></a>

### SizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.map.SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| size | [int32](#int32) |  |  |





 


<a name="atomix.map.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| INSERTED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="atomix.map.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| NOOP | 1 |  |
| WRITE_LOCK | 2 |  |
| PRECONDITION_FAILED | 3 |  |


 

 


<a name="atomix.map.MapService"></a>

### MapService
Map service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.map.CreateRequest) | [CreateResponse](#atomix.map.CreateResponse) |  |
| KeepAlive | [KeepAliveRequest](#atomix.map.KeepAliveRequest) | [KeepAliveResponse](#atomix.map.KeepAliveResponse) |  |
| Close | [CloseRequest](#atomix.map.CloseRequest) | [CloseResponse](#atomix.map.CloseResponse) |  |
| Size | [SizeRequest](#atomix.map.SizeRequest) | [SizeResponse](#atomix.map.SizeResponse) |  |
| Exists | [ExistsRequest](#atomix.map.ExistsRequest) | [ExistsResponse](#atomix.map.ExistsResponse) |  |
| Put | [PutRequest](#atomix.map.PutRequest) | [PutResponse](#atomix.map.PutResponse) |  |
| Replace | [ReplaceRequest](#atomix.map.ReplaceRequest) | [ReplaceResponse](#atomix.map.ReplaceResponse) |  |
| Get | [GetRequest](#atomix.map.GetRequest) | [GetResponse](#atomix.map.GetResponse) |  |
| Remove | [RemoveRequest](#atomix.map.RemoveRequest) | [RemoveResponse](#atomix.map.RemoveResponse) |  |
| Clear | [ClearRequest](#atomix.map.ClearRequest) | [ClearResponse](#atomix.map.ClearResponse) |  |
| Events | [EventRequest](#atomix.map.EventRequest) | [EventResponse](#atomix.map.EventResponse) stream |  |
| Entries | [EntriesRequest](#atomix.map.EntriesRequest) | [EntriesResponse](#atomix.map.EntriesResponse) stream |  |

 



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

