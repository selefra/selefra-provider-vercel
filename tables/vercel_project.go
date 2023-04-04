package tables

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/project"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelProjectGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelProjectGenerator{}

func (x *TableVercelProjectGenerator) GetTableName() string {
	return "vercel_project"
}

func (x *TableVercelProjectGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelProjectGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelProjectGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelProjectGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := project.ListProjectsRequest{Limit: 100}
			for {

				res, err := conn.Project.List(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range res.Projects {
					resultChannel <- i
				}

				if res.Pagination.Next == 0 {
					break
				}
				opts.Until = res.Pagination.Next
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableVercelProjectGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelProjectGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("framework").ColumnType(schema.ColumnTypeString).Description("Framework used in the project, e.g. nextjs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("analytics").ColumnType(schema.ColumnTypeJSON).Description("Analytics information, if enabled for the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("directory_listing").ColumnType(schema.ColumnTypeBool).Description("If true then the project is listed in the Vercel directory.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("accountid").ColumnType(schema.ColumnTypeString).Description("Account ID for the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_files_outside_root_directory").ColumnType(schema.ColumnTypeBool).Description("If true then source files are outside the root directory.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("targets").ColumnType(schema.ColumnTypeJSON).Description("Targets of the build.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_deployments").ColumnType(schema.ColumnTypeJSON).Description("Information about the latest deployments of the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_protection").ColumnType(schema.ColumnTypeJSON).Description("Password protection information, if enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("link").ColumnType(schema.ColumnTypeJSON).Description("Details of the link from this project to a source code repository.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sso_protection").ColumnType(schema.ColumnTypeJSON).Description("SSO protection information, if enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("install_command").ColumnType(schema.ColumnTypeString).Description("The install command for this project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_version").ColumnType(schema.ColumnTypeString).Description("Node version used by the project, e.g. 16.x.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_expose_system_envs").ColumnType(schema.ColumnTypeBool).Description("If true then system environment variables are exposed for use.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dev_command").ColumnType(schema.ColumnTypeString).Description("The dev command for this project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("env").ColumnType(schema.ColumnTypeJSON).Description("Environment variables for the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the project.").
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("live").ColumnType(schema.ColumnTypeBool).Description("If true, the project is live.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_directory").ColumnType(schema.ColumnTypeString).Description("Directory where output of the build will go.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_source").ColumnType(schema.ColumnTypeBool).Description("If true, the project is linked to a public source.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serverless_function_region").ColumnType(schema.ColumnTypeString).Description("Region where serverless functions will be deployed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).Description("Time when the project was last updated.").
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeJSON).Description("Permissions settings.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("root_directory").ColumnType(schema.ColumnTypeString).Description("Root directory for the build process.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_command").ColumnType(schema.ColumnTypeString).Description("The build command for this project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Description("Time when the project was created.").
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
	}
}

func (x *TableVercelProjectGenerator) GetSubTables() []*schema.Table {
	return nil
}
