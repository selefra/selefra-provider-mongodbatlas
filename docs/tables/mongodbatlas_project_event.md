# Table: mongodbatlas_project_event

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| metric_name | string | X | √ | The name of the metric associated to the alertId. | 
| username | string | X | √ | The username for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the publicKey field. | 
| invoice_id | string | X | √ | The unique identifier of the invoice associated to the event. | 
| hostname | string | X | √ | The hostname of the Atlas host machine associated to the event. | 
| port | int | X | √ | The port on which the mongod or mongos listens. | 
| public_key | string | X | √ | Public key associated with the API Key that triggered the event. If this field is present in the response, Atlas does not return the 'username' field. | 
| alert_id | string | X | √ | Unique identifier for the alert associated with the event. | 
| database | string | X | √ | Name of the database on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD. | 
| replica_set_name | string | X | √ | The name of the replica set associated to the event. | 
| whitelist_entry | string | X | √ | The white list entry of the API Key targeted by the event. | 
| alert_config_id | string | X | √ | Unique identifier for the alert configuration associated to the alertId. | 
| target_public_key | string | X | √ | The public key of the API Key targeted by the event. | 
| payment_id | string | X | √ | The unique identifier of the invoice payment associated to the event. | 
| event_type_name | string | X | √ | Human-readable label that indicates the type of event. | 
| org_id | string | X | √ | The unique identifier for the organization in which the event occurred. | 
| created | timestamp | X | √ | UTC date when the event occurred. | 
| user_id | string | X | √ | The unique identifier for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the apiKeyId field. | 
| current_value | json | X | √ | Describes the value of the metricName at the time of the event. | 
| op_type | string | X | √ | Type of operation that occurred. This field is present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD. | 
| shard_name | string | X | √ | The name of the shard associated to the event. | 
| target_username | string | X | √ | The username for the Atlas user targeted by the event. | 
| api_key_id | string | X | √ | Unique identifier for the API Key that triggered the event. If this field is present in the response, Atlas does not return the userId field. | 
| collection | string | X | √ | Name of the collection on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD. | 
| project_id | string | X | √ | The unique identifier for the project in which the event occurred. | 
| is_global_admin | bool | X | √ | Indicates whether the user who triggered the event is a MongoDB employee. | 
| links | json | X | √ | One or more uniform resource locators that link to sub-resources and/or related resources. The Web Linking Specification explains the relation-types between URLs. | 
| remote_address | string | X | √ | IP address of the userId Atlas user who triggered the event. | 
| team_id | string | X | √ | The unique identifier for the Atlas team associated to the event. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | Unique identifier for the event. | 


