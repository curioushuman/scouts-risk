package model

import (
	"github.com/a-h/templ"
)

// NOTE: we'll create another for theme colours
// i.e. base, primary, slate, etc
// Poor mans string enum
// ! NOTE: it does not prevent a.Icon = UiColour("garbage")
type UiColour string
const (
	UiColourDefault 	UiColour = "arrow"
	UiColourError   	UiColour = "error"
	UiColourNeutral 	UiColour = "neutral"
	UiColourSuccess 	UiColour = "success"
	UiColourWarning 	UiColour = "warning"
	UiColourPrimary 	UiColour = "primary"
	UiColourSecondary UiColour = "secondary"
	UiColourAccent    UiColour = "accent"
	UiColourInfo    	UiColour = "info"
)

type AccordionIcon string
const (
	AccordionIconArrow AccordionIcon = "arrow"
	AccordionIconPlus  AccordionIcon = "plus"
)

type Accordion struct {
	Label       any // string or templ.Component
	Id		  		string
	GroupId		  string
	Description	string
	TitleBadges []templ.Component
	Icon        AccordionIcon
	Multi			  bool
	Border			bool
	ClassBg		  string
	ClassBgOpen	string
}

type ActionListAction string
const (
	ActionListActionAdd    		ActionListAction = "add"
	ActionListActionEdit   		ActionListAction = "edit"
	ActionListActionDuplicate ActionListAction = "duplicate"
	ActionListActionDelete 		ActionListAction = "delete"
)

type ActionListLink struct {
	LinkTitle string
	LinkHref string
}

type ActionList struct {
	Id 			string
	Actions map[ActionListAction]string
	Links 	[]ActionListLink
}

type Anchor struct {
	Href  string
	Label string
	Icon  templ.Component
	Attrs templ.Attributes
}

type ButtonFloatPlace string
const (
	ButtonFloatPlaceTop    		ButtonFloatPlace = "top"
	ButtonFloatPlaceBottom   		ButtonFloatPlace = "bottom"
	ButtonFloatPlaceLeft ButtonFloatPlace = "left"
	ButtonFloatPlaceRight 		ButtonFloatPlace = "right"
)

type ButtonFloating struct {
	Colour UiColour
	Attrs  templ.Attributes
	Y 		 ButtonFloatPlace // top, bottom; default bottom
	X			 ButtonFloatPlace // left, right; default right
}

type Card struct {
	Title   string
	Content string
	Source  string
	Alt     string
}

type Checkbox struct {
	Label   string
	Name    string
	Checked bool
	Class   string
	Attrs   templ.Attributes
}

type CompanyInfo struct {
	Icon        templ.Component
	Name        string
	Description string
	Copyright   string
}

type DropdownItem struct {
	Label string
	Attrs templ.Attributes
}

type Feature struct {
	Icon        templ.Component
	Title       string
	Description string
	URL         string
}

type HeadingHelp struct {
	Label string
	Level int
	Help  any
}

type Image struct {
	Source string
	Alt    string
}

type Input struct {
	Label   string
	Name    string
	Value   string
	Err     string
	Small   bool
	Attrs   templ.Attributes
	Classes string
	Icon    templ.Component
}

type PaginationItem struct {
	URL      string
	Page     int
	Low      int
	High     int
	MaxPages int
}

type Price struct {
	Title            string
	Description      string
	Price            string
	Per              string
	IncludedFeatures []string
	ExcludedFeatures []string
	Footer           templ.Component
}

type Range struct {
	Label string
	Name  string
	Value int
	Min   int
	Max   int
	Step  int
	Class string
}

type Rating struct {
	Name  string
	Min   int
	Max   int
	Class string
}

type Script struct {
	Source string
	Defer  bool
}

type Select struct {
	Label   string
	Name    string
	Options []SelectOption
	Attrs   templ.Attributes
}

type SelectOption struct {
	Label    string
	Value    string
	Selected bool
	Disabled bool
}

type Stat struct {
	Title       string
	Value       string
	Description string
}

type Textarea struct {
	Label string
	Name  string
	Value string
	Rows  int
	Err   string
	Class string
	Attrs templ.Attributes
}

type TimelineItem struct {
	Start  string
	Middle templ.Component
	End    string
}

type Toggle struct {
	Before    string
	After     string
	Name      string
	Checked   bool
	Class     string
	Highlight bool
	Attrs     templ.Attributes
}
