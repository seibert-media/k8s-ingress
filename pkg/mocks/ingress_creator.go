// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/seibert-media/k8s-ingress/pkg/domain"
	"k8s.io/api/extensions/v1beta1"
)

type IngressCreator struct {
	CreateStub        func(domains []domain.Domain) *v1beta1.Ingress
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		domains []domain.Domain
	}
	createReturns struct {
		result1 *v1beta1.Ingress
	}
	createReturnsOnCall map[int]struct {
		result1 *v1beta1.Ingress
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IngressCreator) Create(domains []domain.Domain) *v1beta1.Ingress {
	var domainsCopy []domain.Domain
	if domains != nil {
		domainsCopy = make([]domain.Domain, len(domains))
		copy(domainsCopy, domains)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		domains []domain.Domain
	}{domainsCopy})
	fake.recordInvocation("Create", []interface{}{domainsCopy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(domains)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *IngressCreator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *IngressCreator) CreateArgsForCall(i int) []domain.Domain {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].domains
}

func (fake *IngressCreator) CreateReturns(result1 *v1beta1.Ingress) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1beta1.Ingress
	}{result1}
}

func (fake *IngressCreator) CreateReturnsOnCall(i int, result1 *v1beta1.Ingress) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1beta1.Ingress
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1beta1.Ingress
	}{result1}
}

func (fake *IngressCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IngressCreator) recordInvocation(key string, args []interface{}) {
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
