package main

type Config struct {
	Repo     string `required:"true"`
	Ref      string `default:"master"`
	Username string
	Password string
	RootDir  string `split_words:"true" default:"/tmp"`
	DestDir  string `split_words:"true"`
	Timeout  uint32 `default:"120"`
	Debug    bool
}
