package participle

import (
	"github.com/gansidui/trie"
	"unicode/utf8"
)

type Participle struct {
	tr *trie.Trie
}

func NewParticiple() *Participle {
	return &Participle{
		tr: trie.New(),
	}
}

// 插入一个词语
func (p *Participle) InsertWord(word string) {
	p.tr.Insert(word, nil)
}

// 正向最大匹配分词
// 如果是逆向最大匹配则完全相反，反转word后再插入，分词也先将sentence反转，查询得到结果数组，反转结果数组，再反转单个的word
func (p *Participle) ForwardMaximumMatchSplit(sentence string) []string {
	length := utf8.RuneCountInString(sentence)
	runes := make([]rune, length)
	i := 0
	for _, v := range sentence {
		runes[i] = v
		i++
	}

	ret := make([]string, 0)
	flag, index, left, right := false, 0, 0, length

	for left < right {
		flag, _, index = p.tr.FindByRunes(runes[left:right])
		index += left

		if flag {
			ret = append(ret, string(runes[left:right]))
			left = right
			right = length

		} else {
			if index == left {
				left++
			} else {
				right = index
			}
		}
	}

	return ret
}
