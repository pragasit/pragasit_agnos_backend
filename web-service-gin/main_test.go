package main

import (
	"fmt"
	"testing"
)

// func Test_password(t *testing.T) {
// 	// type args struct {
// 	// 	c *gin.Context
// 	// }
// 	// tests := []struct {
// 	// 	name string
// 	// 	args args
// 	// }{}
// 	// for _, tt := range tests {
// 	// 	t.Run(tt.name, func(t *testing.T) {
// 	// 		password(tt.args.c)
// 	// 	})
// 	// }
// 	newPassword := main.requestStruct{
// 		pass: "aA1",
// 	}
// 	newPasswordJSON, _ := json.Marshal(newPassword)

// 	request := httptest.NewRequest("POST", "/strong_password_steps", bytes.NewBuffer(newPasswordJSON))
// 	writer := httptest.NewRecorder()

//		testRoute := gin.Default()
//		testRoute.POST("/strong_password_steps", password)
//		testRoute.ServeHTTP(writer, request)
//	}
func TestIntMinTableDriven(t testing.T) {
	var tests = []struct {
		Init_password string
		want          int
	}{
		{"aA1", 3},
		{"1445D1cd", 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.Init_password, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := password(tt.init_password)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
