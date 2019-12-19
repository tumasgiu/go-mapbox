package styles

import "github.com/tumasgiu/go-mapbox/lib/base"

type Anchor string

const (
	AnchorMap      Anchor = "map"
	AnchorViewport Anchor = "anchor"
)

// https://docs.mapbox.com/mapbox-gl-js/style-spec/#light
type Light struct {
	Anchor    Anchor    `json:"anchor,omitempty"`
	Color     string    `json:"color,omitempty"`
	Intensity float64   `json:"intensity,omitempty"`
	Position  []float64 `json:"position,omitempty"`
}

// https://docs.mapbox.com/mapbox-gl-js/style-spec/#transition
type Transition struct {
	Duration int `json:"duration,omitempty"`
	Delay    int `json:"delay,omitempty"`
}

type LayerType string

const (
	LayerTypeFill          LayerType = "fill"
	LayerTypeLine          LayerType = "line"
	LayerTypeSymbol        LayerType = "symbol"
	LayerTypeCircle        LayerType = "circle"
	LayerTypeHeatmap       LayerType = "heatmap"
	LayerTypeFillExtrusion LayerType = "fill-extrusion"
	LayerTypeRaster        LayerType = "raster"
	LayerTypeHillshade     LayerType = "hillshade"
	LayerTypeBackground    LayerType = "background"
)

// TODO: https://docs.mapbox.com/mapbox-gl-js/style-spec/#expressions
type Expresion interface{}

type Visibility string

const (
	VisibilityVisible Visibility = "visible"
	VisibilityNone    Visibility = "none"
)

type LineCap string

const (
	LineCapButt   LineCap = "butt"
	LineCapRound  LineCap = "round"
	LineCapSquare LineCap = "square"
)

type LineJoin string

const (
	LineJoinBevel LineJoin = "bevel"
	LineJoinRound LineJoin = "round"
	LineJoinMiter LineJoin = "miter"
)

type SymbolPlacement string

const (
	SymbolPlacementPoint      SymbolPlacement = "point"
	SymbolPlacementLine       SymbolPlacement = "line"
	SymbolPlacementLienCenter SymbolPlacement = "line-center"
)

type SymbolZOrder string

const (
	SymbolZOrderAuto      SymbolZOrder = "auto"
	SymbolZOrderViewportY SymbolZOrder = "viewport-y"
	SymbolZOrderSource    SymbolZOrder = "source"
)

type Alignment string

const (
	AlignmentMap      Alignment = "map"
	AlignmentViewport Alignment = "viewport"
	AlignmentAuto     Alignment = "auto"
)

type IconTextFit string

const (
	IconTextFitNone   IconTextFit = "none"
	IconTextFitWidth  IconTextFit = "width"
	IconTextFitHeight IconTextFit = "height"
	IconTextFitBoth   IconTextFit = "both"
)

type VariableAnchor string

const (
	VariableAnchorCenter      VariableAnchor = "center"
	VariableAnchorLeft        VariableAnchor = "left"
	VariableAnchorRight       VariableAnchor = "right"
	VariableAnchorTop         VariableAnchor = "top"
	VariableAnchorBottom      VariableAnchor = "bottom"
	VariableAnchorTopLeft     VariableAnchor = "top-left"
	VariableAnchorTopRight    VariableAnchor = "top-right"
	VariableAnchorBottomLeft  VariableAnchor = "bottom-left"
	VariableAnchorBottomRight VariableAnchor = "bottom-right"
)

type Justify string

const (
	JustifyAuto   Justify = "auto"
	JustifyLeft   Justify = "left"
	JustifyCenter Justify = "center"
	JustifyRight  Justify = "right"
)

type WritingMode string

const (
	WritingModeHorizontal WritingMode = "horizontal"
	WritingModeVertical   WritingMode = "vertical"
)

type TextTransform string

