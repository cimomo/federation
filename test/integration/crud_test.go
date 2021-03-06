/*
Copyright 2017 The Kubernetes Authors.

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

package integration

import (
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkgruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/storage/names"
	federationapi "k8s.io/federation/apis/federation/v1beta1"
	"k8s.io/federation/pkg/federatedtypes"
	"k8s.io/federation/pkg/federatedtypes/crudtester"
	"k8s.io/federation/test/integration/framework"
)

// This test includes tests for resource deletion within a namespace that was reworked from existing
// e2e tests.

// TestFederationCRUD validates create/read/update/delete operations for federated resource types.
func TestFederationCRUD(t *testing.T) {
	fedFixture := framework.FederationFixture{DesiredClusterCount: 2}
	fedFixture.SetUp(t)
	defer fedFixture.TearDown(t)

	federatedTypes := federatedtypes.FederatedTypes()
	for kind, fedType := range federatedTypes {
		t.Run(kind, func(t *testing.T) {
			fixture, crudTester, obj, _ := initCRUDTest(t, &fedFixture, fedType.AdapterFactory, kind)
			defer fixture.TearDown(t)

			crudTester.CheckLifecycle(obj)
		})
	}

	// The following tests target a single type since the underlying logic is common across all types.
	kind := federatedtypes.SecretKind
	adapterFactory := federatedtypes.NewSecretAdapter

	// Validate deletion handling where orphanDependents is true or nil
	orphanedDependents := true
	testCases := map[string]*bool{
		"Resource should not be deleted from underlying clusters when OrphanDependents is true": &orphanedDependents,
		"Resource should not be deleted from underlying clusters when OrphanDependents is nil":  nil,
	}
	for testName, orphanDependents := range testCases {
		t.Run(testName, func(t *testing.T) {
			fixture, crudTester, obj, _ := initCRUDTest(t, &fedFixture, adapterFactory, kind)
			defer fixture.TearDown(t)

			updatedObj := crudTester.CheckCreate(obj)
			crudTester.CheckDelete(updatedObj, orphanDependents)
		})
	}

	t.Run("Resource should be propagated to a newly added cluster", func(t *testing.T) {
		fixture, crudTester, obj, _ := initCRUDTest(t, &fedFixture, adapterFactory, kind)
		defer fixture.TearDown(t)

		updatedObj := crudTester.CheckCreate(obj)
		// Start a new cluster and validate that the resource is propagated to it.
		fedFixture.StartCluster(t)
		// Check propagation to the new cluster by providing the updated set of clients
		objectExpected := true
		crudTester.CheckPropagationForClients(updatedObj, fedFixture.ClusterClients, objectExpected)
	})

	t.Run("Resource should only be propagated to the cluster with a matching selector", func(t *testing.T) {
		fixture, crudTester, obj, adapter := initCRUDTest(t, &fedFixture, adapterFactory, kind)
		defer fixture.TearDown(t)

		// Set an annotation to specify that the object is isolated to cluster 1.
		federatedtypes.SetAnnotation(adapter, obj, federationapi.FederationClusterSelectorAnnotation, `[{"key": "cluster", "operator": "==", "values": ["1"]}]`)

		updatedObj := crudTester.Create(obj)

		// Check propagation to the first cluster
		objectExpected := true
		crudTester.CheckPropagationForClients(updatedObj, fedFixture.ClusterClients[0:1], objectExpected)

		// Verify the object is not sent to the second cluster
		objectExpected = false
		crudTester.CheckPropagationForClients(updatedObj, fedFixture.ClusterClients[1:2], objectExpected)

	})

	// Rewrite of existing e2e namespace test
	t.Run("ReplicaSets should be deleted when namespace is deleted", func(t *testing.T) {

		// Create namespace crudtester
		fixtureOne, namespaceTester, nsObj, _ := initCRUDTest(t, &fedFixture, federatedtypes.NewNamespaceAdapter, federatedtypes.NamespaceKind)
		defer fixtureOne.TearDown(t)

		// Create namespace
		namespaceObj := namespaceTester.CheckCreate(nsObj)
		nsName := namespaceObj.(*apiv1.Namespace).Name

		// Create replicaset
		config := fedFixture.APIFixture.NewConfig()
		client := fedFixture.APIFixture.NewClient(fmt.Sprintf("crud-test-%s", federatedtypes.ReplicaSetKind))
		rsAdapter := adapterFactory(client, config, nil)
		obj := rsAdapter.NewTestObject(nsName)

		// Create replicaset inside of namespace
		replicasetObj, err := rsAdapter.FedCreate(obj)
		if err != nil {
			t.Logf("Failed to create replicasets %v in namespace %s, err: %s", &replicasetObj, nsName, err)
		}

		objectExpected := true
		namespaceTester.CheckDelete(namespaceObj, &objectExpected)

	})

	// Rewrite of existing e2e event test
	t.Run("Events should be deleted when namespace is deleted", func(t *testing.T) {
		// Create namespace crudtester
		fixtureOne, namespaceTester, nsObj, _ := initCRUDTest(t, &fedFixture, federatedtypes.NewNamespaceAdapter, federatedtypes.NamespaceKind)
		defer fixtureOne.TearDown(t)

		// Create namespace
		namespaceObj := namespaceTester.CheckCreate(nsObj)
		nsName := namespaceObj.(*apiv1.Namespace).Name

		// Create event
		event := apiv1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Name:      names.SimpleNameGenerator.GenerateName("integration-namespace-test-event"),
				Namespace: nsName,
			},
			InvolvedObject: apiv1.ObjectReference{
				Kind:      "Pod",
				Namespace: nsName,
				Name:      "sample-pod",
			},
		}

		client := fedFixture.APIFixture.NewClient(fmt.Sprintf("crud-test-%s", kind))
		// Create event inside of namespace
		eventObj, err := client.Core().Events(nsName).Create(&event)
		if err != nil {
			t.Logf("Failed to create event %v in namespace %s, err: %s", &eventObj, nsName, err)
		}

		objectExpected := true
		namespaceTester.CheckDelete(namespaceObj, &objectExpected)
	})
}

// initCRUDTest initializes common elements of a crud test
func initCRUDTest(t *testing.T, fedFixture *framework.FederationFixture, adapterFactory federatedtypes.AdapterFactory, kind string) (
	*framework.ControllerFixture, *crudtester.FederatedTypeCRUDTester, pkgruntime.Object, federatedtypes.FederatedTypeAdapter) {
	config := fedFixture.APIFixture.NewConfig()
	fixture := framework.NewControllerFixture(t, kind, adapterFactory, config)

	client := fedFixture.APIFixture.NewClient(fmt.Sprintf("crud-test-%s", kind))
	adapter := adapterFactory(client, config, nil)

	crudTester := framework.NewFederatedTypeCRUDTester(t, adapter, fedFixture.ClusterClients)

	obj := adapter.NewTestObject(uuid.New())

	return fixture, crudTester, obj, adapter
}
