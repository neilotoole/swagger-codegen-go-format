package hce

import (
)

// Models a generic pipeline trigger event.\n
type PipelineTrigger struct {
    // The (HCE) id of the instance.
    Id  int32  `json:"id,omitempty"`
    
    ProjectId  int32  `json:"project_id,omitempty"`
    
    UserId  int32  `json:"user_id,omitempty"`
    
    CommitRef  string  `json:"commit_ref,omitempty"`
}
