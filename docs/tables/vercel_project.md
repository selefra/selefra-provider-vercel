# Table: vercel_project

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| framework | string | X | √ | Framework used in the project, e.g. nextjs. | 
| analytics | json | X | √ | Analytics information, if enabled for the project. | 
| directory_listing | bool | X | √ | If true then the project is listed in the Vercel directory. | 
| accountid | string | X | √ | Account ID for the project. | 
| alias | json | X | √ |  | 
| source_files_outside_root_directory | bool | X | √ | If true then source files are outside the root directory. | 
| targets | json | X | √ | Targets of the build. | 
| latest_deployments | json | X | √ | Information about the latest deployments of the project. | 
| password_protection | json | X | √ | Password protection information, if enabled. | 
| link | json | X | √ | Details of the link from this project to a source code repository. | 
| sso_protection | json | X | √ | SSO protection information, if enabled. | 
| install_command | string | X | √ | The install command for this project. | 
| node_version | string | X | √ | Node version used by the project, e.g. 16.x. | 
| auto_expose_system_envs | bool | X | √ | If true then system environment variables are exposed for use. | 
| dev_command | string | X | √ | The dev command for this project. | 
| env | json | X | √ | Environment variables for the project. | 
| id | string | X | √ | ID of the project. | 
| live | bool | X | √ | If true, the project is live. | 
| output_directory | string | X | √ | Directory where output of the build will go. | 
| public_source | bool | X | √ | If true, the project is linked to a public source. | 
| serverless_function_region | string | X | √ | Region where serverless functions will be deployed. | 
| updated_at | timestamp | X | √ | Time when the project was last updated. | 
| name | string | X | √ | Name of the project. | 
| permissions | json | X | √ | Permissions settings. | 
| root_directory | string | X | √ | Root directory for the build process. | 
| build_command | string | X | √ | The build command for this project. | 
| created_at | timestamp | X | √ | Time when the project was created. | 


