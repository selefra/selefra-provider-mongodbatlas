# Table: mongodbatlas_x509_authentication_database_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| months_until_expiration | int | X | √ | The number of months that the created certificate is valid for before expiry. | 
| project_id | string | X | √ | Unique identifier of the Atlas project to which this certificate belongs. | 
| title | string | X | √ | Title of the resource. | 
| id | int | X | √ | Serial number of this certificate. | 
| subject | string | X | √ | Fully distinguished name of the database user to which this certificate belongs. | 
| created_at | timestamp | X | √ | Time when Atlas created this X.509 certificate. | 
| not_after | timestamp | X | √ | Time when this certificate expires. | 


