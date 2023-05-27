package ginbe

type Controller struct {
	version string
	path    string
	apis    []resource
}

func NewController(version string, path string) *Controller {
	return &Controller{
		version: version,
		path:    path,
	}
}

func (c *Controller) AttachAPI(a *resource) *Controller {
	c.apis = append(c.apis, *a)
	return c
}
