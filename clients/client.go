package clients

type Client interface {
	GetHostForCheckFromUser() (ip string)
	GetOpenPortsUserInfo() (host string, from, limit int)
}
