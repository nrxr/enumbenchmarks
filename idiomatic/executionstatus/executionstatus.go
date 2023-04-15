package executionstatus

type ExecutionStatus uint8

const (
	ExecutionStatusCreated ExecutionStatus = iota
	ExecutionStatusInitVUs
	ExecutionStatusInitExecutors
	ExecutionStatusInitDone
	ExecutionStatusPausedBeforeRun
	ExecutionStatusStarted
	ExecutionStatusSetup
	ExecutionStatusRunning
	ExecutionStatusTeardown
	ExecutionStatusEnded
	ExecutionStatusInterrupted
)

// idiomatic way to build this
// ref: https://cs.opensource.google/go/go/+/refs/tags/go1.20.3:src/net/http/status.go;l=9
func (es ExecutionStatus) String() string {
	switch es {
	case ExecutionStatusCreated:
		return "Created"
	case ExecutionStatusInitVUs:
		return "InitVUs"
	case ExecutionStatusInitExecutors:
		return "InitExecutors"
	case ExecutionStatusInitDone:
		return "InitDone"
	case ExecutionStatusPausedBeforeRun:
		return "PausedBeforeRun"
	case ExecutionStatusStarted:
		return "Started"
	case ExecutionStatusSetup:
		return "Setup"
	case ExecutionStatusRunning:
		return "Running"
	case ExecutionStatusTeardown:
		return "Teardown"
	case ExecutionStatusEnded:
		return "Ended"
	case ExecutionStatusInterrupted:
		return "Interrupted"
	default:
		return ""
	}
}

func ExecutionStatusString(s string) (ExecutionStatus, error) {
	switch s {
	case "Created":
		return ExecutionStatusCreated, nil
	case "InitVUs":
		return ExecutionStatusInitVUs, nil
	case "InitExecutors":
		return ExecutionStatusInitExecutors, nil
	case "InitDone":
		return ExecutionStatusInitDone, nil
	case "PausedBeforeRun":
		return ExecutionStatusPausedBeforeRun, nil
	case "Started":
		return ExecutionStatusStarted, nil
	case "Setup":
		return ExecutionStatusSetup, nil
	case "Running":
		return ExecutionStatusRunning, nil
	case "Teardown":
		return ExecutionStatusTeardown, nil
	case "Ended":
		return ExecutionStatusEnded, nil
	case "Interrupted":
		return ExecutionStatusInterrupted, nil
	default:
		return 0, nil
	}
}
