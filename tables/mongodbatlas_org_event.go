package tables

import (
	"context"
	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasOrgEventGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasOrgEventGenerator{}

func (x *TableMongodbatlasOrgEventGenerator) GetTableName() string {
	return "mongodbatlas_org_event"
}

func (x *TableMongodbatlasOrgEventGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasOrgEventGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasOrgEventGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasOrgEventGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			org := task.ParentRawResult.(*mongodbatlas.Organization)
			client, err := mongodbatlas_client.GetMongoDBAtlasClient(ctx, taskClient.(*mongodbatlas_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			itemsPerPage := int64(500)

			pageNumber := 1
			orgId := org.ID

			for {
				listOptions := &mongodbatlas.EventListOptions{
					ListOptions: mongodbatlas.ListOptions{
						PageNum:      pageNumber,
						ItemsPerPage: int(itemsPerPage),
					},
				}

				orgEvents, response, err := client.Events.ListOrganizationEvents(ctx, orgId, listOptions)

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, projectEvent := range orgEvents.Results {
					resultChannel <- projectEvent

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

func (x *TableMongodbatlasOrgEventGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasOrgEventGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("remote_address").ColumnType(schema.ColumnTypeString).Description("IP address of the userId Atlas user who triggered the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_public_key").ColumnType(schema.ColumnTypeString).Description("The public key of the API Key targeted by the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Description("The username for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the publicKey field.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_key_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the API Key that triggered the event. If this field is present in the response, Atlas does not return the userId field.").
			Extractor(column_value_extractor.StructSelector("APIKeyID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_value").ColumnType(schema.ColumnTypeJSON).Description("Describes the value of the metricName at the time of the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric associated to the alertId.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invoice_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the invoice associated to the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the apiKeyId field.").
			Extractor(column_value_extractor.StructSelector("UserID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("op_type").ColumnType(schema.ColumnTypeString).Description("Type of operation that occurred. This field is present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payment_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier of the invoice payment associated to the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_set_name").ColumnType(schema.ColumnTypeString).Description("The name of the replica set associated to the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shard_name").ColumnType(schema.ColumnTypeString).Description("The name of the shard associated to the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("whitelist_entry").ColumnType(schema.ColumnTypeString).Description("The white list entry of the API Key targeted by the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the alert associated with the event.").
			Extractor(column_value_extractor.StructSelector("AlertID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the project in which the event occurred.").
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_global_admin").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the user who triggered the event is a MongoDB employee.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_config_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the alert configuration associated to the alertId.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("org_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the organization in which the event occurred.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeString).Description("The unique identifier for the Atlas team associated to the event.").
			Extractor(column_value_extractor.StructSelector("TeamID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_username").ColumnType(schema.ColumnTypeString).Description("The username for the Atlas user targeted by the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hostname").ColumnType(schema.ColumnTypeString).Description("The hostname of the Atlas host machine associated to the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("links").ColumnType(schema.ColumnTypeJSON).Description("One or more uniform resource locators that link to sub-resources and/or related resources. The Web Linking Specification explains the relation-types between URLs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_key").ColumnType(schema.ColumnTypeString).Description("Public key associated with the API Key that triggered the event. If this field is present in the response, Atlas does not return the 'username' field.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_type_name").ColumnType(schema.ColumnTypeString).Description("Human-readable label that indicates the type of event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collection").ColumnType(schema.ColumnTypeString).Description("Name of the collection on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("UTC date when the event occurred.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database").ColumnType(schema.ColumnTypeString).Description("Name of the database on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeInt).Description("The port on which the mongod or mongos listens.").Build(),
	}
}

func (x *TableMongodbatlasOrgEventGenerator) GetSubTables() []*schema.Table {
	return nil
}
