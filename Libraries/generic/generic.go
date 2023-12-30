package generic

import (
	"strconv"
)

func ToInts(strings []string) (ints []int, err error) {

	for i := 0; i < len(strings); i++ {
		v, err := strconv.Atoi(strings[i])
		if err != nil {
			return nil, err
		}

		ints = append(ints, v)
	}

	return ints, nil

}