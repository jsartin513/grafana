// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     CommonSchemaJenny
//
// Run 'make gen-cue' from repository root to regenerate.


/**
 *  See also:
 *  https://github.com/grafana/grafana-plugin-sdk-go/blob/main/data/frame_type.go
 */
export enum DataFrameType {
  DirectoryListing = 'directory-listing',
  HeatmapCells = 'heatmap-cells',
  HeatmapRows = 'heatmap-rows',
  Histogram = 'histogram',
  TimeSeriesLong = 'timeseries-long',
  TimeSeriesMany = 'timeseries-many',
  TimeSeriesMulti = 'timeseries-multi',
  TimeSeriesWide = 'timeseries-wide',
}

/**
 * TODO extends #QueryResultBase
 */
export interface DataFrame {
  /**
   * All fields of equal length
   */
  fields: Array<Field>;
  /**
   * The number of rows
   */
  length: number;
  name?: string;
}

export const defaultDataFrame: Partial<DataFrame> = {
  fields: [],
};

/**
 * TODO Field<T = any, V = Vector<T>>
 */
export interface Field {
  /**
   * Meta info about how field and how to display it
   */
  config: FieldConfig;
  /**
   * Convert a value for display TODO extend in veneer
   */
  display?: unknown;
  /**
   * Get value data links with variables interpolated. Extended in veneer
   */
  getLinks?: unknown;
  labels?: Labels;
  /**
   * Name of the field (column)
   */
  name: string;
  /**
   * Cached values with appropriate display and id values TODO | null
   */
  state?: FieldState;
  /**
   * Field value type (string, number, etc)
   */
  type: FieldType;
  /**
   * The raw field values. Extended in veneer
   */
  values: Record<string, unknown>;
}

export enum FieldType {
  bool = 'bool',
  geo = 'geo',
  number = 'number',
  other = 'other',
  string = 'string',
  time = 'time',
  trace = 'trace',
}

/**
 * Every property is optional
 * Plugins may extend this with additional properties. Something like series overrides
 */
export interface FieldConfig {
  /**
   * Significant digits (for display) TODO this should be a separate type
   */
  decimals?: number;
  /**
   * Human readable field metadata
   */
  description?: string;
  /**
   * The display value for this field.  This supports template variables blank is auto
   */
  displayName?: string;
  /**
   * This can be used by data sources that return and explicit naming structure for values and labels
   * When this property is configured, this value is used rather than the default naming strategy.
   */
  displayNameFromDS?: string;
  /**
   * True if data source field supports ad-hoc filters
   */
  filterable?: boolean;
  /**
   * Interval indicates the expected regular step between values in the series.
   * When an interval exists, consumers can identify "missing" values when the expected value is not present.
   * The grafana timeseries visualization will render disconnected values when missing values are found it the time field.
   * The interval uses the same units as the values.  For time.Time, this is defined in milliseconds.
   * TODO | null
   */
  interval?: number;
  /**
   * Convert input values into a display string
   */
  mappings?: Array<ValueMapping>;
  /**
   * TODO | null
   */
  max?: number;
  /**
   * TODO | null
   */
  min?: number;
  /**
   * An explict path to the field in the datasource.  When the frame meta includes a path,
   * This will default to `${frame.meta.path}/${field.name}
   * When defined, this value can be used as an identifier within the datasource scope, and
   * may be used to update the results
   */
  path?: string;
  /**
   * Map numeric values to states
   */
  thresholds?: Array<ThresholdsConfig>;
  /**
   * Numeric Options
   */
  unit?: string;
  /**
   * True if data source can write a value to the path.  Auth/authz are supported separately
   */
  writeable?: boolean;
}

export const defaultFieldConfig: Partial<FieldConfig> = {
  mappings: [],
  thresholds: [],
};

