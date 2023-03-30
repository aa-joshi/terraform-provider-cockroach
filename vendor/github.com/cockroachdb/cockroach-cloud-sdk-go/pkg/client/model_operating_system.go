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

// OperatingSystem the model 'OperatingSystem'.
type OperatingSystem string

// List of OperatingSystem.
const (
	OPERATINGSYSTEM_MAC     OperatingSystem = "OS_MAC"
	OPERATINGSYSTEM_LINUX   OperatingSystem = "OS_LINUX"
	OPERATINGSYSTEM_WINDOWS OperatingSystem = "OS_WINDOWS"
)

// All allowed values of OperatingSystem enum.
var AllowedOperatingSystemEnumValues = []OperatingSystem{
	"OS_MAC",
	"OS_LINUX",
	"OS_WINDOWS",
}

func (v *OperatingSystem) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperatingSystem(value)
	for _, existing := range AllowedOperatingSystemEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperatingSystem", value)
}

// NewOperatingSystemFromValue returns a pointer to a valid OperatingSystem
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOperatingSystemFromValue(v string) (*OperatingSystem, error) {
	ev := OperatingSystem(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperatingSystem: valid values are %v", v, AllowedOperatingSystemEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OperatingSystem) IsValid() bool {
	for _, existing := range AllowedOperatingSystemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperatingSystem value.
func (v OperatingSystem) Ptr() *OperatingSystem {
	return &v
}
