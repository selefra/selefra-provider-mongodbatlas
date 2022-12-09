# Table: mongodbatlas_org_event

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| remote_address | string | X | √ | IP address of the userId Atlas user who triggered the event. | 
| target_public_key | string | X | √ | The public key of the API Key targeted by the event. | 
| username | string | X | √ | The username for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the publicKey field. | 
| api_key_id | string | X | √ | Unique identifier for the API Key that triggered the event. If this field is present in the response, Atlas does not return the userId field. | 
| current_value | json | X | √ | Describes the value of the metricName at the time of the event. | 
| metric_name | string | X | √ | The name of the metric associated to the alertId. | 
| invoice_id | string | X | √ | The unique identifier of the invoice associated to the event. | 
| user_id | string | X | √ | The unique identifier for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the apiKeyId field. | 
| op_type | string | X | √ | Type of operation that occurred. This field is present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD. | 
| payment_id | string | X | √ | The unique identifier of the invoice payment associated to the event. | 
| replica_set_name | string | X | √ | The name of the replica set associated to the event. | 
| shard_name | string | X | √ | The name of the shard associated to the event. | 
| whitelist_entry | string | X | √ | The white list entry of the API Key targeted by the event. | 
| alert_id | string | X | √ | Unique identifier for the alert associated with the event. | 
| project_id | string | X | √ | The unique identifier for the project in which the event occurred. | 
| is_global_admin | bool | X | √ | Indicates whether the user who triggered the event is a MongoDB employee. | 
| alert_config_id | string | X | √ | Unique identifier for the alert configuration associated to the alertId. | 
| org_id | string | X | √ | The unique identifier for the organization in which the event occurred. | 
| team_id | string | X | √ | The unique identifier for the Atlas team associated to the event. | 
| target_username | string | X | √ | The username for the Atlas user targeted by the event. | 
| title | string | X | √ | Title of the resource. | 
| hostname | string | X | √ | The hostname of the Atlas host machine associated to the event. | 
| links | json | X | √ | One or more uniform resource locators that link to sub-resources and/or related resources. The Web Linking Specification explains the relation-types between URLs. | 
| public_key | string | X | √ | Public key associated with the API Key that triggered the event. If this field is present in the response, Atlas does not return the 'username' field. | 
| event_type_name | string | X | √ | Human-readable label that indicates the type of event. | 
| id | string | X | √ | Unique identifier for the event. | 
| collection | string | X | √ | Name of the collection on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD. | 
| created | timestamp | X | √ | UTC date when the event occurred. | 
| database | string | X | √ | Name of the database on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD. | 
| port | int | X | √ | The port on which the mongod or mongos listens. | 


