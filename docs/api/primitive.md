# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/primitive/primitive.proto](#atomix/primitive/primitive.proto)
    - [GetPrimitivesRequest](#atomix.primitive.GetPrimitivesRequest)
    - [GetPrimitivesResponse](#atomix.primitive.GetPrimitivesResponse)
    - [Name](#atomix.primitive.Name)
    - [PrimitiveInfo](#atomix.primitive.PrimitiveInfo)
  
  
  
    - [PrimitiveService](#atomix.primitive.PrimitiveService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/primitive/primitive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/primitive/primitive.proto



<a name="atomix.primitive.GetPrimitivesRequest"></a>

### GetPrimitivesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  |  |
| type | [string](#string) |  |  |






<a name="atomix.primitive.GetPrimitivesResponse"></a>

### GetPrimitivesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primitives | [PrimitiveInfo](#atomix.primitive.PrimitiveInfo) | repeated |  |






<a name="atomix.primitive.Name"></a>

### Name
Namespaced primitive name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |






<a name="atomix.primitive.PrimitiveInfo"></a>

### PrimitiveInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| name | [Name](#atomix.primitive.Name) |  |  |





 

 

 


<a name="atomix.primitive.PrimitiveService"></a>

### PrimitiveService
Primitive service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPrimitives | [GetPrimitivesRequest](#atomix.primitive.GetPrimitivesRequest) | [GetPrimitivesResponse](#atomix.primitive.GetPrimitivesResponse) | GetPrimitives returns a list of primitives |

 



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

