package figma

import (
	"encoding/json"
)

// DefaultID is the root id in a figma file (the document).
const DefaultID = "0:0"

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

// Style is the name of a style.
type Style string

// StyleType is the type of a style.
type StyleType string

const (
	styleTypeFILL   StyleType = "FILL"
	StyleTypeSTROKE           = "STROKE"
	StyleTypeTEXT             = "TEXT"
	StyleTypeEFFECT           = "EFFECT"
	StyleTypeGRID             = "GRID"
)

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
	Name          string                        `json:"name,omitempty"`
	LastModified  string                        `json:"lastModified,omitempty"`
	ThumbnailURL  string                        `json:"thumbnailUrl,omitempty"`
	Document      DocumentNode                  `json:"document,omitempty"`
	SchemaVersion int                           `json:"schemaVersion"`
	Styles        map[StyleType]Style           `json:"styles"`
	Components    map[string]ComponentReference `json:"components,omitempty"`
}

// VectorOrFrameOffset contains the fields from Vector and FrameOffset.
type VectorOrFrameOffset struct {
	X          float64 `json:"x,omitempty"`
	Y          float64 `json:"y,omitempty"`
	NodeID     string  `json:"node_id,omitempty"`
	NodeOffset *Vector `json:"node_offset,omitempty"`
}

