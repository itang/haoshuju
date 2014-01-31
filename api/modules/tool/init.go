package tool

import (
	"github.com/itang/haoshuju/open"
)

func init() {
	open.RegistModule(GetModule())
}
