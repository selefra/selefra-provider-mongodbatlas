package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"go.mongodb.org/atlas/mongodbatlas"
)

type TableMongodbatlasContainerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasContainerGenerator{}

func (x *TableMongodbatlasContainerGenerator) GetTableName() string {
	return "mongodbatlas_container"
}

func (x *TableMongodbatlasContainerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasContainerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasContainerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasContainerGenerator) GetDataSource() *schema.DataSource {
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
				listOptions := &mongodbatlas.ContainersListOptions{
					ListOptions: mongodbatlas.ListOptions{
						PageNum:      pageNumber,
						ItemsPerPage: int(itemsPerPage),
					},
				}

				containers, response, err := client.Containers.List(ctx, project.ID, listOptions)

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, container := range containers {
					c := rowContainer{
						Container: &container,
						ProjectId: project.ID,
					}
					resultChannel <- c

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

type rowContainer struct {
	*mongodbatlas.Container
	ProjectId string
}

func hasNextPage(r *mongodbatlas.Response) bool {
	for _, l := range r.Links {
		if l.Rel == "next" {
			return true
		}
	}
	return false
}

func (x *TableMongodbatlasContainerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasContainerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the project's Network Peering container.").
			Extractor(column_value_extractor.StructSelector("VPCID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("NetworkName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the project.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_subscription_id").ColumnType(schema.ColumnTypeString).Description("Unique identifer of the Azure subscription in which the VNet resides.").
			Extractor(column_value_extractor.StructSelector("AzureSubscriptionID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the Google Cloud project in which the network peer resides. Returns null until a peering connection is created.").
			Extractor(column_value_extractor.StructSelector("GCPProjectID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_name").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the Network Peering connection in the Atlas project. Returns null until a peering connection is created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("AWS region where the VCP resides or Azure region where the VNet resides.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vnet_name").ColumnType(schema.ColumnTypeString).Description("Unique identifier of your Azure VNet. The value is null if there are no network peering connections in the container.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the container.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_name").ColumnType(schema.ColumnTypeString).Description("Cloud provider for this Network Peering connection.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("atlas_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("CIDR block that Atlas uses for your clusters.").
			Extractor(column_value_extractor.StructSelector("AtlasCIDRBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioned").ColumnType(schema.ColumnTypeBool).Description("Flag that indicates if the project has clusters deployed in the Network Peering container or Azure VNet.").Build(),
	}
}

func (x *TableMongodbatlasContainerGenerator) GetSubTables() []*schema.Table {
	return nil
}
