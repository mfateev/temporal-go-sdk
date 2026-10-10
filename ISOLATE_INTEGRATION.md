# Additional SDK integration bindings

This branch starts from upstream v1.49.0. Ordinary Go SDK workflows retain their
existing APIs and execution path. The patch adds four `internalbindings` helpers
for external workflow dispatchers:

- `GetLastFailure` reads previous-run failure metadata without exposing a private
  Go struct layout. Its borrowed protobuf is read-only and must be copied.
- `GenerateChildWorkflowID` reserves the SDK's reset-aware child ID before an
  external dispatcher encodes input with a workflow serialization context.
- `UseMemoDataConverter` preserves the recorded SDK memo encoding flag without
  exporting or duplicating its numeric value.
- `GetSearchAttributeUpdates` copies typed values, retaining unset markers.

`internalbindings.IntegrationVersion` is 1. Integrations pin an audited module
revision and validate this version. Upstream `internalbindings` still makes no
compatibility promise; these additions are supported by this fork only. They
are small enough to propose upstream independently of the isolate runtime.

These helpers do not change history commands, scheduling, codecs, or default
behavior. Call sequence/flag helpers only while processing a workflow task.

The branch's `Isolate integration bindings` CI uses standard Go 1.26.7 for the
canonical quality checks, focused binding tests and reset/typed-attribute
integration tests with default and zero workflow cache. Custom-toolchain builds,
ownership checks and saved-history replay run in the Go isolates repository's
native Linux/macOS arm64/amd64 CI.
