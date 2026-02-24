package types

type TemplateCommandInterface interface {
}

type TemplateCommandCG struct {
	videoChannel int
	layer        int // defaults to 9999
}

// TemplateCommandCGAdd Prepares a template for displaying.
// It won't show until you call CG PLAY (unless you supply the play-on-load flag, 1 for true).
// Data is either inline XML or a reference to a saved dataset.
type TemplateCommandCGAdd struct {
	TemplateCommandCG

	cgLayer  int
	template string

	playOnLoad bool
	data       *string // JSON or XML inline string
}

// TemplateCommandCGPlay plays and displays the template in the specified layer.
type TemplateCommandCGPlay struct {
	TemplateCommandCG

	cgLayer int
}

// TemplateCommandCGStop stops the template in the specified layer.
// This is different from CG REMOVE in that the template gets a chance to animate out when it is stopped.
type TemplateCommandCGStop struct {
	TemplateCommandCG

	cgLayer int
}

// TemplateCommandCGNext triggers a "continue" in the template on the specified layer.
// This is used to control animations that has multiple discreet steps.
type TemplateCommandCGNext struct {
	TemplateCommandCG

	cgLayer int
}

// TemplateCommandCGRemove removes the template from the specified layer.
type TemplateCommandCGRemove struct {
	TemplateCommandCG

	cgLayer int
}

// TemplateCommandCGClear removes all templates on a video layer. The entire cg producer will be removed.
type TemplateCommandCGClear struct {
	TemplateCommandCG
}

// TemplateCommandCGUpdate sends new data to the template on specified layer.
// Data is either inline XML or a reference to a saved dataset.
type TemplateCommandCGUpdate struct {
	TemplateCommandCG

	cgLayer int
	data    string // JSON or XML inline string
}

// TemplateCommandCGInvoke invokes the given method on the template on the specified layer.
type TemplateCommandCGInvoke struct {
	TemplateCommandCG

	cgLayer int
	method  string
}

// Retrieves information about the template on the specified layer.
// If `cg_layer` is not given, information about the template host is given instead.
type TemplateCommandCGInfo struct {
	TemplateCommandCG

	cgLayer *int
}
