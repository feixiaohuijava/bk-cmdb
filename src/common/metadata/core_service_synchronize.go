/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017,-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metadata

import (
	"configcenter/src/common/mapstr"
)

// SynchronizeOperateType synchronize data operate type
type SynchronizeOperateType int64

const (
	// SynchronizeOperateTypeAdd synchronize data add
	SynchronizeOperateTypeAdd SynchronizeOperateType = iota + 1
	// SynchronizeOperateTypeUpdate synchronize data update
	SynchronizeOperateTypeUpdate
	// SynchronizeOperateTypeRepalce synchronize data add or update
	SynchronizeOperateTypeRepalce
	// SynchronizeOperateTypeDelete synchronize data delete
	SynchronizeOperateTypeDelete
)

// SynchronizeOperateDataType synchronize data operate type
type SynchronizeOperateDataType int64

const (
	// SynchronizeOperateDataTypeInstance synchronize data is instance
	SynchronizeOperateDataTypeInstance SynchronizeOperateDataType = iota + 1
	// SynchronizeOperateDataTypeModel synchronize data is model
	SynchronizeOperateDataTypeModel
	//SynchronizeOperateDataTypeAssociation synchronize data is association
	SynchronizeOperateDataTypeAssociation
)

// SynchronizeDataInfo synchronize instance data http request parameter
type SynchronizeDataInfo struct {
	OperateDataType SynchronizeOperateDataType `json:"operate_data_type"`
	// OperateDataType = SynchronizeOperateDataTypeInstance,
	// DataClassify = object_id,  eg:host,plat,module,proc etc.
	// OperateDataType = SynchronizeOperateDataTypeModel,
	// DataClassify = common.SynchronizeModelDescTypeGroupInfo,common.SynchronizeModelDescTypeModuleAttribute etc
	// OperateDataType = SynchronizeOperateDataTypeAssociation
	// DataClassify = common.SynchronizeAssociationTypeModelHost etc.
	DataClassify string          `json:"data_classify"`
	InfoArray    []mapstr.MapStr `json:"instance_info_array"`
	// OffSet current data offset  start location
	Offset int64 `json:"offset"`
	// Count total data count
	Count           int64  `json:"count"`
	Version         string `json:"version"`
	SynchronizeFlag string `json:"synchronize_flag"`
}

// SynchronizeParameter synchronize instance data http request parameter
type SynchronizeParameter struct {
	OperateType SynchronizeOperateType `json:"op_type"`
	// synchronize data type
	OperateDataType SynchronizeOperateDataType `json:"operate_data_type"`
	// OperateDataType = SynchronizeOperateDataTypeInstance,
	// DataClassify = object_id,  eg:host,plat,module,proc etc.
	// OperateDataType = SynchronizeOperateDataTypeModel,
	// DataClassify = common.SynchronizeModelDescTypeGroupInfo,common.SynchronizeModelDescTypeModuleAttribute etc
	// OperateDataType = SynchronizeOperateDataTypeAssociation
	// DataClassify = common.SynchronizeAssociationTypeModelHost etc.
	DataClassify    string             `json:"data_classify"`
	InfoArray       []*SynchronizeItem `json:"instance_info_array"`
	Version         string             `json:"version"`
	SynchronizeFlag string             `json:"synchronize_flag"`
}

// SynchronizeItem synchronize data information
type SynchronizeItem struct {
	Info mapstr.MapStr `json:"info"`
	ID   int64         `json:"id"`
}

// SynchronizeFetchInfoParameter synchronize  data fetch data http request parameter
type SynchronizeFetchInfoParameter struct {
	DataType     SynchronizeOperateDataType `json:"data_type"`
	DataClassify string                     `json:"data_classify"`
	Condition    mapstr.MapStr              `json:"condition"`
	Start        uint64                     `json:"start"`
	Limit        uint64                     `json:"limit"`
}

// SynchronizeResult synchronize result
type SynchronizeResult struct {
	BaseResp `json:",inline"`
	Data     SetDataResult `json:"data"`
}

// SynchronizeDataResult common Synchronize result definition
type SynchronizeDataResult struct {
	//Created    []CreatedDataResult `json:"created"`
	//Updated    []UpdatedDataResult `json:"updated"`
	Exceptions []ExceptionResult `json:"exception"`
}
