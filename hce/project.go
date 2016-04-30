package hce

import (
)

// Models an individual project.\n
type Project struct {
    // The internal project id.
    Id  int32  `json:"id,omitempty"`
    // The user-provided label for the project.
    Name  string  `json:"name,omitempty"`
    // The platform type of the project, e.g. 'nodejs' or 'java', etc.
    Type_  string  `json:"type,omitempty"`
    // The internal user id of the project owner.
    UserId  int32  `json:"user_id,omitempty"`
    // The (semi-secret) code that an invited user users to join the project.
    JoinCode  string  `json:"joinCode,omitempty"`
    // The GitHub token to use (presumably???)
    Token  string  `json:"token,omitempty"`
    // The branch to use (FIXME: This is a duplicate of what's in \"Repo\")\n
    BranchRefName  string  `json:"branchRefName,omitempty"`
    
    Repo  Repo  `json:"repo,omitempty"`
    // An array of target deployment environments (such as a CloudFoundry or AWS instance) that the project is deployed to.
    Targets  []DeploymentTarget  `json:"targets,omitempty"`
}
