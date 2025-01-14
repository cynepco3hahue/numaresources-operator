/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rte_e2e

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"

	_ "github.com/k8stopologyawareschedwg/resource-topology-exporter/test/e2e/rte"
	_ "github.com/k8stopologyawareschedwg/resource-topology-exporter/test/e2e/topology_updater"
)

func TestMain(m *testing.M) {
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()

	framework.AfterReadingAllFlags(&framework.TestContext)

	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestRTE(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RTE Test Suite")
}
