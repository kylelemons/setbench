// Copyright 2018 Kyle Lemons
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

// Package setbench implements a simple benchmark for a string set.
//
// Use `go test` to run the benchmarks.
package setbench

// A Set stores a list of strings.
type Set interface {
	Add(string)
	Contains(string) bool
}

// A Map is a dumb implementation of a Set that uses a Go map.
type Map map[string]bool

// Add adds a string to the set.
func (m *Map) Add(s string) {
	if *m == nil {
		*m = make(map[string]bool)
	}
	(*m)[s] = true
}

// Contains returns whether a string is in the set.
func (m *Map) Contains(s string) bool {
	return (*m)[s]
}
