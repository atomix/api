# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/headers/headers.proto](#atomix/headers/headers.proto)
    - [RequestHeader](#atomix.headers.RequestHeader)
    - [ResponseHeader](#atomix.headers.ResponseHeader)
    - [StreamHeader](#atomix.headers.StreamHeader)
  
    - [ResponseStatus](#atomix.headers.ResponseStatus)
    - [ResponseType](#atomix.headers.ResponseType)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/headers/headers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/headers/headers.proto



<a name="atomix.headers.RequestHeader"></a>

### RequestHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [atomix.primitive.Name](#atomix.primitive.Name) |  |  |
| session_id | [uint64](#uint64) |  |  |
| request_id | [uint64](#uint64) |  |  |
| index | [uint64](#uint64) |  |  |
| streams | [StreamHeader](#atomix.headers.StreamHeader) | repeated |  |






<a name="atomix.headers.ResponseHeader"></a>

### ResponseHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [uint64](#uint64) |  |  |
| stream_id | [uint64](#uint64) |  |  |
| response_id | [uint64](#uint64) |  |  |
| index | [uint64](#uint64) |  |  |
| leader | [string](#string) |  |  |
| status | [ResponseStatus](#atomix.headers.ResponseStatus) |  |  |
| type | [ResponseType](#atomix.headers.ResponseType) |  |  |






<a name="atomix.headers.StreamHeader"></a>

### StreamHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stream_id | [uint64](#uint64) |  |  |
| response_id | [uint64](#uint64) |  |  |





 


<a name="atomix.headers.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| ERROR | 1 |  |
| NOT_LEADER | 2 |  |



<a name="atomix.headers.ResponseType"></a>

### ResponseType


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESPONSE | 0 |  |
| OPEN_STREAM | 1 |  |
| CLOSE_STREAM | 2 |  |


 

 

 



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

