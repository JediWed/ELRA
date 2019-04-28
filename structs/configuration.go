package structs

// Configuration for config params set by user in config.json
type Configuration struct {
	Port        int
	AllowedHost string
	SigningKey  string
	Pepper      string
	Database    string
	Macaroon    string
}
