package ff_rule

import (
	"fmt"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"testing"
)

type TestRule struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	PId   int    `json:"p_id"`
	Level int    `json:"level"`
}

func TestTree(t *testing.T) {
	var treeArr []*TestRule
	treeArr = append(
		treeArr,
		&TestRule{Id: 1, Title: "a", PId: 0},
		&TestRule{Id: 2, Title: "b", PId: 0},
		&TestRule{Id: 3, Title: "c", PId: 0},
		&TestRule{Id: 4, Title: "d", PId: 1},
		&TestRule{Id: 5, Title: "e", PId: 2},
		&TestRule{Id: 6, Title: "f", PId: 3},
		&TestRule{Id: 7, Title: "g", PId: 4},
	)
	list := GetTree(treeArr, 0, 0)
	for _, v := range list {
		fmt.Println(ff_json.MarshalToStringNoError(v))
	}
}

var TreeList []*TestRule

func GetTree(treeArr []*TestRule, Pid int, level int) []*TestRule {
	for _, v := range treeArr {
		if v.PId == Pid {
			v.Level = level + 1
			TreeList = append(TreeList, v)
			GetTree(treeArr, v.Id, level+1)
		}
	}
	return TreeList
}
