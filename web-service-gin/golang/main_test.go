package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_password(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{"aA1", args{}},
		{"1445D1cd", args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password(tt.args.c)
		})
	}
}
