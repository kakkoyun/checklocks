// Copyright 2020 The gVisor Authors.
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

package test

func testDeferValidUnlock(tc *oneGuardStruct) {
	tc.mu.Lock()
	tc.guardedField = 1
	defer tc.mu.Unlock()
}

func testDeferValidAccess(tc *oneGuardStruct) {
	tc.mu.Lock()
	defer func() {
		tc.guardedField = 1
		tc.mu.Unlock()
	}()
}

func testMultipleDefersValidAccess(tc *oneGuardStruct) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	defer func() {
		tc.guardedField = 1
	}()
}

func testDeferInvalidAccess(tc *oneGuardStruct) {
	tc.mu.Lock()
	defer func() {
		// N.B. Executed after tc.mu.Unlock().
		tc.guardedField = 1 // +checklocksfail
	}()
	tc.mu.Unlock()
}

func testDeferNonInlinedClosure(tc *oneGuardStruct) {
	tc.mu.Lock()
	// Store closure in a variable to prevent inlining
	cleanup := func() {
		tc.guardedField = 1
		tc.mu.Unlock()
	}
	defer cleanup()
}

func testDeferNonInlinedClosureInvalid(tc *oneGuardStruct) {
	tc.mu.Lock()
	// Store closure in a variable to prevent inlining
	cleanup := func() {
		tc.guardedField = 1 // +checklocksfail
	}
	defer cleanup()
	tc.mu.Unlock()
}
