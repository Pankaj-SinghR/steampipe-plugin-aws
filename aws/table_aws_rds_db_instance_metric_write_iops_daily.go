package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRdsInstanceMetricWriteIopsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_instance_metric_write_iops_daily",
		Description: "AWS RDS DB Instance Cloudwatch Metrics - Write IOPS (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listRDSDBInstances,
			Hydrate:       listRdsInstanceMetricWriteIopsDaily,
		},
		GetMatrixItemFunc: CloudWatchRegionsMatrix,
		Columns: awsRegionalColumns(cwMetricColumns(
			[]*plugin.Column{
				{
					Name:        "db_instance_identifier",
					Description: "The friendly name to identify the DB Instance.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			})),
	}
}

func listRdsInstanceMetricWriteIopsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	instance := h.Item.(types.DBInstance)
	return listCWMetricStatistics(ctx, d, "DAILY", "AWS/RDS", "ReadIOPS", "DBInstanceIdentifier", *instance.DBInstanceIdentifier)
}
