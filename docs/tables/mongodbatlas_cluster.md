# Table: mongodbatlas_cluster

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| replication_spec | json | X | √ | Configuration of each region in the cluster. Each element in this object represents a region where Atlas deploys your cluster. | 
| project_id | string | X | √ | Unique identifier of the project that this cluster belongs to. | 
| encryption_at_rest_provider | string | X | √ | Cloud service provider that offers Encryption at Rest. | 
| mongo_uri | string | X | √ | Base connection string for the cluster. | 
| provider_settings | json | X | √ | Configuration for the provisioned hosts on which MongoDB runs. The available options are specific to the cloud service provider. | 
| mongodb_major_version | string | X | √ | MongoDB Version of the cluster. | 
| srv_address | string | X | √ | Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS. The mongoURI parameter lists additional options. | 
| title | string | X | √ | Title of the resource. | 
| paused | bool | X | √ | Flag that indicates whether the cluster has been paused. | 
| connection_strings | json | X | √ | Set of connection strings that your applications use to connect to this cluster. | 
| mongo_uri_updated | timestamp | X | √ | Timestamp when the connection string was last updated. | 
| num_shards | int | X | √ | Positive integer that specifies the number of shards for a sharded cluster. If this is set to 1, the cluster is a replica set. If this is set to 2 or higher, the cluster is a sharded cluster with the number of shards specified. | 
| state_name | string | X | √ | Condition in which the API resource finds the cluster when you called the resource. The resource returns one of the following states: IDLE, CREATING, UPDATING, DELETING, DELETED, REPAIRING. | 
| name | string | X | √ | The name of the cluster as it appears in Atlas. | 
| replication_factor | int | X | √ | Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. | 
| pit_enabled | bool | X | √ | Flag that indicates whether the cluster uses continuous cloud backups. More information is available at https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#continuous-cloud-backups. | 
| id | string | X | √ | Unique 24-hexadecimal digit string that identifies the cluster. | 
| bi_connector | json | X | √ | Configuration settings applied to BI Connector for Atlas on this cluster. | 
| labels | json | X | √ | Collection of key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. | 
| mongodb_version | string | X | √ | Version of MongoDB that the cluster is running, in X.Y.Z format. | 
| provider_backup_enabled | bool | X | √ | Flag that indicates if the cluster uses Back Up Your Database Deployment for backups. | 
| replication_specs | json | X | √ | Configuration for each zone in a Global Cluster. Each object in this array represents a zone where Atlas deploys nodes for your Global Cluster. | 
| cluster_type | string | X | √ | Type of the cluster. | 
| disk_size_gb | float | X | √ | Capacity, in gigabytes, of the host's root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive number. | 
| version_release_system | string | X | √ | Release cadence that Atlas uses for this cluster. | 
| auto_scaling | json | X | √ | Collection of settings that configures auto-scaling information for the cluster. | 
| mongo_uri_with_options | string | X | √ | Connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. | 


