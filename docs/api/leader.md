# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/leader/latch.proto](#atomix/leader/latch.proto)
    - [CloseRequest](#atomix.leader.CloseRequest)
    - [CloseResponse](#atomix.leader.CloseResponse)
    - [CreateRequest](#atomix.leader.CreateRequest)
    - [CreateResponse](#atomix.leader.CreateResponse)
    - [EventRequest](#atomix.leader.EventRequest)
    - [EventResponse](#atomix.leader.EventResponse)
    - [GetRequest](#atomix.leader.GetRequest)
    - [GetResponse](#atomix.leader.GetResponse)
    - [KeepAliveRequest](#atomix.leader.KeepAliveRequest)
    - [KeepAliveResponse](#atomix.leader.KeepAliveResponse)
    - [Latch](#atomix.leader.Latch)
    - [LatchRequest](#atomix.leader.LatchRequest)
    - [LatchResponse](#atomix.leader.LatchResponse)
  
    - [EventResponse.Type](#atomix.leader.EventResponse.Type)
  
  
    - [LeaderLatchService](#atomix.leader.LeaderLatchService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/leader/latch.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/leader/latch.proto



<a name="atomix.leader.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.leader.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.leader.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.leader.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.leader.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.leader.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.leader.EventResponse.Type) |  |  |
| latch | [Latch](#atomix.leader.Latch) |  |  |






<a name="atomix.leader.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.leader.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| latch | [Latch](#atomix.leader.Latch) |  |  |






<a name="atomix.leader.KeepAliveRequest"></a>

### KeepAliveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.leader.KeepAliveResponse"></a>

### KeepAliveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.leader.Latch"></a>

### Latch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| leader | [string](#string) |  |  |
| participants | [string](#string) | repeated |  |






<a name="atomix.leader.LatchRequest"></a>

### LatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| participant_id | [string](#string) |  |  |






<a name="atomix.leader.LatchResponse"></a>

### LatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| latch | [Latch](#atomix.leader.Latch) |  |  |





 


<a name="atomix.leader.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANGED | 0 |  |


 

 


<a name="atomix.leader.LeaderLatchService"></a>

### LeaderLatchService
Leader latch service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.leader.CreateRequest) | [CreateResponse](#atomix.leader.CreateResponse) |  |
| KeepAlive | [KeepAliveRequest](#atomix.leader.KeepAliveRequest) | [KeepAliveResponse](#atomix.leader.KeepAliveResponse) |  |
| Close | [CloseRequest](#atomix.leader.CloseRequest) | [CloseResponse](#atomix.leader.CloseResponse) |  |
| Latch | [LatchRequest](#atomix.leader.LatchRequest) | [LatchResponse](#atomix.leader.LatchResponse) |  |
| Get | [GetRequest](#atomix.leader.GetRequest) | [GetResponse](#atomix.leader.GetResponse) |  |
| Events | [EventRequest](#atomix.leader.EventRequest) | [EventResponse](#atomix.leader.EventResponse) stream |  |

 



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

