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
)

// Resource struct for Resource.
type Resource struct {
	Id   *string      `json:"id,omitempty"`
	Type ResourceType `json:"type"`
}

// NewResource instantiates a new Resource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(type_ ResourceType) *Resource {
	p := Resource{}
	p.Type = type_
	return &p
}

// NewResourceWithDefaults instantiates a new Resource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	p := Resource{}
	return &p
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Resource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Resource) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *Resource) GetType() ResourceType {
	if o == nil {
		var ret ResourceType
		return ret
	}

	return o.Type
}

// SetType sets field value.
func (o *Resource) SetType(v ResourceType) {
	o.Type = v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}
