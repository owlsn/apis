package api

import (
	"fmt"

	"github.com/kataras/iris/v12/sessions"

	"github.com/kataras/iris/v12/mvc"
)

// If controller's fields (or even its functions) expecting an interface
// but a struct value is binded then it will check
// if that struct value implements
// the interface and if true then it will add this to the
// available bindings, as expected, before the server ran of course,
// remember? Iris always uses the best possible way to reduce load
// on serving web resources.

// LoggerService : Logger Service
type LoggerService interface {
	Log(string)
}
// PrefixedLogger : Prefixed Logger
type PrefixedLogger struct {
	Prefix string
}
// Log : Log
func (s *PrefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.Prefix, msg)
}
// BasicController : Basic Controller
type BasicController struct {
	Logger LoggerService

	Session *sessions.Session
}
// BeforeActivation : Before Activation
func (c *BasicController) BeforeActivation(b mvc.BeforeActivation) {
	b.HandleMany("GET", "/custom /custom2", "Custom")
}
// AfterActivation : After Activation
func (c *BasicController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
	}
}
// Get : Get
func (c *BasicController) Get() string {
	count := c.Session.Increment("count", 1)

	body := fmt.Sprintf("Hello from basicController\nTotal visits from you: %d", count)
	c.Logger.Log(body)
	return body
}
// Custom : Custom
func (c *BasicController) Custom() string {
	return "custom"
}
// BasicSubController : Basic Sub Controller
type BasicSubController struct {
	Session *sessions.Session
}
// Get : Get
func (c *BasicSubController) Get() string {
	count := c.Session.GetIntDefault("count", 1)
	return fmt.Sprintf("Hello from basicSubController.\nRead-only visits count: %d", count)
}
