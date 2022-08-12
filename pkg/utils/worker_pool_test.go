package utils

import (
	"fmt"
	"strconv"
	"testing"
)

type mockJob struct {
	payload string
}

func (m *mockJob) Execute() {
	fmt.Println(m.payload)
}

func TestWorkerPool(t *testing.T) {
	var request = 1000

	d := NewDispatcher(100, 100)
	d.Run()

	for i := 0; i < request; i++ {
		j := &mockJob{
			payload: strconv.Itoa(i),
		}
		JobQueue <- j
		t.Log("Operation successful")
	}
}