export interface FieldState {
  /**
   * Cache of reduced values
   */
  calcs?: Record<string, unknown>;
  /**
   * An appropriate name for the field (does not include frame info) TODO | null
   */
  displayName?: string;
  /**
   * Boolean value is true if field is in a larger data set with multiple frames.
   * This is only related to the cached displayName property above.
   */
  multipleFrames?: boolean;
  /**
   * Boolean value is true if a null filling threshold has been applied
   * against the frame of the field. This is used to avoid cases in which
   * this would applied more than one time.
   */
  nullThresholdApplied?: boolean;
  /**
   * Location of this field within the context frames results
   * @internal -- we will try to make this unnecessary
   */
  origin?: DataFrameFieldIndex;
  /**
   * The numeric range for values in this field.  This value will respect the min/max
   * set in field config, or when set to `auto` this will have the min/max for all data
   * in the response
   */
  range?: NumericRange;
  /**
   * Appropriate values for templating
   */
  scopedVars?: ScopedVars;
  /**
   * Series index is index for this field in a larger data set that can span multiple DataFrames
   * Useful for assigning color to series by looking up a color in a palette using this index
   */
  seriesIndex?: number;
}

/**
 * TODO docs
 */
export interface NumericRange {
  delta: number;
  max?: number;
  min?: number;
}

export interface DataFrameFieldIndex {
  fieldIndex: number;
  frameIndex: number;
}

/**
 * TODO Duplicate declaration
 */
export interface ThresholdsConfig {
  mode: ThresholdsMode;
  /**
   * Must be sorted by 'value', first value is always -Infinity
   */
  steps: Array<Threshold>;
}

export const defaultThresholdsConfig: Partial<ThresholdsConfig> = {
  steps: [],
};

/**
 * TODO Duplicate declaration
 */
export interface Threshold {
  /**
   * TODO docs
   */
  color: string;
  /**
   * TODO docs
   * TODO are the values here enumerable into a disjunction?
   * Some seem to be listed in typescript comment
   */
  state?: string;
  /**
   * TODO docs
   * FIXME the corresponding typescript field is required/non-optional, but nulls currently appear here when serializing -Infinity to JSON
   */
  value?: number;
}

/**
 * TODO Duplicate declaration
 */
export enum ThresholdsMode {
  Absolute = 'absolute',
  Percentage = 'percentage',
}

/**
 * TODO docs | Duplicate declaration
 */
export type ValueMapping = (ValueMap | RangeMap | RegexMap | SpecialValueMap);

/**
 * TODO docs | Duplicate declaration
 */
export enum MappingType {
  RangeToText = 'range',
  RegexToText = 'regex',
  SpecialValue = 'special',
  ValueToText = 'value',
}

/**
 * TODO docs | Duplicate declaration
 */
export interface ValueMap {
  options: Record<string, ValueMappingResult>;
  type: MappingType.ValueToText;
}

/**
 * TODO docs | Duplicate declaration
 */
export interface RangeMap {
  options: {
    /**
     * to and from are `number | null` in current ts, really not sure what to do
     */
    from: number;
    to: number;
    result: ValueMappingResult;
  };
  type: MappingType.RangeToText;
}

/**
 * TODO docs | Duplicate declaration
 */
export interface RegexMap {
  options: {
    pattern: string;
    result: ValueMappingResult;
  };
  type: MappingType.RegexToText;
}

/**
 * TODO docs | Duplicate declaration
 */
export interface SpecialValueMap {
  options: {
    match: ('true' | 'false');
    pattern: string;
    result: ValueMappingResult;
  };
  type: MappingType.SpecialValue;
}

export interface DataSourceJsonData {
  alertmanagerUid?: string;
  authType?: string;
  defaultRegion?: string;
  manageAlerts?: boolean;
  profile?: string;
}

/**
 * These are the common properties available to all queries in all datasources.
 * Specific implementations will *extend* this interface, adding the required
 * properties for the given context.
 */
