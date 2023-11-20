// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package service

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/api"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/api/define"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/globalconfig"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/metadata/models/resulttable"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/metadata/models/storage"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/store/mysql"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

// BkDataStorageSvc bkdata storage service
type BkDataStorageSvc struct {
	*storage.BkDataStorage
}

func NewBkDataStorageSvc(obj *storage.BkDataStorage) BkDataStorageSvc {
	return BkDataStorageSvc{
		BkDataStorage: obj,
	}
}

func (s BkDataStorageSvc) CreateDatabusClean(rt *resulttable.ResultTable) error {
	var kafkaStorage storage.KafkaStorage
	if err := storage.NewKafkaStorageQuerySet(mysql.GetDBSession().DB.New()).TableIDEq(rt.TableId).One(&kafkaStorage); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logger.Errorf("result table [%s] data not write into mq", rt.TableId)
		}
		return err
	}
	// 增加接入部署计划
	svc := NewKafkaStorageSvc(&kafkaStorage)
	storageCluster, err := svc.StorageCluster()
	if err != nil {
		return err
	}
	consulConfig := NewClusterInfoSvc(storageCluster).ConsulConfig()
	domain := consulConfig.ClusterConfig.DomainName
	port := consulConfig.ClusterConfig.Port
	// kafka broker_url 以实际配置为准，如果没有配置，再使用默认的 broker url
	brokerUrl := viper.GetString(globalconfig.BkdataKafkaBrokerUrlPath)
	if domain != "" && port != 0 {
		brokerUrl = fmt.Sprintf("%s:%v", domain, port)
	}
	isSasl := consulConfig.ClusterConfig.IsSslVerify
	user := consulConfig.AuthInfo.Username
	passwd := consulConfig.AuthInfo.Password
	// 采用结果表区分消费组
	KafkaConsumerGroupName := GenBkdataRtIdWithoutBizId(rt.TableId)
	// 计算平台要求，raw_data_name不能超过50个字符
	rtId := strings.ReplaceAll(rt.TableId, ".", "__")
	index := len(rtId) - 50
	if index < 0 {
		index = 0
	}
	rtId = rtId[index:]
	rawDataName := fmt.Sprintf("%s_%s", viper.GetString(globalconfig.BkdataRtIdPrefixPath), rtId)

	params := map[string]interface{}{
		"data_scenario": "queue",
		"bk_biz_id":     viper.GetInt(globalconfig.BkdataBkBizIdPath),
		"description":   "",
		"access_raw_data": map[string]interface{}{
			"raw_data_name":    rawDataName,
			"maintainer":       viper.GetInt(globalconfig.BkdataProjectMaintainerPath),
			"raw_data_alias":   rt.TableNameZh,
			"data_source":      "kafka",
			"data_encoding":    "UTF-8",
			"sensitivity":      "private",
			"description":      fmt.Sprintf("接入配置 (%s)", rt.TableNameZh),
			"tags":             []interface{}{},
			"data_source_tags": []string{"src_kafka"},
		},
		"access_conf_info": map[string]interface{}{
			"collection_model": map[string]interface{}{"collection_type": "incr", "start_at": 1, "period": "-1"},
			"resource": map[string]interface{}{
				"type": "kafka",
				"scope": []map[string]interface{}{
					{
						"master":            brokerUrl,
						"group":             KafkaConsumerGroupName,
						"topic":             svc.Topic,
						"tasks":             svc.Partition,
						"use_sasl":          isSasl,
						"security_protocol": "SASL_PLAINTEXT",
						"sasl_mechanism":    "SCRAM-SHA-512",
						"user":              user,
						"password":          passwd,
					},
				},
			},
		},
	}
	bkdataApi, err := api.GetBkdataApi()
	if err != nil {
		return err
	}
	var resp define.APICommonMapResp
	if _, err := bkdataApi.AccessDeployPlan().SetBody(params).SetResult(&resp).Request(); err != nil {
		return errors.Wrapf(err, "access to bkdata failed, params [%#v]", params)
	}
	logger.Infof("access to bkdata, result [%#v]", resp)
	id, ok := resp.Data["raw_data_id"].(float64)
	if !ok {
		return fmt.Errorf("parse AccessDeployPlan resp error, %v", err)
	}
	s.RawDataID = int(id)
	if err := s.Update(mysql.GetDBSession().DB, storage.BkDataStorageDBSchema.RawDataID); err != nil {
		return err
	}
	return nil
}

func GenBkdataRtIdWithoutBizId(tableId string) string {
	tableIdConv := strings.ReplaceAll(tableId, ".", "_")
	index := len(tableIdConv) - 32
	if index < 0 {
		index = 0
	}
	tableIdConv = tableIdConv[index:]
	rtId := strings.ToLower(fmt.Sprintf("%s_%s", viper.GetString(globalconfig.BkdataRtIdPrefixPath), tableIdConv))
	return strings.TrimLeft(rtId, "_")
}