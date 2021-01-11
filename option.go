package gowooco

import "fmt"

type Option func(c *Client)

// WithVersion version config option
func WithVersion(apiVersion string) Option {
	return func(c *Client) {
		pathPrefix := defaultApiPathPrefix
		if len(apiVersion) > 0 && apiVersionRegex.MatchString(apiVersion) {
			c.pathPrefix = fmt.Sprintf("%s/%s", defaultApiPathPrefix, apiVersion)
		}
		c.version = apiVersion
		c.pathPrefix = pathPrefix
	}
}

// WithRetry Timeout config option
func WithRetry(retries int) Option {
	return func(c *Client) {
		c.retries = retries
	}
}

// WithLog log config option
func WithLog(logger LeveledLoggerInterface) Option {
	return func(c *Client) {
		c.log = logger
	}
}
