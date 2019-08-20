// Copyright 2019 Jeff Nickoloff "jeff@allingeek.com"
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		os.Stderr.Write([]byte("expected exactly 2 arguments\n"))
		os.Exit(1)
		return
	}
	envVarName := os.Args[1]
	envVarValue := os.Getenv(envVarName)
	if len(envVarValue) <= 0 {
		os.Stderr.Write([]byte("Input variable: " + envVarName + " is unset: " + envVarValue + "\n"))
		os.Exit(1)
		return
	}

	outFile := os.Args[2]

	inDir := filepath.Dir(outFile)
	os.MkdirAll(inDir, 0777)

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		os.Stderr.Write([]byte(err.Error() + "\n"))
		os.Exit(2)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte(envVarValue))
	if err == nil && n < len(envVarValue) {
		os.Stderr.Write([]byte("unable to write complete value to file\n"))
		os.Exit(3)
		return
	}
	if err != nil {
		os.Stderr.Write([]byte(err.Error() + "\n"))
		os.Exit(4)
		return
	}
}
