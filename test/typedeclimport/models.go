package typedeclimport

import subpkg "github.com/temporalio/gogo-protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
