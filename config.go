package urlLookup

// Config holds the system-wide configuration data structure
type Config struct {
	RedisDB `json:"redis"`
}

// RedisDB holds the configuration types of a redis client connection
type RedisDB struct {
	// The network type, either tcp or unix.
	Network string `json:"network,omitempty"`

	// host:port address
	Address string `json:"address,omitempty"`

	// Optional password. Must match the password specified in the
	// require pass server configuration option.
	Password string `json:"password,omitempty"`

	// Database to be selected after connecting to the server.
	DB int `json:"db,omitempty"`
}
