package hce

import (
)

// Models an update to an existing credential. Note that it is not possible to update the credential_type, owner_id, timestamps etc,\nso those fields are omitted.\n
type CredentialUpdate struct {
    // The id of the credential to update.
    CredentialId  int32  `json:"credential_id,omitempty"`
    // The credential key, e.g. the username, or AWS Access Key ID, etc. May be null.
    CredentialKey  string  `json:"credential_key,omitempty"`
    // The credential value, e.g. the password, or AWS Secret Key.
    CredentialValue  string  `json:"credential_value,omitempty"`
    // An optional additional (potentially secret) field, e.g. the user email in the case of a DockerHub authentication.
    CredentialExtra  string  `json:"credential_extra,omitempty"`
    // An optional label for the credential, e.g. \"Alice's GitHub OAuth token\". May be null.
    Label  string  `json:"label,omitempty"`
}
