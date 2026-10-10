package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
	failurepb "go.temporal.io/api/failure/v1"
)

type integrationEnvironment struct {
	WorkflowEnvironment
	info     WorkflowInfo
	sequence int64
	useUser  bool
	flags    []sdkFlag
}

func (e *integrationEnvironment) WorkflowInfo() *WorkflowInfo { return &e.info }
func (e *integrationEnvironment) GenerateSequence() int64     { e.sequence++; return e.sequence }
func (e *integrationEnvironment) TryUse(flag sdkFlag) bool {
	e.flags = append(e.flags, flag)
	return e.useUser
}

func TestWorkflowIntegrationHistoryMetadata(t *testing.T) {
	e := &integrationEnvironment{}
	require.Nil(t, GetLastFailureForIntegration(e))
	e.info.lastFailure = &failurepb.Failure{Message: "previous failure"}
	require.Same(t, e.info.lastFailure, GetLastFailureForIntegration(e))
	e.info.currentRunID, e.info.OriginalRunID = "current", "original"
	require.Equal(t, "current_1", GenerateChildWorkflowIDForIntegration(e))
	e.info.currentRunID = "reset"
	require.Equal(t, "reset_2", GenerateChildWorkflowIDForIntegration(e))
	require.False(t, UseMemoDataConverterForIntegration(e))
	e.useUser = true
	require.True(t, UseMemoDataConverterForIntegration(e))
	require.Equal(t, []sdkFlag{SDKFlagMemoUserDCEncode, SDKFlagMemoUserDCEncode}, e.flags)
}

func TestWorkflowIntegrationSearchAttributeUpdates(t *testing.T) {
	list, removed := NewSearchAttributeKeyKeywordList("list"), NewSearchAttributeKeyInt64("removed")
	attributes := NewSearchAttributes(list.ValueSet([]string{"first"}), removed.ValueUnset())
	values := GetSearchAttributeUpdatesForIntegration(attributes)
	value, present := values[removed]
	require.True(t, present)
	require.Nil(t, value)
	values[list].([]string)[0] = "changed"
	delete(values, removed)
	original := GetSearchAttributeUpdatesForIntegration(attributes)
	require.Equal(t, []string{"first"}, original[list])
	_, present = original[removed]
	require.True(t, present)
}
