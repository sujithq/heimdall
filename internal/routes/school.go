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

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moducate/heimdall/internal/env"
	"github.com/moducate/heimdall/internal/graph/model"
	"net/http"
)

func School(e *env.Env, r *gin.RouterGroup) {
	r.GET("/", getAllSchools(e))
	r.POST("/", createSchool(e))
	r.PUT("/:id", updateSchool(e))
	r.DELETE("/:id", deleteSchool(e))
}

func getAllSchools(e *env.Env) func(g *gin.Context) {
	return func(g *gin.Context) {
		schools, err := e.Services.School.GetAll()

		if err != nil {
			_ = g.Error(err)
			return
		}

		g.JSON(http.StatusOK, schools)
	}
}

func createSchool(e *env.Env) func(g *gin.Context) {
	return func(g *gin.Context) {
		input := model.NewSchool{}

		if err := g.BindJSON(&input); err != nil {
			_ = g.Error(err)
			return
		}

		school, err := e.Services.School.Create(input.Name)

		if err != nil {
			_ = g.Error(err)
			return
		}

		g.JSON(http.StatusOK, school)
	}
}

func updateSchool(e *env.Env) func(g *gin.Context) {
	return func(g *gin.Context) {
		input := model.NewSchool{}

		if err := g.BindJSON(&input); err != nil {
			_ = g.Error(err)
			return
		}

		school, err := e.Services.School.Update(g.Param("id"), input.Name)

		if err != nil {
			_ = g.Error(err)
			return
		}

		g.JSON(http.StatusOK, school)
	}
}

func deleteSchool(e *env.Env) func(g *gin.Context) {
	return func(g *gin.Context) {
		input := model.NewSchool{}

		if err := g.BindJSON(&input); err != nil {
			_ = g.Error(err)
			return
		}

		school, err := e.Services.School.Delete(g.Param("id"))

		if err != nil {
			_ = g.Error(err)
			return
		}

		g.JSON(http.StatusOK, school)
	}
}
