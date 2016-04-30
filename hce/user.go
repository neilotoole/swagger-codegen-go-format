package hce

import (
)

// Models a user object.
type User struct {
    // The HCE id of the user.
    Id  int32  `json:"id,omitempty"`
    // The GitHub user id.
    UserId  string  `json:"userId,omitempty"`
    // The GitHub login name of the user
    Login  string  `json:"login,omitempty"`
    // The version control system, (eg. github).
    Vcs  string  `json:"vcs,omitempty"`
    // The login token.
    Secret  string  `json:"secret,omitempty"`
}
