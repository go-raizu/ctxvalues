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

package ctxvalues_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-raizu/ctxvalues"
)

func Test_Collision(t *testing.T) {
	type UserID string

	var UserIDCtxKey = ctxvalues.NewKey[UserID]()

	type AdminID string

	var AdminIDCtxKey = ctxvalues.NewKey[AdminID]()

	ctx := context.Background()

	ctx = UserIDCtxKey.WithValue(ctx, "foo")
	ctx = AdminIDCtxKey.WithValue(ctx, "foo")

	assert.Equal(t, UserID("foo"), UserIDCtxKey.GetOrDefault(ctx, ""))
	assert.Equal(t, AdminID("foo"), AdminIDCtxKey.GetOrDefault(ctx, ""))
}

func Test_Individual(t *testing.T) {
	type UserID string

	var UserIDCtxKey = ctxvalues.NewKey[UserID]()

	type AdminID string

	var AdminIDCtxKey = ctxvalues.NewKey[AdminID]()

	ctx := context.Background()
	ctx = UserIDCtxKey.WithValue(ctx, "foo")
	ctx = AdminIDCtxKey.WithValue(ctx, "baa")

	assert.Equal(t, UserID("foo"), UserIDCtxKey.GetOrDefault(ctx, ""))
	assert.Equal(t, AdminID("baa"), AdminIDCtxKey.GetOrDefault(ctx, ""))
}

func Test_Interface(t *testing.T) {
	type InfoWriter io.Writer
	type ErrorWriter io.Writer

	var InfoWriterCtxKey = ctxvalues.NewKey[InfoWriter]()
	var ErrorWriterCtxKey = ctxvalues.NewKey[ErrorWriter]()

	var bufA bytes.Buffer
	var bufB bytes.Buffer

	ctx := context.Background()
	ctx = InfoWriterCtxKey.WithValue(ctx, &bufA)
	ctx = ErrorWriterCtxKey.WithValue(ctx, &bufB)

	assert.Same(t, &bufA, InfoWriterCtxKey.GetOrDefault(ctx, nil))
	assert.Same(t, &bufB, ErrorWriterCtxKey.GetOrDefault(ctx, nil))
}
