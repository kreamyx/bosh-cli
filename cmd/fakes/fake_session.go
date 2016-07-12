// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/cmd"
	cmdconf "github.com/cloudfoundry/bosh-init/cmd/config"
	boshdir "github.com/cloudfoundry/bosh-init/director"
	boshuaa "github.com/cloudfoundry/bosh-init/uaa"
)

type FakeSession struct {
	TargetStub        func() string
	targetMutex       sync.RWMutex
	targetArgsForCall []struct{}
	targetReturns     struct {
		result1 string
	}
	CredentialsStub        func() cmdconf.Creds
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct{}
	credentialsReturns     struct {
		result1 cmdconf.Creds
	}
	UAAStub        func() (boshuaa.UAA, error)
	uAAMutex       sync.RWMutex
	uAAArgsForCall []struct{}
	uAAReturns     struct {
		result1 boshuaa.UAA
		result2 error
	}
	DirectorStub        func() (boshdir.Director, error)
	directorMutex       sync.RWMutex
	directorArgsForCall []struct{}
	directorReturns     struct {
		result1 boshdir.Director
		result2 error
	}
	AnonymousDirectorStub        func() (boshdir.Director, error)
	anonymousDirectorMutex       sync.RWMutex
	anonymousDirectorArgsForCall []struct{}
	anonymousDirectorReturns     struct {
		result1 boshdir.Director
		result2 error
	}
	DeploymentStub        func() (boshdir.Deployment, error)
	deploymentMutex       sync.RWMutex
	deploymentArgsForCall []struct{}
	deploymentReturns     struct {
		result1 boshdir.Deployment
		result2 error
	}
}

func (fake *FakeSession) Target() string {
	fake.targetMutex.Lock()
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct{}{})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub()
	} else {
		return fake.targetReturns.result1
	}
}

func (fake *FakeSession) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeSession) TargetReturns(result1 string) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSession) Credentials() cmdconf.Creds {
	fake.credentialsMutex.Lock()
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct{}{})
	fake.credentialsMutex.Unlock()
	if fake.CredentialsStub != nil {
		return fake.CredentialsStub()
	} else {
		return fake.credentialsReturns.result1
	}
}

func (fake *FakeSession) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeSession) CredentialsReturns(result1 cmdconf.Creds) {
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 cmdconf.Creds
	}{result1}
}

func (fake *FakeSession) UAA() (boshuaa.UAA, error) {
	fake.uAAMutex.Lock()
	fake.uAAArgsForCall = append(fake.uAAArgsForCall, struct{}{})
	fake.uAAMutex.Unlock()
	if fake.UAAStub != nil {
		return fake.UAAStub()
	} else {
		return fake.uAAReturns.result1, fake.uAAReturns.result2
	}
}

func (fake *FakeSession) UAACallCount() int {
	fake.uAAMutex.RLock()
	defer fake.uAAMutex.RUnlock()
	return len(fake.uAAArgsForCall)
}

func (fake *FakeSession) UAAReturns(result1 boshuaa.UAA, result2 error) {
	fake.UAAStub = nil
	fake.uAAReturns = struct {
		result1 boshuaa.UAA
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) Director() (boshdir.Director, error) {
	fake.directorMutex.Lock()
	fake.directorArgsForCall = append(fake.directorArgsForCall, struct{}{})
	fake.directorMutex.Unlock()
	if fake.DirectorStub != nil {
		return fake.DirectorStub()
	} else {
		return fake.directorReturns.result1, fake.directorReturns.result2
	}
}

func (fake *FakeSession) DirectorCallCount() int {
	fake.directorMutex.RLock()
	defer fake.directorMutex.RUnlock()
	return len(fake.directorArgsForCall)
}

func (fake *FakeSession) DirectorReturns(result1 boshdir.Director, result2 error) {
	fake.DirectorStub = nil
	fake.directorReturns = struct {
		result1 boshdir.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) AnonymousDirector() (boshdir.Director, error) {
	fake.anonymousDirectorMutex.Lock()
	fake.anonymousDirectorArgsForCall = append(fake.anonymousDirectorArgsForCall, struct{}{})
	fake.anonymousDirectorMutex.Unlock()
	if fake.AnonymousDirectorStub != nil {
		return fake.AnonymousDirectorStub()
	} else {
		return fake.anonymousDirectorReturns.result1, fake.anonymousDirectorReturns.result2
	}
}

func (fake *FakeSession) AnonymousDirectorCallCount() int {
	fake.anonymousDirectorMutex.RLock()
	defer fake.anonymousDirectorMutex.RUnlock()
	return len(fake.anonymousDirectorArgsForCall)
}

func (fake *FakeSession) AnonymousDirectorReturns(result1 boshdir.Director, result2 error) {
	fake.AnonymousDirectorStub = nil
	fake.anonymousDirectorReturns = struct {
		result1 boshdir.Director
		result2 error
	}{result1, result2}
}

func (fake *FakeSession) Deployment() (boshdir.Deployment, error) {
	fake.deploymentMutex.Lock()
	fake.deploymentArgsForCall = append(fake.deploymentArgsForCall, struct{}{})
	fake.deploymentMutex.Unlock()
	if fake.DeploymentStub != nil {
		return fake.DeploymentStub()
	} else {
		return fake.deploymentReturns.result1, fake.deploymentReturns.result2
	}
}

func (fake *FakeSession) DeploymentCallCount() int {
	fake.deploymentMutex.RLock()
	defer fake.deploymentMutex.RUnlock()
	return len(fake.deploymentArgsForCall)
}

func (fake *FakeSession) DeploymentReturns(result1 boshdir.Deployment, result2 error) {
	fake.DeploymentStub = nil
	fake.deploymentReturns = struct {
		result1 boshdir.Deployment
		result2 error
	}{result1, result2}
}

var _ cmd.Session = new(FakeSession)