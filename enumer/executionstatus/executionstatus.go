package executionstatus

import (
	"fmt"
)

type ExecutionStatus uint32

// Possible execution status values
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

const _ExecutionStatusName = "CreatedInitVUsInitExecutorsInitDonePausedBeforeRunStartedSetupRunningTeardownEndedInterrupted"

var _ExecutionStatusIndex = [...]uint8{0, 7, 14, 27, 35, 50, 57, 62, 69, 77, 82, 93}

func (i ExecutionStatus) String() string {
	if i >= ExecutionStatus(len(_ExecutionStatusIndex)-1) {
		return fmt.Sprintf("ExecutionStatus(%d)", i)
	}
	return _ExecutionStatusName[_ExecutionStatusIndex[i]:_ExecutionStatusIndex[i+1]]
}

var _ExecutionStatusValues = []ExecutionStatus{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var _ExecutionStatusNameToValueMap = map[string]ExecutionStatus{
	_ExecutionStatusName[0:7]:   0,
	_ExecutionStatusName[7:14]:  1,
	_ExecutionStatusName[14:27]: 2,
	_ExecutionStatusName[27:35]: 3,
	_ExecutionStatusName[35:50]: 4,
	_ExecutionStatusName[50:57]: 5,
	_ExecutionStatusName[57:62]: 6,
	_ExecutionStatusName[62:69]: 7,
	_ExecutionStatusName[69:77]: 8,
	_ExecutionStatusName[77:82]: 9,
	_ExecutionStatusName[82:93]: 10,
}

// ExecutionStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ExecutionStatusString(s string) (ExecutionStatus, error) {
	if val, ok := _ExecutionStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ExecutionStatus values", s)
}

// ExecutionStatusValues returns all values of the enum
func ExecutionStatusValues() []ExecutionStatus {
	return _ExecutionStatusValues
}

// IsAExecutionStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ExecutionStatus) IsAExecutionStatus() bool {
	for _, v := range _ExecutionStatusValues {
		if i == v {
			return true
		}
	}
	return false
}
