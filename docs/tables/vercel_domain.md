# Table: vercel_domain

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| config_verified_at | timestamp | X | √ | Time when the domain configuration was verified. | 
| id | string | X | √ | ID of the domain. | 
| transferred_at | timestamp | X | √ | Time when the domain was created. | 
| txt_verified_at | timestamp | X | √ | Time when the domain was created. | 
| cdn_enabled | bool | X | √ | If true, then the Content Delivery Network is enabled for this domain. | 
| name_servers | json | X | √ | Name servers for the domain. | 
| creator | json | X | √ | Creator of the domain. | 
| name | string | X | √ | Name of the domain. | 
| zone | bool | X | √ | Zone of the domain. | 
| verification_record | string | X | √ | Verification record for the domain. | 
| verified | bool | X | √ | True if the domain is verified. | 
| renew | bool | X | √ | True if the domain should auto-renew. | 
| ns_verified_at | timestamp | X | √ | Time when the name server was verified. | 
| created_at | timestamp | X | √ | Time when the domain was created. | 
| expires_at | timestamp | X | √ | Time when the domain expires. | 
| service_type | string | X | √ | Service provided by the domain, e.g. external. | 
| intended_name_servers | json | X | √ | Intended name servers for the domain. | 
| bought_at | timestamp | X | √ | Time when the domain was bought. | 


