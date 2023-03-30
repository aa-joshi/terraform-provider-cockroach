// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2022-09-20

package client

import (
	"encoding/json"
	"fmt"
)

// ClusterStateType  - LOCKED: An exclusive operation is being performed on this cluster. Other operations should not proceed if they did not set a cluster into the LOCKED state.
type ClusterStateType string

// List of ClusterStateType.
const (
	CLUSTERSTATETYPE_CREATING        ClusterStateType = "CREATING"
	CLUSTERSTATETYPE_CREATED         ClusterStateType = "CREATED"
	CLUSTERSTATETYPE_CREATION_FAILED ClusterStateType = "CREATION_FAILED"
	CLUSTERSTATETYPE_DELETED         ClusterStateType = "DELETED"
	CLUSTERSTATETYPE_LOCKED          ClusterStateType = "LOCKED"
)

// All allowed values of ClusterStateType enum.
var AllowedClusterStateTypeEnumValues = []ClusterStateType{
	"CREATING",
	"CREATED",
	"CREATION_FAILED",
	"DELETED",
	"LOCKED",
}

func (v *ClusterStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterStateType(value)
	for _, existing := range AllowedClusterStateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterStateType", value)
}

// NewClusterStateTypeFromValue returns a pointer to a valid ClusterStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterStateTypeFromValue(v string) (*ClusterStateType, error) {
	ev := ClusterStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterStateType: valid values are %v", v, AllowedClusterStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterStateType) IsValid() bool {
	for _, existing := range AllowedClusterStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterStateType value.
func (v ClusterStateType) Ptr() *ClusterStateType {
	return &v
}
