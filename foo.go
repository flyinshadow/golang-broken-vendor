package broken

import (
	"a"
	"b"
)

func broken() {
	var it a.A
	it = "foo"

	b.Do(it)
}
