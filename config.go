package main

type Config struct {
	Input  string `json:"input"`
	Output string `json:"output"`
	Cookie string `json:"cookie"`
	Token  string `json:"token"`
	URL    string `json:"url"`
}