const (
	TextTransformNone      TextTransform = "none"
	TextTransformUppercase TextTransform = "uppercase"
	TextTransformLowercase TextTransform = "lowercase"
)

// https://docs.mapbox.com/mapbox-gl-js/style-spec/#layers-background
type Layout struct {
	Visibility            Visibility      `json:"visibility,omitempty"`
	FillSortKey           float64         `json:"fill-sort-key,omitempty"`
	LineCap               LineCap         `json:"line-cap,omitempty"`
	LineJoin              LineJoin        `json:"line-join,omitempty"`
	LineMiterLimit        float64         `json:"line-miter-limit,omitempty"`
	LineRoundLimit        float64         `json:"line-round-limit,omitempty"`
	LineSortKey           int             `json:"line-sort-key,omitempty"`
	SymbolPlacement       SymbolPlacement `json:"symbol-placement,omitempty"`
	SymbolSpacing         int             `json:"symbol-spacing,omitempty"`
	SymbolAvoidEdge       bool            `json:"symbol-avoid-edge,omitempty"`
	SymbolSortKey         int             `json:"symbol-sort-key,omitempty"`
	SymbolZOrder          SymbolZOrder    `json:"symbol-z-order,omitempty"`
	IconAllowOverlap      bool            `json:"icon-allow-overlap,omitempty"`
	IconIgnorePlacement   bool            `json:"icon-ignore-placement,omitempty"`
	IconOptional          bool            `json:"icon-optional,omitempty"`
	IconRotationAlignment Alignment       `json:"icon-rotation-alignment,omitempty"`
	IconSize              int             `json:"icon-size,omitempty"`
	IconTextFit           IconTextFit     `json:"icon-text-fit,omitempty"`
	IconImage             interface{}     `json:"icon-image,omitempty"`
	IconRotate            int             `json:"icon-rotate,omitempty"`
	IconPadding           int             `json:"icon-padding,omitempty"`
	IconKeepUpright       bool            `json:"icon-keep-upright,omitempty"`
	IconOffset            base.Point      `json:"icon-offset,omitempty"`
	IconAnchor            VariableAnchor  `json:"icon-anchor,omitempty"`
	IconPitchAlignment    Alignment       `json:"icon-pitch-alignment,omitempty"`
	TextPitchAlignment    Alignment       `json:"text-pitch-alignment,omitempty"`
	TextRotationAlignment Alignment       `json:"text-rotation-alignment,omitempty"`
	TextField             string          `json:"text-field,omitempty"`
	TextFont              []string        `json:"text-font,omitempty"`
	TextSize              int             `json:"text-size,omitempty"`
	TextMaxWidth          int             `json:"text-max-width,omitempty"`
	TextLineHeight        float64         `json:"text-line-height,omitempty"`
	TextLetterSpacing     float64         `json:"text-letter-spacing,omitempty"`
	TextJustify           Justify         `json:"text-justify,omitempty"`
	TextRadialOffset      float64         `json:"text-radial-offset,omitempty"`
	TextVariableAnchor    VariableAnchor  `json:"text-variable-anchor,omitempty"`
	TextAnchor            VariableAnchor  `json:"text-anchor,omitempty"`
	TextMaxAngle          int             `json:"text-max-angle,omitempty"`
	TextWritingMode       WritingMode     `json:"text-writing-mode,omitempty"`
	TextRotate            int             `json:"text-rotate,omitempty"`
	TextPadding           int             `json:"text-padding,omitempty"`
	TextKeepUpright       bool            `json:"text-keep-upright,omitempty"`
	TextTransform         TextTransform   `json:"text-transform,omitempty"`
	TextOffset            base.Point      `json:"text-offset,omitempty"`
	TextAllowOverlap      bool            `json:"text-allow-overlap,omitempty"`
	TextIgnorePlacement   bool            `json:"text-ignore-placement,omitempty"`
	TextOptional          bool            `json:"text-optional,omitempty"`
	CircleSortKey         int             `json:"circle-sort-key,omitempty"`
}

