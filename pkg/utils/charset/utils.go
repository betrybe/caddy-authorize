// Copyright 2020 Paul Greenberg greenpau@outlook.com
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

package utils

import (
	"strings"
)

// TokenChars characters allowed in base64 encoded string.
const TokenChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/-_="

// ContainsTokenCharset checks if a string contains non-base64 encoding
// characters.
func ContainsTokenCharset(s string) bool {
	dots := 0
	for _, c := range s {
		if !strings.ContainsRune(TokenChars, c) {
			return false
		}
		// Match dot
		if c == 46 {
			dots++
		}
	}
	if dots != 2 {
		return false
	}
	return true
}