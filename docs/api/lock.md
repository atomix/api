# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/lock/lock.proto](#atomix/lock/lock.proto)
    - [CloseRequest](#atomix.lock.CloseRequest)
    - [CloseResponse](#atomix.lock.CloseResponse)
    - [CreateRequest](#atomix.lock.CreateRequest)
    - [CreateResponse](#atomix.lock.CreateResponse)
    - [IsLockedRequest](#atomix.lock.IsLockedRequest)
    - [IsLockedResponse](#atomix.lock.IsLockedResponse)
    - [KeepAliveRequest](#atomix.lock.KeepAliveRequest)
    - [KeepAliveResponse](#atomix.lock.KeepAliveResponse)
    - [LockRequest](#atomix.lock.LockRequest)
    - [LockResponse](#atomix.lock.LockResponse)
    - [UnlockRequest](#atomix.lock.UnlockRequest)
    - [UnlockResponse](#atomix.lock.UnlockResponse)
  
  
  
    - [LockService](#atomix.lock.LockService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/lock/lock.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/lock/lock.proto



<a name="atomix.lock.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.lock.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.lock.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.lock.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.lock.IsLockedRequest"></a>

### IsLockedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| version | [uint64](#uint64) |  |  |






<a name="atomix.lock.IsLockedResponse"></a>

### IsLockedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| isLocked | [bool](#bool) |  |  |






<a name="atomix.lock.KeepAliveRequest"></a>

### KeepAliveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.lock.KeepAliveResponse"></a>

### KeepAliveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.lock.LockRequest"></a>

### LockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.lock.LockResponse"></a>

### LockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| version | [uint64](#uint64) |  |  |






<a name="atomix.lock.UnlockRequest"></a>

### UnlockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| version | [uint64](#uint64) |  |  |






<a name="atomix.lock.UnlockResponse"></a>

### UnlockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| unlocked | [bool](#bool) |  |  |





 

 

 


<a name="atomix.lock.LockService"></a>

### LockService
Lock service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.lock.CreateRequest) | [CreateResponse](#atomix.lock.CreateResponse) |  |
| KeepAlive | [KeepAliveRequest](#atomix.lock.KeepAliveRequest) | [KeepAliveResponse](#atomix.lock.KeepAliveResponse) |  |
| Close | [CloseRequest](#atomix.lock.CloseRequest) | [CloseResponse](#atomix.lock.CloseResponse) |  |
| Lock | [LockRequest](#atomix.lock.LockRequest) | [LockResponse](#atomix.lock.LockResponse) |  |
| Unlock | [UnlockRequest](#atomix.lock.UnlockRequest) | [UnlockResponse](#atomix.lock.UnlockResponse) |  |
| IsLocked | [IsLockedRequest](#atomix.lock.IsLockedRequest) | [IsLockedResponse](#atomix.lock.IsLockedResponse) |  |

 



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

