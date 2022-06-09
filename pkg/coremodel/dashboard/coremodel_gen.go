// This file is autogenerated. DO NOT EDIT.
//
// To regenerate, run "make gen-cue" from repository root.
//
// Derived from the Thema lineage at pkg/coremodel/dashboard

package dashboard

import (
	"embed"
	"path/filepath"

	"github.com/grafana/grafana/pkg/cuectx"
	"github.com/grafana/thema"
)

// Defines values for DashboardGraphTooltip.
const (
	ModelGraphTooltipN0 ModelGraphTooltip = 0

	ModelGraphTooltipN1 ModelGraphTooltip = 1

	ModelGraphTooltipN2 ModelGraphTooltip = 2
)

// Defines values for DashboardStyle.
const (
	ModelStyleDark ModelStyle = "dark"

	ModelStyleLight ModelStyle = "light"
)

// Defines values for DashboardTimezone.
const (
	ModelTimezoneBrowser ModelTimezone = "browser"

	ModelTimezoneEmpty ModelTimezone = ""

	ModelTimezoneUtc ModelTimezone = "utc"
)

// Defines values for DashboardDashboardCursorSync.
const (
	ModelDashboardCursorSyncN0 ModelDashboardCursorSync = 0

	ModelDashboardCursorSyncN1 ModelDashboardCursorSync = 1

	ModelDashboardCursorSyncN2 ModelDashboardCursorSync = 2
)

// Defines values for DashboardDashboardLinkType.
const (
	ModelDashboardLinkTypeDashboards ModelDashboardLinkType = "dashboards"

	ModelDashboardLinkTypeLink ModelDashboardLinkType = "link"
)

// Defines values for DashboardFieldColorModeId.
const (
	ModelFieldColorModeIdContinuousGrYlRd ModelFieldColorModeId = "continuous-GrYlRd"

	ModelFieldColorModeIdFixed ModelFieldColorModeId = "fixed"

	ModelFieldColorModeIdPaletteClassic ModelFieldColorModeId = "palette-classic"

	ModelFieldColorModeIdPaletteSaturated ModelFieldColorModeId = "palette-saturated"

	ModelFieldColorModeIdThresholds ModelFieldColorModeId = "thresholds"
)

// Defines values for DashboardFieldColorSeriesByMode.
const (
	ModelFieldColorSeriesByModeLast ModelFieldColorSeriesByMode = "last"

	ModelFieldColorSeriesByModeMax ModelFieldColorSeriesByMode = "max"

	ModelFieldColorSeriesByModeMin ModelFieldColorSeriesByMode = "min"
)

// Defines values for DashboardGraphPanelType.
const (
	ModelGraphPanelTypeGraph ModelGraphPanelType = "graph"
)

// Defines values for DashboardHeatmapPanelType.
const (
	ModelHeatmapPanelTypeHeatmap ModelHeatmapPanelType = "heatmap"
)

// Defines values for DashboardPanelRepeatDirection.
const (
	ModelPanelRepeatDirectionH ModelPanelRepeatDirection = "h"

	ModelPanelRepeatDirectionV ModelPanelRepeatDirection = "v"
)

// Defines values for DashboardRowPanelType.
const (
	ModelRowPanelTypeRow ModelRowPanelType = "row"
)

// Defines values for DashboardThresholdsConfigMode.
const (
	ModelThresholdsConfigModeAbsolute ModelThresholdsConfigMode = "absolute"

	ModelThresholdsConfigModePercentage ModelThresholdsConfigMode = "percentage"
)

// Defines values for DashboardThresholdsMode.
const (
	ModelThresholdsModeAbsolute ModelThresholdsMode = "absolute"

	ModelThresholdsModePercentage ModelThresholdsMode = "percentage"
)

// Defines values for DashboardVariableModelType.
const (
	ModelVariableModelTypeAdhoc ModelVariableModelType = "adhoc"

	ModelVariableModelTypeConstant ModelVariableModelType = "constant"

	ModelVariableModelTypeCustom ModelVariableModelType = "custom"

	ModelVariableModelTypeDatasource ModelVariableModelType = "datasource"

	ModelVariableModelTypeInterval ModelVariableModelType = "interval"

	ModelVariableModelTypeQuery ModelVariableModelType = "query"

	ModelVariableModelTypeSystem ModelVariableModelType = "system"

	ModelVariableModelTypeTextbox ModelVariableModelType = "textbox"
)

