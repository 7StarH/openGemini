/*
Copyright 2022 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pool_test

import (
	"runtime/debug"
	"testing"

	"github.com/openGemini/openGemini/lib/pool"
	"github.com/stretchr/testify/assert"
)

func TestUint32Array(t *testing.T) {
	// disable GC
	debug.SetGCPercent(-1)
	defer debug.SetGCPercent(100)

	p := pool.NewUint32Array()
	s := p.Get(16)
	assert.Equal(t, 16, len(s))

	p.Put(s)
	s = p.Get(32)
	assert.Equal(t, 32, len(s))

	p.Put(s)
	s = p.Get(16)
	assert.Equal(t, 16, len(s))

	assert.Equal(t, 333, int(p.HitRatio()*1000))
}
