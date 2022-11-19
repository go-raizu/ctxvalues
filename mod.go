// Copyright(c) 2022 individual contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package ctxvalues

import (
	"context"
)

type Key[T any] struct{}

// WithValue returns a copy of the parent context in which the value
// associated with key type is val.
func (k Key[T]) WithValue(ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, (*T)(nil), val)
}

// Get returns the value and a boolean value indicating whenever
// the value was found within the context or not.
func (k Key[T]) Get(ctx context.Context) (v T, ok bool) {
	v, ok = ctx.Value((*T)(nil)).(T)
	return
}

// GetOrElse returns the value if the value was found within the
// context or the provided default value otherwise.
func (k Key[T]) GetOrElse(ctx context.Context, val T) T {
	if v, ok := k.Get(ctx); ok {
		val = v
	}
	return val
}

func NewKey[T any]() Key[T] {
	return Key[T]{}
}
