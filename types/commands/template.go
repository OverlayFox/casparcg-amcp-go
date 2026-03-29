package commands

type CGCommand struct {
	VideoChannel int
	Layer        *int // defaults to 9999
	CgLayer      *int // optional, only used for layer-specific commands
}

type TemplateCGAdd struct {
	CGCommand

	Template string

	PlayOnLoad bool
	Data       *string // JSON or XML inline string
}

func (c TemplateCGAdd) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("ADD"))
	cmd = appendInt(cmd, c.CgLayer)
	cmd = appendQuotedString(cmd, ptr(c.Template))
	cmd = appendBool(cmd, &c.PlayOnLoad)
	return appendQuotedString(cmd, c.Data)
}

type TemplateCGPlay struct {
	CGCommand
}

func (c TemplateCGPlay) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("PLAY"))
	return appendInt(cmd, c.CgLayer)
}

type TemplateCGStop struct {
	CGCommand
}

func (c TemplateCGStop) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("STOP"))
	return appendInt(cmd, c.CgLayer)
}

type TemplateCGNext struct {
	CGCommand
}

func (c TemplateCGNext) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("NEXT"))
	return appendInt(cmd, c.CgLayer)
}

type TemplateCGRemove struct {
	CGCommand
}

func (c TemplateCGRemove) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("REMOVE"))
	return appendInt(cmd, c.CgLayer)
}

type TemplateCGClear struct {
	CGCommand
}

func (c TemplateCGClear) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("CLEAR"))
	return cmd
}

type TemplateCGUpdate struct {
	CGCommand

	Data string // JSON or XML inline string
}

func (c TemplateCGUpdate) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("UPDATE"))
	cmd = appendInt(cmd, c.CgLayer)
	return appendQuotedString(cmd, ptr(c.Data))
}

type TemplateCGInvoke struct {
	CGCommand

	Method string
}

func (c TemplateCGInvoke) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("INVOKE"))
	cmd = appendInt(cmd, c.CgLayer)
	return appendQuotedString(cmd, ptr(c.Method))
}

type TemplateCGInfo struct {
	CGCommand
}

func (c TemplateCGInfo) String() string {
	cmd := baseCommand("CG", c.VideoChannel, c.Layer)
	cmd = appendString(cmd, ptr("INFO"))
	return appendInt(cmd, c.CgLayer)
}
