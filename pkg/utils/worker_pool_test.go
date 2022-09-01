package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockJob struct {
	payload []int
	index   int
}

func (m *mockJob) Execute() {
	m.payload[m.index] = 1
}

func Test_WorkerPool(t *testing.T) {
	var request = 1000

	d := NewDispatcher(100, 100)
	d.Run()

	a := make([]int, request)
	for i := 0; i < request; i++ {
		j := &mockJob{
			payload: a,
			index:   i,
		}
		d.JobQueue <- j
	}

	for _, val := range a {
		assert.Equal(t, val, 1)
	}
}
func Benchmark_Worker_Pool(b *testing.B) {
	d := NewDispatcher(100, 100)
	d.Run()

	a := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		j := &mockJob{
			payload: a,
			index:   i,
		}
		d.JobQueue <- j
	}
}
