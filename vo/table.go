package vo

import modelcore "github.com/sweetrpg/model-core.go/vo"

// Table value object.
// This value object is a serializable representation of the Table model.
type TableVO struct {
	ID           string            `json:"id" jsonapi:"primary,table"`
	UserID       string            `json:"user_id" jsonapi:"relation,user"`
	Name         string            `json:"name" jsonapi:"attr,name"`
	VolumeIDs    []string          `json:"volume_ids" jsonapi:"relation,volume"`
	VolumeTitles map[string]string `json:"volume_titles" jsonapi:"attr,volume_titles"`
	Visibility   string            `json:"visibility" jsonapi:"attr,visibility"`
	modelcore.AuditableVO
}
