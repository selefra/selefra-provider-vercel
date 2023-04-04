package tables

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/team"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelTeamGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelTeamGenerator{}

func (x *TableVercelTeamGenerator) GetTableName() string {
	return "vercel_team"
}

func (x *TableVercelTeamGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelTeamGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelTeamGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelTeamGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := team.ListTeamsRequest{Limit: 100}
			for {

				res, err := conn.Team.ListTeams(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range res.Teams {
					resultChannel <- i
				}

				if res.Pagination.Next == 0 {
					break
				}
				opts.Until = int(res.Pagination.Next)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableVercelTeamGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelTeamGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("slug").ColumnType(schema.ColumnTypeString).Description("Slug of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeString).Description("Time when the team was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preview_deployment_suffix").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("staging_prefix").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invite_code").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("soft_block").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_project_transfers").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("avatar").ColumnType(schema.ColumnTypeString).Description("Avatar for the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("membership").ColumnType(schema.ColumnTypeJSON).Description("Membership of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profiles").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator_id").ColumnType(schema.ColumnTypeString).Description("ID of the user who created the team.").
			Extractor(column_value_extractor.StructSelector("Creatorid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).Description("Time when the team was last updated.").
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
	}
}

func (x *TableVercelTeamGenerator) GetSubTables() []*schema.Table {
	return nil
}
