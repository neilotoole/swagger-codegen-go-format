package hce

import (
)

// Models a pipeline event, such as the start of a build; tests passing; or a deployment, etc.\n
type PipelineEvent struct {
    // Pipeline event id.
    Id  int32  `json:"id,omitempty"`
    // Pipeline event name.
    Name  string  `json:"name,omitempty"`
    
    Type_  string  `json:"type,omitempty"`
    
    State  string  `json:"state,omitempty"`
    
    StartDate  string  `json:"startDate,omitempty"`
    
    EndDate  string  `json:"endDate,omitempty"`
    // Build id.
    ExecutionId  int32  `json:"execution_id,omitempty"`
    // Artifact id.
    ArtifactId  int32  `json:"artifact_id,omitempty"`
}
