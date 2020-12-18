package pinyin

import (
	"regexp"
	"strings"
)

type Pinyin struct {
	Tune    int        // 带声调
	Normal  int        // 不带声调
	Initial int        // 仅声母
	First   int        // 仅首字母

	chinese  string    // 中文内容
	data     []string  // 拼音
}

const (
	Tune = iota    // 带声调
	Normal         // 不带声调
	Initial        // 仅声母
	First          // 仅首字母
)

const (
	Tune0 = iota   // 轻声
	Tune1 = iota   // 第一声
	Tune2 = iota   // 第二声
	Tune3 = iota   // 第三声
	Tune4 = iota   // 第四声
)

/* 初始化 */
func NewPinyin() *Pinyin {
	var py Pinyin = Pinyin{}
	py.Tune = Tune
	py.Normal = Normal
	py.Initial = Initial
	py.First = First

	return &py
}

/* 注册新的多音字的词语 */
func (py *Pinyin) Register(word string, pinyin []string) {
	words_dict[word] = pinyin
	if len(word) > max_word_dict_key_len {
		max_word_dict_key_len = len(word)
	}
}

/* 解析汉族生产拼音 */
func (py *Pinyin) Parse(chinese string) {
	py.chinese = chinese
	py.data = []string{}

	var chinese_rune = []rune(chinese)
	for ; len(chinese_rune) > 0; {
		var found bool = false
		/* 查词典 找多音字 */
		for i:=1; i < max_word_dict_key_len && i < len(chinese_rune); i++ {
			var word = string(chinese_rune[:i])
			var pinyin, exist = words_dict[word]
			if exist {
				for j := range pinyin {
					py.data = append(py.data, pinyin[j])
				}
				found = true
				chinese_rune = chinese_rune[i:]
			}
		}
		
		/* 词典里没查到的 查字典 */
		if !found {
			var pinyin, exist = chars_dict[chinese_rune[0]]
			if exist {
				var data = strings.Split(pinyin, ",")
				py.data = append(py.data, data[0])
			}
			chinese_rune = chinese_rune[1:] // 无论是否查到都删除第一个字符
		}
	}
}

/* 返回离散的原始拼音数据 */
func (py *Pinyin) Array() []string {
	var data = make([]string, len(py.data))
	copy(data, py.data)
	return data
}

/* 返回拼音的字符串 */
func (py *Pinyin) String(interval string, flag int) string {
	switch flag {

	case Tune:
		return strings.Join(py.data, interval)

	case Normal:
		var data []string = make([]string, len(py.data))
		copy(data, py.data)
		normal(data)
		return strings.Join(data, interval)

	case Initial:
		var data []string = make([]string, len(py.data))
		copy(data, py.data)
		initial(data)
		return strings.Join(data, interval)

	case First:
		var data []string = make([]string, len(py.data))
		copy(data, py.data)
		first(data)
		return strings.Join(data, interval)

	default:
		return ""
	}

}

/* 删除声调 */
func normal(pinyin []string)  {
	var finals_a, _ = regexp.Compile("[āáǎà]") 
	var finals_o, _ = regexp.Compile("[ōóǒò]") 
	var finals_e, _ = regexp.Compile("[ēéěè]") 
	var finals_i, _ = regexp.Compile("[īíǐì]") 
	var finals_u, _ = regexp.Compile("[ūúǔù]") 
	var finals_v, _ = regexp.Compile("[üǖǘǚǜ]") 

	/* 正则替换 */
	for i := range pinyin {
		pinyin[i] = finals_a.ReplaceAllString(pinyin[i], "a")
		pinyin[i] = finals_o.ReplaceAllString(pinyin[i], "o")
		pinyin[i] = finals_e.ReplaceAllString(pinyin[i], "e")
		pinyin[i] = finals_i.ReplaceAllString(pinyin[i], "i")
		pinyin[i] = finals_u.ReplaceAllString(pinyin[i], "u")
		pinyin[i] = finals_v.ReplaceAllString(pinyin[i], "v")  // 用v表示ü
	}
}


/* 只留声母 */
func initial(pinyin []string) {
	var initial_list []string = []string {
		"b", "p", "m", "f", "d", "t", "n", 
		"l", "g", "k", "h", "j", "q", "x", 
		"zh", "ch", "sh", "r", "z", "c", 
		"s", "y", "w",
	}

	for i := range pinyin {
		var found = false
		for j := range initial_list {
			if strings.HasPrefix(pinyin[i], initial_list[j]) {
				pinyin[i] = initial_list[j]
				found = true
				break
			}
		}

		if !found {
			pinyin[i] = " "
		}
	}
}


/* 只留首字母 */
func first(pinyin []string) {
	normal(pinyin)
	for i := range pinyin {
		pinyin[i] = pinyin[i][:1]
	}
}