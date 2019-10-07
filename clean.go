package clean_sec_group

import (
	"fmt"
)

func remove_sec(edge [2]string, node string) (bool, error) {
	if node == edge[len(edge)-1] {
		return true, nil
	} else {
		return false, fmt.Errorf("%s is not at the end or it may not exist", node)
	}
}
