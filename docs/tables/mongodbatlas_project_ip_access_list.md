# Table: mongodbatlas_project_ip_access_list

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| delete_after_date | timestamp | X | √ | Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the temporary access list entry. Atlas returns this field if you specified an expiration date when creating this access list entry. | 
| ip_address | ip | X | √ | Entry using an IP address in this access list entry. | 
| project_id | string | X | √ | Unique identifier of the project to which this access list entry applies. | 
| title | string | X | √ | Title of the resource. | 
| aws_security_group | string | X | √ | Unique identifier of AWS security group in this access list entry. | 
| cidr_block | ip | X | √ | Range of IP addresses in CIDR notation in this access list entry. | 
| comment | string | X | √ | Comment associated with this access list entry. | 


