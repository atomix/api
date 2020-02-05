# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/counter/counter.proto](#atomix/counter/counter.proto)
    - [CheckAndSetRequest](#atomix.counter.CheckAndSetRequest)
    - [CheckAndSetResponse](#atomix.counter.CheckAndSetResponse)
    - [CloseRequest](#atomix.counter.CloseRequest)
    - [CloseResponse](#atomix.counter.CloseResponse)
    - [CreateRequest](#atomix.counter.CreateRequest)
    - [CreateResponse](#atomix.counter.CreateResponse)
    - [DecrementRequest](#atomix.counter.DecrementRequest)
    - [DecrementResponse](#atomix.counter.DecrementResponse)
    - [GetRequest](#atomix.counter.GetRequest)
    - [GetResponse](#atomix.counter.GetResponse)
    - [IncrementRequest](#atomix.counter.IncrementRequest)
    - [IncrementResponse](#atomix.counter.IncrementResponse)
    - [SetRequest](#atomix.counter.SetRequest)
    - [SetResponse](#atomix.counter.SetResponse)
  
  
  
    - [CounterService](#atomix.counter.CounterService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/counter/counter.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/counter/counter.proto



<a name="atomix.counter.CheckAndSetRequest"></a>

### CheckAndSetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| expect | [int64](#int64) |  |  |
| update | [int64](#int64) |  |  |






<a name="atomix.counter.CheckAndSetResponse"></a>

### CheckAndSetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| succeeded | [bool](#bool) |  |  |






<a name="atomix.counter.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.counter.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.counter.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.counter.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.counter.DecrementRequest"></a>

### DecrementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delta | [int64](#int64) |  |  |






<a name="atomix.counter.DecrementResponse"></a>

### DecrementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| previous_value | [int64](#int64) |  |  |
| next_value | [int64](#int64) |  |  |






<a name="atomix.counter.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.counter.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| value | [int64](#int64) |  |  |






<a name="atomix.counter.IncrementRequest"></a>

### IncrementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delta | [int64](#int64) |  |  |






<a name="atomix.counter.IncrementResponse"></a>

### IncrementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| previous_value | [int64](#int64) |  |  |
| next_value | [int64](#int64) |  |  |






<a name="atomix.counter.SetRequest"></a>

### SetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| value | [int64](#int64) |  |  |






<a name="atomix.counter.SetResponse"></a>

### SetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| previous_value | [int64](#int64) |  |  |





 

 

 


<a name="atomix.counter.CounterService"></a>

### CounterService
CounterService implements a distributed counter

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.counter.CreateRequest) | [CreateResponse](#atomix.counter.CreateResponse) | Create creates a counter |
| Close | [CloseRequest](#atomix.counter.CloseRequest) | [CloseResponse](#atomix.counter.CloseResponse) | Close closes the counter |
| Set | [SetRequest](#atomix.counter.SetRequest) | [SetResponse](#atomix.counter.SetResponse) | Set sets the counter value |
| Get | [GetRequest](#atomix.counter.GetRequest) | [GetResponse](#atomix.counter.GetResponse) | Get gets the current counter value |
| Increment | [IncrementRequest](#atomix.counter.IncrementRequest) | [IncrementResponse](#atomix.counter.IncrementResponse) | Increment increments the counter value |
| Decrement | [DecrementRequest](#atomix.counter.DecrementRequest) | [DecrementResponse](#atomix.counter.DecrementResponse) | Decrement decrements the counter value |
| CheckAndSet | [CheckAndSetRequest](#atomix.counter.CheckAndSetRequest) | [CheckAndSetResponse](#atomix.counter.CheckAndSetResponse) | CheckAndSet performs a check-and-set operation on the counter value |

 



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

