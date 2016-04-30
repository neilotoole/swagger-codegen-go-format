package hce

import (

)

// Used to specify/store a credential.
type CredentialDetail struct {
	// The user id of the owner of this credential. May be null.
	OwnerId int32 `json:"owner_id,omitempty"`
	// The type of credential.
	CredentialType string `json:"credential_type,omitempty"`
	// The credential key, e.g. the username, or AWS Access Key ID, etc. May be null.
	CredentialKey string `json:"credential_key,omitempty"`
	// The credential value, e.g. the password, or AWS Secret Key.
	CredentialValue string `json:"credential_value,omitempty"`
	// An optional additional (secret) field in the case of a three-value credential, e.g. the user email of a DockerHub credential.
	CredentialExtra string `json:"credential_extra,omitempty"`
	// An optional label for the credential, e.g. \"Alice's GitHub OAuth token\". May be null.
	Label string `json:"label,omitempty"`
}
