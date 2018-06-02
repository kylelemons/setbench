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

package setbench

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"
)

// LoadWords loads the word list from data
func LoadWords() ([]string, error) {
	file, err := os.Open("data/words.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open word list: %s", err)
	}

	var words []string

	lines := bufio.NewScanner(file)
	for lines.Scan() {
		words = append(words, lines.Text())
	}
	if lines.Err() != nil {
		return nil, fmt.Errorf("failed to read full list: %s", err)
	}
	return words, nil
}

func TestLoadWords(t *testing.T) {
	words, err := LoadWords()
	if err != nil {
		t.Fatalf("LoadWords failed: %s", err)
	}
	valid, err := regexp.Compile("^[a-z]*$")
	if err != nil {
		t.Fatalf("BUG: failed to compile valid-word regexp: %s", err)
	}
	seen := make(map[string]bool)
	for _, w := range words {
		if seen[w] {
			t.Errorf("duplicate word %q", w)
		}
		seen[w] = true
		if valid.MatchString(w) {
			continue
		}
		t.Errorf("found invalid word %q", w)
	}
}

func TestSets(t *testing.T) {
	implementations := []struct {
		name string
		s    Set
	}{
		{"Map", new(Map)},
	}

	words, err := LoadWords()
	if err != nil {
		t.Fatalf("LoadWords failed: %s", err)
	}
	toAdd, missing := words[:len(words)/2], words[len(words)/2:]

	for _, imp := range implementations {
		t.Run(fmt.Sprintf("%s.Add", imp.name), func(t *testing.T) {
			for _, w := range toAdd {
				imp.s.Add(w)
			}
		})
		t.Run(fmt.Sprintf("%s.Contains", imp.name), func(t *testing.T) {
			for _, w := range toAdd {
				if !imp.s.Contains(w) {
					t.Errorf("set should contain %q", w)
				}
			}
			for _, w := range missing {
				if imp.s.Contains(w) {
					t.Errorf("set should NOT contain %q", w)
				}
			}
		})
	}
}

func BenchmarkSets(b *testing.B) {
	implementations := []struct {
		name string
		s    Set
	}{
		{"Map", new(Map)},
	}

	words, err := LoadWords()
	if err != nil {
		b.Fatalf("LoadWords failed: %s", err)
	}

	for _, imp := range implementations {
		b.Run(fmt.Sprintf("%s.Add", imp.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				imp.s.Add(words[i%len(words)])
			}
		})
		b.Run(fmt.Sprintf("%s.Contains", imp.name), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if !imp.s.Contains(words[i%len(words)]) {
					b.Errorf("set should contain %q", words[i%len(words)])
				}
			}
		})
	}
}
