package singleton

import "testing"

func BenchmarkSingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetSingleton() != GetSingleton() {
				b.Errorf("test fail")
			}
		}
	})
}

func BenchmarkLayzSingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
