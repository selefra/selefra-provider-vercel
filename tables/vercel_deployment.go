package tables

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/deployment"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelDeploymentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelDeploymentGenerator{}

func (x *TableVercelDeploymentGenerator) GetTableName() string {
	return "vercel_deployment"
}

func (x *TableVercelDeploymentGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelDeploymentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelDeploymentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelDeploymentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := deployment.ListDeploymentsRequest{Limit: 100}

			total := 0
			for {
				res, err := conn.Deployment.List(req)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range res.Deployments {
					resultChannel <- i
					total++

				}
				if res.Pagination.Next == 0 {
					break
				}
				req.Until = res.Pagination.Next
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableVercelDeploymentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelDeploymentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("meta").ColumnType(schema.ColumnTypeJSON).Description("GitHub metadata associated with the deployment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the deployment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("URL of the deployment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("One of: BUILDING, ERROR, INITIALIZING, QUEUED, READY, CANCELED.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Description("Time when the deployment was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeJSON).Description("Creator of the deployment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("building_at").ColumnType(schema.ColumnTypeString).Description("Time when deployment started to build.").
			Extractor(column_value_extractor.StructSelector("BuildingAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ready").ColumnType(schema.ColumnTypeString).Description("Time when deployment is ready to view.").
			Extractor(column_value_extractor.StructSelector("Ready")).Build(),
	}
}

func (x *TableVercelDeploymentGenerator) GetSubTables() []*schema.Table {
	return nil
}
