##participle

正向最大匹配分词

将该package改为逆向最大匹配分词的方法如下：

（1）在InsertWord的时候反转word，再插入

（2）先反转sentence，再分词，再将得到的结果数组反转（先反转数组，再反转数组元素）


基于 https://github.com/gansidui/trie

对于内存吃紧的可以改用double array trie实现。

~~~go
package main

import (
	"github.com/gansidui/participle"
	"log"
)

func main() {
	p := participle.NewParticiple()

	p.InsertWord("喜欢")
	p.InsertWord("谷歌")
	p.InsertWord("Golang")

	ret := p.ForwardMaximumMatchSplit("我喜欢Golang，谷歌大法好")
	if len(ret) != 3 || ret[0] != "喜欢" || ret[1] != "Golang" || ret[2] != "谷歌" {
		log.Fatal()
	}
}

~~~


##LICENSE

MIT