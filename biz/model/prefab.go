package model

// Prefab 预制体
type Prefab struct {
	MODEL
	PrefabID   int64  `json:"prefab_id"`
	PrefabName string `json:"prefab_name"`
	PrefabURL  string `json:"prefab_url"`
}
