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

package db

import (
	"embed"
	"log"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(dsn string) (int, error) {
	conn, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return 0, err
	}

	defer func() {
		if conn != nil {
			if err := conn.Close(); err != nil {
				log.Fatal(err.Error())
			}
		}
	}()

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations,
		Root:       "migrations",
	}

	return migrate.Exec(conn.DB, "postgres", migrations, migrate.Up)
}
