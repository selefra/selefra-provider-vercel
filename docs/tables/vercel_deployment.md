# Table: vercel_deployment

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| meta | json | X | √ | GitHub metadata associated with the deployment. | 
| name | string | X | √ | Name of the deployment. | 
| url | string | X | √ | URL of the deployment. | 
| state | string | X | √ | One of: BUILDING, ERROR, INITIALIZING, QUEUED, READY, CANCELED. | 
| created_at | timestamp | X | √ | Time when the deployment was created. | 
| creator | json | X | √ | Creator of the deployment. | 
| building_at | timestamp | X | √ | Time when deployment started to build. | 
| ready | timestamp | X | √ | Time when deployment is ready to view. | 


