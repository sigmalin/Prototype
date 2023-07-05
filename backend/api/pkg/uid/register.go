package uid

import (
	"fmt"
)

var uuids = make(map[string]Generator)

func Register(key string, generator Generator) {
	if generator == nil {
		panic("uid: cannot register uuid with nil value")
	}

	if _, exist := uuids[key]; exist {
		panic(fmt.Errorf("uid: cannot register the same response %s", key))
	}

	uuids[key] = generator
}

func GetGenerator(key string) Generator {
	gen, exist := uuids[key]
	if !exist {
		panic(fmt.Errorf("uid: cannot find generator %s", key))
	}
	return gen
}
