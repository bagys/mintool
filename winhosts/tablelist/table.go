package tablelist

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type Table struct {
	/**
	 * 存放 每段数据最大长度
	 */
	contentFieldMaxLen []int
	/**
	 * 存放 每行数据
	 */
	contentLine [][]string
	/**
	 * 存放 第一行tab数据
	 */
	tab []string
	/**
	 * 存放 tab的间距
	 */
	spacing int
	/**
	 * 存放 tab的前缀
	 */
	prefixTab string
	/**
	 * 存放 每行数据的前缀
	 */
	prefixContent string
	/**
	 * 默认字段数量
	 */
	tablen int
}

func NewTable() *Table {
	t := &Table{
		spacing:       10,
		prefixTab:     " - ",
		prefixContent: " * ",
		tablen:        5,
	}
	t.initContentMaxLen(t.tablen)
	return t
}

func (t *Table) SetPrefixTab(Prefix string) {
	t.prefixTab = Prefix
	t.prefixContent = Prefix
}

func (t *Table) SetPrefixContent(Prefix string) {
	t.prefixContent = Prefix
}

func (t *Table) SetSpacing(spacing int) {
	t.spacing = spacing
}

func (t *Table) SetTab(tab []string) {
	t.tab = tab
	if len(t.contentFieldMaxLen) > len(tab) {
		t.tab = append(t.tab, make([]string, len(t.contentFieldMaxLen)-len(tab))...)
	}
}

func (t *Table) initContentMaxLen(l int) {
	t.contentFieldMaxLen = make([]int, l)
}

func (t *Table) SetData(data []string) {
	t.contentLine = append(t.contentLine, data)
}

func (t *Table) printLine() {
	for k, contentSlice := range t.contentLine {
		lineStr := ""
		// 判断是否有前缀
		if t.prefixContent != "" {
			// 判断是否有tab
			if k == 0 && len(t.tab) > 0 {
				lineStr += t.prefixTab
			} else {
				lineStr += t.prefixContent
			}
		}
		for index, val := range contentSlice {
			// 当列最长 - 当前长度 + 间距
			space := t.contentFieldMaxLen[index] - len(val) + t.spacing
			lineStr += fmt.Sprintf("%s%s",
				val,
				strings.Repeat(" ", space),
			)
		}
		//lineStr += "\n"

		if k == 0 && len(t.tab) > 0 {
			fmt.Println(lineStr)
		} else {
			color.Green(lineStr)
		}
	}
}

func (t *Table) readData() {
	for _, v1 := range t.contentLine {
		for k, v2 := range v1 {
			if len(v2) > t.contentFieldMaxLen[k] {
				t.contentFieldMaxLen[k] = len(v2)
			}
		}
	}
}

/**
 *	打印
 */
func (t *Table) Print() {
	if len(t.tab) > 0 {
		t.contentLine = append([][]string{t.tab}, t.contentLine...)
	}
	t.readData()

	t.printLine()
}
