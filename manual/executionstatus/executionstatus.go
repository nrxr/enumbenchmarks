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

var valuesInMap = map[string]ExecutionStatus{
	"Created":         ExecutionStatusCreated,
	"InitVUs":         ExecutionStatusInitVUs,
	"InitExecutors":   ExecutionStatusInitExecutors,
	"InitDone":        ExecutionStatusInitDone,
	"PausedBeforeRun": ExecutionStatusPausedBeforeRun,
	"Started":         ExecutionStatusStarted,
	"Setup":           ExecutionStatusSetup,
	"Running":         ExecutionStatusRunning,
	"Teardown":        ExecutionStatusTeardown,
	"Ended":           ExecutionStatusEnded,
	"Interrupted":     ExecutionStatusInterrupted,
}

func (es ExecutionStatus) String() string {
	return valuesInSlice[es]
}

func ExecutionStatusString(s string) (ExecutionStatus, error) {
	if v, ok := valuesInMap[s]; ok {
		return v, nil
	}

	return 0, nil
}
