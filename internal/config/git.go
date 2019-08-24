package config

type GitConfig struct {
	U string       `yaml:"url"`
	C Credential   `yaml:"credential"`
	R []Repository `yaml:"repositories"`
}

type Repository struct {
	U string     `yaml:"url"`
	N string     `yaml:"name"`
	C Credential `yaml:"credential"`
}

type Credential struct {
	Us string `yaml:"username"`
	P  string `yaml:"password"`
}

func NewRepository(url, name, username, password string) Repository {
	return Repository{U: url, N: name, C: Credential{Us: username, P: password}}
}

func (c *GitConfig) Url() string {
	return c.U
}

func (c *GitConfig) Password() string {
	return c.C.P
}

func (c *GitConfig) Username() string {
	return c.C.Us
}

func (c *GitConfig) Repositories() []Repository {
	return c.R
}

func (r *Repository) Name() string {
	return r.N
}

func (r *Repository) Url() string {
	return r.U
}

func (r *Repository) Username() string {
	return r.C.Us
}

func (r *Repository) Password() string {
	return r.C.P
}
