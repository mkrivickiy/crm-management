package rest

type ConfigInterface interface {
	GetURL() string
	GetContentType() string
}

type Config struct {
	URL         string
	ContentType string
}

func (c *Config) GetURL() string {
	return c.URL
}

func (c *Config) GetContentType() string {
	return c.ContentType
}
