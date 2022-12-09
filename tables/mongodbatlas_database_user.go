package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasDatabaseUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasDatabaseUserGenerator{}

func (x *TableMongodbatlasDatabaseUserGenerator) GetTableName() string {
	return "mongodbatlas_database_user"
}

func (x *TableMongodbatlasDatabaseUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasDatabaseUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasDatabaseUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasDatabaseUserGenerator) GetDataSource() *schema.DataSource {
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
				databaseUsers, response, err := client.DatabaseUsers.List(ctx, project.ID, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, databaseUser := range databaseUsers {
					resultChannel <- databaseUser

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

func (x *TableMongodbatlasDatabaseUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasDatabaseUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Description("Username needed to authenticate to the MongoDB database or collection.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("roles").ColumnType(schema.ColumnTypeJSON).Description("List that contains key-value pairs that tag and categorize the database user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scopes").ColumnType(schema.ColumnTypeJSON).Description("List that contains key-value pairs that tag and categorize the database user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("DatabaseName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_after_date").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the temporary access list entry. Atlas returns this field if you specified an expiration date when creating this access list entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the project to which this access list entry applies.").
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Description("List that contains key-value pairs that tag and categorize the database user.").Build(),
	}
}

func (x *TableMongodbatlasDatabaseUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
