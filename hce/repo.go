package hce

import (
)

// Models a Version Control System repository, such as a GitHub repo, and references a specific branch.
type Repo struct {
    // The repo type, e.g. `github`
    Vcs  string  `json:"vcs,omitempty"`
    // The full name of the repo (typically `owner/repo`), e.g. `neilotoole/hello-node`.
    FullName  string  `json:"full_name,omitempty"`
    // The owner of the repo (as determined by the repo itself), e.g. `neilotoole`
    Owner  string  `json:"owner,omitempty"`
    // The name of the repo, e.g. `hello-node`.
    Name  string  `json:"name,omitempty"`
    // For GitHub repos, the unique GitHub identifier.
    GithubRepoId  string  `json:"githubRepoId,omitempty"`
    // The repo branch being referenced, e.g. `master`.
    Branch  string  `json:"branch,omitempty"`
    // The HTTPS URL used to clone the repo.
    CloneUrl  string  `json:"cloneUrl,omitempty"`
    // The SSH URL used to clone the repo.
    SshUrl  string  `json:"sshUrl,omitempty"`
    // The repo home page.
    HttpUrl  string  `json:"httpUrl,omitempty"`
    // On GitHub, this is the unique id of the webhook for this repo that calls back to HelionCE.
    WebHookId  string  `json:"webHookId,omitempty"`
}
