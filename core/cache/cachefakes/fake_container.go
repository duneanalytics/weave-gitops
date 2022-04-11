// Code generated by counterfeiter. DO NOT EDIT.
package cachefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/core/cache"
	v1 "k8s.io/api/core/v1"
)

type FakeContainer struct {
	ForceRefreshStub        func(cache.StorageType)
	forceRefreshMutex       sync.RWMutex
	forceRefreshArgsForCall []struct {
		arg1 cache.StorageType
	}
	NamespacesStub        func() map[string][]v1.Namespace
	namespacesMutex       sync.RWMutex
	namespacesArgsForCall []struct {
	}
	namespacesReturns struct {
		result1 map[string][]v1.Namespace
	}
	namespacesReturnsOnCall map[int]struct {
		result1 map[string][]v1.Namespace
	}
	StartStub        func(context.Context)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainer) ForceRefresh(arg1 cache.StorageType) {
	fake.forceRefreshMutex.Lock()
	fake.forceRefreshArgsForCall = append(fake.forceRefreshArgsForCall, struct {
		arg1 cache.StorageType
	}{arg1})
	stub := fake.ForceRefreshStub
	fake.recordInvocation("ForceRefresh", []interface{}{arg1})
	fake.forceRefreshMutex.Unlock()
	if stub != nil {
		fake.ForceRefreshStub(arg1)
	}
}

func (fake *FakeContainer) ForceRefreshCallCount() int {
	fake.forceRefreshMutex.RLock()
	defer fake.forceRefreshMutex.RUnlock()
	return len(fake.forceRefreshArgsForCall)
}

func (fake *FakeContainer) ForceRefreshCalls(stub func(cache.StorageType)) {
	fake.forceRefreshMutex.Lock()
	defer fake.forceRefreshMutex.Unlock()
	fake.ForceRefreshStub = stub
}

func (fake *FakeContainer) ForceRefreshArgsForCall(i int) cache.StorageType {
	fake.forceRefreshMutex.RLock()
	defer fake.forceRefreshMutex.RUnlock()
	argsForCall := fake.forceRefreshArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainer) Namespaces() map[string][]v1.Namespace {
	fake.namespacesMutex.Lock()
	ret, specificReturn := fake.namespacesReturnsOnCall[len(fake.namespacesArgsForCall)]
	fake.namespacesArgsForCall = append(fake.namespacesArgsForCall, struct {
	}{})
	stub := fake.NamespacesStub
	fakeReturns := fake.namespacesReturns
	fake.recordInvocation("Namespaces", []interface{}{})
	fake.namespacesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeContainer) NamespacesCallCount() int {
	fake.namespacesMutex.RLock()
	defer fake.namespacesMutex.RUnlock()
	return len(fake.namespacesArgsForCall)
}

func (fake *FakeContainer) NamespacesCalls(stub func() map[string][]v1.Namespace) {
	fake.namespacesMutex.Lock()
	defer fake.namespacesMutex.Unlock()
	fake.NamespacesStub = stub
}

func (fake *FakeContainer) NamespacesReturns(result1 map[string][]v1.Namespace) {
	fake.namespacesMutex.Lock()
	defer fake.namespacesMutex.Unlock()
	fake.NamespacesStub = nil
	fake.namespacesReturns = struct {
		result1 map[string][]v1.Namespace
	}{result1}
}

func (fake *FakeContainer) NamespacesReturnsOnCall(i int, result1 map[string][]v1.Namespace) {
	fake.namespacesMutex.Lock()
	defer fake.namespacesMutex.Unlock()
	fake.NamespacesStub = nil
	if fake.namespacesReturnsOnCall == nil {
		fake.namespacesReturnsOnCall = make(map[int]struct {
			result1 map[string][]v1.Namespace
		})
	}
	fake.namespacesReturnsOnCall[i] = struct {
		result1 map[string][]v1.Namespace
	}{result1}
}

func (fake *FakeContainer) Start(arg1 context.Context) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub(arg1)
	}
}

func (fake *FakeContainer) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeContainer) StartCalls(stub func(context.Context)) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeContainer) StartArgsForCall(i int) context.Context {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainer) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub()
	}
}

func (fake *FakeContainer) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeContainer) StopCalls(stub func()) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.forceRefreshMutex.RLock()
	defer fake.forceRefreshMutex.RUnlock()
	fake.namespacesMutex.RLock()
	defer fake.namespacesMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cache.Container = new(FakeContainer)
