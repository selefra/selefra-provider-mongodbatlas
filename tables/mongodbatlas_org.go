package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasOrgGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasOrgGenerator{}

func (x *TableMongodbatlasOrgGenerator) GetTableName() string {
	return "mongodbatlas_org"
}

func (x *TableMongodbatlasOrgGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasOrgGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasOrgGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasOrgGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			itemsPerPage := int64(500)

			pageNumber := 1

			for {
				orgs, response, _ := client.Organizations.List(ctx, &mongodbatlas.OrganizationsListOptions{
					ListOptions: mongodbatlas.ListOptions{
						PageNum:      pageNumber,
						ItemsPerPage: int(itemsPerPage),
					},
				})

				for _, org := range orgs.Results {
					resultChannel <- org

				}

				if hasNextPage(response) {
					pageNumber++
					continue
				}

				break
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableMongodbatlasOrgGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasOrgGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_deleted").ColumnType(schema.ColumnTypeBool).Description("Flag indicating if the organization is deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableMongodbatlasOrgGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableMongodbatlasTeamGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasOrgEventGenerator{}),
	}
}
