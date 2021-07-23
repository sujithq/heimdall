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
	"database/sql"
	"embed"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	Sqlx *sqlx.DB
}

func New(dsn string) (*DB, error) {
	conn, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return &DB{
		Sqlx: conn,
	}, err
}

//go:embed queries/*
var queries embed.FS

func load(filename string) (string, error) {
	f, err := queries.Open(path.Join("queries", fmt.Sprintf("%s.sql", filename)))

	if err != nil {
		return "", err
	}

	var closeError error = nil

	defer func() {
		if err := f.Close(); err != nil {
			closeError = err
		}
	}()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		return "", err
	}

	return string(b), closeError
}

func (db DB) Selectf(dest interface{}, queryFile string, args ...interface{}) error {
	query, err := load(queryFile)

	if err != nil {
		return err
	}

	return db.Sqlx.Select(dest, query, args...)
}

func (db DB) Getf(dest interface{}, queryFile string, args ...interface{}) (bool, error) {
	query, err := load(queryFile)

	if err != nil {
		return false, err
	}

	err = db.Sqlx.Get(dest, query, args...)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
