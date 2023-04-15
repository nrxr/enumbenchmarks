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

var valuesInSlice = []string{
	"Created",
	"InitVUs",
	"InitExecutors",
	"InitDone",
	"PausedBeforeRun",
	"Started",
	"Setup",
	"Running",
	"Teardown",
	"Ended",
	"Interrupted",
}

func (es ExecutionStatus) String() string {
	return valuesInSlice[es]
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
