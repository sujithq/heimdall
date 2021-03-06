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

package env

import (
	"github.com/moducate/heimdall/internal/db"
	"github.com/moducate/heimdall/internal/services"
)

type Env struct {
	DB       *db.DB
	Services struct {
		School services.SchoolService
	}
}

func New(db *db.DB) *Env {
	env := &Env{
		DB: db,
	}

	env.Services.School = services.SchoolServiceSql{DB: db}

	return env
}
