package clients

type Client interface {
	GetHost() (ip string)
	GetOpenPortsUserInfo() (host string, from, limit int)
	GetFilePath() (filepath string)
	GetHash() (hash string)
}
