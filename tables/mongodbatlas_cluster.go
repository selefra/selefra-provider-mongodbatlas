package tables

import (
	"context"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/selefra/selefra-provider-mongodbatlas/mongodbatlas_client"
	"github.com/selefra/selefra-provider-mongodbatlas/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableMongodbatlasClusterGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableMongodbatlasClusterGenerator{}

func (x *TableMongodbatlasClusterGenerator) GetTableName() string {
	return "mongodbatlas_cluster"
}

func (x *TableMongodbatlasClusterGenerator) GetTableDescription() string {
	return ""
}

func (x *TableMongodbatlasClusterGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableMongodbatlasClusterGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableMongodbatlasClusterGenerator) GetDataSource() *schema.DataSource {
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
				clusters, response, err := client.Clusters.List(ctx, project.ID, &mongodbatlas.ListOptions{
					PageNum:      pageNumber,
					ItemsPerPage: int(itemsPerPage),
				})

				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, cluster := range clusters {
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

func (x *TableMongodbatlasClusterGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableMongodbatlasClusterGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("replication_spec").ColumnType(schema.ColumnTypeJSON).Description("Configuration of each region in the cluster. Each element in this object represents a region where Atlas deploys your cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the project that this cluster belongs to.").
			Extractor(column_value_extractor.StructSelector("GroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_at_rest_provider").ColumnType(schema.ColumnTypeString).Description("Cloud service provider that offers Encryption at Rest.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mongo_uri").ColumnType(schema.ColumnTypeString).Description("Base connection string for the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_settings").ColumnType(schema.ColumnTypeJSON).Description("Configuration for the provisioned hosts on which MongoDB runs. The available options are specific to the cloud service provider.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mongodb_major_version").ColumnType(schema.ColumnTypeString).Description("MongoDB Version of the cluster.").
			Extractor(column_value_extractor.StructSelector("MongoDBMajorVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("srv_address").ColumnType(schema.ColumnTypeString).Description("Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS. The mongoURI parameter lists additional options.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paused").ColumnType(schema.ColumnTypeBool).Description("Flag that indicates whether the cluster has been paused.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_strings").ColumnType(schema.ColumnTypeJSON).Description("Set of connection strings that your applications use to connect to this cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mongo_uri_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the connection string was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_shards").ColumnType(schema.ColumnTypeInt).Description("Positive integer that specifies the number of shards for a sharded cluster. If this is set to 1, the cluster is a replica set. If this is set to 2 or higher, the cluster is a sharded cluster with the number of shards specified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_name").ColumnType(schema.ColumnTypeString).Description("Condition in which the API resource finds the cluster when you called the resource. The resource returns one of the following states: IDLE, CREATING, UPDATING, DELETING, DELETED, REPAIRING.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the cluster as it appears in Atlas.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_factor").ColumnType(schema.ColumnTypeInt).Description("Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pit_enabled").ColumnType(schema.ColumnTypeBool).Description("Flag that indicates whether the cluster uses continuous cloud backups. More information is available at https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#continuous-cloud-backups.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique 24-hexadecimal digit string that identifies the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bi_connector").ColumnType(schema.ColumnTypeJSON).Description("Configuration settings applied to BI Connector for Atlas on this cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Description("Collection of key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mongodb_version").ColumnType(schema.ColumnTypeString).Description("Version of MongoDB that the cluster is running, in X.Y.Z format.").
			Extractor(column_value_extractor.StructSelector("MongoDBVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_backup_enabled").ColumnType(schema.ColumnTypeBool).Description("Flag that indicates if the cluster uses Back Up Your Database Deployment for backups.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_specs").ColumnType(schema.ColumnTypeJSON).Description("Configuration for each zone in a Global Cluster. Each object in this array represents a zone where Atlas deploys nodes for your Global Cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_type").ColumnType(schema.ColumnTypeString).Description("Type of the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_size_gb").ColumnType(schema.ColumnTypeFloat).Description("Capacity, in gigabytes, of the host's root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive number.").
			Extractor(column_value_extractor.StructSelector("DiskSizeGB")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_release_system").ColumnType(schema.ColumnTypeString).Description("Release cadence that Atlas uses for this cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling").ColumnType(schema.ColumnTypeJSON).Description("Collection of settings that configures auto-scaling information for the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mongo_uri_with_options").ColumnType(schema.ColumnTypeString).Description("Connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.").Build(),
	}
}

func (x *TableMongodbatlasClusterGenerator) GetSubTables() []*schema.Table {
	return nil
}
