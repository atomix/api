# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/controller/controller.proto](#atomix/controller/controller.proto)
    - [CreatePartitionGroupRequest](#atomix.controller.CreatePartitionGroupRequest)
    - [CreatePartitionGroupResponse](#atomix.controller.CreatePartitionGroupResponse)
    - [DeletePartitionGroupRequest](#atomix.controller.DeletePartitionGroupRequest)
    - [DeletePartitionGroupResponse](#atomix.controller.DeletePartitionGroupResponse)
    - [GetPartitionGroupsRequest](#atomix.controller.GetPartitionGroupsRequest)
    - [GetPartitionGroupsResponse](#atomix.controller.GetPartitionGroupsResponse)
    - [PartitionElectionRequest](#atomix.controller.PartitionElectionRequest)
    - [PartitionElectionResponse](#atomix.controller.PartitionElectionResponse)
    - [PrimaryTerm](#atomix.controller.PrimaryTerm)
  
  
  
    - [ControllerService](#atomix.controller.ControllerService)
  

- [atomix/controller/partition.proto](#atomix/controller/partition.proto)
    - [NodeConfig](#atomix.controller.NodeConfig)
    - [Partition](#atomix.controller.Partition)
    - [PartitionConfig](#atomix.controller.PartitionConfig)
    - [PartitionEndpoint](#atomix.controller.PartitionEndpoint)
    - [PartitionGroup](#atomix.controller.PartitionGroup)
    - [PartitionGroupId](#atomix.controller.PartitionGroupId)
    - [PartitionGroupSpec](#atomix.controller.PartitionGroupSpec)
    - [PartitionId](#atomix.controller.PartitionId)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="atomix/controller/controller.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/controller/controller.proto



<a name="atomix.controller.CreatePartitionGroupRequest"></a>

### CreatePartitionGroupRequest
Creates a new partition group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [PartitionGroupId](#atomix.controller.PartitionGroupId) |  |  |
| spec | [PartitionGroupSpec](#atomix.controller.PartitionGroupSpec) |  |  |






<a name="atomix.controller.CreatePartitionGroupResponse"></a>

### CreatePartitionGroupResponse
Returns the status of a partition group






<a name="atomix.controller.DeletePartitionGroupRequest"></a>

### DeletePartitionGroupRequest
Deletes a partition group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [PartitionGroupId](#atomix.controller.PartitionGroupId) |  |  |






<a name="atomix.controller.DeletePartitionGroupResponse"></a>

### DeletePartitionGroupResponse
Returns the result of deleting a partition group






<a name="atomix.controller.GetPartitionGroupsRequest"></a>

### GetPartitionGroupsRequest
Gets a list of partition groups in the given namespace


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [PartitionGroupId](#atomix.controller.PartitionGroupId) |  |  |






<a name="atomix.controller.GetPartitionGroupsResponse"></a>

### GetPartitionGroupsResponse
Returns a list of partition groups


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [PartitionGroup](#atomix.controller.PartitionGroup) | repeated |  |






<a name="atomix.controller.PartitionElectionRequest"></a>

### PartitionElectionRequest
Enters a primary election


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partition_id | [PartitionId](#atomix.controller.PartitionId) |  |  |
| member | [string](#string) |  |  |






<a name="atomix.controller.PartitionElectionResponse"></a>

### PartitionElectionResponse
Enter response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [PrimaryTerm](#atomix.controller.PrimaryTerm) |  |  |






<a name="atomix.controller.PrimaryTerm"></a>

### PrimaryTerm
Primary term


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [int64](#int64) |  |  |
| primary | [string](#string) |  |  |
| candidates | [string](#string) | repeated |  |





 

 

 


<a name="atomix.controller.ControllerService"></a>

### ControllerService
Atomix controller service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreatePartitionGroup | [CreatePartitionGroupRequest](#atomix.controller.CreatePartitionGroupRequest) | [CreatePartitionGroupResponse](#atomix.controller.CreatePartitionGroupResponse) | Creates a partition group |
| DeletePartitionGroup | [DeletePartitionGroupRequest](#atomix.controller.DeletePartitionGroupRequest) | [DeletePartitionGroupResponse](#atomix.controller.DeletePartitionGroupResponse) | Deletes a partition group |
| GetPartitionGroups | [GetPartitionGroupsRequest](#atomix.controller.GetPartitionGroupsRequest) | [GetPartitionGroupsResponse](#atomix.controller.GetPartitionGroupsResponse) | Gets a list of active partition groups |
| EnterElection | [PartitionElectionRequest](#atomix.controller.PartitionElectionRequest) | [PartitionElectionResponse](#atomix.controller.PartitionElectionResponse) stream | Enters a primary election for a specific partition |

 



<a name="atomix/controller/partition.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/controller/partition.proto



<a name="atomix.controller.NodeConfig"></a>

### NodeConfig
Node configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| host | [string](#string) |  |  |
| port | [int32](#int32) |  |  |
| api_port | [int32](#int32) |  |  |
| protocol_port | [int32](#int32) |  |  |






<a name="atomix.controller.Partition"></a>

### Partition
Partition info


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partition_id | [int32](#int32) |  |  |
| endpoints | [PartitionEndpoint](#atomix.controller.PartitionEndpoint) | repeated |  |






<a name="atomix.controller.PartitionConfig"></a>

### PartitionConfig
Partition configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partition | [PartitionId](#atomix.controller.PartitionId) |  |  |
| controller | [NodeConfig](#atomix.controller.NodeConfig) |  |  |
| members | [NodeConfig](#atomix.controller.NodeConfig) | repeated |  |






<a name="atomix.controller.PartitionEndpoint"></a>

### PartitionEndpoint
Partition endpoint


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |
| port | [int32](#int32) |  |  |






<a name="atomix.controller.PartitionGroup"></a>

### PartitionGroup
Partition group info


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [PartitionGroupId](#atomix.controller.PartitionGroupId) |  |  |
| spec | [PartitionGroupSpec](#atomix.controller.PartitionGroupSpec) |  |  |
| partitions | [Partition](#atomix.controller.Partition) | repeated |  |






<a name="atomix.controller.PartitionGroupId"></a>

### PartitionGroupId
Partition group name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| namespace | [string](#string) |  |  |






<a name="atomix.controller.PartitionGroupSpec"></a>

### PartitionGroupSpec
Partition group specification


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partitions | [uint32](#uint32) |  |  |
| partition_size | [uint32](#uint32) |  |  |
| protocol | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="atomix.controller.PartitionId"></a>

### PartitionId
Partition identifier


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| partition | [int32](#int32) |  |  |
| group | [PartitionGroupId](#atomix.controller.PartitionGroupId) |  |  |





 

 

 

 



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

