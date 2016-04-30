package hce

import (

)

// Models a reference to a container image, e.g. a Docker image.\n
type Image struct {
	// The (HCE) id of this instance.
	ImageId int32 `json:"image_id,omitempty"`
	// The (HCE) id of the registry instance where this container is hosted.
	ImageRegistryId int32 `json:"image_registry_id,omitempty"`
	// The id of the credentials used to access this container in the registry.\n
	RegistryCredentialId int32 `json:"registry_credential_id,omitempty"`
	// The name of the image's repo in the container registry.
	ImageRepo string `json:"image_repo,omitempty"`
	// The image's tag in the container registry.
	ImageTag string `json:"image_tag,omitempty"`
	// The user-friendly label for this instance. Must be unique.
	ImageLabel string `json:"image_label,omitempty"`
}