export interface DataQuery {
  /**
   * For mixed data sources the selected datasource is on the query level.
   * For non mixed scenarios this is undefined.
   * TODO find a better way to do this ^ that's friendly to schema
   * TODO this shouldn't be unknown but DataSourceRef | null
   */
  datasource?: unknown;
  /**
   * true if query is disabled (ie should not be returned to the dashboard)
   */
  hide?: boolean;
  /**
   * Unique, guid like, string used in explore mode
   */
  key?: string;
  /**
   * Specify the query flavor
   * TODO make this required and give it a default
   */
  queryType?: string;
  /**
   * A - Z
   */
  refId: string;
}

/**
 * TODO docs
 */
export enum AxisPlacement {
  Auto = 'auto',
  Bottom = 'bottom',
  Hidden = 'hidden',
  Left = 'left',
  Right = 'right',
  Top = 'top',
}

/**
 * TODO docs
 */
export enum AxisColorMode {
  Series = 'series',
  Text = 'text',
}

/**
 * TODO docs
 */
export enum VisibilityMode {
  Always = 'always',
  Auto = 'auto',
  Never = 'never',
}

/**
 * TODO docs
 */
export enum GraphDrawStyle {
  Bars = 'bars',
  Line = 'line',
  Points = 'points',
}

/**
 * TODO docs
 */
export enum GraphTransform {
  Constant = 'constant',
  NegativeY = 'negative-Y',
}

/**
 * TODO docs
 */
export enum LineInterpolation {
  Linear = 'linear',
  Smooth = 'smooth',
  StepAfter = 'stepAfter',
  StepBefore = 'stepBefore',
}

/**
 * TODO docs
 */
export enum ScaleDistribution {
  Linear = 'linear',
  Log = 'log',
  Ordinal = 'ordinal',
  Symlog = 'symlog',
}

/**
 * TODO docs
 */
export enum GraphGradientMode {
  Hue = 'hue',
  None = 'none',
  Opacity = 'opacity',
  Scheme = 'scheme',
}

/**
 * TODO docs
 */
export enum StackingMode {
  None = 'none',
  Normal = 'normal',
  Percent = 'percent',
}

/**
 * TODO docs
 */
export enum BarAlignment {
  After = 1,
  Before = -1,
  Center = 0,
}

/**
 * TODO docs
 */
export enum ScaleOrientation {
  Horizontal = 0,
  Vertical = 1,
}

/**
 * TODO docs
 */
export enum ScaleDirection {
  Down = -1,
  Left = -1,
  Right = 1,
  Up = 1,
}

/**
 * TODO docs
 */
export interface LineStyle {
  dash?: Array<number>;
  fill?: ('solid' | 'dash' | 'dot' | 'square');
}

export const defaultLineStyle: Partial<LineStyle> = {
  dash: [],
};

/**
 * TODO docs
 */
export interface LineConfig {
  lineColor?: string;
  lineInterpolation?: LineInterpolation;
  lineStyle?: LineStyle;
  lineWidth?: number;
  /**
   * Indicate if null values should be treated as gaps or connected.
   * When the value is a number, it represents the maximum delta in the
   * X axis that should be considered connected.  For timeseries, this is milliseconds
   */
  spanNulls?: (boolean | number);
}

/**
 * TODO docs
 */
export interface BarConfig {
  barAlignment?: BarAlignment;
  barMaxWidth?: number;
  barWidthFactor?: number;
}

/**
 * TODO docs
 */
export interface FillConfig {
  fillBelowTo?: string;
  fillColor?: string;
  fillOpacity?: number;
}

/**
 * TODO docs
 */
export interface PointsConfig {
  pointColor?: string;
  pointSize?: number;
  pointSymbol?: string;
  showPoints?: VisibilityMode;
}

/**
 * TODO docs
 */
export interface ScaleDistributionConfig {
  linearThreshold?: number;
  log?: number;
  type: ScaleDistribution;
}

/**
 * TODO docs
 */
