package model

// Prefab 预制体
type Prefab struct {
	MODEL
	PrefabName string `json:"prefab_name"`
	PrefabURL  string `json:"prefab_url"`
}
