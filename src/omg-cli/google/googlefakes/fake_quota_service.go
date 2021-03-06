// Code generated by counterfeiter. DO NOT EDIT.
package googlefakes

import (
	"omg-cli/google"
	"sync"
)

type FakeQuotaService struct {
	ProjectStub        func() (map[string]google.Quota, error)
	projectMutex       sync.RWMutex
	projectArgsForCall []struct {
	}
	projectReturns struct {
		result1 map[string]google.Quota
		result2 error
	}
	projectReturnsOnCall map[int]struct {
		result1 map[string]google.Quota
		result2 error
	}
	RegionStub        func(string) (map[string]google.Quota, error)
	regionMutex       sync.RWMutex
	regionArgsForCall []struct {
		arg1 string
	}
	regionReturns struct {
		result1 map[string]google.Quota
		result2 error
	}
	regionReturnsOnCall map[int]struct {
		result1 map[string]google.Quota
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeQuotaService) Project() (map[string]google.Quota, error) {
	fake.projectMutex.Lock()
	ret, specificReturn := fake.projectReturnsOnCall[len(fake.projectArgsForCall)]
	fake.projectArgsForCall = append(fake.projectArgsForCall, struct {
	}{})
	fake.recordInvocation("Project", []interface{}{})
	fake.projectMutex.Unlock()
	if fake.ProjectStub != nil {
		return fake.ProjectStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.projectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeQuotaService) ProjectCallCount() int {
	fake.projectMutex.RLock()
	defer fake.projectMutex.RUnlock()
	return len(fake.projectArgsForCall)
}

func (fake *FakeQuotaService) ProjectCalls(stub func() (map[string]google.Quota, error)) {
	fake.projectMutex.Lock()
	defer fake.projectMutex.Unlock()
	fake.ProjectStub = stub
}

func (fake *FakeQuotaService) ProjectReturns(result1 map[string]google.Quota, result2 error) {
	fake.projectMutex.Lock()
	defer fake.projectMutex.Unlock()
	fake.ProjectStub = nil
	fake.projectReturns = struct {
		result1 map[string]google.Quota
		result2 error
	}{result1, result2}
}

func (fake *FakeQuotaService) ProjectReturnsOnCall(i int, result1 map[string]google.Quota, result2 error) {
	fake.projectMutex.Lock()
	defer fake.projectMutex.Unlock()
	fake.ProjectStub = nil
	if fake.projectReturnsOnCall == nil {
		fake.projectReturnsOnCall = make(map[int]struct {
			result1 map[string]google.Quota
			result2 error
		})
	}
	fake.projectReturnsOnCall[i] = struct {
		result1 map[string]google.Quota
		result2 error
	}{result1, result2}
}

func (fake *FakeQuotaService) Region(arg1 string) (map[string]google.Quota, error) {
	fake.regionMutex.Lock()
	ret, specificReturn := fake.regionReturnsOnCall[len(fake.regionArgsForCall)]
	fake.regionArgsForCall = append(fake.regionArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Region", []interface{}{arg1})
	fake.regionMutex.Unlock()
	if fake.RegionStub != nil {
		return fake.RegionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.regionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeQuotaService) RegionCallCount() int {
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	return len(fake.regionArgsForCall)
}

func (fake *FakeQuotaService) RegionCalls(stub func(string) (map[string]google.Quota, error)) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = stub
}

func (fake *FakeQuotaService) RegionArgsForCall(i int) string {
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	argsForCall := fake.regionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeQuotaService) RegionReturns(result1 map[string]google.Quota, result2 error) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = nil
	fake.regionReturns = struct {
		result1 map[string]google.Quota
		result2 error
	}{result1, result2}
}

func (fake *FakeQuotaService) RegionReturnsOnCall(i int, result1 map[string]google.Quota, result2 error) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = nil
	if fake.regionReturnsOnCall == nil {
		fake.regionReturnsOnCall = make(map[int]struct {
			result1 map[string]google.Quota
			result2 error
		})
	}
	fake.regionReturnsOnCall[i] = struct {
		result1 map[string]google.Quota
		result2 error
	}{result1, result2}
}

func (fake *FakeQuotaService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.projectMutex.RLock()
	defer fake.projectMutex.RUnlock()
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeQuotaService) recordInvocation(key string, args []interface{}) {
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

var _ google.QuotaService = new(FakeQuotaService)
