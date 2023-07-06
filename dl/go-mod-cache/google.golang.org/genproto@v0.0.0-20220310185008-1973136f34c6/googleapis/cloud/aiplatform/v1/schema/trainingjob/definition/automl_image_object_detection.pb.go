// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1/schema/trainingjob/definition/automl_image_object_detection.proto

package definition

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AutoMlImageObjectDetectionInputs_ModelType int32

const (
	// Should not be set.
	AutoMlImageObjectDetectionInputs_MODEL_TYPE_UNSPECIFIED AutoMlImageObjectDetectionInputs_ModelType = 0
	// A model best tailored to be used within Google Cloud, and which cannot
	// be exported. Expected to have a higher latency, but should also have a
	// higher prediction quality than other cloud models.
	AutoMlImageObjectDetectionInputs_CLOUD_HIGH_ACCURACY_1 AutoMlImageObjectDetectionInputs_ModelType = 1
	// A model best tailored to be used within Google Cloud, and which cannot
	// be exported. Expected to have a low latency, but may have lower
	// prediction quality than other cloud models.
	AutoMlImageObjectDetectionInputs_CLOUD_LOW_LATENCY_1 AutoMlImageObjectDetectionInputs_ModelType = 2
	// A model that, in addition to being available within Google
	// Cloud can also be exported (see ModelService.ExportModel) and
	// used on a mobile or edge device with TensorFlow afterwards.
	// Expected to have low latency, but may have lower prediction
	// quality than other mobile models.
	AutoMlImageObjectDetectionInputs_MOBILE_TF_LOW_LATENCY_1 AutoMlImageObjectDetectionInputs_ModelType = 3
	// A model that, in addition to being available within Google
	// Cloud can also be exported (see ModelService.ExportModel) and
	// used on a mobile or edge device with TensorFlow afterwards.
	AutoMlImageObjectDetectionInputs_MOBILE_TF_VERSATILE_1 AutoMlImageObjectDetectionInputs_ModelType = 4
	// A model that, in addition to being available within Google
	// Cloud, can also be exported (see ModelService.ExportModel) and
	// used on a mobile or edge device with TensorFlow afterwards.
	// Expected to have a higher latency, but should also have a higher
	// prediction quality than other mobile models.
	AutoMlImageObjectDetectionInputs_MOBILE_TF_HIGH_ACCURACY_1 AutoMlImageObjectDetectionInputs_ModelType = 5
)

// Enum value maps for AutoMlImageObjectDetectionInputs_ModelType.
var (
	AutoMlImageObjectDetectionInputs_ModelType_name = map[int32]string{
		0: "MODEL_TYPE_UNSPECIFIED",
		1: "CLOUD_HIGH_ACCURACY_1",
		2: "CLOUD_LOW_LATENCY_1",
		3: "MOBILE_TF_LOW_LATENCY_1",
		4: "MOBILE_TF_VERSATILE_1",
		5: "MOBILE_TF_HIGH_ACCURACY_1",
	}
	AutoMlImageObjectDetectionInputs_ModelType_value = map[string]int32{
		"MODEL_TYPE_UNSPECIFIED":    0,
		"CLOUD_HIGH_ACCURACY_1":     1,
		"CLOUD_LOW_LATENCY_1":       2,
		"MOBILE_TF_LOW_LATENCY_1":   3,
		"MOBILE_TF_VERSATILE_1":     4,
		"MOBILE_TF_HIGH_ACCURACY_1": 5,
	}
)

func (x AutoMlImageObjectDetectionInputs_ModelType) Enum() *AutoMlImageObjectDetectionInputs_ModelType {
	p := new(AutoMlImageObjectDetectionInputs_ModelType)
	*p = x
	return p
}

func (x AutoMlImageObjectDetectionInputs_ModelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoMlImageObjectDetectionInputs_ModelType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_enumTypes[0].Descriptor()
}

func (AutoMlImageObjectDetectionInputs_ModelType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_enumTypes[0]
}

