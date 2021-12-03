// Copyright 2021 Red Hat, Inc.
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

package ossm

import "os"

var (
	meshNamespace    string = getenv("MESHNAMESPACE", "istio-system")
	smcpName         string = getenv("SMCPNAME", "basic")
	smcpV21_template        = "../templates/smcp-templates/v2.1/cr_2.1_default_template.yaml"
	smcpV20_template        = "../templates/smcp-templates/v2.0/cr_2.0_default_template.yaml"
	smcpV11_template        = "../templates/smcp-templates/v1.1/cr_1.1_default_template.yaml"
	smcpV21                 = "../templates/smcp-templates/v2.1/cr_2.1_default.yaml"
	smcpV20                 = "../templates/smcp-templates/v2.0/cr_2.0_default.yaml"
	smcpV11                 = "../templates/smcp-templates/v1.1/cr_1.1_default.yaml"
	smmr                    = "../templates/smmr-templates/smmr_default.yaml"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
