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
	"fmt"
	"time"

	"github.com/couchbaselabs/query/value"
)

type Node interface {
	fmt.Stringer
}

type Expression interface {
	Node

	Evaluate(item value.Value, context Context) (value.Value, error)

	// Is this Expression equivalent to another
	EquivalentTo(other Expression) bool

	// A list of other Expressions on which this depends
	Dependencies() ExpressionList

	// Copy
	Copy() Expression
}

type ExpressionList []Expression

type Path Expression

type Context interface {
	Now() time.Time
	Argument(parameter string) value.Value
	EvaluateSubquery(query *SelectNode, item value.Value) (value.Value, error)
}
