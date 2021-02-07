package main

import (
	"fmt"
)

/**
1.5.12
使用路径压缩的 quick-union 算法。根据路径压缩修改 quick-union 算法(请见 1.5.2.3 节)，在 find() 方法中添加一个循环来将从 p 到根节点的路径上的每个触点都连接到根节点。给出一列 输入，使该方法能够产生一条长度为 4 的路径。
注意:该算法的所有操作的均摊成本已知为对 数级别
*/
func main() {
	QK := NewQuickUnionPathCompression(10)

	QK.Union(4, 3)
	QK.Union(3, 8)
	QK.Union(6, 5)
	QK.Union(9, 4)
	QK.Union(2, 1)
	QK.Union(8, 9)
	QK.Union(5, 0)
	QK.Union(7, 2)
	QK.Union(6, 1)
	QK.Union(1, 0)
	QK.Union(6, 7)

	fmt.Println(QK.Count())
}
