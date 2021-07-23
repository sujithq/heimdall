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
	"github.com/moducate/heimdall/internal/db"
	"github.com/spf13/cobra"
)

func MakeMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate <DSN>",
		Short: "Performs Heimdall's PostgreSQL database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			n, err := db.Migrate(args[0])

			if err != nil {
				cmd.PrintErrln(err)
			} else if n == 0 {
				cmd.Println("No migrations could be applied: the provided database is already up to date!")
			} else {
				cmd.Printf("Applied %d migrations successfully!\n", n)
			}
		},
	}
}

var migrateCmd = MakeMigrateCmd()

func init() {
	rootCmd.AddCommand(migrateCmd)
}
