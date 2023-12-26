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
	"context"
	"strings"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/metadata/models"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/metadata/models/storage"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/store/elasticsearch"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/store/mysql"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/utils/mocker"
)

func TestEsSnapshotRestoreSvc_GetCompleteDocCount(t *testing.T) {
	mocker.InitTestDBConfig("../../../bmw_test.yaml")
	db := mysql.GetDBSession().DB
	restore := storage.EsSnapshotRestore{
		TableID:          "restore_test_table_id",
		StartTime:        time.Now(),
		EndTime:          time.Now(),
		ExpiredTime:      time.Now(),
		Indices:          "index1,index2,index3",
		CompleteDocCount: 5,
		TotalDocCount:    20,
		CreateTime:       time.Now().Add(-time.Hour),
		LastModifyTime:   time.Now(),
	}
	db.Delete(&restore, "table_id = ?", restore.TableID)
	err := restore.Create(db)
	assert.NoError(t, err)

	cluster := storage.ClusterInfo{
		ClusterName:      "es_test_default",
		ClusterType:      models.StorageTypeES,
		DomainName:       "127.0.0.1",
		Port:             9200,
		IsDefaultCluster: true,
		Schema:           "http",
		Version:          "7",
		RegisteredSystem: "bkmonitor",
	}
	db.Delete(&cluster, "cluster_name = ?", cluster.ClusterName)
	err = cluster.Create(db)
	assert.NoError(t, err)

	ess := storage.ESStorage{
		TableID:          restore.TableID,
		StorageClusterID: cluster.ClusterID,
	}
	db.Delete(&ess, "table_id = ?", ess.TableID)
	err = ess.Create(db)
	assert.NoError(t, err)
	gomonkey.ApplyFunc(elasticsearch.Elasticsearch.CatIndices, func(es elasticsearch.Elasticsearch, ctx context.Context, indices []string, format string) (*elasticsearch.Response, error) {
		reader := strings.NewReader(`[{"health":"yellow","status":"open","index":"restore_index1","uuid":"xxx","pri":"1","rep":"1","docs.count":"6","docs.deleted":"0","store.size":"7.9kb","pri.store.size":"7.9kb"},{"health":"yellow","status":"open","index":"restore_index2","uuid":"xx","pri":"1","rep":"1","docs.count":"3","docs.deleted":"0","store.size":"7.9kb","pri.store.size":"7.9kb"},{"health":"yellow","status":"open","index":"restore_index4","uuid":"x","pri":"1","rep":"1","docs.count":"3","docs.deleted":"0","store.size":"7.9kb","pri.store.size":"7.9kb"}]`)
		return &elasticsearch.Response{StatusCode: 200, Body: ioutils.NewReadCloserWrapper(reader, func() error { return nil })}, nil
	})
	svc := NewEsSnapshotRestoreSvc(&restore)
	count, err := svc.GetCompleteDocCount(context.Background())
	assert.NoError(t, err)
	assert.True(t, assert.Equal(t, 6+3, count))
}
