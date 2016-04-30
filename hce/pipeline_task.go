package hce

import (
)

// Models a instance of a task to perform in a specific project pipeline. For example, a\npost-deployment task, such as execution of StormRunner on \"My First Project\".\n
type PipelineTask struct {
    // The (HCE) id of this task instance.
    Id  int32  `json:"id,omitempty"`
    // The name of the task, e.g. \"StormRunner\".
    Name  string  `json:"name,omitempty"`
    // The (HCE) id of the project to execute this task on.
    ProjectId  int32  `json:"project_id,omitempty"`
}
