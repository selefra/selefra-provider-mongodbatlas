# Table: mongodbatlas_container

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vpc_id | string | X | √ | Unique identifier of the project's Network Peering container. | 
| title | string | X | √ | Title of the resource. | 
| project_id | string | X | √ | Unique identifier for the project. | 
| azure_subscription_id | string | X | √ | Unique identifer of the Azure subscription in which the VNet resides. | 
| gcp_project_id | string | X | √ | Unique identifier of the Google Cloud project in which the network peer resides. Returns null until a peering connection is created. | 
| network_name | string | X | √ | Unique identifier of the Network Peering connection in the Atlas project. Returns null until a peering connection is created. | 
| region | string | X | √ | AWS region where the VCP resides or Azure region where the VNet resides. | 
| vnet_name | string | X | √ | Unique identifier of your Azure VNet. The value is null if there are no network peering connections in the container. | 
| id | string | X | √ | Unique identifier for the container. | 
| provider_name | string | X | √ | Cloud provider for this Network Peering connection. | 
| atlas_cidr_block | cidr | X | √ | CIDR block that Atlas uses for your clusters. | 
| provisioned | bool | X | √ | Flag that indicates if the project has clusters deployed in the Network Peering container or Azure VNet. | 


