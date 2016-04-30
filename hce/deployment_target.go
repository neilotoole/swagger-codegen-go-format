package hce

import (

)

// Models a deployment target/environment (e.g. a CloudFoundry instance) that a project can be deployed to.
type DeploymentTarget struct {
	// The internal id of this target instance.
	Id int32 `json:"id,omitempty"`
	// The owner's user id.
	UserId int32 `json:"user_id,omitempty"`
	// The type of target environment, e.g. `cloudfoundry`, `aws`, etc..
	Type_ string `json:"type,omitempty"`
	// The URL endpoint that the target is accessible at.
	Url string `json:"url,omitempty"`
	// The username to authenticate to the target with.
	UserName string `json:"userName,omitempty"`
	// The password to authenticate to the target with.
	Password string `json:"password,omitempty"`
	// The user-provided label for this target.
	Name string `json:"name,omitempty"`
	// The organization under which a project will be deployed to this target.
	Organization string `json:"organization,omitempty"`
	// The space within an organization that a project will be deployed under on this target.
	Space string `json:"space,omitempty"`
}
