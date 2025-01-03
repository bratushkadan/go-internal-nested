package booger

import (
	booger "github.com/bratushkadan/internal-nested/client/v3/internal"
)

func Boo(name string) {
	booger.BooImpl(name)
}