type Resampling string

const (
	ResamplingLinear  Resampling = "linear"
	ResamplingNearest Resampling = "nearest"
)

// https://docs.mapbox.com/mapbox-gl-js/style-spec/#layers-background
type Paint struct {
	BackgroundColor                string      `json:"background-color,omitempty"`
	BackgroundPattern              interface{} `json:"background-pattern,omitempty"`
	BackgroundOpacity              float64     `json:"background-opacity,omitempty"`
	FillAntialias                  bool        `json:"fill-antialias,omitempty"`
	FillOpacity                    float64     `json:"fill-opacity,omitempty"`
	FillColor                      string      `json:"fill-color,omitempty"`
	FillOutlineColor               string      `json:"fill-outline-color,omitempty"`
	FillTranslate                  base.Point  `json:"fill-translate,omitempty"`
	FillTranslateAnchor            Anchor      `json:"fill-translate-anchor,omitempty"`
	FillPattern                    interface{} `json:"fill-pattern,omitempty"`
	LineOpacity                    float64     `json:"line-opacity,omitempty"`
	LineColor                      string      `json:"line-color,omitempty"`
	LineTranslate                  base.Point  `json:"line-translate,omitempty"`
	LineTranslateAnchor            Anchor      `json:"line-translate-anchor,omitempty"`
	LineWidth                      int         `json:"line-width,omitempty"`
	LineGapWidth                   int         `json:"line-gap-width,omitempty"`
	LineOffset                     int         `json:"line-offset,omitempty"`
	LineBlur                       float64     `json:"line-blur,omitempty"`
	LineDashArray                  []int       `json:"line-dash-array,omitempty"`
	LinePattern                    interface{} `json:"line-pattern,omitempty"`
	LineGradient                   string      `json:"line-gradient,omitempty"`
	IconOpacity                    float64     `json:"icon-opacity,omitempty"`
	IconColor                      string      `json:"icon-color,omitempty"`
	IconHaloColor                  string      `json:"icon-halo-color,omitempty"`
	IconHaloWidth                  int         `json:"icon-halo-width,omitempty"`
	IconHaloBlur                   float64     `json:"icon-halo-blur,omitempty"`
	IconTranslatable               base.Point  `json:"icon-translatable,omitempty"`
	IconTranslateAnchor            Anchor      `json:"icon-translate-anchor,omitempty"`
	TextOpacity                    float64     `json:"text-opacity,omitempty"`
	TextColor                      string      `json:"text-color,omitempty"`
	TextHaloColor                  string      `json:"text-halo-color,omitempty"`
	TextHaloWidth                  int         `json:"text-halo-width,omitempty"`
	TextHaloBlur                   float64     `json:"text-halo-blur,omitempty"`
	TextTranslate                  base.Point  `json:"text-translate,omitempty"`
	TextTranslateAnchor            Anchor      `json:"text-translate-anchor,omitempty"`
	RasterOpacity                  float64     `json:"raster-opacity,omitempty"`
	RasterHueRotate                int         `json:"raster-hue-rotate,omitempty"`
	RasterBrightnessMin            float64     `json:"raster-brightness-min,omitempty"`
	RasterBrightnessMax            float64     `json:"raster-brightness-max,omitempty"`
	RasterSaturation               float64     `json:"raster-saturation,omitempty"`
	RasterContrast                 float64     `json:"raster-contrast,omitempty"`
	RasterResampling               Resampling  `json:"raster-resampling,omitempty"`
	RasterFadeDuration             int         `json:"raster-fade-duration,omitempty"`
	CircleRadius                   int         `json:"circle-radius,omitempty"`
	CircleColor                    string      `json:"circle-color,omitempty"`
	CircleBlur                     float64     `json:"circle-blur,omitempty"`
	CircleOpacity                  float64     `json:"circle-opacity,omitempty"`
	CircleTranslate                base.Point  `json:"circle-translate,omitempty"`
	CircleTranslateAnchor          Anchor      `json:"circle-translate-anchor,omitempty"`
	CirclePitchScale               Anchor      `json:"circle-pitch-scale,omitempty"`
	CirclePitchAlignment           Anchor      `json:"circle-pitch-alignment,omitempty"`
	CircleStrokeWidth              int         `json:"circle-stroke-width,omitempty"`
	CircleStrokeColor              string      `json:"circle-stroke-color,omitempty"`
	CircleStrokeOpacity            float64     `json:"circle-stroke-opacity,omitempty"`
	FillExtrusionOpacity           float64     `json:"fill-extrusion-opacity,omitempty"`
	FillExtrusionColor             string      `json:"fill-extrusion-color,omitempty"`
	FillExtrusionTranslate         base.Point  `json:"fill-extrusion-translate,omitempty"`
	FillExtrusionTranslateAnchor   Anchor      `json:"fill-extrusion-translate-anchor,omitempty"`
	FillExtrusionPattern           interface{} `json:"fill-extrusion-pattern,omitempty"`
	FillExtrusionHeight            int         `json:"fill-extrusion-height,omitempty"`
	FillExtrusionBase              int         `json:"fill-extrusion-base,omitempty"`
	FillExtrusionVerticalGradient  bool        `json:"fill-extrusion-vertical-gradient,omitempty"`
	HeatmapRadius                  int         `json:"heatmap-radius,omitempty"`
	HeatmapWeight                  float64     `json:"heatmap-weight,omitempty"`
	HeatmapIntensity               float64     `json:"heatmap-intensity,omitempty"`
	HeatmapColor                   Expresion   `json:"heatmap-color,omitempty"`
	HeatmapOpacity                 float64     `json:"heatmap-opacity,omitempty"`
	HillshadeIlluminationDirection int         `json:"hillshade-illumination-direction,omitempty"`
	HillshadeIlluminationAnchor    Anchor      `json:"hillshade-illumination-anchor,omitempty"`
	HillshadeExaggeration          float64     `json:"hillshade-exaggeration,omitempty"`
	HillshadeShadowColor           string      `json:"hillshade-shadow-color,omitempty"`
	HillshadeHighlightColor        string      `json:"hillshade-highlight-color,omitempty"`
	HillshadeAccentColor           string      `json:"hillshade-accent-color,omitempty"`
}

