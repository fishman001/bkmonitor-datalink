// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package globalconfig

import "github.com/spf13/viper"

const (
	TimeSeriesMetricExpiredDaysPath  = "global_config.time_series_metric_expired_days"  // 自定义指标过期时间
	IsRestrictDsBelongSpacePath      = "global_config.is_restrict_ds_belong_space"      // 是否限制数据源归属具体空间
	DefaultBkdataBizIdPath           = "global_config.default_bkdata_biz_id"            // 接入计算平台使用的业务 ID
	DefaultKafkaStorageClusterIdPath = "global_config.default_kafka_storage_cluster_id" // 默认 kafka 存储集群ID
	BkdataKafkaBrokerUrlPath         = "global_config.bkdata_kafka_broker_url"          // 与计算平台对接的消息队列BROKER地址
	BkappDeployPlatformPath          = "global_config.bkapp_deploy_platform"            // 监控平台版本
	BkdataRtIdPrefixPath             = "global_config.bkdata_rt_id_prefix"              // 监控在计算平台的数据表前缀
	BkdataBkBizIdPath                = "global_config.bkdata_bk_biz_id"                 // 监控在计算平台使用的公共业务ID
	BkdataProjectMaintainerPath      = "global_config.bkdata_project_maintainer"        // 计算平台项目的维护人员
)

func init() {
	viper.SetDefault(TimeSeriesMetricExpiredDaysPath, 30)
	viper.SetDefault(IsRestrictDsBelongSpacePath, true)
	viper.SetDefault(DefaultBkdataBizIdPath, 0)
	viper.SetDefault(BkappDeployPlatformPath, "enterprise")
	viper.SetDefault(BkdataRtIdPrefixPath, viper.GetString(BkappDeployPlatformPath))
	viper.SetDefault(BkdataBkBizIdPath, 2)
	viper.SetDefault(BkdataProjectMaintainerPath, "admin")
}
