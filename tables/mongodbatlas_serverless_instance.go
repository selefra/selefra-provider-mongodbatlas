package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasServerlessInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasServerlessInstanceGenerator{}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetTableName() string {
	return "mongodbatlas_serverless_instance"
}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			project := task.ParentRawResult.(*mongodbatlas.Project)
			client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			itemsPerPage := int64(500)

			pageNumber := 1

			for {
				serverlessInstances, response, err := client.ServerlessInstances.List(ctx, project.ID, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, cluster := range serverlessInstances.Results {
					resultChannel <- cluster

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

func (x *TableMongodbatlasServerlessInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("connection_strings").ColumnType(schema.ColumnTypeJSON).Description("Set of connection strings that your applications use to connect to this cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the cluster as it appears in Atlas.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the project that this cluster belongs to.").
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mongodb_version").ColumnType(schema.ColumnTypeString).Description("Version of MongoDB that the serverless instance runs, in <major version>.<minor version> format.").
			Extractor(column_value_extractor.StructSelector("MongoDBVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_settings").ColumnType(schema.ColumnTypeJSON).Description("Configuration for the provisioned hosts on which MongoDB runs. The available options are specific to the cloud service provider.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_name").ColumnType(schema.ColumnTypeString).Description("Stage of deployment of this serverless instance when the resource made its request.").Build(),
	}
}

func (x *TableMongodbatlasServerlessInstanceGenerator) GetSubTables() []*schema.Table {
	return nil
}
