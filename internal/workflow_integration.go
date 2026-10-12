package internal

import (
	"strconv"

	failurepb "go.temporal.io/api/failure/v1"
)

// GetLastFailureForIntegration returns borrowed, read-only history metadata.
func GetLastFailureForIntegration(env WorkflowEnvironment) *failurepb.Failure {
	return env.WorkflowInfo().lastFailure
}

// GetCurrentRunIDForIntegration returns the history run identity at this task.
// Reset changes it at the reset point, independently of OriginalRunID.
func GetCurrentRunIDForIntegration(env WorkflowEnvironment) string {
	return env.WorkflowInfo().currentRunID
}

// GenerateChildWorkflowIDForIntegration reserves the same ID used by
// ExecuteChildWorkflow. Pass the returned ID explicitly when scheduling; the
// current run seed can change during reset, unlike OriginalRunID.
func GenerateChildWorkflowIDForIntegration(env WorkflowEnvironment) string {
	if env.WorkflowInfo().currentRunID == "" {
		panic("workflow integration: current run ID is unavailable; specify an explicit child WorkflowID")
	}
	return env.WorkflowInfo().currentRunID + "_" + strconv.FormatInt(env.GenerateSequence(), 10)
}

// UseMemoDataConverterForIntegration records/reads the SDK's memo encoding flag.
// Call during a workflow task before serializing memo values outside the SDK.
func UseMemoDataConverterForIntegration(env WorkflowEnvironment) bool {
	return env.TryUse(SDKFlagMemoUserDCEncode)
}

// GetSearchAttributeUpdatesForIntegration copies values including nil/unset
// updates. GetUntypedValues intentionally omits those deletion markers.
func GetSearchAttributeUpdatesForIntegration(attributes SearchAttributes) map[SearchAttributeKey]any {
	values := make(map[SearchAttributeKey]any, len(attributes.untypedValue))
	for key, value := range attributes.untypedValue {
		if list, ok := value.([]string); ok {
			value = append([]string(nil), list...)
		}
		values[key] = value
	}
	return values
}
