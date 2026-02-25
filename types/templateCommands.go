package types

import (
	"fmt"
	"strconv"
)

type TemplateCommandInterface interface {
	String() string
}

type TemplateCommandCG struct {
	VideoChannel int
	Layer        int // defaults to 9999
}

// TemplateCommandCGAdd Prepares a template for displaying.
// It won't show until you call CG PLAY (unless you supply the play-on-load flag, 1 for true).
// Data is either inline XML or a reference to a saved dataset.
type TemplateCommandCGAdd struct {
	TemplateCommandCG

	CgLayer  int
	Template string

	PlayOnLoad bool
	Data       *string // JSON or XML inline string
}

func (c TemplateCommandCGAdd) String() string {
	cmd := fmt.Sprintf("CG %d-%d ADD %d %s", c.VideoChannel, c.Layer, c.CgLayer, quote(c.Template))
	if c.PlayOnLoad {
		cmd += " 1"
	} else {
		cmd += " 0"
	}
	if c.Data != nil {
		cmd += " " + quote(*c.Data)
	}
	return cmd
}

// TemplateCommandCGPlay plays and displays the template in the specified layer.
type TemplateCommandCGPlay struct {
	TemplateCommandCG

	CgLayer int
}

func (c TemplateCommandCGPlay) String() string {
	return fmt.Sprintf("CG %d-%d PLAY %d", c.VideoChannel, c.Layer, c.CgLayer)
}

// TemplateCommandCGStop stops the template in the specified layer.
// This is different from CG REMOVE in that the template gets a chance to animate out when it is stopped.
type TemplateCommandCGStop struct {
	TemplateCommandCG

	CgLayer int
}

func (c TemplateCommandCGStop) String() string {
	return fmt.Sprintf("CG %d-%d STOP %d", c.VideoChannel, c.Layer, c.CgLayer)
}

// TemplateCommandCGNext triggers a "continue" in the template on the specified layer.
// This is used to control animations that has multiple discreet steps.
type TemplateCommandCGNext struct {
	TemplateCommandCG

	CgLayer int
}

func (c TemplateCommandCGNext) String() string {
	return fmt.Sprintf("CG %d-%d NEXT %d", c.VideoChannel, c.Layer, c.CgLayer)
}

// TemplateCommandCGRemove removes the template from the specified layer.
type TemplateCommandCGRemove struct {
	TemplateCommandCG

	CgLayer int
}

func (c TemplateCommandCGRemove) String() string {
	return fmt.Sprintf("CG %d-%d REMOVE %d", c.VideoChannel, c.Layer, c.CgLayer)
}

// TemplateCommandCGClear removes all templates on a video layer. The entire cg producer will be removed.
type TemplateCommandCGClear struct {
	TemplateCommandCG
}

func (c TemplateCommandCGClear) String() string {
	return fmt.Sprintf("CG %d-%d CLEAR", c.VideoChannel, c.Layer)
}

// TemplateCommandCGUpdate sends new data to the template on specified layer.
// Data is either inline XML or a reference to a saved dataset.
type TemplateCommandCGUpdate struct {
	TemplateCommandCG

	CgLayer int
	Data    string // JSON or XML inline string
}

func (c TemplateCommandCGUpdate) String() string {
	return fmt.Sprintf("CG %d-%d UPDATE %d %s", c.VideoChannel, c.Layer, c.CgLayer, quote(c.Data))
}

// TemplateCommandCGInvoke invokes the given method on the template on the specified layer.
type TemplateCommandCGInvoke struct {
	TemplateCommandCG

	CgLayer int
	Method  string
}

func (c TemplateCommandCGInvoke) String() string {
	return fmt.Sprintf("CG %d-%d INVOKE %d %s", c.VideoChannel, c.Layer, c.CgLayer, quote(c.Method))
}

// Retrieves information about the template on the specified layer.
// If `cg_layer` is not given, information about the template host is given instead.
type TemplateCommandCGInfo struct {
	TemplateCommandCG

	CgLayer *int
}

func (c TemplateCommandCGInfo) String() string {
	cmd := fmt.Sprintf("CG %d-%d INFO", c.VideoChannel, c.Layer)
	if c.CgLayer != nil {
		cmd += " " + strconv.Itoa(*c.CgLayer)
	}
	return cmd
}
