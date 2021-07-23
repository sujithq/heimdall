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

package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	sql "github.com/moducate/heimdall/internal/db"
	"github.com/moducate/heimdall/internal/env"
	"github.com/moducate/heimdall/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Gin *gin.Engine
	Env *env.Env
}

func newGin() *gin.Engine {
	gin.ForceConsoleColor()
	g := gin.Default()
	return g
}

func New(dsn string) *Server {
	srv := &Server{
		Gin: newGin(),
	}

	db, err := sql.New(dsn)
	if err != nil {
		log.Fatalf("Heimdall server failed to connect to PostgreSQL: %s\n", err)
	}

	srv.Env = env.New(db)

	routes.School(srv.Env, srv.Gin.Group("/school"))
	routes.Graphql(srv.Env, srv.Gin.Group("/graphql"))

	return srv
}

func (s *Server) ListenAndServe(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: s.Gin,
	}

	go func() {
		log.Println(fmt.Sprintf("Heimdall server now listening on %s", addr))
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Println("Heimdall server closed")
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Heimdall server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Heimdall server forced to shutdown: %s\n", err)
	}
}
