package singleton

import "sync"

/*
1. 单例模式
*/

// 饿汉单例模式
type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetSingleton() *Singleton {
	return singleton
}

type SingletonV1 struct {
}

// 懒汉单例模式
var (
	lazySingleton *SingletonV1
	once          = &sync.Once{}
)

func GetLazySingleton() *SingletonV1 {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &SingletonV1{}
		})
	}
	return lazySingleton
}
