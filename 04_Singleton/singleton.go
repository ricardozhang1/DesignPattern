package _4_Singleton

import "sync"

/*
保证一个类仅有一个实例，并提供一个访问它的全局访问点。
*/

type SingleData struct {
	data int
}

var data *SingleData
// 通过该类型可以实现单例模式，虽然是多次赋值。
// 但是只执行一次
// 一个对象多次实例化，但只有一个共享对象地址
var once sync.Once

func getInstance(i int) *SingleData {
	once.Do(func() {
		data = &SingleData{i}
	})
	return data
}





