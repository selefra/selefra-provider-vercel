package tables

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/dns"
	"github.com/chronark/vercel-go/endpoints/domain"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

type TableVercelDnsRecordGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableVercelDnsRecordGenerator{}

func (x *TableVercelDnsRecordGenerator) GetTableName() string {
	return "vercel_dns_record"
}

func (x *TableVercelDnsRecordGenerator) GetTableDescription() string {
	return ""
}

func (x *TableVercelDnsRecordGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableVercelDnsRecordGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableVercelDnsRecordGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := vercel_client.Connect(ctx, taskClient.(*vercel_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			parentDomain := task.ParentRawResult.(domain.Domain)
			opts := dns.ListDnsRequest{Domain: parentDomain.Name, Limit: 100}

			for {

				res, err := conn.Dns.List(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range res.Records {
					resultChannel <- dnsRecordRow{parentDomain, i}
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

type dnsRecordRow struct {
	Domain domain.Domain
	dns.Record
}

func (x *TableVercelDnsRecordGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableVercelDnsRecordGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).Description("Priority of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).Description("Time when the DNS record was created.").
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Description("Time when the DNS record was created.").
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeString).Description("Creator of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mx_priority").ColumnType(schema.ColumnTypeInt).Description("MX priority of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("slug").ColumnType(schema.ColumnTypeString).Description("Slug of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).Description("Domain name the record belongs to.").
			Extractor(column_value_extractor.StructSelector("Domain.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Type of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Description("Type of the DNS record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl").ColumnType(schema.ColumnTypeInt).Description("Time To Live of the DNS record.").Build(),
	}
}

func (x *TableVercelDnsRecordGenerator) GetSubTables() []*schema.Table {
	return nil
}
