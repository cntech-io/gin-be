package ginbe

type controller struct {
	version string
	path    string
	apis    []api
}

func NewController(version string, path string) *controller {
	return &controller{
		version: version,
		path:    path,
	}
}

func (c *controller) AttachAPI(a *api) *controller {
	c.apis = append(c.apis, *a)
	return c
}
