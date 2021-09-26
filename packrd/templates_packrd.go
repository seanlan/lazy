package packrd

import (
	"github.com/gobuffalo/packr/v2"
)

func init() {
	packr.New("lazy", "../templates")
}
