package hce

import (
)

// Models a deployment.
type Deployment struct {
    // The internal deployment id.
    Id  int32  `json:"id,omitempty"`
    // The build id to be deployed.
    BuildId  int32  `json:"build_id,omitempty"`
    // The project id to be deployed.
    ProjectId  int32  `json:"project_id,omitempty"`
    // The name of the deployment.
    Name  string  `json:"name,omitempty"`
    // Date time when deployment is created.
    CreatedDate  string  `json:"createdDate,omitempty"`
    // The url where the build is deployed.
    BrowseUrl  string  `json:"browseUrl,omitempty"`
}
