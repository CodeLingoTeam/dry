package ui

import (
	"bytes"
	"encoding/json"

	godocker "github.com/fsouza/go-dockerclient"
)

type inspectRenderer struct {
	container *godocker.Container
}

//NewDockerInspectRenderer creates renderer for inspect information
func NewDockerInspectRenderer(container *godocker.Container) Renderer {
	return &inspectRenderer{
		container: container,
	}
}

func (r *inspectRenderer) Render() string {
	c, _ := json.Marshal(r.container)
	buf := new(bytes.Buffer)
	buf.WriteString("[\n")
	if err := json.Indent(buf, c, "", "    "); err == nil {
		if buf.Len() > 1 {
			// Remove trailing ','
			buf.Truncate(buf.Len() - 1)
		}
	} else {
		buf.WriteString("There was an error inspecting container information")
	}
	buf.WriteString("]\n")
	return buf.String()
}
