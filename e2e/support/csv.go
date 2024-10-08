//go:build integration
// +build integration

// To enable compilation of this file in Goland, go to "Settings -> Go -> Vendoring & Build Tags -> Custom Tags" and add "integration"

/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package support

import (
	"context"
	"strings"
	"testing"
	"unsafe"

	ctrl "sigs.k8s.io/controller-runtime/pkg/client"

	olm "github.com/operator-framework/api/pkg/operators/v1alpha1"
)

func ClusterServiceVersion(t *testing.T, ctx context.Context, conditions func(olm.ClusterServiceVersion) bool, ns string) func() *olm.ClusterServiceVersion {
	return func() *olm.ClusterServiceVersion {
		lst := olm.ClusterServiceVersionList{}
		if err := TestClient(t).List(ctx, &lst, ctrl.InNamespace(ns)); err != nil {
			panic(err)
		}
		for _, s := range lst.Items {
			if strings.Contains(s.Name, "camel-k") && conditions(s) {
				return &s
			}
		}
		return nil
	}
}

func ClusterServiceVersionPhase(t *testing.T, ctx context.Context, conditions func(olm.ClusterServiceVersion) bool, ns string) func() olm.ClusterServiceVersionPhase {
	return func() olm.ClusterServiceVersionPhase {
		if csv := ClusterServiceVersion(t, ctx, conditions, ns)(); csv != nil && unsafe.Sizeof(csv.Status) > 0 {
			return csv.Status.Phase
		}
		return ""
	}
}
