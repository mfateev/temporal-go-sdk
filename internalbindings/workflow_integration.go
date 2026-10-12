package internalbindings

import (
	failurepb "go.temporal.io/api/failure/v1"
	"go.temporal.io/sdk/internal"
)

// IntegrationVersion versions the additional bindings maintained by this SDK
// fork. It does not promise stability for upstream internalbindings APIs.
const IntegrationVersion = 1

// GetLastFailure returns borrowed, read-only previous-run history metadata.
// Integrations must copy/serialize it before exposing it to another owner.
func GetLastFailure(env WorkflowEnvironment) *failurepb.Failure {
	return internal.GetLastFailureForIntegration(env)
}

// GetCurrentRunID returns the deterministic, reset-aware run identity. Read it
// before admitting workflow code on each task; it does not consume a sequence ID.
func GetCurrentRunID(env WorkflowEnvironment) string {
	return internal.GetCurrentRunIDForIntegration(env)
}

// GenerateChildWorkflowID reserves a child ID using the current reset-aware run
// seed and SDK sequence. Use that ID in ExecuteWorkflowParams.WorkflowID.
func GenerateChildWorkflowID(env WorkflowEnvironment) string {
	return internal.GenerateChildWorkflowIDForIntegration(env)
}

// UseMemoDataConverter selects memo serialization using recorded SDK flags.
// Call during the workflow task; it may record a flag on new histories.
func UseMemoDataConverter(env WorkflowEnvironment) bool {
	return internal.UseMemoDataConverterForIntegration(env)
}

// GetSearchAttributeUpdates returns a private copy including nil deletion
// markers. Keyword-list values are copied; keys are immutable SDK descriptors.
func GetSearchAttributeUpdates(attributes internal.SearchAttributes) map[internal.SearchAttributeKey]any {
	return internal.GetSearchAttributeUpdatesForIntegration(attributes)
}