export interface AxisConfig {
  axisCenteredZero?: boolean;
  axisColorMode?: AxisColorMode;
  axisGridShow?: boolean;
  axisLabel?: string;
  axisPlacement?: AxisPlacement;
  axisSoftMax?: number;
  axisSoftMin?: number;
  axisWidth?: number;
  scaleDistribution?: ScaleDistributionConfig;
}

/**
 * TODO docs
 */
export interface HideSeriesConfig {
  legend: boolean;
  tooltip: boolean;
  viz: boolean;
}

/**
 * TODO docs
 */
export interface StackingConfig {
  group?: string;
  mode?: StackingMode;
}

/**
 * TODO docs
 */
export interface StackableFieldConfig {
  stacking?: StackingConfig;
}

/**
 * TODO docs
 */
export interface HideableFieldConfig {
  hideFrom?: HideSeriesConfig;
}

/**
 * TODO docs
 */
export enum GraphTresholdsStyleMode {
  Area = 'area',
  Dashed = 'dashed',
  DashedAndArea = 'dashed+area',
  Line = 'line',
  LineAndArea = 'line+area',
  Off = 'off',
  Series = 'series',
}

/**
 * TODO docs
 */
export interface GraphThresholdsStyleConfig {
  mode: GraphTresholdsStyleMode;
}

/**
 * TODO docs
 */
export type LegendPlacement = ('bottom' | 'right');

/**
 * TODO docs
 * Note: "hidden" needs to remain as an option for plugins compatibility
 */
export enum LegendDisplayMode {
  Hidden = 'hidden',
  List = 'list',
  Table = 'table',
}

/**
 * TODO docs
 */
export interface SingleStatBaseOptions extends OptionsWithTextFormatting {
  orientation: VizOrientation;
  reduceOptions: ReduceDataOptions;
}

/**
 * TODO docs
 */
export interface ReduceDataOptions {
  /**
   * When !values, pick one value for the whole field
   */
  calcs: Array<string>;
  /**
   * Which fields to show.  By default this is only numeric fields
   */
  fields?: string;
  /**
   * if showing all values limit
   */
  limit?: number;
  /**
   * If true show each row value
   */
  values?: boolean;
}

export const defaultReduceDataOptions: Partial<ReduceDataOptions> = {
  calcs: [],
};

/**
 * TODO docs
 */
export enum VizOrientation {
  Auto = 'auto',
  Horizontal = 'horizontal',
  Vertical = 'vertical',
}

/**
 * TODO docs
 */
export interface OptionsWithTooltip {
  tooltip: VizTooltipOptions;
}

/**
 * TODO docs
 */
export interface OptionsWithLegend {
  legend: VizLegendOptions;
}

/**
 * TODO docs
 */
export interface OptionsWithTimezones {
  timezone?: Array<TimeZone>;
}

export const defaultOptionsWithTimezones: Partial<OptionsWithTimezones> = {
  timezone: [],
};

/**
 * TODO docs
 */
export interface OptionsWithTextFormatting {
  text?: VizTextDisplayOptions;
}

/**
 * TODO docs
 */
export enum BigValueColorMode {
  Background = 'background',
  None = 'none',
  Value = 'value',
}

/**
 * TODO docs
 */
export enum BigValueGraphMode {
  Area = 'area',
  Line = 'line',
  None = 'none',
}

/**
 * TODO docs
 */
export enum BigValueJustifyMode {
  Auto = 'auto',
  Center = 'center',
}

/**
 * TODO docs
 */
export enum BigValueTextMode {
  Auto = 'auto',
  Name = 'name',
  None = 'none',
  Value = 'value',
  ValueAndName = 'value_and_name',
}

/**
 * TODO -- should not be table specific!
 * TODO docs
 */
export type FieldTextAlignment = ('auto' | 'left' | 'right' | 'center');

/**
 * TODO docs
 */
export interface VizTextDisplayOptions {
  /**
   * Explicit title text size
   */
  titleSize?: number;
  /**
   * Explicit value text size
   */
  valueSize?: number;
}

/**
 * TODO docs
 */
