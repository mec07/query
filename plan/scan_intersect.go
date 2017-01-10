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

	"github.com/couchbase/query/expression"
	"github.com/couchbase/query/value"
)

// IntersectScan scans multiple indexes and intersects the results.
type IntersectScan struct {
	readonly
	scans []SecondaryScan
}

func NewIntersectScan(scans ...SecondaryScan) *IntersectScan {
	return &IntersectScan{
		scans: scans,
	}
}

func (this *IntersectScan) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitIntersectScan(this)
}

func (this *IntersectScan) New() Operator {
	return &IntersectScan{}
}

func (this *IntersectScan) Covers() expression.Covers {
	return this.scans[0].Covers()
}

func (this *IntersectScan) FilterCovers() map[*expression.Cover]value.Value {
	return this.scans[0].FilterCovers()
}

func (this *IntersectScan) Covering() bool {
	return this.scans[0].Covering()
}

func (this *IntersectScan) Scans() []SecondaryScan {
	return this.scans
}

func (this *IntersectScan) String() string {
	bytes, _ := this.MarshalJSON()
	return string(bytes)
}

func (this *IntersectScan) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.MarshalBase(nil))
}

func (this *IntersectScan) MarshalBase(f func(map[string]interface{})) map[string]interface{} {
	r := map[string]interface{}{"#operator": "IntersectScan"}
	if f != nil {
		f(r)
	} else {
		r["scans"] = this.scans
	}
	return r
}

func (this *IntersectScan) UnmarshalJSON(body []byte) error {
	var _unmarshalled struct {
		_     string            `json:"#operator"`
		Scans []json.RawMessage `json:"scans"`
	}

	err := json.Unmarshal(body, &_unmarshalled)
	if err != nil {
		return err
	}

	this.scans = make([]SecondaryScan, 0, len(_unmarshalled.Scans))

	for _, raw_scan := range _unmarshalled.Scans {
		var scan_type struct {
			Operator string `json:"#operator"`
		}

		err = json.Unmarshal(raw_scan, &scan_type)
		if err != nil {
			return err
		}

		scan_op, err := MakeOperator(scan_type.Operator, raw_scan)
		if err != nil {
			return err
		}

		this.scans = append(this.scans, scan_op.(SecondaryScan))
	}

	return nil
}
