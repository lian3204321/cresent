package main

type treeNode struct {
	matchType    string
	name         string
	handle       map[string]func()
	methodList   []string
	parentNode   *treeNode
	chirdrenNode []*treeNode
}

//在父节点后加子节点
func (p *treeNode) addSubNode() {

}

type tree struct {
	base *treeNode
}

func (p *tree)new(){
	
}

//在树中添加一个节点
func (p *tree) addNode() {

}



func main() {

}