// https://docs.mapbox.com/mapbox-gl-js/style-spec/#layers
type Layer struct {
	Id          string      `json:"id"`
	Type        LayerType   `json:"type"`
	Metadata    interface{} `json:"metadata,omitempty"`
	Source      string      `json:"source,omitempty"`
	SourceLayer string      `json:"source-layer,omitempty"`
	MinZoom     int         `json:"minzoom,omitempty"`
	MaxZoom     int         `json:"maxzoom,omitempty"`
	Filter      Expresion   `json:"filter,omitempty"`
	Layout      interface{} `json:"layout,omitempty"`
	Paint       interface{} `json:"paint,omitempty"`
}

// The Style Object
//https://docs.mapbox.com/mapbox-gl-js/style-spec/
type Style struct {
	Version    int         `json:"version"`
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Metadata   interface{} `json:"metadata,omitempty"`
	Center     base.Point  `json:"center,omitempty"`
	Zoom       float64     `json:"zoom,omitempty"`
	Bearing    float64     `json:"bearing,omitempty"`
	Pitch      float64     `json:"pitch,omitempty"`
	Light      Light       `json:"light,omitempty"`
	Sources    interface{} `json:"sources"`
	Sprite     string      `json:"sprite,omitempty"`
	Glyphs     string      `json:"glyphs,omitempty"`
	Transition Transition  `json:"transition,omitempty"`
	Layers     []Layer     `json:"layers"`
}
