package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_updateSessonTime(t *testing.T) {
	type args struct {
		sessionID string
	}
	tests := []struct {
		name      string
		sessionID string
	}{
		{"first", "4785495499858954"},
		{"second", "894478994894444"},
	}
	fmt.Println("Test_updateSessonTime")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateSessonTime(tt.sessionID)
			for i, t := range TimeGroup {
				fmt.Println(i, t)
			}
			time.Sleep(2 * time.Second)
			checkSessionAlive()
			for i, t := range TimeGroup {
				fmt.Println(i, t)
			}
		})
	}
}
