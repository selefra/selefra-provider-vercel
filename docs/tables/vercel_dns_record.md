# Table: vercel_dns_record

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Name of the DNS record. | 
| id | string | X | √ | ID of the DNS record. | 
| priority | int | X | √ | Priority of the DNS record. | 
| updated_at | timestamp | X | √ | Time when the DNS record was created. | 
| created_at | timestamp | X | √ | Time when the DNS record was created. | 
| creator | string | X | √ | Creator of the DNS record. | 
| mx_priority | int | X | √ | MX priority of the DNS record. | 
| slug | string | X | √ | Slug of the DNS record. | 
| domain_name | string | X | √ | Domain name the record belongs to. | 
| type | string | X | √ | Type of the DNS record. | 
| value | string | X | √ | Type of the DNS record. | 
| ttl | int | X | √ | Time To Live of the DNS record. | 


