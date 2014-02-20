//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package algebra

import (
	"github.com/couchbaselabs/query/value"
)

type Reciprocate struct {
	unaryBase
}

func NewReciprocate(operand Expression) Expression {
	return &Reciprocate{
		unaryBase{
			operand: operand,
		},
	}
}

func (this *Reciprocate) evaluate(operand value.Value) (value.Value, error) {
	if operand.Type() == value.NUMBER {
		a := operand.Actual().(float64)
		if a == 0.0 {
			return _NULL_VALUE, nil
		}
		return value.NewValue(1.0 / a), nil
	} else if operand.Type() == value.MISSING {
		return _MISSING_VALUE, nil
	} else {
		return _NULL_VALUE, nil
	}
}
