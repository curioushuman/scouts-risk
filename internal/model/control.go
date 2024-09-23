package model

// Poor mans string enum
// ! NOTE: it does not prevent c.When = ControlWhen("garbage")
type ControlWhen string
const (
	ControlWhenBefore   ControlWhen = "before"
	ControlWhenDuring   ControlWhen = "during"
	ControlWhenRegular  ControlWhen = "regular"
	ControlWhenAfter    ControlWhen = "after"
	ControlWhenResponse ControlWhen = "response"
)

type Control struct {
	Label string
	Id		string
	When	ControlWhen
	// Tags  []Tag
}
