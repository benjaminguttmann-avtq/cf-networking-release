// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type EgressPolicyRepo struct {
	CreateTerminalStub        func(tx db.Transaction) (int64, error)
	createTerminalMutex       sync.RWMutex
	createTerminalArgsForCall []struct {
		tx db.Transaction
	}
	createTerminalReturns struct {
		result1 int64
		result2 error
	}
	createTerminalReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateAppStub        func(tx db.Transaction, sourceTerminalID int64, appGUID string) (int64, error)
	createAppMutex       sync.RWMutex
	createAppArgsForCall []struct {
		tx               db.Transaction
		sourceTerminalID int64
		appGUID          string
	}
	createAppReturns struct {
		result1 int64
		result2 error
	}
	createAppReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateIPRangeStub        func(tx db.Transaction, destinationTerminalID int64, startIP, endIP, protocol string) (int64, error)
	createIPRangeMutex       sync.RWMutex
	createIPRangeArgsForCall []struct {
		tx                    db.Transaction
		destinationTerminalID int64
		startIP               string
		endIP                 string
		protocol              string
	}
	createIPRangeReturns struct {
		result1 int64
		result2 error
	}
	createIPRangeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateEgressPolicyStub        func(tx db.Transaction, sourceTerminalID, destinationTerminalID int64) (int64, error)
	createEgressPolicyMutex       sync.RWMutex
	createEgressPolicyArgsForCall []struct {
		tx                    db.Transaction
		sourceTerminalID      int64
		destinationTerminalID int64
	}
	createEgressPolicyReturns struct {
		result1 int64
		result2 error
	}
	createEgressPolicyReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetTerminalByAppGUIDStub        func(tx db.Transaction, appGUID string) (int64, error)
	getTerminalByAppGUIDMutex       sync.RWMutex
	getTerminalByAppGUIDArgsForCall []struct {
		tx      db.Transaction
		appGUID string
	}
	getTerminalByAppGUIDReturns struct {
		result1 int64
		result2 error
	}
	getTerminalByAppGUIDReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetAllPoliciesStub        func() ([]store.EgressPolicy, error)
	getAllPoliciesMutex       sync.RWMutex
	getAllPoliciesArgsForCall []struct{}
	getAllPoliciesReturns     struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getAllPoliciesReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByGuidsStub        func(ids []string) ([]store.EgressPolicy, error)
	getByGuidsMutex       sync.RWMutex
	getByGuidsArgsForCall []struct {
		ids []string
	}
	getByGuidsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByGuidsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyRepo) CreateTerminal(tx db.Transaction) (int64, error) {
	fake.createTerminalMutex.Lock()
	ret, specificReturn := fake.createTerminalReturnsOnCall[len(fake.createTerminalArgsForCall)]
	fake.createTerminalArgsForCall = append(fake.createTerminalArgsForCall, struct {
		tx db.Transaction
	}{tx})
	fake.recordInvocation("CreateTerminal", []interface{}{tx})
	fake.createTerminalMutex.Unlock()
	if fake.CreateTerminalStub != nil {
		return fake.CreateTerminalStub(tx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTerminalReturns.result1, fake.createTerminalReturns.result2
}

func (fake *EgressPolicyRepo) CreateTerminalCallCount() int {
	fake.createTerminalMutex.RLock()
	defer fake.createTerminalMutex.RUnlock()
	return len(fake.createTerminalArgsForCall)
}

func (fake *EgressPolicyRepo) CreateTerminalArgsForCall(i int) db.Transaction {
	fake.createTerminalMutex.RLock()
	defer fake.createTerminalMutex.RUnlock()
	return fake.createTerminalArgsForCall[i].tx
}

func (fake *EgressPolicyRepo) CreateTerminalReturns(result1 int64, result2 error) {
	fake.CreateTerminalStub = nil
	fake.createTerminalReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateTerminalReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateTerminalStub = nil
	if fake.createTerminalReturnsOnCall == nil {
		fake.createTerminalReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createTerminalReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateApp(tx db.Transaction, sourceTerminalID int64, appGUID string) (int64, error) {
	fake.createAppMutex.Lock()
	ret, specificReturn := fake.createAppReturnsOnCall[len(fake.createAppArgsForCall)]
	fake.createAppArgsForCall = append(fake.createAppArgsForCall, struct {
		tx               db.Transaction
		sourceTerminalID int64
		appGUID          string
	}{tx, sourceTerminalID, appGUID})
	fake.recordInvocation("CreateApp", []interface{}{tx, sourceTerminalID, appGUID})
	fake.createAppMutex.Unlock()
	if fake.CreateAppStub != nil {
		return fake.CreateAppStub(tx, sourceTerminalID, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createAppReturns.result1, fake.createAppReturns.result2
}

func (fake *EgressPolicyRepo) CreateAppCallCount() int {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return len(fake.createAppArgsForCall)
}

func (fake *EgressPolicyRepo) CreateAppArgsForCall(i int) (db.Transaction, int64, string) {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return fake.createAppArgsForCall[i].tx, fake.createAppArgsForCall[i].sourceTerminalID, fake.createAppArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) CreateAppReturns(result1 int64, result2 error) {
	fake.CreateAppStub = nil
	fake.createAppReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateAppReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateAppStub = nil
	if fake.createAppReturnsOnCall == nil {
		fake.createAppReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createAppReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateIPRange(tx db.Transaction, destinationTerminalID int64, startIP string, endIP string, protocol string) (int64, error) {
	fake.createIPRangeMutex.Lock()
	ret, specificReturn := fake.createIPRangeReturnsOnCall[len(fake.createIPRangeArgsForCall)]
	fake.createIPRangeArgsForCall = append(fake.createIPRangeArgsForCall, struct {
		tx                    db.Transaction
		destinationTerminalID int64
		startIP               string
		endIP                 string
		protocol              string
	}{tx, destinationTerminalID, startIP, endIP, protocol})
	fake.recordInvocation("CreateIPRange", []interface{}{tx, destinationTerminalID, startIP, endIP, protocol})
	fake.createIPRangeMutex.Unlock()
	if fake.CreateIPRangeStub != nil {
		return fake.CreateIPRangeStub(tx, destinationTerminalID, startIP, endIP, protocol)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIPRangeReturns.result1, fake.createIPRangeReturns.result2
}

func (fake *EgressPolicyRepo) CreateIPRangeCallCount() int {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return len(fake.createIPRangeArgsForCall)
}

func (fake *EgressPolicyRepo) CreateIPRangeArgsForCall(i int) (db.Transaction, int64, string, string, string) {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return fake.createIPRangeArgsForCall[i].tx, fake.createIPRangeArgsForCall[i].destinationTerminalID, fake.createIPRangeArgsForCall[i].startIP, fake.createIPRangeArgsForCall[i].endIP, fake.createIPRangeArgsForCall[i].protocol
}

func (fake *EgressPolicyRepo) CreateIPRangeReturns(result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	fake.createIPRangeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateIPRangeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	if fake.createIPRangeReturnsOnCall == nil {
		fake.createIPRangeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createIPRangeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicy(tx db.Transaction, sourceTerminalID int64, destinationTerminalID int64) (int64, error) {
	fake.createEgressPolicyMutex.Lock()
	ret, specificReturn := fake.createEgressPolicyReturnsOnCall[len(fake.createEgressPolicyArgsForCall)]
	fake.createEgressPolicyArgsForCall = append(fake.createEgressPolicyArgsForCall, struct {
		tx                    db.Transaction
		sourceTerminalID      int64
		destinationTerminalID int64
	}{tx, sourceTerminalID, destinationTerminalID})
	fake.recordInvocation("CreateEgressPolicy", []interface{}{tx, sourceTerminalID, destinationTerminalID})
	fake.createEgressPolicyMutex.Unlock()
	if fake.CreateEgressPolicyStub != nil {
		return fake.CreateEgressPolicyStub(tx, sourceTerminalID, destinationTerminalID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createEgressPolicyReturns.result1, fake.createEgressPolicyReturns.result2
}

func (fake *EgressPolicyRepo) CreateEgressPolicyCallCount() int {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return len(fake.createEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) CreateEgressPolicyArgsForCall(i int) (db.Transaction, int64, int64) {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return fake.createEgressPolicyArgsForCall[i].tx, fake.createEgressPolicyArgsForCall[i].sourceTerminalID, fake.createEgressPolicyArgsForCall[i].destinationTerminalID
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturns(result1 int64, result2 error) {
	fake.CreateEgressPolicyStub = nil
	fake.createEgressPolicyReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateEgressPolicyStub = nil
	if fake.createEgressPolicyReturnsOnCall == nil {
		fake.createEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createEgressPolicyReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUID(tx db.Transaction, appGUID string) (int64, error) {
	fake.getTerminalByAppGUIDMutex.Lock()
	ret, specificReturn := fake.getTerminalByAppGUIDReturnsOnCall[len(fake.getTerminalByAppGUIDArgsForCall)]
	fake.getTerminalByAppGUIDArgsForCall = append(fake.getTerminalByAppGUIDArgsForCall, struct {
		tx      db.Transaction
		appGUID string
	}{tx, appGUID})
	fake.recordInvocation("GetTerminalByAppGUID", []interface{}{tx, appGUID})
	fake.getTerminalByAppGUIDMutex.Unlock()
	if fake.GetTerminalByAppGUIDStub != nil {
		return fake.GetTerminalByAppGUIDStub(tx, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTerminalByAppGUIDReturns.result1, fake.getTerminalByAppGUIDReturns.result2
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDCallCount() int {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return len(fake.getTerminalByAppGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDArgsForCall(i int) (db.Transaction, string) {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return fake.getTerminalByAppGUIDArgsForCall[i].tx, fake.getTerminalByAppGUIDArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturns(result1 int64, result2 error) {
	fake.GetTerminalByAppGUIDStub = nil
	fake.getTerminalByAppGUIDReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturnsOnCall(i int, result1 int64, result2 error) {
	fake.GetTerminalByAppGUIDStub = nil
	if fake.getTerminalByAppGUIDReturnsOnCall == nil {
		fake.getTerminalByAppGUIDReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.getTerminalByAppGUIDReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPolicies() ([]store.EgressPolicy, error) {
	fake.getAllPoliciesMutex.Lock()
	ret, specificReturn := fake.getAllPoliciesReturnsOnCall[len(fake.getAllPoliciesArgsForCall)]
	fake.getAllPoliciesArgsForCall = append(fake.getAllPoliciesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPolicies", []interface{}{})
	fake.getAllPoliciesMutex.Unlock()
	if fake.GetAllPoliciesStub != nil {
		return fake.GetAllPoliciesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllPoliciesReturns.result1, fake.getAllPoliciesReturns.result2
}

func (fake *EgressPolicyRepo) GetAllPoliciesCallCount() int {
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	return len(fake.getAllPoliciesArgsForCall)
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetAllPoliciesStub = nil
	fake.getAllPoliciesReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetAllPoliciesStub = nil
	if fake.getAllPoliciesReturnsOnCall == nil {
		fake.getAllPoliciesReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getAllPoliciesReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGuids(ids []string) ([]store.EgressPolicy, error) {
	var idsCopy []string
	if ids != nil {
		idsCopy = make([]string, len(ids))
		copy(idsCopy, ids)
	}
	fake.getByGuidsMutex.Lock()
	ret, specificReturn := fake.getByGuidsReturnsOnCall[len(fake.getByGuidsArgsForCall)]
	fake.getByGuidsArgsForCall = append(fake.getByGuidsArgsForCall, struct {
		ids []string
	}{idsCopy})
	fake.recordInvocation("GetByGuids", []interface{}{idsCopy})
	fake.getByGuidsMutex.Unlock()
	if fake.GetByGuidsStub != nil {
		return fake.GetByGuidsStub(ids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByGuidsReturns.result1, fake.getByGuidsReturns.result2
}

func (fake *EgressPolicyRepo) GetByGuidsCallCount() int {
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	return len(fake.getByGuidsArgsForCall)
}

func (fake *EgressPolicyRepo) GetByGuidsArgsForCall(i int) []string {
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	return fake.getByGuidsArgsForCall[i].ids
}

func (fake *EgressPolicyRepo) GetByGuidsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetByGuidsStub = nil
	fake.getByGuidsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGuidsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetByGuidsStub = nil
	if fake.getByGuidsReturnsOnCall == nil {
		fake.getByGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByGuidsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTerminalMutex.RLock()
	defer fake.createTerminalMutex.RUnlock()
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyRepo) recordInvocation(key string, args []interface{}) {
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
