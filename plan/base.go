//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package plan

import (
	"encoding/json"
)

type readonly struct {
}

func (this *readonly) Readonly() bool {
	return true
}

func (this *readonly) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"readonly": true})
}

func (this *readonly) UnmarshalJSON([]byte) error {
	// TODO: implement
	return nil
}

type readwrite struct {
}

func (this *readwrite) Readonly() bool {
	return false
}

func (this *readwrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{"readonly": false})
}

func (this *readwrite) UnmarshalJSON([]byte) error {
	// TODO: implement
	return nil
}
