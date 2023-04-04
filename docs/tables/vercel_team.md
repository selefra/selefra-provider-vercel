# Table: vercel_team

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| slug | string | X | √ | Slug of the team. | 
| created | timestamp | X | √ | Time when the team was created. | 
| platform_version | string | X | √ |  | 
| preview_deployment_suffix | json | X | √ |  | 
| resource_config | json | X | √ |  | 
| staging_prefix | string | X | √ |  | 
| name | string | X | √ | Name of the team. | 
| description | string | X | √ |  | 
| invite_code | string | X | √ |  | 
| soft_block | json | X | √ |  | 
| id | string | X | √ | Unique identifier of the team. | 
| allow_project_transfers | bool | X | √ |  | 
| avatar | string | X | √ | Avatar for the team. | 
| billing | json | X | √ |  | 
| membership | json | X | √ | Membership of the team. | 
| profiles | json | X | √ |  | 
| creator_id | string | X | √ | ID of the user who created the team. | 
| updated_at | timestamp | X | √ | Time when the team was last updated. | 