// Comment is a comment or reply left by a user.
type Comment struct {
	// Unique identifier for comment.
	ID string `json:"id"`
	// The content of the comment.
	Message string `json:"message"`
	// The position of the comment. Either the absolute coordinates on the canvas or a relative offset within a frame.
	ClientMeta VectorOrFrameOffset `json:"client_meta,omitempty"`
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

// Types

// Color is an RGBA Color.
type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

// FormatType describes the possible image formats.
type FormatType string

const (
	FormatTypeJPG FormatType = "JPG"
	FormatTypePNG            = "PNG"
	FormatTypeSVG            = "SVG"
)

// ExportSetting describes export settings for a Figma object.
type ExportSetting struct {
	Suffix     string     `json:"suffix"`
	Format     FormatType `json:"format"`
	Constraint Constraint `json:"constraint"`
}

// ConstraintType describes a type of Constraint.
type ConstraintType string

const (
	ConstraintTypeSCALE  ConstraintType = "SCALE"
	ConstraintTypeWIDTH                 = "WIDTH"
	ConstraintTypeHEIGHT                = "HEIGHT"
)

// A Constraint is a sizing constraint for exports.
type Constraint struct {
	Type  ConstraintType `json:"type"`
	Value float64        `json:"value"`
}

// A Rectangle expresses a bounding box in absolute coordinates
type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// BlendMode describes how layer blends with layers below.
type BlendMode string

const (
	BlendModePASS_THROUGH BlendMode = "PASS_THROUGH"
	BlendModeNORMAL                 = "NORMAL"
	BlendModeDARKEN                 = "DARKEN"
	BlendModeMULTIPLY               = "MULTIPLY"
	BlendModeLINEAR_BURN            = "LINEAR_BURN"
	BlendModeCOLOR_BURN             = "COLOR_BURN"
	BlendModeLIGHTEN                = "LIGHTEN"
	BlendModeSCREEN                 = "SCREEN"
	BlendModeLINEAR_DODGE           = "LINEAR_DODGE"
	BlendModeCOLOR_DODGE            = "COLOR_DODGE"
	BlendModeOVERLAY                = "OVERLAY"
	BlendModeSOFT_LIGHT             = "SOFT_LIGHT"
	BlendModeHARD_LIGHT             = "HARD_LIGHT"
	BlendModeDIFFERENCE             = "DIFFERENCE"
	BlendModeEXCLUSION              = "EXCLUSION"
	BlendModeHUE                    = "HUE"
	BlendModeSATURATION             = "SATURATION"
	BlendModeCOLOR                  = "COLOR"
	BlendModeLUMINOSITY             = "LUMINOSITY"
)

type LayoutConstraintVertical string

const (
	LayoutConstraintVerticalTOP        LayoutConstraintVertical = "TOP"
	LayoutConstraintVerticalBOTTOM                              = "BOTTOM"
	LayoutConstraintVerticalCENTER                              = "CENTER"
	LayoutConstraintVerticalTOP_BOTTOM                          = "TOP_BOTTOM"
	LayoutConstraintVerticalSCALE                               = "SCALE"
)

type LayoutConstraintHorizontal string

const (
	LayoutConstraintHoritontalLEFT       LayoutConstraintHorizontal = "LEFT"
	LayoutConstraintHoritontalRIGHT                                 = "RIGHT"
	LayoutConstraintHoritontalCENTER                                = "CENTER"
	LayoutConstraintHoritontalLEFT_RIGHT                            = "LEFT_RIGHT"
	LayoutConstraintHoritontalSCALE                                 = "SCALE"
)

type LayoutConstraint struct {
	Vertical   LayoutConstraintVertical   `json:"vertical,omitempty"`
	Horizontal LayoutConstraintHorizontal `json:"horizontal,omitempty"`
}

type LayoutGridPattern string

const (
	LayoutGridPatternCOLUMNS LayoutGridPattern = "COLUMNS"
	LayoutGridPatternROWS                      = "ROWS"
	LayoutGridPatternGRID                      = "GRID"
)

type LayoutGridAlignment string

const (
	LayoutGridAlignmentMIN    LayoutGridAlignment = "MIN"
	LayoutGridAlignmentMAX                        = "MAX"
	LayoutGridAlignmentCENTER                     = "CENTER"
)

type LayoutGrid struct {
	Pattern     LayoutGridPattern   `json:"pattern,omitempty"`
	SectionSize float64             `json:"sectionSize,omitempty"`
	Color       Color               `json:"color,omitempty"`
	Alignment   LayoutGridAlignment `json:"alignment,omitempty"`
	GutterSize  int                 `json:"gutterSize"`
	Count       int                 `json:"count,omitempty"`
	Offset      int                 `json:"offset"`
	Visible     *bool               `json:"visible,omitempty"`
}

type EffectType string

const (
	EffectTypeINNER_SHADOW    EffectType = "INNER_SHADOW"
	EffectTypeDROP_SHADOW                = "DROP_SHADOW"
	EffectTypeLAYER_BLUR                 = "LAYER_BLUR"
	EffectTypeBACKGROUND_BLUR            = "BACKGROUND_BLUR"
)

type Effect struct {
	Type      EffectType `json:"type"`
	Visible   *bool      `json:"visible,omitempty"`
	Color     Color      `json:"color,omitempty"`
	BlendMode string     `json:"blendMode,omitempty"`
	Radius    float64    `json:"radius,omitempty"`
	Offset    Vector     `json:"offset,omitempty"`
}

type PaintType string

const (
	PaintTypeSOLID            PaintType = "SOLID"
	PaintTypeGRADIENT_LINEAR            = "GRADIENT_LINEAR"
	PaintTypeGRADIENT_RADIAL            = "GRADIENT_RADIAL"
	PaintTypeGRADIENT_ANGULAR           = "GRADIENT_ANGULAR"
	PaintTypeGRADIENT_DIAMOND           = "GRADIENT_DIAMOND"
	PaintTypeIMAGE                      = "IMAGE"
	PaintTypeEMOJI                      = "EMOJI"
)

type ScaleMode string

const (
	ScaleModeFILL    ScaleMode = "FILL"
	ScaleModeFIT               = "FIT"
	ScaleModeTILE              = "TILE"
	ScaleModeSTRETCH           = "STRETCH"
)

// Paint is a solid color, gradient, or image texture that can be applied as fills or strokes.
type Paint struct {
	Type PaintType `json:"type"`
	// Is the paint enabled?. default: true.
	Visible *bool `json:"visible,omitempty"`
	// Overall opacity of paint (colors within the paint can also have opacity values which would blend with this). default: 1.
	Opacity float64 `json:"opacity,omitempty"`
	// Solid color of the paint.
	Color *Color `json:"color,omitempty"`
	// This field contains three vectors, each of which are a position in normalized object space (normalized object space is if the top left corner of the bounding box of the object is (0, 0) and the bottom right is (1,1)). The first position corresponds to the start of the gradient (value 0 for the purposes of calculating gradient stops), the second position is the end of the gradient (value 1), and the third handle position determines the width of the gradient (only relevant for non-linear gradients). See image examples below:.
	GradientHandlePositions []Vector `json:"gradientHandlePositions,omitempty"`

	// Positions of key points along the gradient axis with the colors anchored there. Colors along the gradient are interpolated smoothly between neighboring gradient stops..
	GradientStops []ColorStop `json:"gradientStops,omitempty"`
	// Image scaling mode.
	ScaleMode string `json:"scaleMode,omitempty"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Transform [][]float64

type Path struct {
	Path        string `json:"path"`
	WindingRule string `json:"winding_rule"`
}

type FrameOffset struct {
	NodeID     string `json:"node_id"`
	NodeOffset Vector `json:"node_offset"`
}

type ColorStop struct {
	Position float64 `json:"position"`
	Color    Color   `json:"color,omitempty"`
}

type TypeStyle struct {
	FontFamily          string  `json:"fontFamily,omitempty"`
	FontPostScriptName  string  `json:"fontPostScriptName,omitempty"`
	Italic              bool    `json:"italic,omitempty"`
	FontWeight          float64 `json:"fontWeight,omitempty"`
	FontSize            float64 `json:"fontSize,omitempty"`
	TextAlignHorizontal string  `json:"textAlignHorizontal,omitempty"`
	TextAlignVertical   string  `json:"textAlignVertical,omitempty"`
	LetterSpacing       float64 `json:"letterSpacing"`
	LineHeightPercent   float64 `json:"lineHeightPercent,omitempty"`
	LineHeightPx        float64 `json:"lineHeightPx,omitempty"`
	Fills               []Paint `json:"fills,omitempty"`
}

// Node Types

type NodeType string

const (
	NodeTypeDOCUMENT        NodeType = "DOCUMENT"
	NodeTypeCANVAS                   = "CANVAS"
	NodeTypeFRAME                    = "FRAME"
	NodeTypeGROUP                    = "GROUP"
	NodeTypeVECTOR                   = "VECTOR"
	NodeTypeBOOLEAN                  = "BOOLEAN"
	NodeTypeSTAR                     = "STAR"
	NodeTypeLINE                     = "LINE"
	NodeTypeELLIPSE                  = "ELLIPSE"
	NodeTypeREGULAR_POLYGON          = "REGULAR_POLYGON"
	NodeTypeRECTANGLE                = "RECTANGLE"
	NodeTypeTEXT                     = "TEXT"
	NodeTypeSLICE                    = "SLICE"
	NodeTypeCOMPONENT                = "COMPONENT"
	NodeTypeINSTANCE                 = "INSTANCE"
)

type Children []Node

func (c *Children) UnmarshalJSON(data []byte) error {
	var v []json.RawMessage
	type justType struct {
		Type NodeType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	var result Children
	for _, n := range v {
		var jt justType
		err := json.Unmarshal([]byte(n), &jt)
		if err != nil {
			return err
		}

		var v Node
		switch jt.Type {
		case NodeTypeDOCUMENT:
			v = &DocumentNode{}
		case NodeTypeCANVAS:
			v = &CanvasNode{}
		case NodeTypeFRAME:
			v = &FrameNode{}
		case NodeTypeGROUP:
			v = &GroupNode{}
		case NodeTypeVECTOR:
			v = &VectorNode{}
		case NodeTypeBOOLEAN:
			v = &BooleanNode{}
		case NodeTypeSTAR:
			v = &StarNode{}
		case NodeTypeLINE:
			v = &LineNode{}
		case NodeTypeELLIPSE:
			v = &EllipseNode{}
		case NodeTypeREGULAR_POLYGON:
			v = &RegularPolygonNode{}
		case NodeTypeRECTANGLE:
			v = &RectangleNode{}
		case NodeTypeTEXT:
			v = &TextNode{}
		case NodeTypeSLICE:
			v = &SliceNode{}
		case NodeTypeCOMPONENT:
			v = &ComponentNode{}
		case NodeTypeINSTANCE:
			v = &InstanceNode{}
		default:
			v = &UnknownNode{}
		}
		err = json.Unmarshal([]byte(n), &v)
		result = append(result, v)
	}
	*c = result
	return nil
}

type Node interface {
	GetID() string
	GetName() string
	GetType() NodeType
	GetVisible() *bool
}

type Parent interface {
	GetChildren() Children
}

type UnknownNode struct {
	BaseNode
}

// BaseNode contains common fields for every Node type.
type BaseNode struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    NodeType `json:"type,omitempty"`
	Visible *bool    `json:"visible,omitempty"`
}

// BaseParentNode adds Children to BaseNode.
type BaseParentNode struct {
	BaseNode
	// An array of canvases attached to the document
	Children Children `json:"children,omitempty"`
}

func (b *BaseNode) GetID() string {
	return b.ID
}

func (b *BaseNode) GetName() string {
	return b.Name
}

func (b *BaseNode) GetType() NodeType {
	return b.Type
}

func (b *BaseNode) GetVisible() *bool {
	return b.Visible
}

func (b *BaseParentNode) GetChildren() Children {
	return b.Children
}

// Node Types

// Document is the root node.
type DocumentNode struct {
	BaseParentNode
}

// Canvas is represents a single page.
type CanvasNode struct {
	BaseParentNode
	// Background color of the canvas.
	BackgroundColor Color `json:"backgroundColor,omitempty"`
	// An array of export settings representing images to export from the canvas.
	ExportSettings []ExportSetting `json:"exportSettings"`
}

// Frame is a node of fixed size containing other nodes.
type FrameNode struct {
	BaseParentNode
	// Background color of the node.
	BackgroundColor Color `json:"backgroundColor"`
	// An array of export settings representing images to export from node.
	ExportSettings []ExportSetting `json:"exportSettings,omitEmpty"`
	// How this node blends with nodes behind it in the scene (see blend mode section for more details).
	BlendMode BlendMode `json:"blendMode,omitempty"`
	// Keep height and width constrained to same ratio. default: false.
	PreserveRatio bool `json:"preserveRatio,omitempty"`
	// Horizontal and vertical layout constraints for node.
	Constraints LayoutConstraint `json:"constraints,omitempty"`
	// Node ID of node to transition to in prototyping. default: null.
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
	// Opacity of the node. default: 1.
	Opacity float64 `json:"opacity,omitempty"`
	// Bounding box of the node in absolute space coordinates.
	AbsoluteBoundingBox Rectangle `json:"absoluteBoundingBox,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if geometry=paths is passed.
	Size *Vector `json:"size,omitempty"`
	// The top two rows of a matrix that represents the 2D transform of this node relative to its parent. The bottom row of the matrix is implicitly always (0, 0, 1). Use to transform coordinates in geometry. Only present if geometry=paths is passed.
	RelativeTransform Transform `json:"relativeTransform,omitempty"`
	// Does this node clip content outside of its bounds?.
	ClipsContent bool `json:"clipsContent"`
	// An array of layout grids attached to this node (see layout grids section for more details). GROUP nodes do not have this attribute.
	LayoutGrids []LayoutGrid `json:"layoutGrids,omitempty"`
	// An array of effects attached to this node (see effects section for more details).
	Effects []Effect `json:"effects"`
	// Does this node mask sibling nodes in front of it?. default: false.
	IsMask bool `json:"isMask,omitempty"`
}

// Group is a logical grouping of nodes.
type GroupNode FrameNode

type StrokeAlignType string

const (
	StrokeAlignTypeINSIDE  StrokeAlignType = "INSIDE"
	StrokeAlignTypeOUTSIDE                 = "OUTSIDE"
	StrokeAlignTypeCENTER                  = "CENTER"
)

// VectorNode is a vector network, consisting of vertices and edges.
type VectorNode struct {
	BaseNode
	// An array of export settings representing images to export from node
	ExportSettings []ExportSetting `json:"exportSettings"`
	// How this node blends with nodes behind it in the scene (see blend mode section for more details)
	BlendMode BlendMode `json:"blendMode,omitempty"`
	// Keep height and width constrained to same ratio default: false.
	PreserveRatio bool `json:"preserveRatio,omitempty"`
	// Horizontal and vertical layout constraints for node
	Constraints LayoutConstraint `json:"constraints,omitempty"`
	// Node ID of node to transition to in prototyping default: null.
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
	// Opacity of the node default: 1.
	Opacity float64 `json:"opacity,omitempty"`
	// Bounding box of the node in absolute space coordinates
	AbsoluteBoundingBox Rectangle `json:"absoluteBoundingBox,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if geometry=paths is passed
	Size *Vector `json:"size,omitempty"`
	// The top two rows of a matrix that represents the 2D transform of this node relative to its parent. The bottom row of the matrix is implicitly always (0, 0, 1). Use to transform coordinates in geometry. Only present if geometry=paths is passed
	RelativeTransform Transform `json:"relativeTransform,omitempty"`
	// An array of effects attached to this node (see effects section for more details)
	Effects []Effect `json:"effects"`
	// Does this node mask sibling nodes in front of it? default: false.
	IsMask bool `json:"isMask,omitempty"`
	// An array of fill paints applied to the node
	Fills []Paint `json:"fills,omitempty"`
	// Only specified if parameter geometry=paths is used. An array of paths representing the object fill
	FillGeometry []Path `json:"fillGeometry,omitempty"`
	// An array of stroke paints applied to the node
	Strokes []Paint `json:"strokes"`
	// The weight of strokes on the node
	StrokeWeight float64 `json:"strokeWeight"`
	// Only specified if parameter geometry=paths is used. An array of paths representing the object stroke
	StrokeGeometry []Path `json:"strokeGeometry,omitempty"`
	// Where stroke is drawn relative to the vector outline as a string enum
	StrokeAlign StrokeAlignType `json:"strokeAlign,omitempty"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles map[StyleType]string `json:"styles,omitempty"`
}

// Boolean is a group that has a boolean operation applied to it.
type BooleanNode struct {
	VectorNode
	Children Children `json:"children,omitempty"`
}

func (b *BooleanNode) GetChildren() Children {
	return b.Children
}

// Star is a regular star shape.
type StarNode VectorNode

// Line is a straight line.
type LineNode VectorNode

// Ellipse is an ellipse.
type EllipseNode VectorNode

// RegularPolygon is a regular n-sided polygon.
type RegularPolygonNode VectorNode

// Rectangle is a rectangle.
type RectangleNode struct {
	VectorNode
	// Radius of each corner of the rectangle.
	CornerRadius float64 `json:"cornerRadius,omitempty"`
}

// Text is a text box.
type TextNode struct {
	VectorNode
	// Text contained within text box
	Characters string `json:"characters,omitempty"`
	// Style of text including font family and weight (see type style section for more information)
	Style TypeStyle `json:"style,omitempty"`
	// Array with same number of elements as characeters in text box, each element is a reference to the styleOverrideTable defined below and maps to the corresponding character in the characters field. Elements with value 0 have the default type style
	CharacterStyleOverrides []int `json:"characterStyleOverrides"`
	// Map from ID to TypeStyle for looking up style overrides
	StyleOverrideTable map[int]TypeStyle `json:"styleOverrideTable"`
}

// Slice is a rectangular region of the canvas that can be exported.
type SliceNode struct {
	BaseNode
	// An array of export settings representing images to export from this node
	ExportSettings []ExportSetting `json:"exportSettings"`
	// Bounding box of the node in absolute space coordinates
	AbsoluteBoundingBox Rectangle `json:"absoluteBoundingBox,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if geometry=paths is passed
	Size Vector `json:"size,omitempty"`
	// The top two rows of a matrix that represents the 2D transform of this node relative to its parent. The bottom row of the matrix is implicitly always (0, 0, 1). Use to transform coordinates in geometry. Only present if geometry=paths is passed
	RelativeTransform Transform `json:"relativeTransform,omitempty"`
}

// Component is a node that can have instances created of it that share the same properties.
type ComponentNode FrameNode

// Instance is an instance of a component, changes to the component result in the same changes applied to the instance.
type InstanceNode struct {
	FrameNode
	// ID of component that this instance came from, refers to components table.
	ComponentID string `json:"componentId,omitempty"`
}
