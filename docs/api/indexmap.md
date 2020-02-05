# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/indexedmap/indexedmap.proto](#atomix/indexedmap/indexedmap.proto)
    - [ClearRequest](#atomix.indexedmap.ClearRequest)
    - [ClearResponse](#atomix.indexedmap.ClearResponse)
    - [CloseRequest](#atomix.indexedmap.CloseRequest)
    - [CloseResponse](#atomix.indexedmap.CloseResponse)
    - [CreateRequest](#atomix.indexedmap.CreateRequest)
    - [CreateResponse](#atomix.indexedmap.CreateResponse)
    - [EntriesRequest](#atomix.indexedmap.EntriesRequest)
    - [EntriesResponse](#atomix.indexedmap.EntriesResponse)
    - [EventRequest](#atomix.indexedmap.EventRequest)
    - [EventResponse](#atomix.indexedmap.EventResponse)
    - [ExistsRequest](#atomix.indexedmap.ExistsRequest)
    - [ExistsResponse](#atomix.indexedmap.ExistsResponse)
    - [FirstEntryRequest](#atomix.indexedmap.FirstEntryRequest)
    - [FirstEntryResponse](#atomix.indexedmap.FirstEntryResponse)
    - [GetRequest](#atomix.indexedmap.GetRequest)
    - [GetResponse](#atomix.indexedmap.GetResponse)
    - [LastEntryRequest](#atomix.indexedmap.LastEntryRequest)
    - [LastEntryResponse](#atomix.indexedmap.LastEntryResponse)
    - [NextEntryRequest](#atomix.indexedmap.NextEntryRequest)
    - [NextEntryResponse](#atomix.indexedmap.NextEntryResponse)
    - [PrevEntryRequest](#atomix.indexedmap.PrevEntryRequest)
    - [PrevEntryResponse](#atomix.indexedmap.PrevEntryResponse)
    - [PutRequest](#atomix.indexedmap.PutRequest)
    - [PutResponse](#atomix.indexedmap.PutResponse)
    - [RemoveRequest](#atomix.indexedmap.RemoveRequest)
    - [RemoveResponse](#atomix.indexedmap.RemoveResponse)
    - [ReplaceRequest](#atomix.indexedmap.ReplaceRequest)
    - [ReplaceResponse](#atomix.indexedmap.ReplaceResponse)
    - [SizeRequest](#atomix.indexedmap.SizeRequest)
    - [SizeResponse](#atomix.indexedmap.SizeResponse)
  
    - [EventResponse.Type](#atomix.indexedmap.EventResponse.Type)
    - [ResponseStatus](#atomix.indexedmap.ResponseStatus)
  
  
    - [IndexedMapService](#atomix.indexedmap.IndexedMapService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/indexedmap/indexedmap.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/indexedmap/indexedmap.proto



<a name="atomix.indexedmap.ClearRequest"></a>

### ClearRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.indexedmap.ClearResponse"></a>

### ClearResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.indexedmap.CloseRequest"></a>

### CloseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| delete | [bool](#bool) |  |  |






<a name="atomix.indexedmap.CloseResponse"></a>

### CloseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.indexedmap.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.indexedmap.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |






<a name="atomix.indexedmap.EntriesRequest"></a>

### EntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.indexedmap.EntriesResponse"></a>

### EntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| key | [string](#string) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| replay | [bool](#bool) |  |  |
| key | [string](#string) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.indexedmap.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| type | [EventResponse.Type](#atomix.indexedmap.EventResponse.Type) |  |  |
| key | [string](#string) |  |  |
| index | [int64](#int64) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.ExistsRequest"></a>

### ExistsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| key | [string](#string) |  |  |






<a name="atomix.indexedmap.ExistsResponse"></a>

### ExistsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| contains_key | [bool](#bool) |  |  |






<a name="atomix.indexedmap.FirstEntryRequest"></a>

### FirstEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.indexedmap.FirstEntryResponse"></a>

### FirstEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |






<a name="atomix.indexedmap.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.LastEntryRequest"></a>

### LastEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.indexedmap.LastEntryResponse"></a>

### LastEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.NextEntryRequest"></a>

### NextEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.indexedmap.NextEntryResponse"></a>

### NextEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.PrevEntryRequest"></a>

### PrevEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |






<a name="atomix.indexedmap.PrevEntryResponse"></a>

### PrevEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.PutRequest"></a>

### PutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| ttl | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.indexedmap.PutResponse"></a>

### PutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.indexedmap.ResponseStatus) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |






<a name="atomix.indexedmap.RemoveRequest"></a>

### RemoveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| value | [bytes](#bytes) |  |  |
| version | [int64](#int64) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="atomix.indexedmap.RemoveResponse"></a>

### RemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.indexedmap.ResponseStatus) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |






<a name="atomix.indexedmap.ReplaceRequest"></a>

### ReplaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |
| new_value | [bytes](#bytes) |  |  |
| ttl | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="atomix.indexedmap.ReplaceResponse"></a>

### ReplaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| status | [ResponseStatus](#atomix.indexedmap.ResponseStatus) |  |  |
| index | [int64](#int64) |  |  |
| key | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| previous_value | [bytes](#bytes) |  |  |
| previous_version | [int64](#int64) |  |  |






<a name="atomix.indexedmap.SizeRequest"></a>

### SizeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.RequestHeader](#atomix.headers.RequestHeader) |  |  |






<a name="atomix.indexedmap.SizeResponse"></a>

### SizeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [atomix.headers.ResponseHeader](#atomix.headers.ResponseHeader) |  |  |
| size | [int32](#int32) |  |  |





 


<a name="atomix.indexedmap.EventResponse.Type"></a>

### EventResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| INSERTED | 1 |  |
| UPDATED | 2 |  |
| REMOVED | 3 |  |



<a name="atomix.indexedmap.ResponseStatus"></a>

### ResponseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| NOOP | 1 |  |
| WRITE_LOCK | 2 |  |
| PRECONDITION_FAILED | 3 |  |


 

 


<a name="atomix.indexedmap.IndexedMapService"></a>

### IndexedMapService
IndexedMap service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#atomix.indexedmap.CreateRequest) | [CreateResponse](#atomix.indexedmap.CreateResponse) | Create creates an indexed map |
| Close | [CloseRequest](#atomix.indexedmap.CloseRequest) | [CloseResponse](#atomix.indexedmap.CloseResponse) | Close closes an indexed map |
| Size | [SizeRequest](#atomix.indexedmap.SizeRequest) | [SizeResponse](#atomix.indexedmap.SizeResponse) | Size returns the size of the map |
| Exists | [ExistsRequest](#atomix.indexedmap.ExistsRequest) | [ExistsResponse](#atomix.indexedmap.ExistsResponse) | Exists checks whether a key exists in the map |
| Put | [PutRequest](#atomix.indexedmap.PutRequest) | [PutResponse](#atomix.indexedmap.PutResponse) | Put puts an entry into the map |
| Replace | [ReplaceRequest](#atomix.indexedmap.ReplaceRequest) | [ReplaceResponse](#atomix.indexedmap.ReplaceResponse) | Replace performs a check-and-set operation on an entry in the map |
| Get | [GetRequest](#atomix.indexedmap.GetRequest) | [GetResponse](#atomix.indexedmap.GetResponse) | Get gets the entry for a key |
| FirstEntry | [FirstEntryRequest](#atomix.indexedmap.FirstEntryRequest) | [FirstEntryResponse](#atomix.indexedmap.FirstEntryResponse) | FirstEntry gets the first entry in the map |
| LastEntry | [LastEntryRequest](#atomix.indexedmap.LastEntryRequest) | [LastEntryResponse](#atomix.indexedmap.LastEntryResponse) | LastEntry gets the last entry in the map |
| PrevEntry | [PrevEntryRequest](#atomix.indexedmap.PrevEntryRequest) | [PrevEntryResponse](#atomix.indexedmap.PrevEntryResponse) | PrevEntry gets the previous entry in the map |
| NextEntry | [NextEntryRequest](#atomix.indexedmap.NextEntryRequest) | [NextEntryResponse](#atomix.indexedmap.NextEntryResponse) | NextEntry gets the next entry in the map |
| Remove | [RemoveRequest](#atomix.indexedmap.RemoveRequest) | [RemoveResponse](#atomix.indexedmap.RemoveResponse) | Remove removes an entry from the map |
| Clear | [ClearRequest](#atomix.indexedmap.ClearRequest) | [ClearResponse](#atomix.indexedmap.ClearResponse) | Clear removes all entries from the map |
| Events | [EventRequest](#atomix.indexedmap.EventRequest) | [EventResponse](#atomix.indexedmap.EventResponse) stream | Events listens for change events |
| Entries | [EntriesRequest](#atomix.indexedmap.EntriesRequest) | [EntriesResponse](#atomix.indexedmap.EntriesResponse) stream | Entries lists all entries in the map |

 



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

