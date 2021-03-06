// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/formcontent"
)

type Multipart struct {
	AddFieldStub        func(string, string) error
	addFieldMutex       sync.RWMutex
	addFieldArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addFieldReturns struct {
		result1 error
	}
	addFieldReturnsOnCall map[int]struct {
		result1 error
	}
	AddFileStub        func(string, string) error
	addFileMutex       sync.RWMutex
	addFileArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addFileReturns struct {
		result1 error
	}
	addFileReturnsOnCall map[int]struct {
		result1 error
	}
	FinalizeStub        func() formcontent.ContentSubmission
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct {
	}
	finalizeReturns struct {
		result1 formcontent.ContentSubmission
	}
	finalizeReturnsOnCall map[int]struct {
		result1 formcontent.ContentSubmission
	}
	ResetStub        func()
	resetMutex       sync.RWMutex
	resetArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Multipart) AddField(arg1 string, arg2 string) error {
	fake.addFieldMutex.Lock()
	ret, specificReturn := fake.addFieldReturnsOnCall[len(fake.addFieldArgsForCall)]
	fake.addFieldArgsForCall = append(fake.addFieldArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AddField", []interface{}{arg1, arg2})
	fake.addFieldMutex.Unlock()
	if fake.AddFieldStub != nil {
		return fake.AddFieldStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addFieldReturns
	return fakeReturns.result1
}

func (fake *Multipart) AddFieldCallCount() int {
	fake.addFieldMutex.RLock()
	defer fake.addFieldMutex.RUnlock()
	return len(fake.addFieldArgsForCall)
}

func (fake *Multipart) AddFieldCalls(stub func(string, string) error) {
	fake.addFieldMutex.Lock()
	defer fake.addFieldMutex.Unlock()
	fake.AddFieldStub = stub
}

func (fake *Multipart) AddFieldArgsForCall(i int) (string, string) {
	fake.addFieldMutex.RLock()
	defer fake.addFieldMutex.RUnlock()
	argsForCall := fake.addFieldArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Multipart) AddFieldReturns(result1 error) {
	fake.addFieldMutex.Lock()
	defer fake.addFieldMutex.Unlock()
	fake.AddFieldStub = nil
	fake.addFieldReturns = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) AddFieldReturnsOnCall(i int, result1 error) {
	fake.addFieldMutex.Lock()
	defer fake.addFieldMutex.Unlock()
	fake.AddFieldStub = nil
	if fake.addFieldReturnsOnCall == nil {
		fake.addFieldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFieldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) AddFile(arg1 string, arg2 string) error {
	fake.addFileMutex.Lock()
	ret, specificReturn := fake.addFileReturnsOnCall[len(fake.addFileArgsForCall)]
	fake.addFileArgsForCall = append(fake.addFileArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AddFile", []interface{}{arg1, arg2})
	fake.addFileMutex.Unlock()
	if fake.AddFileStub != nil {
		return fake.AddFileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addFileReturns
	return fakeReturns.result1
}

func (fake *Multipart) AddFileCallCount() int {
	fake.addFileMutex.RLock()
	defer fake.addFileMutex.RUnlock()
	return len(fake.addFileArgsForCall)
}

func (fake *Multipart) AddFileCalls(stub func(string, string) error) {
	fake.addFileMutex.Lock()
	defer fake.addFileMutex.Unlock()
	fake.AddFileStub = stub
}

func (fake *Multipart) AddFileArgsForCall(i int) (string, string) {
	fake.addFileMutex.RLock()
	defer fake.addFileMutex.RUnlock()
	argsForCall := fake.addFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Multipart) AddFileReturns(result1 error) {
	fake.addFileMutex.Lock()
	defer fake.addFileMutex.Unlock()
	fake.AddFileStub = nil
	fake.addFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) AddFileReturnsOnCall(i int, result1 error) {
	fake.addFileMutex.Lock()
	defer fake.addFileMutex.Unlock()
	fake.AddFileStub = nil
	if fake.addFileReturnsOnCall == nil {
		fake.addFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Multipart) Finalize() formcontent.ContentSubmission {
	fake.finalizeMutex.Lock()
	ret, specificReturn := fake.finalizeReturnsOnCall[len(fake.finalizeArgsForCall)]
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct {
	}{})
	fake.recordInvocation("Finalize", []interface{}{})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.finalizeReturns
	return fakeReturns.result1
}

func (fake *Multipart) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *Multipart) FinalizeCalls(stub func() formcontent.ContentSubmission) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = stub
}

func (fake *Multipart) FinalizeReturns(result1 formcontent.ContentSubmission) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 formcontent.ContentSubmission
	}{result1}
}

func (fake *Multipart) FinalizeReturnsOnCall(i int, result1 formcontent.ContentSubmission) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = nil
	if fake.finalizeReturnsOnCall == nil {
		fake.finalizeReturnsOnCall = make(map[int]struct {
			result1 formcontent.ContentSubmission
		})
	}
	fake.finalizeReturnsOnCall[i] = struct {
		result1 formcontent.ContentSubmission
	}{result1}
}

func (fake *Multipart) Reset() {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct {
	}{})
	fake.recordInvocation("Reset", []interface{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		fake.ResetStub()
	}
}

func (fake *Multipart) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *Multipart) ResetCalls(stub func()) {
	fake.resetMutex.Lock()
	defer fake.resetMutex.Unlock()
	fake.ResetStub = stub
}

func (fake *Multipart) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addFieldMutex.RLock()
	defer fake.addFieldMutex.RUnlock()
	fake.addFileMutex.RLock()
	defer fake.addFileMutex.RUnlock()
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Multipart) recordInvocation(key string, args []interface{}) {
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
