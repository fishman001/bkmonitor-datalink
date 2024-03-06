// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/script/diff/redis"
	redisUtils "github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/register/redis"
)

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "diff redis data",
	Long:  "diff redis data",
	Run:   startRedisDiff,
}

func init() {
	redisCmd.Flags().String("srcKey", "", "src redis key")
	redisCmd.Flags().String("bypassKey", "", "bypass redis key")
	redisCmd.Flags().String("keyType", "", "key type [string, hash, list, set]")
	rootCmd.AddCommand(redisCmd)
}

func startRedisDiff(cmd *cobra.Command, args []string) {
	InitConfig()
	initRedisDiffConfig()

	srcKey, _ := cmd.Flags().GetString("srcKey")
	if srcKey != "" {
		// 命令行参数覆盖配置文件参数
		viper.Set("diffRedis.srcKey", srcKey)
	} else {
		srcKey = viper.GetString("diffRedis.srcKey")
		if srcKey == "" {
			fmt.Printf("srcKey can not be empty")
			os.Exit(1)
		}
	}

	bypassKey, _ := cmd.Flags().GetString("bypassKey")
	if bypassKey != "" {
		// 命令行参数覆盖配置文件参数
		viper.SetDefault("diffRedis.bypassKey", bypassKey)
	} else {
		bypassKey = viper.GetString("diffRedis.bypassKey")
		if bypassKey == "" {
			fmt.Printf("bypassKey can not be empty")
			os.Exit(1)
		}
	}

	if keyType, _ := cmd.Flags().GetString("keyType"); keyType != "" {
		// 命令行参数覆盖配置文件参数
		viper.Set("diffRedis.keyType", keyType)
	} else {
		// 命令行参数未配置则使用配置文件参数，默认为hash
		viper.SetDefault("diffRedis.keyType", "hash")
	}

	du := redis.DiffUtil{
		KeyType:      viper.GetString("diffRedis.keyType"),
		SrcKey:       srcKey,
		BypassKey:    bypassKey,
		SrcConfig:    getRedisClientOpt("srcRedis"),
		BypassConfig: getRedisClientOpt("bypassRedis"),
	}
	equal, err := du.Diff()
	if err != nil {
		fmt.Printf("diff key [%s] and [%s] %s data failed, %v", du.SrcKey, du.BypassKey, du.KeyType, err)
		os.Exit(1)
	}
	if equal {
		fmt.Printf("key [%s] and [%s] %s data is equal", du.SrcKey, du.BypassKey, du.KeyType)
	} else {
		fmt.Printf("key [%s] and [%s] %s data is different", du.SrcKey, du.BypassKey, du.KeyType)
	}

}

func initRedisDiffConfig() {
	viper.SetDefault("diffRedis.srcRedis.mod", "standalone")
	viper.SetDefault("diffRedis.srcRedis.sentinel.masterName", "standalone")
	viper.SetDefault("diffRedis.srcRedis.sentinel.address", []string{"127.0.0.1"})
	viper.SetDefault("diffRedis.srcRedis.sentinel.password", "")
	viper.SetDefault("diffRedis.srcRedis.standalone.host", "127.0.0.1")
	viper.SetDefault("diffRedis.srcRedis.standalone.port", 6379)
	viper.SetDefault("diffRedis.srcRedis.standalone.password", "")
	viper.SetDefault("diffRedis.srcRedis.db", 0)
	viper.SetDefault("diffRedis.srcRedis.dialTimeout", 10*time.Second)
	viper.SetDefault("diffRedis.srcRedis.readTimeout", 10*time.Second)

	viper.SetDefault("diffRedis.bypassRedis.mod", "standalone")
	viper.SetDefault("diffRedis.bypassRedis.sentinel.masterName", "standalone")
	viper.SetDefault("diffRedis.bypassRedis.sentinel.address", []string{"127.0.0.1"})
	viper.SetDefault("diffRedis.bypassRedis.sentinel.password", "")
	viper.SetDefault("diffRedis.bypassRedis.standalone.host", "127.0.0.1")
	viper.SetDefault("diffRedis.bypassRedis.standalone.port", 6379)
	viper.SetDefault("diffRedis.bypassRedis.standalone.password", "")
	viper.SetDefault("diffRedis.bypassRedis.db", 0)
	viper.SetDefault("diffRedis.bypassRedis.dialTimeout", 10*time.Second)
	viper.SetDefault("diffRedis.bypassRedis.readTimeout", 10*time.Second)
}

func getRedisClientOpt(redisName string) *redisUtils.Option {
	return &redisUtils.Option{
		Mode:             viper.GetString(fmt.Sprintf("diffRedis.%s.mod", redisName)),
		Host:             viper.GetString(fmt.Sprintf("diffRedis.%s.standalone.host", redisName)),
		Port:             viper.GetInt(fmt.Sprintf("diffRedis.%s.standalone.port", redisName)),
		SentinelAddress:  viper.GetStringSlice(fmt.Sprintf("diffRedis.%s.sentinel.address", redisName)),
		MasterName:       viper.GetString(fmt.Sprintf("diffRedis.%s.sentinel.masterName", redisName)),
		SentinelPassword: viper.GetString(fmt.Sprintf("diffRedis.%s.sentinel.password", redisName)),
		Password:         viper.GetString(fmt.Sprintf("diffRedis.%s.standalone.password", redisName)),
		Db:               viper.GetInt(fmt.Sprintf("diffRedis.%s.db", redisName)),
		DialTimeout:      viper.GetDuration(fmt.Sprintf("diffRedis.%s.dialTimeout", redisName)),
		ReadTimeout:      viper.GetDuration(fmt.Sprintf("diffRedis.%s.readTimeout", redisName)),
	}
}
