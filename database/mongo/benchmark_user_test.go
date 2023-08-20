package mongodb

import (
	// mongodb "main/database/mongo"
	"main/models/user"
	"testing"
)

func Benchmark_CreateUserBatch(b *testing.B) {
	Connect()
	user := userModel.User{
		Name:     "benchmark",
		Nickname: "speedtest",
		Birth:    "2009-04-3",
		Stack: []string{
			"go", "go", "go", "go", "go", "go", "go", "go",
			"go", "go", "go", "go", "go", "go", "go", "go",
		},
	}
	for i := 0; i < b.N; i++ {
		Builder().
			UseDatabase("benchmark").
		
      UseCollection("users").CreateUser(user)
	}
}
