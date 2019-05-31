// Copyright 2017-2019 Lei Ni (nilei81@gmail.com)
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

// +build dragonboat_monkeytest dragonboat_slowtest

package tests

import (
	"time"

	"github.com/lni/dragonboat/logger"
)

var (
	// TestMonkeyEnabled indicates whether we are in monkey test mode
	TestMonkeyEnabled = true
	plog              = logger.GetLogger("tests")
)

// ReadyToReturnTestKnob is a test knob that returns a boolean value indicating
// whether the system is being shutdown. In production, this function always
// return false without check the stopC chan.
func ReadyToReturnTestKnob(stopC <-chan struct{}, pos string) bool {
	if stopC == nil {
		return false
	}
	time.Sleep(1000 * time.Millisecond)
	select {
	case <-stopC:
		plog.Infof("test knob set, returning early before %s", pos)
		return true
	default:
		return false
	}
}
