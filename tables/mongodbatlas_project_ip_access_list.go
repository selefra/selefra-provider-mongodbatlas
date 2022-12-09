package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasProjectIpAccessListGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasProjectIpAccessListGenerator{}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetTableName() string {
	return "mongodbatlas_project_ip_access_list"
}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetDataSource() *schema.DataSource {
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
				projectIpAccessLists, response, err := client.ProjectIPAccessList.List(ctx, project.ID, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, projectIPAccessList := range projectIpAccessLists.Results {
					resultChannel <- projectIPAccessList

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

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("delete_after_date").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the temporary access list entry. Atlas returns this field if you specified an expiration date when creating this access list entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeIp).Description("Entry using an IP address in this access list entry.").
			Extractor(column_value_extractor.StructSelector("IPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the project to which this access list entry applies.").
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("CIDRBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_security_group").ColumnType(schema.ColumnTypeString).Description("Unique identifier of AWS security group in this access list entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_block").ColumnType(schema.ColumnTypeIp).Description("Range of IP addresses in CIDR notation in this access list entry.").
			Extractor(column_value_extractor.StructSelector("CIDRBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment associated with this access list entry.").Build(),
	}
}

func (x *TableMongodbatlasProjectIpAccessListGenerator) GetSubTables() []*schema.Table {
	return nil
}
