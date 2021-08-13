package wind

import "strings"

/**
	预备知识：
		./tree_learn.go
	设计思路：
	路由匹配前缀树，数据结构为 struct node，将其命名为 tree
	1、tree 支持 增加子树，增加子树的过程，就是匹配 parent path，在下面增加 children path
  2、tree 支持 路由参数，通配符
*/

type node struct {
	path     string  //路径
	children []*node //子节点
	isWild   bool    //是否精确匹配，path 含有 : 或 * 时为true
	handle   Handle
}

func (n *node) addRoute(path string, handle Handle) {
	//pathslice := parsePath(path)
}

func (n *node) insertChild(path, fullPath string, handle Handle) {

}

//解析 path
// /post/:id => ["post",":id"]
func parsePath(path string) (paths []string) {
	pathslice := strings.Split(path, "/")
	for _, p := range pathslice {
		if p != "" {
			paths = append(paths, p)
		}
		if p == "*" {
			break
		}
	}
	return
}
