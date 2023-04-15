package executionstatus

import "testing"

var benchmarkString string

func BenchmarkExecutionStatus_String_LastItem(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = ExecutionStatusInterrupted.String()
	}

	benchmarkString = r
}

func BenchmarkExecutionStatus_String_MiddleItem(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = ExecutionStatusStarted.String()
	}

	benchmarkString = r
}

func BenchmarkExecutionStatus_String_FirstItem(b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = ExecutionStatusCreated.String()
	}

	benchmarkString = r
}

var benchmarkValue ExecutionStatus

func BenchmarkExecutionStatusString_LastItem(b *testing.B) {
	var r ExecutionStatus

	for n := 0; n < b.N; n++ {
		r, _ = ExecutionStatusString("Interrupted")
	}

	benchmarkValue = r
}

func BenchmarkExecutionStatusString_MiddleItem(b *testing.B) {
	var r ExecutionStatus

	for n := 0; n < b.N; n++ {
		r, _ = ExecutionStatusString("Started")
	}

	benchmarkValue = r
}

func BenchmarkExecutionStatusString_FirstItem(b *testing.B) {
	var r ExecutionStatus

	for n := 0; n < b.N; n++ {
		r, _ = ExecutionStatusString("Created")
	}

	benchmarkValue = r
}
