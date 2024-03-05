package main

import (
	"testing"
)

const n = 1000000

func BenchmarkAppend1(b *testing.B) {
	tmp := make([]int, n) // 创建一个足够大的切片来模拟实际使用情况
	for i := 0; i < n; i++ {
		tmp[i] = i // 初始化切片
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		append1(tmp)
	}
}

func BenchmarkAppend2(b *testing.B) {
	tmp := make([]int, n) // 创建一个足够大的切片来模拟实际使用情况
	for i := 0; i < n; i++ {
		tmp[i] = i // 初始化切片
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		append2(tmp)
	}
}
