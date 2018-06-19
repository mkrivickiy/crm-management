package rest

type SecretsInterface interface {
	GetKey() string
}

type Secrets struct {
	key string
}

func (s *Secrets) GetKey() string {
	return s.key
}
