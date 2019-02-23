package pinyin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	tones = [][]rune{
		[]rune{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'},
		[]rune{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'},
		[]rune{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'},
		[]rune{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'},
	}
	neutrals = []rune{'a', 'e', 'i', 'o', 'u', 'v', 'A', 'E', 'I', 'O', 'U', 'V'}
	punctuations = []string{
		// 逗号
		"，", ",",
		// 句号
		"。", ".",
		// 感叹号
		"！", "!",
		// 问号
		"？", "?",
		// 冒号
		"：", ":",
		// 分号
		"；", ";",
		// 左/右单引号
		"‘", " '", "’", " '",
		// 左/右双引号
		"“", ` "`, "”", ` "`,
		// 左/右直角引号
		"「", " [", "」", " ]",
		"『", " [", "』", " ]",
		// 左/右括号
		"（", " (", "）", " )",
		"〔", " [", "〕", " ]",
		"【", " [", "】", " ]",
		"{", " {", "}", " }",
		// 省略号
		"……", "...",
		// 破折号
		"——", "-",
		// 连接号
		"—", "-",
		// 左/右斜杆
		"/", " /", "\\", " \\",
		// 波浪线
		"～", "~",
		// 书名号
		"《", " <", "》", " >",
		"〈", " <", "〉", " >",
		// 间隔号
		"·", " ·",
		// 顿号
		"、", ",",
	}
)

var (
	// 从带声调的声母到对应的英文字符的映射
	tonesMap map[rune]rune

	// 从汉字到声调的映射
	numericTonesMap map[rune]int

	// 从汉字到拼音的映射（带声调）
	pinyinMap map[rune]string

	initialized bool

	path = "pinyin.txt"
)

type Mode int

const (
	WithoutTone        Mode = iota + 1 // 默认模式，例如：guo
	Tone                               // 带声调的拼音 例如：guó
	InitialsInCapitals                 // 首字母大写不带声调，例如：Guo
)

type pinyin struct {
	origin string
	split  string
	mode   Mode
}

func init() {
	tonesMap = make(map[rune]rune)
	numericTonesMap = make(map[rune]int)
	pinyinMap = make(map[rune]string)
	for i, runes := range tones {
		for j, tone := range runes {
			tonesMap[tone] = neutrals[j]
			numericTonesMap[tone] = i + 1
		}
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		initialized = false
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "=>")
		if len(strs) < 2 {
			continue
		}
		i, err := strconv.ParseInt(strs[0], 16, 32)
		if err != nil {
			continue
		}
		pinyinMap[rune(i)] = strs[1]
	}
	initialized = true
}

func New(origin string) *pinyin {
	return &pinyin{
		origin:origin,
		split: " ",
		mode:WithoutTone,
	}
}

func (py *pinyin)Split(split string) *pinyin  {
	py.split = split
	return py
}

func (py *pinyin)Mode (mode Mode) *pinyin  {
	py.mode = mode
	return py
}

func (py *pinyin)Convert() string {
	sr := []rune(py.origin)
	words := make([]string, 0)
	for _, s := range sr{
		word, err := getPinyin(s, py.mode)
		if err != nil {
			return err.Error()
		}
		if len(word) > 0{
			words = append(words, word)
		}
	}
	return strings.Join(words, py.split)
}

func getPinyin(hanzi rune, mode Mode) (string, error) {
	if !initialized {
		return "", ErrInitialize
	}

	switch mode {
	case Tone:
		return getTone(hanzi), nil
	case InitialsInCapitals:
        return getInitialsInCapitals(hanzi), nil
	default:
		return getDefault(hanzi), nil
	}
}

func getTone(hanzi rune) string {
	return pinyinMap[hanzi]
}

func getDefault(hanzi rune) string {
	tone := getTone(hanzi)

	if tone == ""{
		return tone
	}

	output := make([]rune, utf8.RuneCountInString(tone))

	count := 0
	for _, t := range tone {
		neutral, found := tonesMap[t]
		if found {
			output[count] = neutral
		} else {
			output[count] = t
		}
		count++
	}
	return string(output)
}

func getInitialsInCapitals(hanzi rune) string {
	def := getDefault(hanzi)
	if def == ""{
		return def
	}
	sr := []rune(def)
	if sr[0] > 32 {
		sr[0] = sr[0] - 32
	}
	return string(sr)
}