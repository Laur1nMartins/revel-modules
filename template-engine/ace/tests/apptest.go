package tests

import (
	"github.com/Laur1nMartins/revel/testing"
)

type ApplicationTest struct {
	testing.TestSuite
}

func (t *ApplicationTest) Before() {
	println("Set up")
}

func (t *ApplicationTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *ApplicationTest) After() {
	println("Tear down")
}
