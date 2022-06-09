// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package exampleoc is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by ygnmi version: (devel): (ygot: v0.23.0)
using the following YANG input files:
	- ../../pathgen/testdata/yang/openconfig-simple.yang
	- ../../pathgen/testdata/yang/openconfig-withlistval.yang
	- ../../pathgen/testdata/yang/openconfig-nested.yang
Imported modules were sourced from:
*/
package exampleoc

import (
	"reflect"

	"github.com/openconfig/ygot/ygot"
)

// ΛEnum is a map, keyed by the name of the type defined for each enum in the
// generated Go code, which provides a mapping between the constant int64 value
// of each value of the enumeration, and the string that is used to represent it
// in the YANG schema. The map is named ΛEnum in order to avoid clash with any
// valid YANG identifier.
var ΛEnum = map[string]map[int64]ygot.EnumDefinition{
	"E_Child_Three": {
		1: {Name: "ONE"},
		2: {Name: "TWO"},
	},
}

// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
func initΛEnumTypes() {
	ΛEnumTypes = map[string][]reflect.Type{
		"/parent/child/state/three": []reflect.Type{
			reflect.TypeOf((E_Child_Three)(0)),
		},
	}
}
