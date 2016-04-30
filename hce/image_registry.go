package hce

import (

)

// Models a reference to a container registry, e.g. DockerHub.\n
type ImageRegistry struct {
	// The (HCE) id of this instance.
	ImageRegistryId int32 `json:"image_registry_id,omitempty"`
	// The container technology type. Currently we only suppot docker.
	ImageType string `json:"image_type,omitempty"`
	// The URL for API operations on this registry.
	RegistryUrl string `json:"registry_url,omitempty"`
	// The user-friendly label for this VCS instance. Must be unique.
	RegistryLabel string `json:"registry_label,omitempty"`
}
