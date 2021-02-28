package main

import "flag"

type args struct {
	ConfigPath string
}

func newArgs() *args {
	c := &args{
		ConfigPath: "",
	}
	c.parse()

	return c
}

func (c *args) parse() {
	flag.StringVar(&c.ConfigPath, "config-path", "configs/config.toml", "path to config file")

	flag.Parse()
}