// Defines values for DashboardVariableType.
const (
	ModelVariableTypeAdhoc ModelVariableType = "adhoc"

	ModelVariableTypeConstant ModelVariableType = "constant"

	ModelVariableTypeCustom ModelVariableType = "custom"

	ModelVariableTypeDatasource ModelVariableType = "datasource"

	ModelVariableTypeInterval ModelVariableType = "interval"

	ModelVariableTypeQuery ModelVariableType = "query"

	ModelVariableTypeSystem ModelVariableType = "system"

	ModelVariableTypeTextbox ModelVariableType = "textbox"
)

// Dashboard defines model for dashboard.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type Model struct {
	Annotations *struct {
		List []ModelAnnotationQuery `json:"list"`
	} `json:"annotations,omitempty"`

	// Description of dashboard.
	Description *string `json:"description,omitempty"`

	// Whether a dashboard is editable or not.
	Editable bool `json:"editable"`

	// TODO docs
	FiscalYearStartMonth *int              `json:"fiscalYearStartMonth,omitempty"`
	GnetId               *string           `json:"gnetId,omitempty"`
	GraphTooltip         ModelGraphTooltip `json:"graphTooltip"`

	// Unique numeric identifier for the dashboard.
	// TODO must isolate or remove identifiers local to a Grafana instance...?
	Id *int64 `json:"id,omitempty"`

	// TODO docs
	Links *[]ModelDashboardLink `json:"links,omitempty"`

	// TODO docs
	LiveNow *bool          `json:"liveNow,omitempty"`
	Panels  *[]interface{} `json:"panels,omitempty"`

	// TODO docs
	Refresh *interface{} `json:"refresh,omitempty"`

	// Version of the JSON schema, incremented each time a Grafana update brings
	// changes to said schema.
	// TODO this is the existing schema numbering system. It will be replaced by Thema's themaVersion
	SchemaVersion int `json:"schemaVersion"`

	// Theme of dashboard.
	Style ModelStyle `json:"style"`

	// Tags associated with dashboard.
	Tags       *[]string `json:"tags,omitempty"`
	Templating *struct {
		List []ModelVariableModel `json:"list"`
	} `json:"templating,omitempty"`

	// Time range for dashboard, e.g. last 6 hours, last 7 days, etc
	Time *struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"time,omitempty"`

	// TODO docs
	// TODO this appears to be spread all over in the frontend. Concepts will likely need tidying in tandem with schema changes
	Timepicker *struct {
		// Whether timepicker is collapsed or not.
		Collapse bool `json:"collapse"`

		// Whether timepicker is enabled or not.
		Enable bool `json:"enable"`

		// Whether timepicker is visible or not.
		Hidden bool `json:"hidden"`

		// Selectable intervals for auto-refresh.
		RefreshIntervals []string `json:"refresh_intervals"`
	} `json:"timepicker,omitempty"`

	// Timezone of dashboard,
	Timezone *ModelTimezone `json:"timezone,omitempty"`

	// Title of dashboard.
	Title *string `json:"title,omitempty"`

	// Unique dashboard identifier that can be generated by anyone. string (8-40)
	Uid *string `json:"uid,omitempty"`

	// Version of the dashboard, incremented each time the dashboard is updated.
	Version *int `json:"version,omitempty"`

	// TODO docs
	WeekStart *string `json:"weekStart,omitempty"`
}

// DashboardGraphTooltip defines model for Dashboard.GraphTooltip.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelGraphTooltip int

// Theme of dashboard.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelStyle string

// Timezone of dashboard,
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelTimezone string

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelAnnotationQuery struct {
	BuiltIn int `json:"builtIn"`

	// Datasource to use for annotation.
	Datasource struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource"`

	// Whether annotation is enabled.
	Enable bool `json:"enable"`

	// Whether to hide annotation.
	Hide *bool `json:"hide,omitempty"`

	// Annotation icon color.
	IconColor *string `json:"iconColor,omitempty"`

	// Name of annotation.
	Name *string `json:"name,omitempty"`

	// Query for annotation data.
	RawQuery *string `json:"rawQuery,omitempty"`
	ShowIn   int     `json:"showIn"`

	// Schema for panel targets is specified by datasource
	// plugins. We use a placeholder definition, which the Go
	// schema loader either left open/as-is with the Base
	// variant of the Dashboard and Panel families, or filled
	// with types derived from plugins in the Instance variant.
	// When working directly from CUE, importers can extend this
	// type directly to achieve the same effect.
	Target *ModelTarget `json:"target,omitempty"`
	Type   string       `json:"type"`
}

// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelDashboardCursorSync int

// FROM public/app/features/dashboard/state/DashboardModels.ts - ish
// TODO docs
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelDashboardLink struct {
	AsDropdown  bool                   `json:"asDropdown"`
	Icon        *string                `json:"icon,omitempty"`
	IncludeVars bool                   `json:"includeVars"`
	KeepTime    bool                   `json:"keepTime"`
	Tags        []string               `json:"tags"`
	TargetBlank bool                   `json:"targetBlank"`
	Title       string                 `json:"title"`
	Tooltip     *string                `json:"tooltip,omitempty"`
	Type        ModelDashboardLinkType `json:"type"`
	Url         *string                `json:"url,omitempty"`
}

// DashboardDashboardLinkType defines model for DashboardDashboardLink.Type.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelDashboardLinkType string

// TODO docs
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelFieldColor struct {
	// Stores the fixed color value if mode is fixed
	FixedColor *string `json:"fixedColor,omitempty"`

	// The main color scheme mode
	Mode interface{} `json:"mode"`

	// TODO docs
	SeriesBy *ModelFieldColorSeriesByMode `json:"seriesBy,omitempty"`
}

// TODO docs
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelFieldColorModeId string

// TODO docs
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelFieldColorSeriesByMode string

// DashboardGraphPanel defines model for dashboard.GraphPanel.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelGraphPanel struct {
	// Support for legacy graph and heatmap panels.
	Type ModelGraphPanelType `json:"type"`
}

// Support for legacy graph and heatmap panels.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelGraphPanelType string

// DashboardGridPos defines model for dashboard.GridPos.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelGridPos struct {
	// Panel
	H int `json:"h"`

	// true if fixed
	Static *bool `json:"static,omitempty"`

	// Panel
	W int `json:"w"`

	// Panel x
	X int `json:"x"`

	// Panel y
	Y int `json:"y"`
}

// DashboardHeatmapPanel defines model for dashboard.HeatmapPanel.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelHeatmapPanel struct {
	Type ModelHeatmapPanelType `json:"type"`
}

// DashboardHeatmapPanelType defines model for DashboardHeatmapPanel.Type.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelHeatmapPanelType string

