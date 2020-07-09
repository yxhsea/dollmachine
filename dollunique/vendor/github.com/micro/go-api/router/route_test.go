package router

import (
	"testing"
)

func TestApiRoute(t *testing.T) {
	namespace := "go.micro.api"

	testData := []struct {
		path    string
		service string
		method  string
	}{
		{
			"/foo/bar",
			namespace + ".foo",
			"Foo.Bar",
		},
		{
			"/foo/foo/bar",
			namespace + ".foo",
			"Foo.Bar",
		},
		{
			"/foo/bar/baz",
			namespace + ".foo",
			"Bar.Baz",
		},
		{
			"/foo/bar/baz-xyz",
			namespace + ".foo",
			"Bar.BazXyz",
		},
		{
			"/foo/bar/baz/cat",
			namespace + ".foo.bar",
			"Baz.Cat",
		},
		{
			"/foo/bar/baz/cat/car",
			namespace + ".foo.bar.baz",
			"Cat.Car",
		},
		{
			"/foo/fooBar/bazCat",
			namespace + ".foo",
			"FooBar.BazCat",
		},
		{
			"/v1/foo/bar",
			namespace + ".v1.foo",
			"Foo.Bar",
		},
		{
			"/v1/foo/bar/baz",
			namespace + ".v1.foo",
			"Bar.Baz",
		},
		{
			"/v1/foo/bar/baz/cat",
			namespace + ".v1.foo.bar",
			"Baz.Cat",
		},
	}

	for _, d := range testData {
		s, m := apiRoute(namespace, d.path)
		if d.service != s {
			t.Fatalf("Expected service: %s for path: %s got: %s", d.service, d.path, s)
		}
		if d.method != m {
			t.Fatalf("Expected service: %s for path: %s got: %s", d.method, d.path, m)
		}
	}
}

func TestProxyRoute(t *testing.T) {
	namespace := "go.micro.api"

	testData := []struct {
		path    string
		service string
	}{
		{
			"/f",
			namespace + ".f",
		},
		{
			"/f-b",
			namespace + ".f-b",
		},
		{
			"/foo/bar",
			namespace + ".foo",
		},
		{
			"/foo-bar",
			namespace + ".foo-bar",
		},
		{
			"/foo-bar-baz",
			namespace + ".foo-bar-baz",
		},
		{
			"/foo/bar/bar",
			namespace + ".foo",
		},
		{
			"/v1/foo/bar",
			namespace + ".v1.foo",
		},
		{
			"/v1/foo/bar/baz",
			namespace + ".v1.foo",
		},
		{
			"/v1/foo/bar/baz/cat",
			namespace + ".v1.foo",
		},
	}

	for _, d := range testData {
		s := proxyRoute(namespace, d.path)
		if d.service != s {
			t.Fatalf("Expected service: %s for path: %s got: %s", d.service, d.path, s)
		}
	}
}
