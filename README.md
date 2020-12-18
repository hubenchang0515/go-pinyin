# pinyin
Go的汉字转拼音库

使用了以下项目中的拼音数据 :  
* [hotoo/pinyin](https://github.com/hotoo/pinyin)
* [mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin)

## Usage
```Go
package main

import "github.com/hubenchang0515/go-pinyin"

func main() {
	var py = pinyin.NewPinyin()
	py.Parse("重庆有很多体重非常重的人")
	println(py.String("-", py.Initial))
	println(py.String("-", py.Normal))
	println(py.String("-", py.Tune))
}
```
---
初始化
```Go
func NewPinyin() *Pinyin
```
---
注册一个词语的读音
```Go
func (py *Pinyin) Register(word string, pinyin []string)
```
* 参数
  * `word` - 汉语词汇
  * `pinyin` - 拼音，`{"chong", "qing"}`
* 示例
  * `py.Register("重庆", {"chong", "qing"})`

---
解析一串中文
```Go
func (py *Pinyin) Parse(chinese string) bool
```
* 参数
  * `chinese` - 中文内容
* 示例
  * `py.Parse("重庆有很多体重非常重的人")`

---
获得拼音的数组
```Go
func (py *Pinyin) Array() []string
```

---
获得拼音的字符串
```Go
func (py *Pinyin) String(interval string, flag int) string 
```
* 参数
  * `interval` - 间隔字符串
  * `flag` - 表示形式
* 示例
  * `py.String("-", py.Normal)`

flag的取值 :  
```Go
const (
	Tune = iota    // 带声调
	Normal         // 不带声调
	Initial        // 仅声母
	First          // 仅首字母(不带声调)
)
```
