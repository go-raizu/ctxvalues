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

	var UserIDCtxKey = ctxvalues.New[UserID]()

	type AdminID string

	var AdminIDCtxKey = ctxvalues.New[AdminID]()

	ctx := context.Background()

	ctx = UserIDCtxKey.WithValue(ctx, "foo")
	ctx = AdminIDCtxKey.WithValue(ctx, "foo")

	assert.Equal(t, UserID("foo"), UserIDCtxKey.GetOrZero(ctx))
	assert.Equal(t, AdminID("foo"), AdminIDCtxKey.GetOrZero(ctx))
}

func Test_Individual(t *testing.T) {
	type UserID string

	var UserIDCtxKey = ctxvalues.New[UserID]()

	type AdminID string

	var AdminIDCtxKey = ctxvalues.New[AdminID]()

	ctx := context.Background()
	ctx = UserIDCtxKey.WithValue(ctx, "foo")
	ctx = AdminIDCtxKey.WithValue(ctx, "baa")

	assert.Equal(t, UserID("foo"), UserIDCtxKey.GetOrZero(ctx))
	assert.Equal(t, AdminID("baa"), AdminIDCtxKey.GetOrZero(ctx))
}

func Test_Interface(t *testing.T) {
	type InfoWriter io.Writer
	type ErrorWriter io.Writer

	var InfoWriterCtxKey = ctxvalues.New[InfoWriter]()
	var ErrorWriterCtxKey = ctxvalues.New[ErrorWriter]()

	var bufA bytes.Buffer
	var bufB bytes.Buffer

	ctx := context.Background()
	ctx = InfoWriterCtxKey.WithValue(ctx, &bufA)
	ctx = ErrorWriterCtxKey.WithValue(ctx, &bufB)

	assert.Same(t, &bufA, InfoWriterCtxKey.GetOrZero(ctx))
	assert.Same(t, &bufB, ErrorWriterCtxKey.GetOrZero(ctx))
}

func Test_MultiType(t *testing.T) {
	type UserID string

	var UserIDCtxKey = ctxvalues.New2[UserID, string]()

	type AdminID string

	var AdminIDCtxKey = ctxvalues.New2[AdminID, string]()

	ctx := context.Background()
	ctx = UserIDCtxKey.WithValue(ctx, "foo")
	ctx = AdminIDCtxKey.WithValue(ctx, "baa")

	assert.Equal(t, "foo", UserIDCtxKey.GetOrZero(ctx))
	assert.Equal(t, "baa", AdminIDCtxKey.GetOrZero(ctx))
}

func TestKey_GetOrElse(t *testing.T) {
	type aType struct{}
	var aKey = ctxvalues.New2[aType, int]()

	type bType struct{}
	var bKey = ctxvalues.New2[bType, int]()

	ctx := aKey.WithValue(context.TODO(), 23)

	t.Run("present", func(t *testing.T) {
		assert.Equal(t, 23, aKey.GetOrElse(ctx, 11))
	})

	t.Run("absent", func(t *testing.T) {
		assert.Equal(t, 11, bKey.GetOrElse(ctx, 11))
	})
}
