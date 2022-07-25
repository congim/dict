/**
 * @Author:
 * @Description:
 * @File:  dict
 * @Date: 2021/5/8 11:52 上午
 */

package dict

import (
	"fmt"
	"regexp"
)

// 词库
type Dict struct {
	trie  *Trie
	Index *Index
	noise *regexp.Regexp // 拓展功能，混淆过滤器，暂时不启用
}

type Index struct {
	Describe string `json:"describe"` // 词库描述
	Name     string `json:"name"`     // niffler key
	Time     int64  `json:"time"`     // 修改&创建时间戳
}

func New() *Dict {
	return &Dict{
		trie:  NewTrie(),
		noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}
}

func (d *Dict) AddWorld(words ...string) {
	d.trie.Add(words...)
}

func (d *Dict) String() string {
	return fmt.Sprintf("describe: %s, name: %s, update: %d", d.Index.Describe, d.Index.Name, d.Index.Time)
}

func (d *Dict) Validate(text string) (bool, string) {
	if d.trie == nil {
		return true, text
	}
	result, sensitiveWord := d.trie.Validate(text)
	return result, sensitiveWord
}

func (d *Dict) Replace(text string, repl rune) string {
	if d.trie == nil {
		return text
	}
	return d.trie.Replace(text, repl)
}