func (x AutoMlImageObjectDetectionInputs_ModelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoMlImageObjectDetectionInputs_ModelType.Descriptor instead.
func (AutoMlImageObjectDetectionInputs_ModelType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescGZIP(), []int{1, 0}
}

type AutoMlImageObjectDetectionMetadata_SuccessfulStopReason int32

const (
	// Should not be set.
	AutoMlImageObjectDetectionMetadata_SUCCESSFUL_STOP_REASON_UNSPECIFIED AutoMlImageObjectDetectionMetadata_SuccessfulStopReason = 0
	// The inputs.budgetMilliNodeHours had been reached.
	AutoMlImageObjectDetectionMetadata_BUDGET_REACHED AutoMlImageObjectDetectionMetadata_SuccessfulStopReason = 1
	// Further training of the Model ceased to increase its quality, since it
	// already has converged.
	AutoMlImageObjectDetectionMetadata_MODEL_CONVERGED AutoMlImageObjectDetectionMetadata_SuccessfulStopReason = 2
)

// Enum value maps for AutoMlImageObjectDetectionMetadata_SuccessfulStopReason.
var (
	AutoMlImageObjectDetectionMetadata_SuccessfulStopReason_name = map[int32]string{
		0: "SUCCESSFUL_STOP_REASON_UNSPECIFIED",
		1: "BUDGET_REACHED",
		2: "MODEL_CONVERGED",
	}
	AutoMlImageObjectDetectionMetadata_SuccessfulStopReason_value = map[string]int32{
		"SUCCESSFUL_STOP_REASON_UNSPECIFIED": 0,
		"BUDGET_REACHED":                     1,
		"MODEL_CONVERGED":                    2,
	}
)

func (x AutoMlImageObjectDetectionMetadata_SuccessfulStopReason) Enum() *AutoMlImageObjectDetectionMetadata_SuccessfulStopReason {
	p := new(AutoMlImageObjectDetectionMetadata_SuccessfulStopReason)
	*p = x
	return p
}

func (x AutoMlImageObjectDetectionMetadata_SuccessfulStopReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoMlImageObjectDetectionMetadata_SuccessfulStopReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_enumTypes[1].Descriptor()
}

func (AutoMlImageObjectDetectionMetadata_SuccessfulStopReason) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_enumTypes[1]
}

func (x AutoMlImageObjectDetectionMetadata_SuccessfulStopReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoMlImageObjectDetectionMetadata_SuccessfulStopReason.Descriptor instead.
func (AutoMlImageObjectDetectionMetadata_SuccessfulStopReason) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescGZIP(), []int{2, 0}
}

