//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package expression

/*
Replacer is used to replace one expr with another
*/

func ReplaceExpr(origExpr, oldExpr, newExpr Expression) (Expression, error) {
	replacer := newReplacer(oldExpr, newExpr)
	replaceExpr, err := replacer.Map(origExpr)
	if err != nil {
		return nil, err
	}

	// reset the value field since expr might have changed
	replaceExpr.ResetValue()

	return replaceExpr, nil
}

type Replacer struct {
	MapperBase

	oldExpr Expression
	newExpr Expression
}

func newReplacer(oldExpr, newExpr Expression) *Replacer {
	rv := &Replacer{
		oldExpr: oldExpr,
		newExpr: newExpr,
	}

	rv.mapFunc = func(expr Expression) (Expression, error) {
		if expr.EquivalentTo(rv.oldExpr) {
			return rv.newExpr, nil
		}

		return expr, expr.MapChildren(rv)
	}

	rv.mapper = rv
	return rv
}
