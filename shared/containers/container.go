package containers

type IContainer interface {
	CreateContainer()
	GetConnHostAndPort() (string, int, error)
	CloseContainer()
}
