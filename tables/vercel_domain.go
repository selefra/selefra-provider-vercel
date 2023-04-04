package tables

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/domain"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelDomainGenerator{}

func (x *TableVercelDomainGenerator) GetTableName() string {
	return "vercel_domain"
}

func (x *TableVercelDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := domain.ListDomainsRequest{Limit: 100}
			for {

				res, err := conn.Domain.List(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range res.Domains {
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

func (x *TableVercelDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelDomainGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("config_verified_at").ColumnType(schema.ColumnTypeString).Description("Time when the domain configuration was verified.").
			Extractor(column_value_extractor.StructSelector("ConfigVerifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the domain.").
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transferred_at").ColumnType(schema.ColumnTypeString).Description("Time when the domain was created.").
			Extractor(column_value_extractor.StructSelector("TransferredAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("txt_verified_at").ColumnType(schema.ColumnTypeString).Description("Time when the domain was created.").
			Extractor(column_value_extractor.StructSelector("TxtVerifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cdn_enabled").ColumnType(schema.ColumnTypeBool).Description("If true, then the Content Delivery Network is enabled for this domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name_servers").ColumnType(schema.ColumnTypeJSON).Description("Name servers for the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeJSON).Description("Creator of the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeBool).Description("Zone of the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verification_record").ColumnType(schema.ColumnTypeString).Description("Verification record for the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verified").ColumnType(schema.ColumnTypeBool).Description("True if the domain is verified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("renew").ColumnType(schema.ColumnTypeBool).Description("True if the domain should auto-renew.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ns_verified_at").ColumnType(schema.ColumnTypeString).Description("Time when the name server was verified.").
			Extractor(column_value_extractor.StructSelector("NsVerifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Description("Time when the domain was created.").
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expires_at").ColumnType(schema.ColumnTypeString).Description("Time when the domain expires.").
			Extractor(column_value_extractor.StructSelector("ExpiresAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_type").ColumnType(schema.ColumnTypeString).Description("Service provided by the domain, e.g. external.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("intended_name_servers").ColumnType(schema.ColumnTypeJSON).Description("Intended name servers for the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bought_at").ColumnType(schema.ColumnTypeString).Description("Time when the domain was bought.").
			Extractor(column_value_extractor.StructSelector("BoughtAt")).Build(),
	}
}

func (x *TableVercelDomainGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableVercelDnsRecordGenerator{}),
	}
}
