package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasTeamGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasTeamGenerator{}

func (x *TableMongodbatlasTeamGenerator) GetTableName() string {
	return "mongodbatlas_team"
}

func (x *TableMongodbatlasTeamGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasTeamGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasTeamGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasTeamGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			org := task.ParentRawResult.(*mongodbatlas.Organization)

			client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			itemsPerPage := int64(500)

			pageNumber := 1

			for {
				teams, response, err := client.Teams.List(ctx, org.ID, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, team := range teams {
					rTeam := rowTeam{
						Team:  team,
						OrgId: org.ID,
					}
					resultChannel <- rTeam

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

type rowTeam struct {
	mongodbatlas.Team
	OrgId string
}

func getTeamUsers(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	data := result.(rowTeam)

	client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
	if err != nil {

		return nil, err
	}

	users, _, err := client.Teams.GetTeamUsersAssigned(ctx, data.OrgId, data.ID)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (x *TableMongodbatlasTeamGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasTeamGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("org_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the organization for this team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the team.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeJSON).Description("Users assigned to the team.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := getTeamUsers(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
	}
}

func (x *TableMongodbatlasTeamGenerator) GetSubTables() []*schema.Table {
	return nil
}
