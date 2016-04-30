package hce

import (
    "time"
)

// Models an artifact such as a log file, or workspace, or binary/executable.
type Artifact struct {
    // The artifact filename
    Name  string  `json:"name,omitempty"`
    // The artifact id.
    Id  int32  `json:"id,omitempty"`
    // The id of the `PipelineExecution` the artifact is associated with.\n
    ExecutionId  int32  `json:"execution_id,omitempty"`
    // The type of artifact, e.g. `BUILD_LOG`.\n
    ArtifactType  string  `json:"artifact_type,omitempty"`
    // Timestamp.
    UploadTimestamp  time.Time  `json:"upload_timestamp,omitempty"`
}
