package nodes

import (
	"encoding/json"

	"github.com/tmc/figma/figmatypes"
)

// NodeType describes the type of a node.
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
			v = &Document{}
		case NodeTypeCANVAS:
			v = &Canvas{}
		case NodeTypeFRAME:
			v = &Frame{}
		case NodeTypeGROUP:
			v = &Group{}
		case NodeTypeVECTOR:
			v = &Vector{}
		case NodeTypeBOOLEAN:
			v = &Boolean{}
		case NodeTypeSTAR:
			v = &Star{}
		case NodeTypeLINE:
			v = &Line{}
		case NodeTypeELLIPSE:
			v = &Ellipse{}
		case NodeTypeREGULAR_POLYGON:
			v = &RegularPolygon{}
		case NodeTypeRECTANGLE:
			v = &Rectangle{}
		case NodeTypeTEXT:
			v = &Text{}
		case NodeTypeSLICE:
			v = &Slice{}
		case NodeTypeCOMPONENT:
			v = &Component{}
		case NodeTypeINSTANCE:
			v = &Instance{}
		default:
			v = &Unknown{}
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

type Unknown struct {
	NodeBase
}

// NodeBase contains common fields for every  type.
type NodeBase struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    NodeType `json:"type,omitempty"`
	Visible *bool    `json:"visible,omitempty"`
}

// ParentNodeBase adds Children to NodeBase.
type ParentNodeBase struct {
	NodeBase
	// An array of canvases attached to the document
	Children Children `json:"children,omitempty"`
}

func (b *NodeBase) GetID() string {
	return b.ID
}

func (b *NodeBase) GetName() string {
	return b.Name
}

func (b *NodeBase) GetType() NodeType {
	return b.Type
}

func (b *NodeBase) GetVisible() *bool {
	return b.Visible
}

func (b *ParentNodeBase) GetChildren() Children {
	return b.Children
}

//  Types

// Document is the root node.
type Document struct {
	ParentNodeBase
}

// Canvas is represents a single page.
type Canvas struct {
	ParentNodeBase
	// Background color of the canvas.
	BackgroundColor figmatypes.Color `json:"backgroundColor,omitempty"`
	// An array of export settings representing images to export from the canvas.
	ExportSettings []figmatypes.ExportSetting `json:"exportSettings"`
}

