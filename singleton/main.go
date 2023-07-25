package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type lazySingleton struct {
	MakeMalloc int
}

var lazySingletonInstance *lazySingleton

func GetLazySingletonInstance() *lazySingleton {
	once.Do(func() {
		lazySingletonInstance = new(lazySingleton)
	})
	return lazySingletonInstance
}

type eagerSingleton struct {
	MakeMalloc int
}

var eagerSingletonInstance = &eagerSingleton{}

func GetEagerSingletonInstance() *eagerSingleton {
	return eagerSingletonInstance
}

func main() {
	eagerS1 := GetEagerSingletonInstance()
	eagerS2 := GetEagerSingletonInstance()
	fmt.Println("eagerS1 == eagerS2:", eagerS1 == eagerS2)

	var lazyS1 *lazySingleton
	var lazyS2 *lazySingleton
	go func() {
		lazyS1 = GetLazySingletonInstance()
	}()
	go func() {
		lazyS2 = GetLazySingletonInstance()
	}()
	fmt.Println("lazyS1 == lazyS2:", lazyS1 == lazyS2)
}
