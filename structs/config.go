package structs

// Retry struct for retry. It has count and interval fields
type Retry struct {
	Count    int `yaml:"count"`
	Interval int `yaml:"interval"`
}

// Header struct for header. It has name and value fields
type Header struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Target struct for target. It has name, url, ssl_verify, method, status, headers, body, timeout and retry fields
type Target struct {
	Name      string   `yaml:"name"`
	Url       string   `yaml:"url"`
	SSLVerify bool     `yaml:"ssl_verify"`
	Method    string   `yaml:"method"`
	Status    int      `yaml:"status"`
	Headers   []Header `yaml:"headers"`
	Body      string   `yaml:"body"`
	Timeout   int      `yaml:"timeout"`
	Retry     Retry    `yaml:"retry"`
}

// Alert struct for alert. It has name, type, token and chat_id fields
type Alert struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Token  string `yaml:"token"`
	ChatId string `yaml:"chat_id"`
}

// Config struct for config. It has target and alert fields
type Config struct {
	Target []Target `yaml:"target"`
	Alert  []Alert  `yaml:"alert"`
}
