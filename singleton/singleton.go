package singleton

import "sync"

type Singleton struct {

}

var singleInstance *Singleton

var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() {
		singleInstance = new(Singleton)
	})
	return singleInstance
}