# Table: mongodbatlas_serverless_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| connection_strings | json | X | √ | Set of connection strings that your applications use to connect to this cluster. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | Unique identifier of the cluster. | 
| name | string | X | √ | The name of the cluster as it appears in Atlas. | 
| project_id | string | X | √ | Unique identifier of the project that this cluster belongs to. | 
| mongodb_version | string | X | √ | Version of MongoDB that the serverless instance runs, in <major version>.<minor version> format. | 
| provider_settings | json | X | √ | Configuration for the provisioned hosts on which MongoDB runs. The available options are specific to the cloud service provider. | 
| state_name | string | X | √ | Stage of deployment of this serverless instance when the resource made its request. | 


