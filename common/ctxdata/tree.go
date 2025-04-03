package ctxdata

type Node struct {
	Id        int64   `json:"id"`
	ParentId  int64   `json:"parent_id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Ancestors string  `json:"ancestors"`
	Children  []*Node `json:"children"`
}

func GetTreeRecursive(list []*Node, parentId int64) []*Node {
	res := make([]*Node, 0)
	for _, v := range list {
		if v.ParentId == parentId {
			v.Children = GetTreeRecursive(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}
