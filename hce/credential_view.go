package hce

import (
	"time"
)

// A view of the non-secret elements of a credential.
type CredentialView struct {
	// The (HCE) id of this credential.
	CredentialId int32 `json:"credential_id,omitempty"`
	// The type of credential.
	CredentialType string `json:"credential_type,omitempty"`
	// The user id of the owner of this credential. May be null.
	OwnerId int32 `json:"owner_id,omitempty"`
	// An optional label for the credential, e.g. \"Alice's GitHub OAuth token\". May be null.
	Label string `json:"label,omitempty"`
	// Timestamp
	Created time.Time `json:"created,omitempty"`
	// Timestamp
	Modified time.Time `json:"modified,omitempty"`
}
