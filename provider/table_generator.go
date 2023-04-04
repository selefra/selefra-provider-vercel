package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
	"github.com/selefra/selefra-provider-vercel/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableVercelDeploymentGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableVercelDomainGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableVercelSecretGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableVercelUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableVercelProjectGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableVercelTeamGenerator{}),
	}
}
