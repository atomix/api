# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/log/log.proto](#atomix/log/log.proto)
    - [AppendRequest](#atomix.log.AppendRequest)
    - [AppendResponse](#atomix.log.AppendResponse)
    - [ClearRequest](#atomix.log.ClearRequest)
    - [ClearResponse](#atomix.log.ClearResponse)
    - [CloseRequest](#atomix.log.CloseRequest)
    - [CloseResponse](#atomix.log.CloseResponse)
    - [CreateRequest](#atomix.log.CreateRequest)
    - [CreateResponse](#atomix.log.CreateResponse)
    - [EntriesRequest](#atomix.log.EntriesRequest)
    - [EntriesResponse](#atomix.log.EntriesResponse)
    - [EventRequest](#atomix.log.EventRequest)
    - [EventResponse](#atomix.log.EventResponse)
    - [ExistsRequest](#atomix.log.ExistsRequest)
    - [ExistsResponse](#atomix.log.ExistsResponse)
    - [FirstEntryRequest](#atomix.log.FirstEntryRequest)
    - [FirstEntryResponse](#atomix.log.FirstEntryResponse)
    - [GetRequest](#atomix.log.GetRequest)
    - [GetResponse](#atomix.log.GetResponse)
    - [LastEntryRequest](#atomix.log.LastEntryRequest)
    - [LastEntryResponse](#atomix.log.LastEntryResponse)
    - [NextEntryRequest](#atomix.log.NextEntryRequest)
    - [NextEntryResponse](#atomix.log.NextEntryResponse)
    - [PrevEntryRequest](#atomix.log.PrevEntryRequest)
    - [PrevEntryResponse](#atomix.log.PrevEntryResponse)
    - [RemoveRequest](#atomix.log.RemoveRequest)
    - [RemoveResponse](#atomix.log.RemoveResponse)
    - [SizeRequest](#atomix.log.SizeRequest)
    - [SizeResponse](#atomix.log.SizeResponse)
  
    - [EventResponse.Type](#atomix.log.EventResponse.Type)
    - [ResponseStatus](#atomix.log.ResponseStatus)
  
  
    - [LogService](#atomix.log.LogService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/log/log.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/log/log.proto



<a name="atomix.log.AppendRequest"></a>

### AppendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |






<a name="atomix.log.AppendResponse"></a>

### AppendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.log.ResponseStatus) |  |  |
| index | [int64](#int64) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.log.ClearResponse"></a>

### ClearResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.log.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.log.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.log.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.log.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.log.EntriesRequest"></a>

### EntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.log.EntriesResponse"></a>

### EntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| replay | [bool](#bool) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.log.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.log.EventResponse.Type) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.ExistsRequest"></a>

### ExistsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [uint64](#uint64) |  |  |






<a name="atomix.log.ExistsResponse"></a>

### ExistsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| contains_index | [bool](#bool) |  |  |






<a name="atomix.log.FirstEntryRequest"></a>

### FirstEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.log.FirstEntryResponse"></a>

### FirstEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.log.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.LastEntryRequest"></a>

### LastEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.log.LastEntryResponse"></a>

### LastEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.NextEntryRequest"></a>

### NextEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.log.NextEntryResponse"></a>

### NextEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.PrevEntryRequest"></a>

### PrevEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.log.PrevEntryResponse"></a>

### PrevEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.log.RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.log.ResponseStatus) |  |  |
| index | [int64](#int64) |  |  |
| previous_value | [bytes](#bytes) |  |  |






<a name="atomix.log.SizeRequest"></a>

### SizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.log.SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| size | [int32](#int32) |  |  |





 


<a name="atomix.log.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| APPENDED | 1 |  |
| REMOVED | 2 |  |



<a name="atomix.log.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| NOOP | 1 |  |
| WRITE_LOCK | 2 |  |
| PRECONDITION_FAILED | 3 |  |


 

 


<a name="atomix.log.LogService"></a>

### LogService
LogService log service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.log.CreateRequest) | [CreateResponse](#atomix.log.CreateResponse) | Create creates a log |
| Close | [CloseRequest](#atomix.log.CloseRequest) | [CloseResponse](#atomix.log.CloseResponse) | Close closes a log |
| Size | [SizeRequest](#atomix.log.SizeRequest) | [SizeResponse](#atomix.log.SizeResponse) | Size returns the size of the log |
| Exists | [ExistsRequest](#atomix.log.ExistsRequest) | [ExistsResponse](#atomix.log.ExistsResponse) | Exists checks whether an index exists in the log |
| Append | [AppendRequest](#atomix.log.AppendRequest) | [AppendResponse](#atomix.log.AppendResponse) | Appends appends an entry into the log |
| Get | [GetRequest](#atomix.log.GetRequest) | [GetResponse](#atomix.log.GetResponse) | Get gets the entry for an index |
| FirstEntry | [FirstEntryRequest](#atomix.log.FirstEntryRequest) | [FirstEntryResponse](#atomix.log.FirstEntryResponse) | FirstEntry gets the first entry in the log |
| LastEntry | [LastEntryRequest](#atomix.log.LastEntryRequest) | [LastEntryResponse](#atomix.log.LastEntryResponse) | LastEntry gets the last entry in the log |
| PrevEntry | [PrevEntryRequest](#atomix.log.PrevEntryRequest) | [PrevEntryResponse](#atomix.log.PrevEntryResponse) | PrevEntry gets the previous entry in the log |
| NextEntry | [NextEntryRequest](#atomix.log.NextEntryRequest) | [NextEntryResponse](#atomix.log.NextEntryResponse) | NextEntry gets the next entry in the log |
| Remove | [RemoveRequest](#atomix.log.RemoveRequest) | [RemoveResponse](#atomix.log.RemoveResponse) | Remove removes an entry from the log |
| Clear | [ClearRequest](#atomix.log.ClearRequest) | [ClearResponse](#atomix.log.ClearResponse) | Clear removes all entries from the log |
| Events | [EventRequest](#atomix.log.EventRequest) | [EventResponse](#atomix.log.EventResponse) stream | Events listens for change events |
| Entries | [EntriesRequest](#atomix.log.EntriesRequest) | [EntriesResponse](#atomix.log.EntriesResponse) stream | Entries lists all entries in the log |

 



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

