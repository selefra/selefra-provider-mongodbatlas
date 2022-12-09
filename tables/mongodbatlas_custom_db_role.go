package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasCustomDbRoleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasCustomDbRoleGenerator{}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetTableName() string {
	return "mongodbatlas_custom_db_role"
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetDataSource() *schema.DataSource {
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
				roles, response, err := client.CustomDBRoles.List(ctx, project.ID, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, role := range *roles {
					r := rowCustomDBRole{
						CustomDBRole: role,
						ProjectID:    project.ID,
					}
					resultChannel <- r

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

type rowCustomDBRole struct {
	mongodbatlas.CustomDBRole
	ProjectID string
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).Description("The name of the role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the project for this role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actions").ColumnType(schema.ColumnTypeJSON).Description("Each object in the actions array represents an individual privilege action granted by the role.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inherited_roles").ColumnType(schema.ColumnTypeJSON).Description("Each object in the inherited_roles array represents a key-value pair indicating the inherited role and the database on which the role is granted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("RoleName")).Build(),
	}
}

func (x *TableMongodbatlasCustomDbRoleGenerator) GetSubTables() []*schema.Table {
	return nil
}
