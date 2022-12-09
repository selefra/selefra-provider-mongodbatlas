# Table: mongodbatlas_custom_db_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| role_name | string | X | √ | The name of the role. | 
| project_id | string | X | √ | The unique identifier of the project for this role. | 
| actions | json | X | √ | Each object in the actions array represents an individual privilege action granted by the role. | 
| inherited_roles | json | X | √ | Each object in the inherited_roles array represents a key-value pair indicating the inherited role and the database on which the role is granted. | 
| title | string | X | √ | Title of the resource. | 


