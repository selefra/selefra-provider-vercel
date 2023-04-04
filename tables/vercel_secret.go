package tables

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/secret"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelSecretGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelSecretGenerator{}

func (x *TableVercelSecretGenerator) GetTableName() string {
	return "vercel_secret"
}

func (x *TableVercelSecretGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelSecretGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelSecretGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelSecretGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := secret.ListSecretsRequest{Limit: 100}
			for {

				res, err := conn.Secret.ListSecrets(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range res.Secrets {
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

func getSecret(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
	if err != nil {

		return nil, err
	}

	var nameOrId string
	if result != nil {
		s := result.(secret.Secret)
		nameOrId = s.Uid
	}

	res, err := conn.Secret.GetSecret(nameOrId)

	return res, err
}

func (x *TableVercelSecretGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelSecretGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Description("Time when the secret was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the team the secret was created for.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the user who created the secret.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getSecret(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the project the secret belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Description("Value of the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("decryptable").ColumnType(schema.ColumnTypeBool).Description("True if the secret value can be decrypted after it is created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the secret.").
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
	}
}

func (x *TableVercelSecretGenerator) GetSubTables() []*schema.Table {
	return nil
}
