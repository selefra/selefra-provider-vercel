package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelUserGenerator{}

func (x *TableVercelUserGenerator) GetTableName() string {
	return "vercel_user"
}

func (x *TableVercelUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			res, err := conn.User.Get()
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- res.User

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableVercelUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("date").ColumnType(schema.ColumnTypeString).Description("Date when the user was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remote_caching").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bio").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Description("Username of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("website").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profiles").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Email address of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("soft_block").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_trial_available").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_flow_git_namespace").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_flow_git_provider").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("staging_prefix").ColumnType(schema.ColumnTypeString).Description("Username of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_version").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_flow_git_namespace_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the user.").
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_config").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableVercelUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