// Dashboard panels. Panels are canonically defined inline
// because they share a version timeline with the dashboard
// schema; they do not evolve independently.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelPanel struct {
	// The datasource used in all targets.
	Datasource *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`
	FieldConfig struct {
		Defaults struct {
			// TODO docs
			Color *ModelFieldColor `json:"color,omitempty"`

			// custom is specified by the PanelFieldConfig field
			// in panel plugin schemas.
			Custom *map[string]interface{} `json:"custom,omitempty"`

			// Significant digits (for display)
			Decimals *float32 `json:"decimals,omitempty"`

			// Human readable field metadata
			Description *string `json:"description,omitempty"`

			// The display value for this field.  This supports template variables blank is auto
			DisplayName *string `json:"displayName,omitempty"`

			// This can be used by data sources that return and explicit naming structure for values and labels
			// When this property is configured, this value is used rather than the default naming strategy.
			DisplayNameFromDS *string `json:"displayNameFromDS,omitempty"`

			// True if data source field supports ad-hoc filters
			Filterable *bool `json:"filterable,omitempty"`

			// // The behavior when clicking on a result
			Links *[]interface{} `json:"links,omitempty"`

			// Convert input values into a display string
			//
			// TODO this one corresponds to a complex type with
			// generics on the typescript side. Ouch. Will
			// either need special care, or we'll just need to
			// accept a very loosely specified schema. It's very
			// unlikely we'll be able to translate cue to
			// typescript generics in the general case, though
			// this particular one *may* be able to work.
			Mappings *[]map[string]interface{} `json:"mappings,omitempty"`
			Max      *float32                  `json:"max,omitempty"`
			Min      *float32                  `json:"min,omitempty"`

			// Alternative to empty string
			NoValue *string `json:"noValue,omitempty"`

			// An explict path to the field in the datasource.  When the frame meta includes a path,
			// This will default to `${frame.meta.path}/${field.name}
			//
			// When defined, this value can be used as an identifier within the datasource scope, and
			// may be used to update the results
			Path       *string                `json:"path,omitempty"`
			Thresholds *ModelThresholdsConfig `json:"thresholds,omitempty"`

			// Numeric Options
			Unit *string `json:"unit,omitempty"`

			// True if data source can write a value to the path.  Auth/authz are supported separately
			Writeable *bool `json:"writeable,omitempty"`
		} `json:"defaults"`
		Overrides []struct {
			Matcher struct {
				Id      string       `json:"id"`
				Options *interface{} `json:"options,omitempty"`
			} `json:"matcher"`
			Properties []struct {
				Id    string       `json:"id"`
				Value *interface{} `json:"value,omitempty"`
			} `json:"properties"`
		} `json:"overrides"`
	} `json:"fieldConfig"`
	GridPos *ModelGridPos `json:"gridPos,omitempty"`

	// TODO docs
	Id *int `json:"id,omitempty"`

	// TODO docs
	// TODO tighter constraint
	Interval *string `json:"interval,omitempty"`

	// Panel links.
	// TODO fill this out - seems there are a couple variants?
	Links *[]ModelDashboardLink `json:"links,omitempty"`

	// TODO docs
	MaxDataPoints *float32 `json:"maxDataPoints,omitempty"`

	// options is specified by the PanelOptions field in panel
	// plugin schemas.
	Options map[string]interface{} `json:"options"`

	// FIXME this almost certainly has to be changed in favor of scuemata versions
	PluginVersion *string `json:"pluginVersion,omitempty"`

	// Name of template variable to repeat for.
	Repeat *string `json:"repeat,omitempty"`

	// Direction to repeat in if 'repeat' is set.
	// "h" for horizontal, "v" for vertical.
	RepeatDirection ModelPanelRepeatDirection `json:"repeatDirection"`

	// TODO docs
	Tags *[]string `json:"tags,omitempty"`

	// TODO docs
	Targets *[]ModelTarget `json:"targets,omitempty"`

	// TODO docs - seems to be an old field from old dashboard alerts?
	Thresholds *[]interface{} `json:"thresholds,omitempty"`

	// TODO docs
	// TODO tighter constraint
	TimeFrom *string `json:"timeFrom,omitempty"`

	// TODO docs
	TimeRegions *[]interface{} `json:"timeRegions,omitempty"`

	// TODO docs
	// TODO tighter constraint
	TimeShift *string `json:"timeShift,omitempty"`

	// Panel title.
	Title           *string `json:"title,omitempty"`
	Transformations []struct {
		Id      string                 `json:"id"`
		Options map[string]interface{} `json:"options"`
	} `json:"transformations"`

	// Whether to display the panel without a background.
	Transparent bool `json:"transparent"`

	// The panel plugin type id. May not be empty.
	Type string `json:"type"`
}

// Direction to repeat in if 'repeat' is set.
// "h" for horizontal, "v" for vertical.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelPanelRepeatDirection string

// Row panel
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelRowPanel struct {
	Collapsed bool `json:"collapsed"`

	// Name of default datasource.
	Datasource *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	GridPos *ModelGridPos `json:"gridPos,omitempty"`
	Id      int           `json:"id"`
	Panels  []interface{} `json:"panels"`

	// Name of template variable to repeat for.
	Repeat *string           `json:"repeat,omitempty"`
	Title  *string           `json:"title,omitempty"`
	Type   ModelRowPanelType `json:"type"`
}

// DashboardRowPanelType defines model for DashboardRowPanel.Type.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelRowPanelType string

