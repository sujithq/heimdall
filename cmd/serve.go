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
	"fmt"
	"github.com/moducate/heimdall/internal/server"
	"github.com/moducate/x/osx"
	"github.com/spf13/cobra"
)

func MakeServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serves Heimdall's HTTP API",
		Run: func(cmd *cobra.Command, args []string) {
			srv := server.New()
			srv.ListenAndServe(fmt.Sprintf(":%s", osx.Getenv("PORT", "1470")))
		},
	}
}

var serveCmd = MakeServeCmd()

func init() {
	rootCmd.AddCommand(serveCmd)
}
