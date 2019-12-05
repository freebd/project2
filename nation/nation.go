package nation

import (
	"fmt"
)

type Nation struct {
	Name string
}

func (a Nation) Speaks(lang string) {
	fmt.Printf("%s says: %s", a.Name, lang)
}
