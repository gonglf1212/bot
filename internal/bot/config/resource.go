package config

import "fmt"

const domain = "api.sgroup.qq.com"
const sandBoxDomain = "sandbox.api.sgroup.qq.com"

const scheme = "https"

type uri string

const (
	GatewayURI uri = "/gateway" //wss接入点
)

func GetURL(endpoint uri, sandbox bool) string {
	d := domain
	if sandbox {
		d = sandBoxDomain
	}
	return fmt.Sprintf("%s://%s%s", scheme, d, endpoint)
}
