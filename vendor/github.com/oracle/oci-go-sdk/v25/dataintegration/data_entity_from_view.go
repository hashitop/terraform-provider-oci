// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// DataEntityFromView The view entity data entity details.
type DataEntityFromView struct {
	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key for the object
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The shape ID.
	ShapeId *string `mandatory:"false" json:"shapeId"`

	Types *TypeLibrary `mandatory:"false" json:"types"`

	// Specifies other type label.
	OtherTypeLabel *string `mandatory:"false" json:"otherTypeLabel"`

	// An array of unique keys.
	UniqueKeys []UniqueKey `mandatory:"false" json:"uniqueKeys"`

	// An array of foreign keys.
	ForeignKeys []ForeignKey `mandatory:"false" json:"foreignKeys"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The entity type.
	EntityType DataEntityFromViewEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

//GetMetadata returns Metadata
func (m DataEntityFromView) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataEntityFromView) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DataEntityFromView) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromView DataEntityFromView
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromView
	}{
		"VIEW_ENTITY",
		(MarshalTypeDataEntityFromView)(m),
	}

	return json.Marshal(&s)
}

// DataEntityFromViewEntityTypeEnum Enum with underlying type: string
type DataEntityFromViewEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromViewEntityTypeEnum
const (
	DataEntityFromViewEntityTypeTable  DataEntityFromViewEntityTypeEnum = "TABLE"
	DataEntityFromViewEntityTypeView   DataEntityFromViewEntityTypeEnum = "VIEW"
	DataEntityFromViewEntityTypeFile   DataEntityFromViewEntityTypeEnum = "FILE"
	DataEntityFromViewEntityTypeQueue  DataEntityFromViewEntityTypeEnum = "QUEUE"
	DataEntityFromViewEntityTypeStream DataEntityFromViewEntityTypeEnum = "STREAM"
	DataEntityFromViewEntityTypeOther  DataEntityFromViewEntityTypeEnum = "OTHER"
)

var mappingDataEntityFromViewEntityType = map[string]DataEntityFromViewEntityTypeEnum{
	"TABLE":  DataEntityFromViewEntityTypeTable,
	"VIEW":   DataEntityFromViewEntityTypeView,
	"FILE":   DataEntityFromViewEntityTypeFile,
	"QUEUE":  DataEntityFromViewEntityTypeQueue,
	"STREAM": DataEntityFromViewEntityTypeStream,
	"OTHER":  DataEntityFromViewEntityTypeOther,
}

// GetDataEntityFromViewEntityTypeEnumValues Enumerates the set of values for DataEntityFromViewEntityTypeEnum
func GetDataEntityFromViewEntityTypeEnumValues() []DataEntityFromViewEntityTypeEnum {
	values := make([]DataEntityFromViewEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromViewEntityType {
		values = append(values, v)
	}
	return values
}