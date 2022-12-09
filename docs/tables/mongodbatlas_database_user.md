# Table: mongodbatlas_database_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| username | string | X | √ | Username needed to authenticate to the MongoDB database or collection. | 
| roles | json | X | √ | List that contains key-value pairs that tag and categorize the database user. | 
| scopes | json | X | √ | List that contains key-value pairs that tag and categorize the database user. | 
| title | string | X | √ | Title of the resource. | 
| database_name | string | X | √ | Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. | 
| delete_after_date | timestamp | X | √ | Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the temporary access list entry. Atlas returns this field if you specified an expiration date when creating this access list entry. | 
| project_id | string | X | √ | Unique identifier of the project to which this access list entry applies. | 
| labels | json | X | √ | List that contains key-value pairs that tag and categorize the database user. | 


