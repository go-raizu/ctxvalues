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

// Key provides a type-safe approach to managing values in [context.Context].
type Key[KT, VT any] struct{}

// WithValue returns a copy of the parent context in which the value
// associated with key type is val.
func (k Key[KT, VT]) WithValue(ctx context.Context, val VT) context.Context {
	return context.WithValue(ctx, (*KT)(nil), val)
}

// Get returns the value and a boolean value indicating whenever
// the value was found within the context or not.
func (k Key[KT, VT]) Get(ctx context.Context) (v VT, ok bool) {
	v, ok = ctx.Value((*KT)(nil)).(VT)
	return
}

// GetOrElse returns the value if the value was found within the
// context or the provided default value otherwise.
func (k Key[KT, VT]) GetOrElse(ctx context.Context, val VT) VT {
	if v, ok := k.Get(ctx); ok {
		val = v
	}
	return val
}

// GetOrZero returns the value if the value was found within the
// context or the zero value of the given type.
func (k Key[KT, VT]) GetOrZero(ctx context.Context) (out VT) {
	if v, ok := k.Get(ctx); ok {
		out = v
	}
	return
}

// New2 creates a new Key with the given key and type value which can differ.
func New2[KT, VT any]() Key[KT, VT] {
	return Key[KT, VT]{}
}

// New creates a new Key with a given type that is used both as the key and value.
func New[T any]() Key[T, T] {
	return Key[T, T]{}
}
