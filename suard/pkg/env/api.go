// /*
// Copyright 2017 The KUAR Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

package env

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/suard/pkg/apiutils"
)

// EnvStatus is returned from a GET to this API endpoing
type EnvStatus struct {
	CommandLine []string          `json:"commandLine"`
	Env         map[string]string `json:"env"`
}

type Env struct{}

func (s *EnvStatus) ToJson() string {
	jsonString := "{"
	for key, value := range s.Env {
		jsonString += fmt.Sprintf(`"%s": "%s",`, key, value)
	}
	jsonString = strings.TrimSuffix(jsonString, ",")
	jsonString += "}"

	return jsonString
}

func New() *Env {
	return &Env{}
}

func (e *Env) AddRoutes(r *chi.Mux) {
	r.Get("/env/api", e.APIGet)
}

func (e *Env) APIGet(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "rootHandler not yet implemented", http.StatusNotImplemented)

	s := EnvStatus{}

	s.CommandLine = os.Args

	s.Env = map[string]string{}
	for k, v := range r.Header {
		s.Env[k] = v[0]
	}

	apiutils.ServeJSON(w, s.ToJson())
}
