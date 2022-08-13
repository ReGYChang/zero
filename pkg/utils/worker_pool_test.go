package utils

import (
	"encoding/json"
	"testing"
)

type mockJob struct {
	payload []byte
}

func (m *mockJob) Execute() {
	b := m.payload
	for i := 0; i < 100; i++ {
		var mt map[string]interface{}

		_ = json.Unmarshal(b, &mt)
		b, _ = json.Marshal(mt)
	}
}

func Test_WorkerPool(t *testing.T) {
	var request = 1000

	d := NewDispatcher(100, 100)
	d.Run()

	for i := 0; i < request; i++ {
		j := &mockJob{
			payload: []byte(`
				"person": {
					"name": {
					  "first": "Leonid",
					  "last": "Bugaev",
					  "fullName": "Leonid Bugaev"
					},
					"github": {
					  "handle": "buger",
					  "followers": 109
					},
					"avatars": [
					  { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
					]
				  },
				  "company": {
					"name": "Acme"
				  }
			`),
		}
		JobQueue <- j
		t.Log("Operation successful")
	}
}

func Benchmark_Worker_Pool(b *testing.B) {
	d := NewDispatcher(100, 100)
	d.Run()

	for i := 0; i < b.N; i++ {
		j := &mockJob{
			payload: []byte(`
				"person": {
					"name": {
					  "first": "Regy",
					  "last": "Chang",
					  "fullName": "Handsome Guy"
					},
					"github": {
					  "handle": "regy",
					  "followers": 3
					},
					"avatars": [
					  { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
					]
				  },
				  "company": {
					"name": "NexAIoT"
				  }
			`),
		}
		JobQueue <- j
	}
}
