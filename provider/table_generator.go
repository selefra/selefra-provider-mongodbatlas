package provider

import (
	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-mongodbatlas/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableMongodbatlasOrgGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableMongodbatlasProjectGenerator{}),
	}
}
