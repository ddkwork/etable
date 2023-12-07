// Code generated by "goki generate"; DO NOT EDIT.

package etview

import (
	"reflect"
	"sync"

	"goki.dev/colors/colormap"
	"goki.dev/etable/v2/etensor"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/giv"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/mat32/v2"
	"goki.dev/ordmap"
)

// SimMatGridType is the [gti.Type] for [SimMatGrid]
var SimMatGridType = gti.AddType(&gti.Type{
	Name:      "goki.dev/etable/v2/etview.SimMatGrid",
	ShortName: "etview.SimMatGrid",
	IDName:    "sim-mat-grid",
	Doc:       "SimMatGrid is a widget that displays a similarity / distance matrix\nwith tensor values as a grid of colored squares, and labels for rows, cols",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"SimMat", &gti.Field{Name: "SimMat", Type: "*goki.dev/etable/v2/simat.SimMat", LocalType: "*simat.SimMat", Doc: "the similarity / distance matrix", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"rowMaxSz", &gti.Field{Name: "rowMaxSz", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"rowMinBlank", &gti.Field{Name: "rowMinBlank", Type: "int", LocalType: "int", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"rowNGps", &gti.Field{Name: "rowNGps", Type: "int", LocalType: "int", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"colMaxSz", &gti.Field{Name: "colMaxSz", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"colMinBlank", &gti.Field{Name: "colMinBlank", Type: "int", LocalType: "int", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"colNGps", &gti.Field{Name: "colNGps", Type: "int", LocalType: "int", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"TensorGrid", &gti.Field{Name: "TensorGrid", Type: "goki.dev/etable/v2/etview.TensorGrid", LocalType: "TensorGrid", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &SimMatGrid{},
})

// NewSimMatGrid adds a new [SimMatGrid] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewSimMatGrid(par ki.Ki, name ...string) *SimMatGrid {
	return par.NewChild(SimMatGridType, name...).(*SimMatGrid)
}

// KiType returns the [*gti.Type] of [SimMatGrid]
func (t *SimMatGrid) KiType() *gti.Type {
	return SimMatGridType
}

// New returns a new [*SimMatGrid] value
func (t *SimMatGrid) New() ki.Ki {
	return &SimMatGrid{}
}

// SetRowMaxSz sets the [SimMatGrid.rowMaxSz]
func (t *SimMatGrid) SetRowMaxSz(v mat32.Vec2) *SimMatGrid {
	t.rowMaxSz = v
	return t
}

// SetRowMinBlank sets the [SimMatGrid.rowMinBlank]
func (t *SimMatGrid) SetRowMinBlank(v int) *SimMatGrid {
	t.rowMinBlank = v
	return t
}

// SetRowNgps sets the [SimMatGrid.rowNGps]
func (t *SimMatGrid) SetRowNgps(v int) *SimMatGrid {
	t.rowNGps = v
	return t
}

// SetColMaxSz sets the [SimMatGrid.colMaxSz]
func (t *SimMatGrid) SetColMaxSz(v mat32.Vec2) *SimMatGrid {
	t.colMaxSz = v
	return t
}

// SetColMinBlank sets the [SimMatGrid.colMinBlank]
func (t *SimMatGrid) SetColMinBlank(v int) *SimMatGrid {
	t.colMinBlank = v
	return t
}

// SetColNgps sets the [SimMatGrid.colNGps]
func (t *SimMatGrid) SetColNgps(v int) *SimMatGrid {
	t.colNGps = v
	return t
}

// SetTooltip sets the [SimMatGrid.Tooltip]
func (t *SimMatGrid) SetTooltip(v string) *SimMatGrid {
	t.Tooltip = v
	return t
}

// SetClass sets the [SimMatGrid.Class]
func (t *SimMatGrid) SetClass(v string) *SimMatGrid {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [SimMatGrid.PriorityEvents]
func (t *SimMatGrid) SetPriorityEvents(v []events.Types) *SimMatGrid {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [SimMatGrid.CustomContextMenu]
func (t *SimMatGrid) SetCustomContextMenu(v func(m *gi.Scene)) *SimMatGrid {
	t.CustomContextMenu = v
	return t
}

// SetDisp sets the [SimMatGrid.Disp]
func (t *SimMatGrid) SetDisp(v TensorDisp) *SimMatGrid {
	t.Disp = v
	return t
}

// SetColorMap sets the [SimMatGrid.ColorMap]
func (t *SimMatGrid) SetColorMap(v *colormap.Map) *SimMatGrid {
	t.ColorMap = v
	return t
}

// TableViewType is the [gti.Type] for [TableView]
var TableViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/etable/v2/etview.TableView",
	ShortName:  "etview.TableView",
	IDName:     "table-view",
	Doc:        "etview.TableView provides a GUI interface for etable.Table's",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Table", &gti.Field{Name: "Table", Type: "*goki.dev/etable/v2/etable.IdxView", LocalType: "*etable.IdxView", Doc: "the idx view of the table that we're a view of", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"TsrDisp", &gti.Field{Name: "TsrDisp", Type: "goki.dev/etable/v2/etview.TensorDisp", LocalType: "TensorDisp", Doc: "overall display options for tensor display", Directives: gti.Directives{}, Tag: ""}},
		{"ColTsrDisp", &gti.Field{Name: "ColTsrDisp", Type: "map[int]*goki.dev/etable/v2/etview.TensorDisp", LocalType: "map[int]*TensorDisp", Doc: "per column tensor display params", Directives: gti.Directives{}, Tag: ""}},
		{"ColTsrBlank", &gti.Field{Name: "ColTsrBlank", Type: "map[int]*goki.dev/etable/v2/etensor.Float64", LocalType: "map[int]*etensor.Float64", Doc: "per column blank tensor values", Directives: gti.Directives{}, Tag: ""}},
		{"NCols", &gti.Field{Name: "NCols", Type: "int", LocalType: "int", Doc: "number of columns in table (as of last update)", Directives: gti.Directives{}, Tag: "inactive:\"+\""}},
		{"SortIdx", &gti.Field{Name: "SortIdx", Type: "int", LocalType: "int", Doc: "current sort index", Directives: gti.Directives{}, Tag: ""}},
		{"SortDesc", &gti.Field{Name: "SortDesc", Type: "bool", LocalType: "bool", Doc: "whether current sort order is descending", Directives: gti.Directives{}, Tag: ""}},
		{"BlankString", &gti.Field{Name: "BlankString", Type: "string", LocalType: "string", Doc: "\tblank values for out-of-range rows", Directives: gti.Directives{}, Tag: ""}},
		{"BlankFloat", &gti.Field{Name: "BlankFloat", Type: "float64", LocalType: "float64", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"SliceViewBase", &gti.Field{Name: "SliceViewBase", Type: "goki.dev/gi/v2/giv.SliceViewBase", LocalType: "giv.SliceViewBase", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TableView{},
})

// NewTableView adds a new [TableView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTableView(par ki.Ki, name ...string) *TableView {
	return par.NewChild(TableViewType, name...).(*TableView)
}

// KiType returns the [*gti.Type] of [TableView]
func (t *TableView) KiType() *gti.Type {
	return TableViewType
}

// New returns a new [*TableView] value
func (t *TableView) New() ki.Ki {
	return &TableView{}
}

// SetTsrDisp sets the [TableView.TsrDisp]:
// overall display options for tensor display
func (t *TableView) SetTsrDisp(v TensorDisp) *TableView {
	t.TsrDisp = v
	return t
}

// SetColTsrDisp sets the [TableView.ColTsrDisp]:
// per column tensor display params
func (t *TableView) SetColTsrDisp(v map[int]*TensorDisp) *TableView {
	t.ColTsrDisp = v
	return t
}

// SetColTsrBlank sets the [TableView.ColTsrBlank]:
// per column blank tensor values
func (t *TableView) SetColTsrBlank(v map[int]*etensor.Float64) *TableView {
	t.ColTsrBlank = v
	return t
}

// SetNcols sets the [TableView.NCols]:
// number of columns in table (as of last update)
func (t *TableView) SetNcols(v int) *TableView {
	t.NCols = v
	return t
}

// SetSortIdx sets the [TableView.SortIdx]:
// current sort index
func (t *TableView) SetSortIdx(v int) *TableView {
	t.SortIdx = v
	return t
}

// SetSortDesc sets the [TableView.SortDesc]:
// whether current sort order is descending
func (t *TableView) SetSortDesc(v bool) *TableView {
	t.SortDesc = v
	return t
}

// SetBlankString sets the [TableView.BlankString]:
//
//	blank values for out-of-range rows
func (t *TableView) SetBlankString(v string) *TableView {
	t.BlankString = v
	return t
}

// SetBlankFloat sets the [TableView.BlankFloat]
func (t *TableView) SetBlankFloat(v float64) *TableView {
	t.BlankFloat = v
	return t
}

// SetTooltip sets the [TableView.Tooltip]
func (t *TableView) SetTooltip(v string) *TableView {
	t.Tooltip = v
	return t
}

// SetClass sets the [TableView.Class]
func (t *TableView) SetClass(v string) *TableView {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [TableView.PriorityEvents]
func (t *TableView) SetPriorityEvents(v []events.Types) *TableView {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [TableView.CustomContextMenu]
func (t *TableView) SetCustomContextMenu(v func(m *gi.Scene)) *TableView {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [TableView.StackTop]
func (t *TableView) SetStackTop(v int) *TableView {
	t.StackTop = v
	return t
}

// SetStripes sets the [TableView.Stripes]
func (t *TableView) SetStripes(v gi.Stripes) *TableView {
	t.Stripes = v
	return t
}

// SetMinRows sets the [TableView.MinRows]
func (t *TableView) SetMinRows(v int) *TableView {
	t.MinRows = v
	return t
}

// SetViewMu sets the [TableView.ViewMu]
func (t *TableView) SetViewMu(v *sync.Mutex) *TableView {
	t.ViewMu = v
	return t
}

// SetSliceNpval sets the [TableView.SliceNPVal]
func (t *TableView) SetSliceNpval(v reflect.Value) *TableView {
	t.SliceNPVal = v
	return t
}

// SetSliceValView sets the [TableView.SliceValView]
func (t *TableView) SetSliceValView(v giv.Value) *TableView {
	t.SliceValView = v
	return t
}

// SetValues sets the [TableView.Values]
func (t *TableView) SetValues(v []giv.Value) *TableView {
	t.Values = v
	return t
}

// SetSelVal sets the [TableView.SelVal]
func (t *TableView) SetSelVal(v any) *TableView {
	t.SelVal = v
	return t
}

// SetSelIdx sets the [TableView.SelIdx]
func (t *TableView) SetSelIdx(v int) *TableView {
	t.SelIdx = v
	return t
}

// SetSelIdxs sets the [TableView.SelIdxs]
func (t *TableView) SetSelIdxs(v map[int]struct{}) *TableView {
	t.SelIdxs = v
	return t
}

// SetInitSelIdx sets the [TableView.InitSelIdx]
func (t *TableView) SetInitSelIdx(v int) *TableView {
	t.InitSelIdx = v
	return t
}

// SetDraggedIdxs sets the [TableView.DraggedIdxs]
func (t *TableView) SetDraggedIdxs(v []int) *TableView {
	t.DraggedIdxs = v
	return t
}

// SetViewPath sets the [TableView.ViewPath]
func (t *TableView) SetViewPath(v string) *TableView {
	t.ViewPath = v
	return t
}

// SetTmpSave sets the [TableView.TmpSave]
func (t *TableView) SetTmpSave(v giv.Value) *TableView {
	t.TmpSave = v
	return t
}

// SetVisRows sets the [TableView.VisRows]
func (t *TableView) SetVisRows(v int) *TableView {
	t.VisRows = v
	return t
}

// SetStartIdx sets the [TableView.StartIdx]
func (t *TableView) SetStartIdx(v int) *TableView {
	t.StartIdx = v
	return t
}

// SetSliceSize sets the [TableView.SliceSize]
func (t *TableView) SetSliceSize(v int) *TableView {
	t.SliceSize = v
	return t
}

// SetConfigIter sets the [TableView.ConfigIter]
func (t *TableView) SetConfigIter(v int) *TableView {
	t.ConfigIter = v
	return t
}

// SetTmpIdx sets the [TableView.TmpIdx]
func (t *TableView) SetTmpIdx(v int) *TableView {
	t.TmpIdx = v
	return t
}

// SetElVal sets the [TableView.ElVal]
func (t *TableView) SetElVal(v reflect.Value) *TableView {
	t.ElVal = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/etable/v2/etview.TensorLayout",
	ShortName: "etview.TensorLayout",
	IDName:    "tensor-layout",
	Doc:       "TensorLayout are layout options for displaying tensors",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"OddRow", &gti.Field{Name: "OddRow", Type: "bool", LocalType: "bool", Doc: "even-numbered dimensions are displayed as Y*X rectangles -- this determines along which dimension to display any remaining odd dimension: OddRow = true = organize vertically along row dimension, false = organize horizontally across column dimension", Directives: gti.Directives{}, Tag: ""}},
		{"TopZero", &gti.Field{Name: "TopZero", Type: "bool", LocalType: "bool", Doc: "if true, then the Y=0 coordinate is displayed from the top-down; otherwise the Y=0 coordinate is displayed from the bottom up, which is typical for emergent network patterns.", Directives: gti.Directives{}, Tag: ""}},
		{"Image", &gti.Field{Name: "Image", Type: "bool", LocalType: "bool", Doc: "display the data as a bitmap image.  if a 2D tensor, then it will be a greyscale image.  if a 3D tensor with size of either the first or last dim = either 3 or 4, then it is a RGB(A) color image", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/etable/v2/etview.TensorDisp",
	ShortName: "etview.TensorDisp",
	IDName:    "tensor-disp",
	Doc:       "TensorDisp are options for displaying tensors",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Range", &gti.Field{Name: "Range", Type: "goki.dev/etable/v2/minmax.Range64", LocalType: "minmax.Range64", Doc: "range to plot", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"MinMax", &gti.Field{Name: "MinMax", Type: "goki.dev/etable/v2/minmax.F64", LocalType: "minmax.F64", Doc: "if not using fixed range, this is the actual range of data", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"ColorMap", &gti.Field{Name: "ColorMap", Type: "goki.dev/gi/v2/giv.ColorMapName", LocalType: "giv.ColorMapName", Doc: "the name of the color map to use in translating values to colors", Directives: gti.Directives{}, Tag: ""}},
		{"Background", &gti.Field{Name: "Background", Type: "image/color.Color", LocalType: "color.Color", Doc: "background color", Directives: gti.Directives{}, Tag: ""}},
		{"GridFill", &gti.Field{Name: "GridFill", Type: "float32", LocalType: "float32", Doc: "what proportion of grid square should be filled by color block -- 1 = all, .5 = half, etc", Directives: gti.Directives{}, Tag: "min:\"0.1\" max:\"1\" step:\"0.1\" def:\"0.9,1\""}},
		{"DimExtra", &gti.Field{Name: "DimExtra", Type: "float32", LocalType: "float32", Doc: "amount of extra space to add at dimension boundaries, as a proportion of total grid size", Directives: gti.Directives{}, Tag: "min:\"0\" max:\"1\" step:\"0.02\" def:\"0.1,0.3\""}},
		{"BotRtSpace", &gti.Field{Name: "BotRtSpace", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "extra space to add at bottom of grid -- needed when included in TableView for example", Directives: gti.Directives{}, Tag: ""}},
		{"GridMinSize", &gti.Field{Name: "GridMinSize", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "minimum size for grid squares -- they will never be smaller than this", Directives: gti.Directives{}, Tag: ""}},
		{"GridMaxSize", &gti.Field{Name: "GridMaxSize", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "maximum size for grid squares -- they will never be larger than this", Directives: gti.Directives{}, Tag: ""}},
		{"TotPrefSize", &gti.Field{Name: "TotPrefSize", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "total preferred display size along largest dimension -- grid squares will be sized to fit within this size, subject to harder GridMin / Max size constraints", Directives: gti.Directives{}, Tag: ""}},
		{"FontSize", &gti.Field{Name: "FontSize", Type: "float32", LocalType: "float32", Doc: "font size in standard point units for labels (e.g., SimMat)", Directives: gti.Directives{}, Tag: ""}},
		{"GridView", &gti.Field{Name: "GridView", Type: "*goki.dev/etable/v2/etview.TensorGrid", LocalType: "*TensorGrid", Doc: "our gridview, for update method", Directives: gti.Directives{}, Tag: "copy:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"TensorLayout", &gti.Field{Name: "TensorLayout", Type: "goki.dev/etable/v2/etview.TensorLayout", LocalType: "TensorLayout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

// TensorGridType is the [gti.Type] for [TensorGrid]
var TensorGridType = gti.AddType(&gti.Type{
	Name:       "goki.dev/etable/v2/etview.TensorGrid",
	ShortName:  "etview.TensorGrid",
	IDName:     "tensor-grid",
	Doc:        "TensorGrid is a widget that displays tensor values as a grid of colored squares.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Tensor", &gti.Field{Name: "Tensor", Type: "goki.dev/etable/v2/etensor.Tensor", LocalType: "etensor.Tensor", Doc: "the tensor that we view", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Disp", &gti.Field{Name: "Disp", Type: "goki.dev/etable/v2/etview.TensorDisp", LocalType: "TensorDisp", Doc: "display options", Directives: gti.Directives{}, Tag: ""}},
		{"ColorMap", &gti.Field{Name: "ColorMap", Type: "*goki.dev/colors/colormap.Map", LocalType: "*colormap.Map", Doc: "the actual colormap", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"WidgetBase", &gti.Field{Name: "WidgetBase", Type: "goki.dev/gi/v2/gi.WidgetBase", LocalType: "gi.WidgetBase", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TensorGrid{},
})

// NewTensorGrid adds a new [TensorGrid] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTensorGrid(par ki.Ki, name ...string) *TensorGrid {
	return par.NewChild(TensorGridType, name...).(*TensorGrid)
}

// KiType returns the [*gti.Type] of [TensorGrid]
func (t *TensorGrid) KiType() *gti.Type {
	return TensorGridType
}

// New returns a new [*TensorGrid] value
func (t *TensorGrid) New() ki.Ki {
	return &TensorGrid{}
}

// SetDisp sets the [TensorGrid.Disp]:
// display options
func (t *TensorGrid) SetDisp(v TensorDisp) *TensorGrid {
	t.Disp = v
	return t
}

// SetColorMap sets the [TensorGrid.ColorMap]:
// the actual colormap
func (t *TensorGrid) SetColorMap(v *colormap.Map) *TensorGrid {
	t.ColorMap = v
	return t
}

// SetTooltip sets the [TensorGrid.Tooltip]
func (t *TensorGrid) SetTooltip(v string) *TensorGrid {
	t.Tooltip = v
	return t
}

// SetClass sets the [TensorGrid.Class]
func (t *TensorGrid) SetClass(v string) *TensorGrid {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [TensorGrid.PriorityEvents]
func (t *TensorGrid) SetPriorityEvents(v []events.Types) *TensorGrid {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [TensorGrid.CustomContextMenu]
func (t *TensorGrid) SetCustomContextMenu(v func(m *gi.Scene)) *TensorGrid {
	t.CustomContextMenu = v
	return t
}

// TensorViewType is the [gti.Type] for [TensorView]
var TensorViewType = gti.AddType(&gti.Type{
	Name:       "goki.dev/etable/v2/etview.TensorView",
	ShortName:  "etview.TensorView",
	IDName:     "tensor-view",
	Doc:        "etview.TensorView provides a GUI interface for etable.Tensor's\nusing a tabular rows-and-columns interface using textfields for editing.\nThis provides an editable complement to the TensorGrid graphical display.",
	Directives: gti.Directives{},
	Fields:     ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"WidgetBase", &gti.Field{Name: "WidgetBase", Type: "goki.dev/gi/v2/gi.WidgetBase", LocalType: "gi.WidgetBase", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &TensorView{},
})

// NewTensorView adds a new [TensorView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewTensorView(par ki.Ki, name ...string) *TensorView {
	return par.NewChild(TensorViewType, name...).(*TensorView)
}

// KiType returns the [*gti.Type] of [TensorView]
func (t *TensorView) KiType() *gti.Type {
	return TensorViewType
}

// New returns a new [*TensorView] value
func (t *TensorView) New() ki.Ki {
	return &TensorView{}
}

// SetTooltip sets the [TensorView.Tooltip]
func (t *TensorView) SetTooltip(v string) *TensorView {
	t.Tooltip = v
	return t
}

// SetClass sets the [TensorView.Class]
func (t *TensorView) SetClass(v string) *TensorView {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [TensorView.PriorityEvents]
func (t *TensorView) SetPriorityEvents(v []events.Types) *TensorView {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [TensorView.CustomContextMenu]
func (t *TensorView) SetCustomContextMenu(v func(m *gi.Scene)) *TensorView {
	t.CustomContextMenu = v
	return t
}
