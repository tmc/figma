package figmatypes

// VectorOrFrameOffset contains the fields from Vector and FrameOffset.
type VectorOrFrameOffset struct {
	X          float64 `json:"x,omitempty"`
	Y          float64 `json:"y,omitempty"`
	NodeID     string  `json:"node_id,omitempty"`
	NodeOffset *Vector `json:"node_offset,omitempty"`
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
