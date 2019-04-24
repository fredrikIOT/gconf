// Copyright 2019 xgfone
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

// +build !go1.10

package gconf

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func printDefaultFlagUsage(fset *flag.FlagSet) {
	if name := filepath.Base(os.Args[0]); name == "" {
		fmt.Fprintf(os.Stderr, "Usage:\n")
	} else {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", name)
	}
	PrintFlagUsage(os.Stderr, fset, false)
}