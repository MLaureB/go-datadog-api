package datadog

import (
	"encoding/json"
	"fmt"
)

const (
	ALERT_GRAPH_WIDGET  = "alert_graph"
	ALERT_VALUE_WIDGET  = "alert_value"
	CHANGE_WIDGET       = "change"
	CHECK_STATUS_WIDGET = "check_status"
	DISTRIBUTION_WIDGET = "distribution"
	GROUP_WIDGET        = "group"
	NOTE_WIDGET         = "note"
	TIMESERIES_WIDGET   = "timeseries"
)

// BoardWidget represents the structure of any widget. However, the widget Definition structure is
// different according to widget type.
type BoardWidget struct {
	Definition interface{}   `json:"definition"`
	Id         *int          `json:"id,omitempty"`
	Layout     *WidgetLayout `json:"layout,omitempty"`
}

// WidgetLayout represents the layout for a widget on a "free" dashboard
type WidgetLayout struct {
	X      *float64 `json:"x,omitempty"`
	Y      *float64 `json:"y,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Width  *float64 `json:"width,omitempty"`
}

func (widget *BoardWidget) GetWidgetType() (string, error) {
	switch widget.Definition.(type) {
	case AlertGraphDefinition:
		return ALERT_GRAPH_WIDGET, nil
	case AlertValueDefinition:
		return ALERT_VALUE_WIDGET, nil
	case ChangeDefinition:
		return CHANGE_WIDGET, nil
	case CheckStatusDefinition:
		return CHECK_STATUS_WIDGET, nil
	case DistributionDefinition:
		return DISTRIBUTION_WIDGET, nil
	case GroupDefinition:
		return GROUP_WIDGET, nil
	case NoteDefinition:
		return NOTE_WIDGET, nil
	case TimeseriesDefinition:
		return TIMESERIES_WIDGET, nil
	default:
		return "", fmt.Errorf("Unsupported widget type")
	}
}

// AlertGraphDefinition represents the definition for an Alert Graph widget
type AlertGraphDefinition struct {
	Type       *string     `json:"type"`
	AlertId    *string     `json:"alert_id"`
	VizType    *string     `json:"viz_type"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// AlertValueDefinition represents the definition for an Alert Value widget
type AlertValueDefinition struct {
	Type       *string `json:"type"`
	AlertId    *string `json:"alert_id"`
	Precision  *int    `json:"precision,omitempty"`
	Unit       *string `json:"unit,omitempty"`
	TextAlign  *string `json:"text_align,omitempty"`
	Title      *string `json:"title,omitempty"`
	TitleSize  *string `json:"title_size,omitempty"`
	TitleAlign *string `json:"title_align,omitempty"`
}

// ChangeDefinition represents the definition for a Change widget
type ChangeDefinition struct {
	Type       *string         `json:"type"`
	Requests   []ChangeRequest `json:"requests"`
	Title      *string         `json:"title,omitempty"`
	TitleSize  *string         `json:"title_size,omitempty"`
	TitleAlign *string         `json:"title_align,omitempty"`
	Time       *WidgetTime     `json:"time,omitempty"`
}
type ChangeRequest struct {
	WidgetRequest
	ChangeType   *string `json:"change_type,omitempty"`
	CompareTo    *string `json:"compare_to,omitempty"`
	IncreaseGood *bool   `json:"increase_good,omitempty"`
	OrderBy      *string `json:"order_by,omitempty"`
	OrderDir     *string `json:"order_dir,omitempty"`
	ShowPresent  *bool   `json:"show_present,omitempty"`
}

// CheckStatusDefinition represents the definition for a Check Status widget
type CheckStatusDefinition struct {
	Type       *string     `json:"type"`
	Check      *string     `json:"check"`
	Grouping   *string     `json:"grouping"`
	Group      *string     `json:"group,omitempty"`
	GroupBy    []string    `json:"group_by,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// DistributionDefinition represents the definition for a Distribution widget
type DistributionDefinition struct {
	Type       *string               `json:"type"`
	Requests   []DistributionRequest `json:"requests"`
	Title      *string               `json:"title,omitempty"`
	TitleSize  *string               `json:"title_size,omitempty"`
	TitleAlign *string               `json:"title_align,omitempty"`
	Time       *WidgetTime           `json:"time,omitempty"`
}
type DistributionRequest struct {
	WidgetRequest
	Style *WidgetRequestStyle `json:"style,omitempty"`
}

// GroupDefinition represents the definition for an Group widget
type GroupDefinition struct {
	Type       *string       `json:"type"`
	LayoutType *string       `json:"layout_type"`
	Widgets    []BoardWidget `json:"widgets"`
	Title      *string       `json:"title,omitempty"`
}

// NoteDefinition represents the definition for a Note widget
type NoteDefinition struct {
	Type            *string `json:"type"`
	Content         *string `json:"content"`
	BackgroundColor *string `json:"background_color,omitempty"`
	FontSize        *string `json:"font_size,omitempty"`
	TextAlign       *string `json:"text_align,omitempty"`
	ShowTick        *bool   `json:"show_tick,omitempty"`
	TickPos         *string `json:"tick_pos,omitempty"`
	TickEdge        *string `json:"tick_edge,omitempty"`
}

// TimeseriesDefinition represents the definition for a Timeseries widget
type TimeseriesDefinition struct {
	Type       *string             `json:"type"`
	Requests   []TimeseriesRequest `json:"requests"`
	Yaxis      *WidgetAxis         `json:"yaxis,omitempty"`
	Events     []WidgetEvent       `json:"events,omitempty"`
	Markers    []WidgetMarker      `json:"markers,omitempty"`
	Title      *string             `json:"title,omitempty"`
	TitleSize  *string             `json:"title_size,omitempty"`
	TitleAlign *string             `json:"title_align,omitempty"`
	ShowLegend *bool               `json:"show_legend,omitempty"`
	LegendSize *string             `json:"legend_size,omitempty"`
	Time       *WidgetTime         `json:"time,omitempty"`
}
type TimeseriesRequest struct {
	WidgetRequest
	Style       *TimeseriesRequestStyle `json:"style,omitempty"`
	Metadata    []WidgetMetadata        `json:"metadata,omitempty"`
	DisplayType *string                 `json:"display_type,omitempty"`
}
type TimeseriesRequestStyle struct {
	WidgetRequestStyle
	LineType  *string `json:"line_type,omitempty"`
	LineWidth *string `json:"line_width,omitempty"`
}

// UnmarshalJSON is a Custom Unmarshal for BoardWidget. If first tries to unmarshal the data in a light
// struct that allows to get the widget type. Then based on the widget type, it will try to unmarshal the
// data using the corresponding widget struct.
func (widget *BoardWidget) UnmarshalJSON(data []byte) error {
	var widgetHandler struct {
		Definition *struct {
			Type *string `json:"type"`
		} `json:"definition"`
		Id     *int          `json:"id,omitempty"`
		Layout *WidgetLayout `json:"layout,omitempty"`
	}
	if err := json.Unmarshal(data, &widgetHandler); err != nil {
		return err
	}

	// Get the widget id
	widget.Id = widgetHandler.Id

	// Get the widget layout
	widget.Layout = widgetHandler.Layout

	// Get the widget definition based on the widget type
	switch *widgetHandler.Definition.Type {
	case ALERT_GRAPH_WIDGET:
		var alertGraphWidget struct {
			Definition AlertGraphDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &alertGraphWidget); err != nil {
			return err
		}
		widget.Definition = alertGraphWidget.Definition
	case ALERT_VALUE_WIDGET:
		var alertValueWidget struct {
			Definition AlertValueDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &alertValueWidget); err != nil {
			return err
		}
		widget.Definition = alertValueWidget.Definition
	case CHANGE_WIDGET:
		var changeWidget struct {
			Definition ChangeDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &changeWidget); err != nil {
			return err
		}
		widget.Definition = changeWidget.Definition
	case CHECK_STATUS_WIDGET:
		var checkStatusWidget struct {
			Definition CheckStatusDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &checkStatusWidget); err != nil {
			return err
		}
		widget.Definition = checkStatusWidget.Definition
	case DISTRIBUTION_WIDGET:
		var distributionWidget struct {
			Definition DistributionDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &distributionWidget); err != nil {
			return err
		}
		widget.Definition = distributionWidget.Definition
	case GROUP_WIDGET:
		var groupWidget struct {
			Definition GroupDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &groupWidget); err != nil {
			return err
		}
		widget.Definition = groupWidget.Definition
	case NOTE_WIDGET:
		var noteWidget struct {
			Definition NoteDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &noteWidget); err != nil {
			return err
		}
		widget.Definition = noteWidget.Definition
	case TIMESERIES_WIDGET:
		var timeseriesWidget struct {
			Definition TimeseriesDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &timeseriesWidget); err != nil {
			return err
		}
		widget.Definition = timeseriesWidget.Definition
	default:
		return fmt.Errorf("Cannot unmarshal widget of type: %s", *widgetHandler.Definition.Type)
	}

	return nil
}

//
// List of structs common to multiple widget definitions
//

type WidgetTime struct {
	LiveSpan *string `json:"live_span,omitempty"`
}

type WidgetAxis struct {
	Label       *string `json:"label,omitempty"`
	Scale       *string `json:"scale,omitempty"`
	Min         *string `json:"min,omitempty"`
	Max         *string `json:"max,omitempty"`
	IncludeZero *bool   `json:"include_zero,omitempty"`
}

type WidgetEvent struct {
	Query *string `json:"q"`
}

type WidgetMarker struct {
	Value       *string `json:"value"`
	DisplayType *string `json:"display_type,omitempty"`
	Label       *string `json:"label,omitempty"`
}

type WidgetMetadata struct {
	Expression *string `json:"expression"`
	AliasName  *string `json:"alias_name,omitempty"`
}

// WidgetRequest represents a request to display on the widget.
// One request object should implement only one type of query.
type WidgetRequest struct {
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}
type WidgetApmOrLogQuery struct {
	Index   *string `json:"index"`
	Compute *struct {
		Aggregation *string `json:"aggregation"`
		Facet       *string `json:"facet,omitempty"`
		Interval    *int    `json:"interval,omitempty"`
	} `json:"compute"`
	Search *struct {
		Query *string `json:"query"`
	} `json:"search,omitempty"`
	GroupBy []struct {
		Facet *string `json:"facet"`
		Limit *int    `json:"limit,omitempty"`
		Sort  *struct {
			Aggregation *string `json:"aggregation"`
			Order       *string `json:"order"`
			Facet       *string `json:"facet,omitempty"`
		} `json:"sort,omitempty"`
	} `json:"group_by,omitempty"`
}
type WidgetProcessQuery struct {
	Metric   *string  `json:"metric"`
	SearchBy *string  `json:"search_by,omitempty"`
	FilterBy []string `json:"filter_by,omitempty"`
	Limit    *int     `json:"limit,omitempty"`
}

// WidgetRequestStyle represents the style that can be apply to a request
type WidgetRequestStyle struct {
	Palette *string `json:"palette,omitempty"`
}
