# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/protocols/raft/raft.proto](#atomix/protocols/raft/raft.proto)
    - [CompactionSpec](#atomix.protocols.raft.CompactionSpec)
    - [RaftProtocol](#atomix.protocols.raft.RaftProtocol)
    - [StorageSpec](#atomix.protocols.raft.StorageSpec)
  
    - [MemberGroupStrategy](#atomix.protocols.raft.MemberGroupStrategy)
    - [StorageLevel](#atomix.protocols.raft.StorageLevel)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/protocols/raft/raft.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/protocols/raft/raft.proto



<a name="atomix.protocols.raft.CompactionSpec"></a>

### CompactionSpec
Partition group compaction configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic | [bool](#bool) |  |  |
| free_disk_buffer | [double](#double) |  |  |






<a name="atomix.protocols.raft.RaftProtocol"></a>

### RaftProtocol
Raft protocol configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| election_timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| heartbeat_interval | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| storage | [StorageSpec](#atomix.protocols.raft.StorageSpec) |  |  |
| compaction | [CompactionSpec](#atomix.protocols.raft.CompactionSpec) |  |  |






<a name="atomix.protocols.raft.StorageSpec"></a>

### StorageSpec
Partition group storage configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_entry_size | [uint32](#uint32) |  |  |
| segment_size | [uint32](#uint32) |  |  |
| level | [StorageLevel](#atomix.protocols.raft.StorageLevel) |  |  |
| flush_on_commit | [bool](#bool) |  |  |





 


<a name="atomix.protocols.raft.MemberGroupStrategy"></a>

### MemberGroupStrategy
Member group strategy

| Name | Number | Description |
| ---- | ------ | ----------- |
| HOST_AWARE | 0 |  |
| RACK_AWARE | 1 |  |
| ZONE_AWARE | 2 |  |



<a name="atomix.protocols.raft.StorageLevel"></a>

### StorageLevel
Storage level

| Name | Number | Description |
| ---- | ------ | ----------- |
| DISK | 0 |  |
| MAPPED | 1 |  |


 

 

 



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