export enum TooltipDisplayMode {
  Multi = 'multi',
  None = 'none',
  Single = 'single',
}

/**
 * TODO docs
 */
export enum SortOrder {
  Ascending = 'asc',
  Descending = 'desc',
  None = 'none',
}

/**
 * TODO docs
 */
export interface GraphFieldConfig extends LineConfig, FillConfig, PointsConfig, AxisConfig, BarConfig, StackableFieldConfig, HideableFieldConfig {
  drawStyle?: GraphDrawStyle;
  gradientMode?: GraphGradientMode;
  thresholdsStyle?: GraphThresholdsStyleConfig;
  transform?: GraphTransform;
}

/**
 * TODO docs
 */
export interface VizLegendOptions {
  asTable?: boolean;
  calcs: Array<string>;
  displayMode: LegendDisplayMode;
  isVisible?: boolean;
  placement: LegendPlacement;
  showLegend: boolean;
  sortBy?: string;
  sortDesc?: boolean;
  width?: number;
}

export const defaultVizLegendOptions: Partial<VizLegendOptions> = {
  calcs: [],
};

/**
 * Enum expressing the possible display modes
 * for the bar gauge component of Grafana UI
 */
export enum BarGaugeDisplayMode {
  Basic = 'basic',
  Gradient = 'gradient',
  Lcd = 'lcd',
}

/**
 * TODO docs
 */
export interface VizTooltipOptions {
  mode: TooltipDisplayMode;
  sort: SortOrder;
}

export interface Labels {}

/**
 * TODO docs | generic type
 */
export interface ScopedVar {
  text: unknown;
  value: unknown;
}

/**
 * TODO docs
 */
export interface ScopedVars {}

/**
 * TODO Should be moved to common data query?
 */
export interface QueryResultMeta {
  /**
   * The path for live stream updates for this frame
   */
  channel?: string;
  /**
   * DataSource Specific Values
   */
  custom?: Record<string, unknown>;
  /**
   * Optionally identify which topic the frame should be assigned to.
   * A value specified in the response will override what the request asked for.
   */
  dataTopic?: DataTopic;
  /**
   * This is the raw query sent to the underlying system.  All macros and templating
   * as been applied.  When metadata contains this value, it will be shown in the query inspector
   */
  executedQueryString?: string;
  instant?: boolean;
  /**
   * Did the query response come from the cache
   */
  isCachedResponse?: boolean;
  /**
   * used to keep track of old json doc values
   */
  json?: boolean;
  /**
   * used by log models and loki
   */
  limit?: number;
  /**
   * Meta notices
   */
  notices?: Array<QueryResultMetaNotice>;
  /**
   * A browsable path on the datasource
   */
  path?: string;
  /**
   * defaults to '/'
   */
  pathSeparator?: string;
  /**
   * Currently used to show results in Explore only in preferred visualisation option
   */
  preferredVisualisationType?: PreferredVisualisationType;
  /**
   * Legacy data source specific, should be moved to custom
   * used by log models and loki
   */
  searchWords?: Array<string>;
  /**
   * Stats
   */
  stats?: Array<QueryResultMetaStat>;
  /**
   * Used to track transformation ids that where part of the processing
   */
  transformations?: Array<string>;
  type?: DataFrameType;
}

export const defaultQueryResultMeta: Partial<QueryResultMeta> = {
  notices: [],
  searchWords: [],
  stats: [],
  transformations: [],
};

/**
 * TODO this is enum with one field
 * Attached to query results (not persisted)
 */
export type DataTopic = ('annotations' | '');

/**
 * TODO extends FieldConfig
 */
export interface QueryResultMetaStat {
  displayName: string;
  value: number;
}

export interface QueryResultMetaNotice {
  /**
   * Optionally suggest an appropriate tab for the panel inspector
   */
  inspect?: ('meta' | 'error' | 'data' | 'stats');
  /**
   * An optional link that may be displayed in the UI.
   * This value may be an absolute URL or relative to grafana root
   */
  link?: string;
  /**
   * Specify the notice severity
   */
  severity: ('info' | 'warning' | 'error');
  /**
   * Notice descriptive text
   */
  text: string;
}

