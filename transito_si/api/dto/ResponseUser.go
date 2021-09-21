package dto

type ResponseUser struct {
	HasRegistry bool `json:"hasRegistry"`
	RegistryID uint `json:"registryID"`
}

