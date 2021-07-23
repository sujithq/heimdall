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

package services

import (
	"github.com/moducate/heimdall/internal/db"
	"github.com/moducate/heimdall/internal/graph/model"
)

type SchoolService interface {
	GetAll() ([]*model.School, error)
}

type SchoolServiceSql struct {
	DB *db.DB
}

func (s SchoolServiceSql) GetAll() ([]*model.School, error) {
	var schools []*model.School
	err := s.DB.Selectf(&schools, "models/school/get_all")

	return schools, err
}
