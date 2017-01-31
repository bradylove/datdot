package config

type Config struct {
	Remote   string            `json:"remote"`
	Dotfiles map[string]string `json:"dotfiles,omitempty"`
}
