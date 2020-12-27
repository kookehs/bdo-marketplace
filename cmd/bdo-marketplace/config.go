package main

type Config struct {
	Files     []FilePair `json:"files"`
	Headers   Headers    `json:"headers"`
	Proxy     string     `json:"proxy"`
	Timeout   string     `json:"timeout"`
	FormToken string     `json:"formToken"`
	URL       string     `json:"url"`
}

type FilePair struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Headers struct {
	CookieToken string `json:"cookieToken"`
	Session     string `json:"session"`
	UserAgent   string `json:"userAgent"`
}
