package ginbe

type Controller struct {
	version   string
	path      string
	resources []resource
}

func NewController(version string, path string) *Controller {
	return &Controller{
		version: version,
		path:    path,
	}
}

func (c *Controller) AttachResource(a *resource) *Controller {
	c.resources = append(c.resources, *a)
	return c
}
