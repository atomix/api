# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/election/election.proto](#atomix/election/election.proto)
    - [AnointRequest](#atomix.election.AnointRequest)
    - [AnointResponse](#atomix.election.AnointResponse)
    - [CloseRequest](#atomix.election.CloseRequest)
    - [CloseResponse](#atomix.election.CloseResponse)
    - [CreateRequest](#atomix.election.CreateRequest)
    - [CreateResponse](#atomix.election.CreateResponse)
    - [EnterRequest](#atomix.election.EnterRequest)
    - [EnterResponse](#atomix.election.EnterResponse)
    - [EventRequest](#atomix.election.EventRequest)
    - [EventResponse](#atomix.election.EventResponse)
    - [EvictRequest](#atomix.election.EvictRequest)
    - [EvictResponse](#atomix.election.EvictResponse)
    - [GetTermRequest](#atomix.election.GetTermRequest)
    - [GetTermResponse](#atomix.election.GetTermResponse)
    - [KeepAliveRequest](#atomix.election.KeepAliveRequest)
    - [KeepAliveResponse](#atomix.election.KeepAliveResponse)
    - [PromoteRequest](#atomix.election.PromoteRequest)
    - [PromoteResponse](#atomix.election.PromoteResponse)
    - [Term](#atomix.election.Term)
    - [WithdrawRequest](#atomix.election.WithdrawRequest)
    - [WithdrawResponse](#atomix.election.WithdrawResponse)
  
    - [EventResponse.Type](#atomix.election.EventResponse.Type)
  
  
    - [LeaderElectionService](#atomix.election.LeaderElectionService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/election/election.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/election/election.proto



<a name="atomix.election.AnointRequest"></a>

### AnointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| candidate_id | [string](#string) |  |  |






<a name="atomix.election.AnointResponse"></a>

### AnointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| term | [Term](#atomix.election.Term) |  |  |






<a name="atomix.election.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.election.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.election.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.election.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.election.EnterRequest"></a>

### EnterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| candidate_id | [string](#string) |  |  |






<a name="atomix.election.EnterResponse"></a>

### EnterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| term | [Term](#atomix.election.Term) |  |  |






<a name="atomix.election.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.election.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.election.EventResponse.Type) |  |  |
| term | [Term](#atomix.election.Term) |  |  |






<a name="atomix.election.EvictRequest"></a>

### EvictRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| candidate_id | [string](#string) |  |  |






<a name="atomix.election.EvictResponse"></a>

### EvictResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| term | [Term](#atomix.election.Term) |  |  |






<a name="atomix.election.GetTermRequest"></a>

### GetTermRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.election.GetTermResponse"></a>

### GetTermResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| term | [Term](#atomix.election.Term) |  |  |






<a name="atomix.election.KeepAliveRequest"></a>

### KeepAliveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.election.KeepAliveResponse"></a>

### KeepAliveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.election.PromoteRequest"></a>

### PromoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| candidate_id | [string](#string) |  |  |






<a name="atomix.election.PromoteResponse"></a>

### PromoteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| term | [Term](#atomix.election.Term) |  |  |






<a name="atomix.election.Term"></a>

### Term



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| leader | [string](#string) |  |  |
| candidates | [string](#string) | repeated |  |






<a name="atomix.election.WithdrawRequest"></a>

### WithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| candidate_id | [string](#string) |  |  |






<a name="atomix.election.WithdrawResponse"></a>

### WithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| term | [Term](#atomix.election.Term) |  |  |





 


<a name="atomix.election.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANGED | 0 |  |


 

 


<a name="atomix.election.LeaderElectionService"></a>

### LeaderElectionService
Leader election service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.election.CreateRequest) | [CreateResponse](#atomix.election.CreateResponse) |  |
| KeepAlive | [KeepAliveRequest](#atomix.election.KeepAliveRequest) | [KeepAliveResponse](#atomix.election.KeepAliveResponse) |  |
| Close | [CloseRequest](#atomix.election.CloseRequest) | [CloseResponse](#atomix.election.CloseResponse) |  |
| Enter | [EnterRequest](#atomix.election.EnterRequest) | [EnterResponse](#atomix.election.EnterResponse) |  |
| Withdraw | [WithdrawRequest](#atomix.election.WithdrawRequest) | [WithdrawResponse](#atomix.election.WithdrawResponse) |  |
| Anoint | [AnointRequest](#atomix.election.AnointRequest) | [AnointResponse](#atomix.election.AnointResponse) |  |
| Promote | [PromoteRequest](#atomix.election.PromoteRequest) | [PromoteResponse](#atomix.election.PromoteResponse) |  |
| Evict | [EvictRequest](#atomix.election.EvictRequest) | [EvictResponse](#atomix.election.EvictResponse) |  |
| GetTerm | [GetTermRequest](#atomix.election.GetTermRequest) | [GetTermResponse](#atomix.election.GetTermResponse) |  |
| Events | [EventRequest](#atomix.election.EventRequest) | [EventResponse](#atomix.election.EventResponse) stream |  |

 



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

