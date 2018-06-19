package rest

type RequestInterfase interface {
	GetConfig() ConfigInterface
	GetContent() ContentInterface
	SetContent(ContentInterface)
	GetSecrets() SecretsInterface
	Do() ResponseInterface
}

type Request struct {
	config   ConfigInterface
	content  ContentInterface
	secrets  SecretsInterface
	response ResponseInterface
}

func (r *Request) GetConfig() ConfigInterface {
	return r.config
}

func (r *Request) GetContent() ContentInterface {
	return r.content
}

func (r *Request) SetContent(content ContentInterface) {
	r.content = content
}

func (r *Request) GetSecrets() SecretsInterface {
	return r.secrets
}

func (r *Request) GetResponse() ResponseInterface {
	return r.response
}
