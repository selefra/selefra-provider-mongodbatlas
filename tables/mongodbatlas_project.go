package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasProjectGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasProjectGenerator{}

func (x *TableMongodbatlasProjectGenerator) GetTableName() string {
	return "mongodbatlas_project"
}

func (x *TableMongodbatlasProjectGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasProjectGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasProjectGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasProjectGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			itemsPerPage := int64(500)

			pageNumber := 1

			for {
				projects, response, _ := client.Projects.GetAllProjects(ctx, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				for _, project := range projects.Results {
					resultChannel <- project

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

func (x *TableMongodbatlasProjectGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasProjectGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("A unique identifier of the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_count").ColumnType(schema.ColumnTypeInt).Description("The number of Atlas clusters deployed in the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the Atlas organization to which the project belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableMongodbatlasProjectGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableMongodbatlasDatabaseUserGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasCustomDbRoleGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasClusterGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasProjectEventGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasContainerGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasServerlessInstanceGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasProjectIpAccessListGenerator{}),
		table_schema_generator.GenTableSchema(&TableMongodbatlasX509AuthenticationDatabaseUserGenerator{}),
	}
}
