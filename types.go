package figma

import (
	"github.com/tmc/figma/figmatypes"
	"github.com/tmc/figma/nodes"
)

// Project refers to a Figma project.
type Project struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ComponentReference is a reference to a component.
type ComponentReference struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// FileMeta stores metadata about a file.
type FileMeta struct {
	Key          string `json:"key,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
	Name         string `json:"name,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
}

// Image is the response to generating an image.
type Image struct {
	Status float64           `json:"status,omitempty"`
	Images map[string]string `json:"images,omitempty"`
	Err    string            `json:"err,omitempty"`
}

// File is a Figma file.
type File struct {
	Name          string                                    `json:"name,omitempty"`
	LastModified  string                                    `json:"lastModified,omitempty"`
	ThumbnailURL  string                                    `json:"thumbnailUrl,omitempty"`
	Document      nodes.Document                            `json:"document,omitempty"`
	SchemaVersion int                                       `json:"schemaVersion"`
	Styles        map[figmatypes.StyleType]figmatypes.Style `json:"styles"`
	Components    map[string]ComponentReference             `json:"components,omitempty"`
}

// Comment is a comment or reply left by a user.
type Comment struct {
	// Unique identifier for comment.
	ID string `json:"id"`
	// The content of the comment.
	Message string `json:"message"`
	// The position of the comment. Either the absolute coordinates on the canvas or a relative offset within a frame.
	ClientMeta figmatypes.VectorOrFrameOffset `json:"client_meta,omitempty"`
	// The file in which the comment lives.
	FileKey string `json:"file_key,omitempty"`
	// If present, the id of the comment to which this is the reply.
	ParentID string `json:"parent_id,omitempty"`
	// The user who left the comment.
	User User `json:"user,omitempty"`
	// The UTC ISO 8601 time at which the comment was left.
	CreatedAt string `json:"created_at,omitempty"`
	// If set, the UTC ISO 8601 time the comment was resolved.
	ResolvedAt string `json:"resolved_at,omitempty"`
	// Only set for top level comments. The number displayed with the comment in the UI.
	OrderID int `json:"order_id,omitempty"`
}

// User is a description of a user.
type User struct {
	// Name of the user.
	Handle string `json:"handle"`
	// URL link to the user's profile image.
	ImgURL string `json:"img_url"`
}

// Version is a version of a file.
type Version struct {
	// Unique identifier for version.
	ID string `json:"id"`
	// The UTC ISO 8601 time at which the version was created.
	CreatedAt string `json:"created_at,omitempty"`
	// The label given to the version in the editor.
	Label string `json:"label,omitempty"`
	// The description of the version as entered in the editor.
	Description string `json:"description,omitempty"`
	// The user that created the version.
	User User `json:"user,omitempty"`
}
