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
LockService implements a distributed lock

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.lock.CreateRequest) | [CreateResponse](#atomix.lock.CreateResponse) | Create creates a lock |
| Close | [CloseRequest](#atomix.lock.CloseRequest) | [CloseResponse](#atomix.lock.CloseResponse) | Close closes a lock |
| Lock | [LockRequest](#atomix.lock.LockRequest) | [LockResponse](#atomix.lock.LockResponse) | Lock attempts to acquire the lock |
| Unlock | [UnlockRequest](#atomix.lock.UnlockRequest) | [UnlockResponse](#atomix.lock.UnlockResponse) | Unlock releases the lock |
| IsLocked | [IsLockedRequest](#atomix.lock.IsLockedRequest) | [IsLockedResponse](#atomix.lock.IsLockedResponse) | IsLocked checks whether the lock is held |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

