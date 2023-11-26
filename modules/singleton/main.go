package main

import (
	"fmt"
	"sync"
)

// Одиночка гарантирует существование только одного объекта определённого класса,
// а также позволяет достучаться до этого объекта из любого места программы

// Вы не сможете просто взять и использовать класс,
// зависящий от одиночки в другой программе.
// Для этого придётся эмулировать присутствие одиночки и там.
// Чаще всего эта проблема проявляется при написании юнит-тестов.

type single struct{}

var (
	singleInstance *single
	once           sync.Once
)

func getInstance(i int) *single {
	once.Do(func() {
		fmt.Printf("Creating single instance now. %d\n", i)
		singleInstance = &single{}
	})
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance(i)
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
