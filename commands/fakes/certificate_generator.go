// This file was generated by counterfeiter
package fakes

import (
	"sync"
)

type CertificateGenerator struct {
	GenerateStub        func() (string, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct{}
	generateReturns     struct {
		result1 string
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CertificateGenerator) Generate() (string, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct{}{})
	fake.recordInvocation("Generate", []interface{}{})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.generateReturns.result1, fake.generateReturns.result2
}

func (fake *CertificateGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *CertificateGenerator) GenerateReturns(result1 string, result2 error) {
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CertificateGenerator) GenerateReturnsOnCall(i int, result1 string, result2 error) {
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *CertificateGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return fake.invocations
}

func (fake *CertificateGenerator) recordInvocation(key string, args []interface{}) {
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
