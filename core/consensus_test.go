// Copyright 2018 BlockCypher
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphWeight(t *testing.T) {
	// initial weights
	assert.Equal(t, GraphWeight(Mainnet, 1, 31), uint64(256*31))
	assert.Equal(t, GraphWeight(Mainnet, 1, 32), uint64(512*32))
	assert.Equal(t, GraphWeight(Mainnet, 1, 33), uint64(1024*33))

	// one year in, 31 starts going down, the rest stays the same
	assert.Equal(t, GraphWeight(Mainnet, YearHeight, 31), uint64(256*30))
	assert.Equal(t, GraphWeight(Mainnet, YearHeight, 32), uint64(512*32))
	assert.Equal(t, GraphWeight(Mainnet, YearHeight, 33), uint64(1024*33))

	// 31 loses one factor per week
	assert.Equal(t, GraphWeight(Mainnet, YearHeight+WeekHeight, 31), uint64(256*29))
	assert.Equal(t, GraphWeight(Mainnet, YearHeight+2*WeekHeight, 31), uint64(256*28))
	assert.Equal(t, GraphWeight(Mainnet, YearHeight+32*WeekHeight, 31), uint64(0))

	// 2 years in, 31 still at 0, 32 starts decreasing
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight, 31), uint64(0))
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight, 32), uint64(512*31))
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight, 33), uint64(1024*33))

	// 32 loses one factor per week
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight+WeekHeight, 32), uint64(512*30))
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight+WeekHeight, 31), uint64(0))
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight+30*WeekHeight, 32), uint64(512))
	assert.Equal(t, GraphWeight(Mainnet, 2*YearHeight+31*WeekHeight, 32), uint64(0))

	// 3 years in, nothing changes
	assert.Equal(t, GraphWeight(Mainnet, 3*YearHeight, 31), uint64(0))
	assert.Equal(t, GraphWeight(Mainnet, 3*YearHeight, 32), uint64(0))
	assert.Equal(t, GraphWeight(Mainnet, 3*YearHeight, 33), uint64(1024*33))

	// 4 years in, 33 starts starts decreasing
	assert.Equal(t, GraphWeight(Mainnet, 4*YearHeight, 31), uint64(0))
	assert.Equal(t, GraphWeight(Mainnet, 4*YearHeight, 32), uint64(0))
	assert.Equal(t, GraphWeight(Mainnet, 4*YearHeight, 33), uint64(1024*32))

}
