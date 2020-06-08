package main

type Config struct {
	Repo           string `required:"true"`
	Ref            string `default:"main"`
	UsernameSecret string `split_words:"true"`
	PasswordSecret string `split_words:"true"`
	RootDir        string `split_words:"true" default:"/tmp"`
	DestDir        string `split_words:"true"`
	Timeout        uint32 `default:"120"`
	Debug          bool
}
