package commands

// CGCommand is the base struct for all CG template commands. It contains the common fields for all CG commands.
type CGCommand struct {
	VideoChannel int
	Layer        *int // defaults to 9999
	CgLayer      *int // optional, only used for layer-specific commands
}

// TemplateCGAdd Prepares a template for displaying.
// It won't show until you call CG PLAY (unless you supply the play-on-load flag, 1 for true).
// Data is either inline XML or a reference to a saved dataset.
type TemplateCGAdd struct {
	CGCommand

	Template string

	PlayOnLoad bool
	Data       *string // JSON or XML inline string
}

func (c TemplateCGAdd) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("ADD"))
	cmd = appendInt(cmd, c.CgLayer)
	cmd = appendQuotedString(cmd, ptr(c.Template))
	cmd = appendBool(cmd, &c.PlayOnLoad)
	return appendQuotedString(cmd, c.Data)
}

// TemplateCGPlay plays and displays the template in the specified layer.
type TemplateCGPlay struct {
	CGCommand
}

func (c TemplateCGPlay) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("PLAY"))
	return appendInt(cmd, c.CgLayer)
}

// TemplateCGStop stops the template in the specified layer.
// This is different from CG REMOVE in that the template gets a chance to animate out when it is stopped.
type TemplateCGStop struct {
	CGCommand
}

func (c TemplateCGStop) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("STOP"))
	return appendInt(cmd, c.CgLayer)
}

// TemplateCGNext triggers a "continue" in the template on the specified layer.
// This is used to control animations that has multiple discreet steps.
type TemplateCGNext struct {
	CGCommand
}

func (c TemplateCGNext) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("NEXT"))
	return appendInt(cmd, c.CgLayer)
}

// TemplateCGRemove removes the template from the specified layer.
type TemplateCGRemove struct {
	CGCommand
}

func (c TemplateCGRemove) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("REMOVE"))
	return appendInt(cmd, c.CgLayer)
}

// TemplateCGClear removes all templates on a video layer. The entire cg producer will be removed.
type TemplateCGClear struct {
	CGCommand
}

func (c TemplateCGClear) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("CLEAR"))
	return cmd
}

// TemplateCGUpdate sends new data to the template on specified layer.
// Data is either inline XML or a reference to a saved dataset.
type TemplateCGUpdate struct {
	CGCommand

	Data string // JSON or XML inline string
}

func (c TemplateCGUpdate) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("UPDATE"))
	cmd = appendInt(cmd, c.CgLayer)
	return appendQuotedString(cmd, ptr(c.Data))
}

// TemplateCGInvoke invokes the given method on the template on the specified layer.
type TemplateCGInvoke struct {
	CGCommand

	Method string
}

func (c TemplateCGInvoke) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("INVOKE"))
	cmd = appendInt(cmd, c.CgLayer)
	return appendQuotedString(cmd, ptr(c.Method))
}

// TemplateCGInfo retrieves information about the template on the specified layer.
// If `cg_layer` is not given, information about the template host is given instead.
type TemplateCGInfo struct {
	CGCommand
}

func (c TemplateCGInfo) String() string {
	cmd := baseLayerCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("INFO"))
	return appendInt(cmd, c.CgLayer)
}
