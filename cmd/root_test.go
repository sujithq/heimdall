// Copyright © 2021 Luke Carr <me+oss@carr.sh>
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

package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestMakeRootCmd(t *testing.T) {
	cmd := MakeRootCmd()

	assert.NotNil(t, cmd)
	assert.IsType(t, &cobra.Command{}, cmd)
	assert.Equal(t, "moducate-heimdall", cmd.Use)
}

func TestExecuteRootCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := MakeRootCmd()
	cmd.SetOut(b)

	assert.NoError(t, cmd.Execute())

	out, err := ioutil.ReadAll(b)
	assert.NoError(t, err)

	assert.Equal(t, "Run and manage Moducate Heimdall", strings.Split(string(out), "\n")[0])
}