/**
 * Internally, this is the "type" of cell that's being displayed
 * in the table such as colored text, JSON, gauge, etc.
 * The color-background-solid, gradient-gauge, and lcd-gauge
 * modes are deprecated in favor of new cell subOptions
 */
export enum TableCellDisplayMode {
  Auto = 'auto',
  BasicGauge = 'basic',
  ColorBackground = 'color-background',
  ColorBackgroundSolid = 'color-background-solid',
  ColorText = 'color-text',
  Gauge = 'gauge',
  GradientGauge = 'gradient-gauge',
  Image = 'image',
  JSONView = 'json-view',
  LcdGauge = 'lcd-gauge',
}

/**
 * Display mode to the "Colored Background" display
 * mode for table cells. Either displays a solid color (basic mode)
 * or a gradient.
 */
export enum TableCellBackgroundDisplayMode {
  Basic = 'basic',
  Gradient = 'gradient',
}

/**
 * TODO docs
 */
export interface TableSortByFieldState {
  desc?: boolean;
  displayName: string;
}

/**
 * Auto mode table cell options
 */
export interface TableAutoCellOptions {
  type: TableCellDisplayMode.Auto;
}

/**
 * Colored text cell options
 */
export interface TableColorTextCellOptions {
  type: TableCellDisplayMode.ColorText;
}

/**
 * Json view cell options
 */
export interface TableJsonViewCellOptions {
  type: TableCellDisplayMode.JSONView;
}

/**
 * Json view cell options
 */
export interface TableImageCellOptions {
  type: TableCellDisplayMode.Image;
}

/**
 * Gauge cell options
 */
export interface TableBarGaugeCellOptions {
  mode?: BarGaugeDisplayMode;
  type: TableCellDisplayMode.Gauge;
}

/**
 * Colored background cell options
 */
export interface TableColoredBackgroundCellOptions {
  mode?: TableCellBackgroundDisplayMode;
  type: TableCellDisplayMode.ColorBackground;
}

/**
 * Table cell options. Each cell has a display mode
 * and other potential options for that display.
 */
export type TableCellOptions = (TableAutoCellOptions | TableBarGaugeCellOptions | TableColoredBackgroundCellOptions | TableColorTextCellOptions | TableImageCellOptions | TableJsonViewCellOptions);

/**
 * Use UTC/GMT timezone
 */
export type TimeZoneUtc = 'utc';

/**
 * Use the timezone defined by end user web browser
 */
export type TimeZoneBrowser = 'browser';

/**
 * TODO duplicate
 */
export interface ValueMappingResult {
  color?: string;
  icon?: string;
  index?: number;
  text?: string;
}

export interface DataSourceRef {
  /**
   * The plugin type-id
   */
  type?: string;
  /**
   * Specific datasource instance
   */
  uid?: string;
}

export type PreferredVisualisationType = ('graph' | 'table' | 'logs' | 'trace' | 'nodeGraph' | 'flamegraph' | 'rawPrometheus');

/**
 * Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
 * Generally defines alignment, filtering capabilties, display options, etc.
 */
export interface TableFieldOptions {
  align: FieldTextAlignment;
  cellOptions: TableCellOptions;
  /**
   * This field is deprecated in favor of using cellOptions
   */
  displayMode?: TableCellDisplayMode;
  filterable?: boolean;
  hidden?: boolean;
  inspect: boolean;
  minWidth?: number;
  width?: number;
}

export const defaultTableFieldOptions: Partial<TableFieldOptions> = {
  align: 'auto',
  inspect: false,
};

/**
 * A specific timezone from https://en.wikipedia.org/wiki/Tz_database
 */
export type TimeZone = (TimeZoneUtc | TimeZoneBrowser | string);

export const defaultTimeZone: TimeZone = 'browser';