// Frame is a node of fixed size containing other nodes.
type Frame struct {
	ParentNodeBase
	// Background color of the node.
	BackgroundColor figmatypes.Color `json:"backgroundColor"`
	// An array of export settings representing images to export from node.
	ExportSettings []figmatypes.ExportSetting `json:"exportSettings,omitEmpty"`
	// How this node blends with nodes behind it in the scene (see blend mode section for more details).
	BlendMode figmatypes.BlendMode `json:"blendMode,omitempty"`
	// Keep height and width constrained to same ratio. default: false.
	PreserveRatio bool `json:"preserveRatio,omitempty"`
	// Horizontal and vertical layout constraints for node.
	Constraints figmatypes.LayoutConstraint `json:"constraints,omitempty"`
	//  ID of node to transition to in prototyping. default: null.
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
	// Opacity of the node. default: 1.
	Opacity float64 `json:"opacity,omitempty"`
	// Bounding box of the node in absolute space coordinates.
	AbsoluteBoundingBox figmatypes.Rectangle `json:"absoluteBoundingBox,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if geometry=paths is passed.
	Size *Vector `json:"size,omitempty"`
	// The top two rows of a matrix that represents the 2D transform of this node relative to its parent. The bottom row of the matrix is implicitly always (0, 0, 1). Use to transform coordinates in geometry. Only present if geometry=paths is passed.
	RelativeTransform figmatypes.Transform `json:"relativeTransform,omitempty"`
	// Does this node clip content outside of its bounds?.
	ClipsContent bool `json:"clipsContent"`
	// An array of layout grids attached to this node (see layout grids section for more details). GROUP nodes do not have this attribute.
	LayoutGrids []figmatypes.LayoutGrid `json:"layoutGrids,omitempty"`
	// An array of effects attached to this node (see effects section for more details).
	Effects []figmatypes.Effect `json:"effects"`
	// Does this node mask sibling nodes in front of it?. default: false.
	IsMask bool `json:"isMask,omitempty"`
}

// Group is a logical grouping of nodes.
type Group Frame

type StrokeAlignType string

const (
	StrokeAlignTypeINSIDE  StrokeAlignType = "INSIDE"
	StrokeAlignTypeOUTSIDE                 = "OUTSIDE"
	StrokeAlignTypeCENTER                  = "CENTER"
)

// Vector is a vector network, consisting of vertices and edges.
type Vector struct {
	NodeBase
	// An array of export settings representing images to export from node
	ExportSettings []figmatypes.ExportSetting `json:"exportSettings"`
	// How this node blends with nodes behind it in the scene (see blend mode section for more details)
	BlendMode figmatypes.BlendMode `json:"blendMode,omitempty"`
	// Keep height and width constrained to same ratio default: false.
	PreserveRatio bool `json:"preserveRatio,omitempty"`
	// Horizontal and vertical layout constraints for node
	Constraints figmatypes.LayoutConstraint `json:"constraints,omitempty"`
	//  ID of node to transition to in prototyping default: null.
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
	// Opacity of the node default: 1.
	Opacity float64 `json:"opacity,omitempty"`
	// Bounding box of the node in absolute space coordinates
	AbsoluteBoundingBox figmatypes.Rectangle `json:"absoluteBoundingBox,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if geometry=paths is passed
	Size *Vector `json:"size,omitempty"`
	// The top two rows of a matrix that represents the 2D transform of this node relative to its parent. The bottom row of the matrix is implicitly always (0, 0, 1). Use to transform coordinates in geometry. Only present if geometry=paths is passed
	RelativeTransform figmatypes.Transform `json:"relativeTransform,omitempty"`
	// An array of effects attached to this node (see effects section for more details)
	Effects []figmatypes.Effect `json:"effects"`
	// Does this node mask sibling nodes in front of it? default: false.
	IsMask bool `json:"isMask,omitempty"`
	// An array of fill paints applied to the node
	Fills []figmatypes.Paint `json:"fills,omitempty"`
	// Only specified if parameter geometry=paths is used. An array of paths representing the object fill
	FillGeometry []figmatypes.Path `json:"fillGeometry,omitempty"`
	// An array of stroke paints applied to the node
	Strokes []figmatypes.Paint `json:"strokes"`
	// The weight of strokes on the node
	StrokeWeight float64 `json:"strokeWeight"`
	// Only specified if parameter geometry=paths is used. An array of paths representing the object stroke
	StrokeGeometry []figmatypes.Path `json:"strokeGeometry,omitempty"`
	// Where stroke is drawn relative to the vector outline as a string enum
	StrokeAlign StrokeAlignType `json:"strokeAlign,omitempty"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles map[figmatypes.StyleType]string `json:"styles,omitempty"`
}

// Boolean is a group that has a boolean operation applied to it.
type Boolean struct {
	Vector
	Children Children `json:"children,omitempty"`
}

func (b *Boolean) GetChildren() Children {
	return b.Children
}

// Star is a regular star shape.
type Star Vector

// Line is a straight line.
type Line Vector

// Ellipse is an ellipse.
type Ellipse Vector

// RegularPolygon is a regular n-sided polygon.
type RegularPolygon Vector

// Rectangle is a rectangle.
type Rectangle struct {
	Vector
	// Radius of each corner of the rectangle.
	CornerRadius float64 `json:"cornerRadius,omitempty"`
}

// Text is a text box.
type Text struct {
	Vector
	// Text contained within text box
	Characters string `json:"characters,omitempty"`
	// Style of text including font family and weight (see type style section for more information)
	Style figmatypes.TypeStyle `json:"style,omitempty"`
	// Array with same number of elements as characeters in text box, each element is a reference to the styleOverrideTable defined below and maps to the corresponding character in the characters field. Elements with value 0 have the default type style
	CharacterStyleOverrides []int `json:"characterStyleOverrides"`
	// Map from ID to TypeStyle for looking up style overrides
	StyleOverrideTable map[int]figmatypes.TypeStyle `json:"styleOverrideTable"`
}

// Slice is a rectangular region of the canvas that can be exported.
type Slice struct {
	NodeBase
	// An array of export settings representing images to export from this node
	ExportSettings []figmatypes.ExportSetting `json:"exportSettings"`
	// Bounding box of the node in absolute space coordinates
	AbsoluteBoundingBox figmatypes.Rectangle `json:"absoluteBoundingBox,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if geometry=paths is passed
	Size Vector `json:"size,omitempty"`
	// The top two rows of a matrix that represents the 2D transform of this node relative to its parent. The bottom row of the matrix is implicitly always (0, 0, 1). Use to transform coordinates in geometry. Only present if geometry=paths is passed
	RelativeTransform figmatypes.Transform `json:"relativeTransform,omitempty"`
}

// Component is a node that can have instances created of it that share the same properties.
type Component Frame

// Instance is an instance of a component, changes to the component result in the same changes applied to the instance.
type Instance struct {
	Frame
	// ID of component that this instance came from, refers to components table.
	ComponentID string `json:"componentId,omitempty"`
}
