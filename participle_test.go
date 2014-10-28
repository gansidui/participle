package participle

import (
	"testing"
)

func TestParticiple(t *testing.T) {
	p := NewParticiple()

	p.InsertWord("我A")
	p.InsertWord("B是")
	p.InsertWord("我AB是D")
	p.InsertWord("我")

	ret := p.ForwardMaximumMatchSplit("我A我AB是DB是D我AB") // 我A/我AB是D/B是/D/我A/B, 其中D和B不是词，会被丢掉
	if len(ret) != 4 || ret[0] != "我A" || ret[1] != "我AB是D" || ret[2] != "B是" || ret[3] != "我A" {
		t.Fatal()
	}

	ret = p.ForwardMaximumMatchSplit("D我AB是D")
	if len(ret) != 1 || ret[0] != "我AB是D" {
		t.Fatal()
	}

	ret = p.ForwardMaximumMatchSplit("D我AB是")
	if len(ret) != 2 || ret[0] != "我A" || ret[1] != "B是" {
		t.Fatal()
	}

	ret = p.ForwardMaximumMatchSplit("我")
	if len(ret) != 1 || ret[0] != "我" {
		t.Fatal()
	}
}

func TestParticiple2(t *testing.T) {
	p := NewParticiple()

	p.InsertWord("喜欢")
	p.InsertWord("谷歌")
	p.InsertWord("Golang")

	ret := p.ForwardMaximumMatchSplit("我喜欢Golang，谷歌大法好")
	if len(ret) != 3 || ret[0] != "喜欢" || ret[1] != "Golang" || ret[2] != "谷歌" {
		t.Fatal()
	}
}
