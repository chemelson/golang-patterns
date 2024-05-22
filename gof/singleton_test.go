package gof

// All implementations consider concurrency.
// Also you can possibly use init() for any file in package

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleton(_ *testing.T) {
	singletonWithMutex := getSingleWithLockInstance()
	fmt.Println(singletonWithMutex)

	singletonWithOnce := getSingleWithOnceInstance()
	fmt.Println(singletonWithOnce)
}

// Singleton with Mutex
var lock = &sync.Mutex{}

type singleWithMutex struct{}

var singleWithMutexInstance *singleWithMutex

func getSingleWithLockInstance() *singleWithMutex {
	if singleWithMutexInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleWithMutexInstance == nil {
			singleWithMutexInstance = &singleWithMutex{}
		}
	}
	return singleWithMutexInstance
}

// Singleton with Once
var once sync.Once

type singleWithOnce struct{}

var singleWithOnceInstance *singleWithOnce

func getSingleWithOnceInstance() *singleWithOnce {
	if singleWithOnceInstance == nil {
		once.Do(
			func() {
				singleWithOnceInstance = &singleWithOnce{}
			})
	}
	return singleWithOnceInstance
}