// A TrainingJob that trains and uploads an AutoML Image Object Detection Model.
type AutoMlImageObjectDetection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The input parameters of this TrainingJob.
	Inputs *AutoMlImageObjectDetectionInputs `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// The metadata information
	Metadata *AutoMlImageObjectDetectionMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *AutoMlImageObjectDetection) Reset() {
	*x = AutoMlImageObjectDetection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoMlImageObjectDetection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoMlImageObjectDetection) ProtoMessage() {}

func (x *AutoMlImageObjectDetection) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoMlImageObjectDetection.ProtoReflect.Descriptor instead.
func (*AutoMlImageObjectDetection) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescGZIP(), []int{0}
}

func (x *AutoMlImageObjectDetection) GetInputs() *AutoMlImageObjectDetectionInputs {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *AutoMlImageObjectDetection) GetMetadata() *AutoMlImageObjectDetectionMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type AutoMlImageObjectDetectionInputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelType AutoMlImageObjectDetectionInputs_ModelType `protobuf:"varint,1,opt,name=model_type,json=modelType,proto3,enum=google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionInputs_ModelType" json:"model_type,omitempty"`
	// The training budget of creating this model, expressed in milli node
	// hours i.e. 1,000 value in this field means 1 node hour. The actual
	// metadata.costMilliNodeHours will be equal or less than this value.
	// If further model training ceases to provide any improvements, it will
	// stop without using the full budget and the metadata.successfulStopReason
	// will be `model-converged`.
	// Note, node_hour  = actual_hour * number_of_nodes_involved.
	// For modelType `cloud`(default), the budget must be between 20,000
	// and 900,000 milli node hours, inclusive. The default value is 216,000
	// which represents one day in wall time, considering 9 nodes are used.
	// For model types `mobile-tf-low-latency-1`, `mobile-tf-versatile-1`,
	// `mobile-tf-high-accuracy-1`
	// the training budget must be between 1,000 and 100,000 milli node hours,
	// inclusive. The default value is 24,000 which represents one day in
	// wall time on a single node that is used.
	BudgetMilliNodeHours int64 `protobuf:"varint,2,opt,name=budget_milli_node_hours,json=budgetMilliNodeHours,proto3" json:"budget_milli_node_hours,omitempty"`
	// Use the entire training budget. This disables the early stopping feature.
	// When false the early stopping feature is enabled, which means that AutoML
	// Image Object Detection might stop training before the entire training
	// budget has been used.
	DisableEarlyStopping bool `protobuf:"varint,3,opt,name=disable_early_stopping,json=disableEarlyStopping,proto3" json:"disable_early_stopping,omitempty"`
}

func (x *AutoMlImageObjectDetectionInputs) Reset() {
	*x = AutoMlImageObjectDetectionInputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoMlImageObjectDetectionInputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoMlImageObjectDetectionInputs) ProtoMessage() {}

func (x *AutoMlImageObjectDetectionInputs) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoMlImageObjectDetectionInputs.ProtoReflect.Descriptor instead.
func (*AutoMlImageObjectDetectionInputs) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescGZIP(), []int{1}
}

func (x *AutoMlImageObjectDetectionInputs) GetModelType() AutoMlImageObjectDetectionInputs_ModelType {
	if x != nil {
		return x.ModelType
	}
	return AutoMlImageObjectDetectionInputs_MODEL_TYPE_UNSPECIFIED
}

func (x *AutoMlImageObjectDetectionInputs) GetBudgetMilliNodeHours() int64 {
	if x != nil {
		return x.BudgetMilliNodeHours
	}
	return 0
}

func (x *AutoMlImageObjectDetectionInputs) GetDisableEarlyStopping() bool {
	if x != nil {
		return x.DisableEarlyStopping
	}
	return false
}

type AutoMlImageObjectDetectionMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual training cost of creating this model, expressed in
	// milli node hours, i.e. 1,000 value in this field means 1 node hour.
	// Guaranteed to not exceed inputs.budgetMilliNodeHours.
	CostMilliNodeHours int64 `protobuf:"varint,1,opt,name=cost_milli_node_hours,json=costMilliNodeHours,proto3" json:"cost_milli_node_hours,omitempty"`
	// For successful job completions, this is the reason why the job has
	// finished.
	SuccessfulStopReason AutoMlImageObjectDetectionMetadata_SuccessfulStopReason `protobuf:"varint,2,opt,name=successful_stop_reason,json=successfulStopReason,proto3,enum=google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionMetadata_SuccessfulStopReason" json:"successful_stop_reason,omitempty"`
}

func (x *AutoMlImageObjectDetectionMetadata) Reset() {
	*x = AutoMlImageObjectDetectionMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoMlImageObjectDetectionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoMlImageObjectDetectionMetadata) ProtoMessage() {}

func (x *AutoMlImageObjectDetectionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoMlImageObjectDetectionMetadata.ProtoReflect.Descriptor instead.
func (*AutoMlImageObjectDetectionMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescGZIP(), []int{2}
}

func (x *AutoMlImageObjectDetectionMetadata) GetCostMilliNodeHours() int64 {
	if x != nil {
		return x.CostMilliNodeHours
	}
	return 0
}

func (x *AutoMlImageObjectDetectionMetadata) GetSuccessfulStopReason() AutoMlImageObjectDetectionMetadata_SuccessfulStopReason {
	if x != nil {
		return x.SuccessfulStopReason
	}
	return AutoMlImageObjectDetectionMetadata_SUCCESSFUL_STOP_REASON_UNSPECIFIED
}

var File_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x1a, 0x41, 0x75, 0x74, 0x6f, 0x4d,
	0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x78, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xca, 0x03, 0x0a, 0x20, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x17, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x14, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x4e, 0x6f, 0x64, 0x65,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x61,
	0x72, 0x6c, 0x79, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0xb2, 0x01, 0x0a, 0x09,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x44,
	0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x5f, 0x41, 0x43, 0x43, 0x55, 0x52, 0x41, 0x43, 0x59, 0x5f, 0x31, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x4c, 0x41,
	0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x42,
	0x49, 0x4c, 0x45, 0x5f, 0x54, 0x46, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x4e,
	0x43, 0x59, 0x5f, 0x31, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45,
	0x5f, 0x54, 0x46, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x31, 0x10,
	0x04, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x46, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x5f, 0x41, 0x43, 0x43, 0x55, 0x52, 0x41, 0x43, 0x59, 0x5f, 0x31, 0x10, 0x05,
	0x22, 0xea, 0x02, 0x0a, 0x22, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x6f, 0x73, 0x74, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0xa7, 0x01, 0x0a, 0x16, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x71, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x14,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x14, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66,
	0x75, 0x6c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x22,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f,
	0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x52,
	0x45, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x44, 0x45,
	0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x47, 0x45, 0x44, 0x10, 0x02, 0x42, 0xfc, 0x02,
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x6a, 0x6f, 0x62, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1f,
	0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x62, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x6a, 0x6f, 0x62, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x38, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x56, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0xca, 0x02, 0x38, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62,
	0x5c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0xea, 0x02, 0x3e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x3a, 0x3a, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62,
	0x3a, 0x3a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_goTypes = []interface{}{
	(AutoMlImageObjectDetectionInputs_ModelType)(0),              // 0: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionInputs.ModelType
	(AutoMlImageObjectDetectionMetadata_SuccessfulStopReason)(0), // 1: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionMetadata.SuccessfulStopReason
	(*AutoMlImageObjectDetection)(nil),                           // 2: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetection
	(*AutoMlImageObjectDetectionInputs)(nil),                     // 3: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionInputs
	(*AutoMlImageObjectDetectionMetadata)(nil),                   // 4: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionMetadata
}
var file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_depIdxs = []int32{
	3, // 0: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetection.inputs:type_name -> google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionInputs
	4, // 1: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetection.metadata:type_name -> google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionMetadata
	0, // 2: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionInputs.model_type:type_name -> google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionInputs.ModelType
	1, // 3: google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionMetadata.successful_stop_reason:type_name -> google.cloud.aiplatform.v1.schema.trainingjob.definition.AutoMlImageObjectDetectionMetadata.SuccessfulStopReason
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_init()
}
func file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoMlImageObjectDetection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoMlImageObjectDetectionInputs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoMlImageObjectDetectionMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto = out.File
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_trainingjob_definition_automl_image_object_detection_proto_depIdxs = nil
}