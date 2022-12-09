package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasX509AuthenticationDatabaseUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasX509AuthenticationDatabaseUserGenerator{}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetTableName() string {
	return "mongodbatlas_x509_authentication_database_user"
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			project := task.ParentRawResult.(*mongodbatlas.Project)

			client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			projectId := project.ID
			dbUsersPageNumber := 1

			for {

				databaseUsers, response, err := client.DatabaseUsers.List(ctx, project.ID, &mongodbatlas.ListOptions{
					PageNum:      dbUsersPageNumber,
					ItemsPerPage: int(500),
				})

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, dbUser := range databaseUsers {
					x509CertsForUser, _, err := client.X509AuthDBUsers.GetUserCertificates(ctx, projectId, dbUser.Username, &mongodbatlas.ListOptions{})
					if err != nil {

						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					for _, uc := range x509CertsForUser {
						resultChannel <- uc

					}
				}

				if hasNextPage(response) {
					dbUsersPageNumber++
					continue
				}

				break

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("months_until_expiration").ColumnType(schema.ColumnTypeInt).Description("The number of months that the created certificate is valid for before expiry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the Atlas project to which this certificate belongs.").
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Subject")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Serial number of this certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject").ColumnType(schema.ColumnTypeString).Description("Fully distinguished name of the database user to which this certificate belongs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("Time when Atlas created this X.509 certificate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("not_after").ColumnType(schema.ColumnTypeTimestamp).Description("Time when this certificate expires.").Build(),
	}
}

func (x *TableMongodbatlasX509AuthenticationDatabaseUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
