// Code generated by counterfeiter. DO NOT EDIT.
package nsaccessfakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/core/nsaccess"
	v1 "k8s.io/api/core/v1"
	v1a "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

type FakeChecker struct {
	FilterAccessibleNamespacesStub        func(context.Context, v1a.AuthorizationV1Interface, []v1.Namespace) ([]v1.Namespace, error)
	filterAccessibleNamespacesMutex       sync.RWMutex
	filterAccessibleNamespacesArgsForCall []struct {
		arg1 context.Context
		arg2 v1a.AuthorizationV1Interface
		arg3 []v1.Namespace
	}
	filterAccessibleNamespacesReturns struct {
		result1 []v1.Namespace
		result2 error
	}
	filterAccessibleNamespacesReturnsOnCall map[int]struct {
		result1 []v1.Namespace
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChecker) FilterAccessibleNamespaces(arg1 context.Context, arg2 v1a.AuthorizationV1Interface, arg3 []v1.Namespace) ([]v1.Namespace, error) {
	var arg3Copy []v1.Namespace
	if arg3 != nil {
		arg3Copy = make([]v1.Namespace, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.filterAccessibleNamespacesMutex.Lock()
	ret, specificReturn := fake.filterAccessibleNamespacesReturnsOnCall[len(fake.filterAccessibleNamespacesArgsForCall)]
	fake.filterAccessibleNamespacesArgsForCall = append(fake.filterAccessibleNamespacesArgsForCall, struct {
		arg1 context.Context
		arg2 v1a.AuthorizationV1Interface
		arg3 []v1.Namespace
	}{arg1, arg2, arg3Copy})
	stub := fake.FilterAccessibleNamespacesStub
	fakeReturns := fake.filterAccessibleNamespacesReturns
	fake.recordInvocation("FilterAccessibleNamespaces", []interface{}{arg1, arg2, arg3Copy})
	fake.filterAccessibleNamespacesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeChecker) FilterAccessibleNamespacesCallCount() int {
	fake.filterAccessibleNamespacesMutex.RLock()
	defer fake.filterAccessibleNamespacesMutex.RUnlock()
	return len(fake.filterAccessibleNamespacesArgsForCall)
}

func (fake *FakeChecker) FilterAccessibleNamespacesCalls(stub func(context.Context, v1a.AuthorizationV1Interface, []v1.Namespace) ([]v1.Namespace, error)) {
	fake.filterAccessibleNamespacesMutex.Lock()
	defer fake.filterAccessibleNamespacesMutex.Unlock()
	fake.FilterAccessibleNamespacesStub = stub
}

func (fake *FakeChecker) FilterAccessibleNamespacesArgsForCall(i int) (context.Context, v1a.AuthorizationV1Interface, []v1.Namespace) {
	fake.filterAccessibleNamespacesMutex.RLock()
	defer fake.filterAccessibleNamespacesMutex.RUnlock()
	argsForCall := fake.filterAccessibleNamespacesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeChecker) FilterAccessibleNamespacesReturns(result1 []v1.Namespace, result2 error) {
	fake.filterAccessibleNamespacesMutex.Lock()
	defer fake.filterAccessibleNamespacesMutex.Unlock()
	fake.FilterAccessibleNamespacesStub = nil
	fake.filterAccessibleNamespacesReturns = struct {
		result1 []v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeChecker) FilterAccessibleNamespacesReturnsOnCall(i int, result1 []v1.Namespace, result2 error) {
	fake.filterAccessibleNamespacesMutex.Lock()
	defer fake.filterAccessibleNamespacesMutex.Unlock()
	fake.FilterAccessibleNamespacesStub = nil
	if fake.filterAccessibleNamespacesReturnsOnCall == nil {
		fake.filterAccessibleNamespacesReturnsOnCall = make(map[int]struct {
			result1 []v1.Namespace
			result2 error
		})
	}
	fake.filterAccessibleNamespacesReturnsOnCall[i] = struct {
		result1 []v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterAccessibleNamespacesMutex.RLock()
	defer fake.filterAccessibleNamespacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeChecker) recordInvocation(key string, args []interface{}) {
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

var _ nsaccess.Checker = new(FakeChecker)