// Schema for panel targets is specified by datasource
// plugins. We use a placeholder definition, which the Go
// schema loader either left open/as-is with the Base
// variant of the Dashboard and Panel families, or filled
// with types derived from plugins in the Instance variant.
// When working directly from CUE, importers can extend this
// type directly to achieve the same effect.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelTarget map[string]interface{}

// TODO docs
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelThreshold struct {
	// TODO docs
	Color string `json:"color"`

	// TODO docs
	// TODO are the values here enumerable into a disjunction?
	// Some seem to be listed in typescript comment
	State *string `json:"state,omitempty"`

	// TODO docs
	// FIXME the corresponding typescript field is required/non-optional, but nulls currently appear here when serializing -Infinity to JSON
	Value *float32 `json:"value,omitempty"`
}

// DashboardThresholdsConfig defines model for dashboard.ThresholdsConfig.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelThresholdsConfig struct {
	Mode ModelThresholdsConfigMode `json:"mode"`

	// Must be sorted by 'value', first value is always -Infinity
	Steps []struct {
		// TODO docs
		Color string `json:"color"`

		// TODO docs
		// TODO are the values here enumerable into a disjunction?
		// Some seem to be listed in typescript comment
		State *string `json:"state,omitempty"`

		// TODO docs
		// FIXME the corresponding typescript field is required/non-optional, but nulls currently appear here when serializing -Infinity to JSON
		Value *float32 `json:"value,omitempty"`
	} `json:"steps"`
}

// DashboardThresholdsConfigMode defines model for DashboardThresholdsConfig.Mode.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelThresholdsConfigMode string

// DashboardThresholdsMode defines model for dashboard.ThresholdsMode.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelThresholdsMode string

// TODO docs
// FIXME this is extremely underspecfied; wasn't obvious which typescript types corresponded to it
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelTransformation struct {
	Id      string                 `json:"id"`
	Options map[string]interface{} `json:"options"`
}

// FROM: packages/grafana-data/src/types/templateVars.ts
// TODO docs
// TODO what about what's in public/app/features/types.ts?
// TODO there appear to be a lot of different kinds of [template] vars here? if so need a disjunction
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelVariableModel struct {
	Label *string                `json:"label,omitempty"`
	Name  string                 `json:"name"`
	Type  ModelVariableModelType `json:"type"`
}

// DashboardVariableModelType defines model for DashboardVariableModel.Type.
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelVariableModelType string

// FROM: packages/grafana-data/src/types/templateVars.ts
// TODO docs
// TODO this implies some wider pattern/discriminated union, probably?
//
// THIS TYPE IS INTENDED FOR INTERNAL USE BY THE GRAFANA BACKEND, AND IS SUBJECT TO BREAKING CHANGES.
// Equivalent Go types at stable import paths are provided in https://github.com/grafana/grok.
type ModelVariableType string

//go:embed lineage.cue
var cueFS embed.FS

// codegen ensures that this is always the latest Thema schema version
var currentVersion = thema.SV(0, 0)

// Lineage returns the Thema lineage representing a Grafana dashboard.
//
// The lineage is the canonical specification of the current dashboard schema,
// all prior schema versions, and the mappings that allow migration between
// schema versions.
func Lineage(lib thema.Library, opts ...thema.BindOption) (thema.Lineage, error) {
	return cuectx.LoadGrafanaInstancesWithThema(filepath.Join("pkg", "coremodel", "dashboard"), cueFS, lib, opts...)
}

var _ thema.LineageFactory = Lineage

// Coremodel contains the foundational schema declaration for dashboards.
type Coremodel struct {
	lin thema.Lineage
}

// Lineage returns the canonical dashboard Lineage.
func (c *Coremodel) Lineage() thema.Lineage {
	return c.lin
}

// CurrentSchema returns the current (latest) dashboard Thema schema.
func (c *Coremodel) CurrentSchema() thema.Schema {
	return thema.SchemaP(c.lin, currentVersion)
}

// GoType returns a pointer to an empty Go struct that corresponds to
// the current Thema schema.
func (c *Coremodel) GoType() interface{} {
	return &Model{}
}

func ProvideCoremodel(lib thema.Library) (*Coremodel, error) {
	lin, err := Lineage(lib)
	if err != nil {
		return nil, err
	}

	return &Coremodel{
		lin: lin,
	}, nil
}
