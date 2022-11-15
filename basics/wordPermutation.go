package main

/* 给定一个词汇和词表，从词表中找出所有该词汇的兄弟词汇。
1、使用词表创建map，其中key为每个词汇的重排词，value为重排词相同的词的slice
2、对于任意给定词，将其重排词去查找map，即可得到对应的全部兄弟词汇slice
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var dic map[string][]string

func init() {
	// 词汇映射表可以在程序初始化时完成
	buildDic()
}

// 构造词汇映射表
func buildDic() {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal("failed to open file word.txt")
	}

	var wordstxt []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordstxt = append(wordstxt, scanner.Text())
	}
	file.Close()

	dic = make(map[string][]string)
	// 词汇重排，并存入map
	for _, word := range wordstxt {
		//词汇重排
		w := alphabetized(word)
		var lst []string

		// 若为已存词，则在基础上加上新词
		if len(dic) > 0 && len(dic[w]) > 0 {
			lst = dic[w]
		} else {
			//否则就重置slice，再加新词
			lst = []string{}
		}
		dic[w] = append(lst, word)
	}
}
func alphabetized(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func output(word string) {
	wd := alphabetized(word)
	fmt.Printf("Permutation group of %s is %s", word, dic[wd])
}
func main() {
	output("aloha")
}
