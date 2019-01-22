// Copyright 2019 Istio Authors
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

// Package dashboard provides testing of the grafana dashboards used in Istio
// to provide mesh monitoring capabilities.

package maistra

import (
	"io/ioutil"
	"testing"
	"time"

	"istio.io/istio/pkg/log"
	"istio.io/istio/tests/util"
)

func cleanup06(namespace, kubeconfig string) {
	log.Infof("# Cleanup. Following error can be ignored...")
	util.KubeDelete(namespace, bookinfoAllv1Yaml, kubeconfig)
	log.Info("Waiting for rules to be cleaned up. Sleep 10 seconds...")
	time.Sleep(time.Duration(10) * time.Second)
}

func setup06(namespace, kubeconfig string) error {
	if err := util.KubeApply(namespace, bookinfoAllv1Yaml, kubeconfig); err != nil {
		return err
	}
	if err := util.KubeApply(namespace, bookinfoRatingDelayv2Yaml, kubeconfig); err != nil {
		return err
	}
	log.Info("Waiting for rules to propagate. Sleep 10 seconds...")
	time.Sleep(time.Duration(10) * time.Second)
	return nil
}

func setTimeout(namespace, kubeconfig string) error {
	log.Infof("# Set request timeouts")
	if err := util.KubeApply(namespace, bookinfoReviewTimeoutYaml, kubeconfig); err != nil {
		return err
	}
	log.Info("Waiting for rules to propagate. Sleep 10 seconds...")
	time.Sleep(time.Duration(10) * time.Second)
	return nil
}

func Test06(t *testing.T) {
	log.Infof("# TC_06 Setting Request Timeouts")
	inspect(setup06(testNamespace, ""), "failed to apply rules", "", t)
	t.Run("A1", func(t *testing.T) {
		inspect(setTimeout(testNamespace, ""), "failed to apply rules", "", t)

		resp, duration, err := getHTTPResponse(productpageURL, testUserJar)
		defer closeResponseBody(resp)
		inspect(err, "failed to get HTTP Response", "", t)
		log.Infof("bookinfo productpage returned in %d ms", duration)
		body, err := ioutil.ReadAll(resp.Body)
		inspect(err, "failed to read response body", "", t)
		inspect(
			compareHTTPResponse(body, "productpage-test-user-v2-review-timeout.html"),
			"Didn't get expected response.",
			"Success. Response timeout matches with expected.",
			t)
	})

	defer cleanup06(testNamespace, "")
}
