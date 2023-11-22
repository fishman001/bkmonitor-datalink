// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package optionx

import (
	"time"
)

// Options 用于传递动态参数
type Options struct {
	params map[string]interface{}
}

func NewOptions(params map[string]interface{}) *Options {
	if params == nil {
		params = map[string]interface{}{}
	}
	return &Options{params: params}
}

func (o *Options) Get(key string) (interface{}, bool) {
	value, ok := o.params[key]
	return value, ok
}

func (o *Options) GetString(key string) (string, bool) {
	value, ok := o.params[key].(string)
	return value, ok

}

func (o *Options) GetBool(key string) (bool, bool) {
	value, ok := o.params[key].(bool)
	return value, ok

}

func (o *Options) GetUint(key string) (uint, bool) {
	value, ok := o.params[key].(uint)
	return value, ok
}

func (o *Options) GetInt(key string) (int, bool) {
	value, ok := o.params[key].(int)
	return value, ok
}

func (o *Options) GetInt8(key string) (int8, bool) {
	value, ok := o.params[key].(int8)
	return value, ok
}

func (o *Options) GetInt64(key string) (int64, bool) {
	value, ok := o.params[key].(int64)
	return value, ok
}

func (o *Options) GetFloat64(key string) (float64, bool) {
	value, ok := o.params[key].(float64)
	return value, ok
}

func (o *Options) GetTime(key string) (time.Time, bool) {
	value, ok := o.params[key].(time.Time)
	return value, ok
}

func (o *Options) GetDuration(key string) (time.Duration, bool) {
	value, ok := o.params[key].(time.Duration)
	return value, ok
}

func (o *Options) GetStringSlice(key string) ([]string, bool) {
	value, ok := o.params[key].([]string)
	return value, ok
}

func (o *Options) GetStringMap(key string) (map[string]interface{}, bool) {
	value, ok := o.params[key].(map[string]interface{})
	return value, ok
}

func (o *Options) GetStringMapString(key string) (map[string]string, bool) {
	value, ok := o.params[key].(map[string]string)
	return value, ok
}

func (o *Options) GetStringMapStringSlice(key string) (map[string][]string, bool) {
	value, ok := o.params[key].(map[string][]string)
	return value, ok
}

func (o *Options) IsSet(key string) bool {
	_, ok := o.params[key]
	return ok
}

func (o *Options) Set(key string, value interface{}) {
	if o.params == nil {
		o.params = make(map[string]interface{})
	}
	o.params[key] = value
}

func (o *Options) AllKeys() []string {
	var keys []string
	for key := range o.params {
		keys = append(keys, key)
	}
	return keys
}
