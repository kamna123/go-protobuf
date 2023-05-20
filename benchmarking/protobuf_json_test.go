package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"go-protobuf/model/message"
	"log"
	"runtime"
	"runtime/debug"
	"testing"
)

const (
	iteration = 100000000 //Number of iterations for the benchmark test
)

func generateDataset() []*message.MyMessage {
	var dataset []*message.MyMessage

	for i := 0; i < iteration; i++ {
		data := &message.MyMessage{
			Email: "johndoe@example.com",
			Name:  "John Doe",
			Id:    int32(i),
		}
		dataset = append(dataset, data)
	}

	return dataset
}

func BenchmarkProtobufSerialisation(b *testing.B) {
	dataset := generateDataset()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range dataset {
			_, err := proto.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	measureMemoryUsage(b)
}

func BenchmarkJSONSerialization(b *testing.B) {
	dataset := generateDataset()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, data := range dataset {
			_, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	measureMemoryUsage(b)

}
func measureMemoryUsage(b *testing.B) {
	debug.FreeOSMemory()
	var mem runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&mem)
	b.ReportMetric(float64(mem.Alloc)/1024/1024, "Memory_MB")
}

func main() {
	// Run the benchmark tests
	testing.Benchmark(BenchmarkProtobufSerialisation)
	testing.Benchmark(BenchmarkJSONSerialization)

}
