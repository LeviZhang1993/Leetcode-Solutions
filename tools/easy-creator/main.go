package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TableRow []string

func (t TableRow) Len() int { return len(t) }

func extractFirstNum(s string) int {
	ans, _ := strconv.Atoi(s[1:6])
	return ans
}

func (t TableRow) Less(i, j int) bool {
	return extractFirstNum(t[i]) < extractFirstNum(t[j])
}

func (t TableRow) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func sortReadMeQuestions() {
	f, err := os.OpenFile("../../README.md", os.O_RDWR, 0666)
	defer f.Close()
	if err != nil {
		log.Fatal("Fail to open README in sortReadMeQuestions")
	}
	table := TableRow{}
	scanner := bufio.NewScanner(f)
	res := make([]string, 0)
	idx := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 && text[0] == byte('|') {
			table = append(table, text)
			idx = append(idx, len(res))
		}
		res = append(res, text)
	}
	table, idx = table[1:], idx[1:]
	sort.Sort(table)
	for i := range idx {
		res[idx[i]] = table[i]
	}
	f.Seek(0, 0)
	for i, text := range res {
		if i != len(res)-1 {
			f.WriteString(text + "\n")
		} else {
			f.WriteString(text)
		}
	}
}

func main() {
	// flag
	// title: question
	var titleFlag = flag.String("title", "", "title of question like to add")
	var link = flag.String("link", "", "link of question'")
	var t = flag.String("type", "create", "type of command")
	flag.Parse()
	if *t == "sort" {
		sortReadMeQuestions()
		return
	}
	if *titleFlag == "" || *link == "" {
		log.Fatal("Expect a non-empty title")
		return
	}
	idx := strings.Index(*titleFlag, ".")
	num, _ := strconv.Atoi((*titleFlag)[0:idx])
	desc := (*titleFlag)[idx+2:]
	title := strings.Replace(desc, " ", "-", -1)
	wholeTitle := fmt.Sprintf("%05d.%s", num, title)
	// create directory
	if _, err := os.Stat(fmt.Sprintf("../../leetcode/%s", wholeTitle)); err == nil {
		log.Println("question exists")
		os.Exit(0)
	}
	for _, dir := range []string{"go-solution", "python-solution"} {
		err := os.MkdirAll(fmt.Sprintf("../../leetcode/%s/%s", wholeTitle, dir), 0700)
		if err != nil {
			log.Println("create " + dir + " failure." + err.Error())
		}
	}
	goSolutionDir := fmt.Sprintf("../../leetcode/%s/go-solution/", wholeTitle)
	pySolutionDir := fmt.Sprintf("../../leetcode/%s/python-solution/", wholeTitle)
	f, err := os.Create(goSolutionDir + "main.go")
	if err != nil {
		log.Println("create " + "main.go" + " failure." + err.Error())
	} else {
		f.Write([]byte("package main\n\nfunc main() {\n}"))
	}
	f.Close()
	var prefix = "github.com/LeviZhang1993/Leetcode-Solutions/leetcode"
	f, err = os.Create(goSolutionDir + "go.mod")
	if err != nil {
		log.Println("create " + "go.mod" + " failure." + err.Error())
	} else {
		f.Write([]byte(fmt.Sprintf("module %s/%s/go-solution\n\ngo 1.20", prefix, wholeTitle)))
	}
	f.Close()
	f, err = os.Create(pySolutionDir + "main.py")
	if err != nil {
		log.Println("create " + "main.py" + " failure." + err.Error())
	} else {
		f.Write([]byte(fmt.Sprintf("class Solution:\n    pass\n\nif __name__ == \"__main__\":\n    solu = Solution()\n")))
	}
	f.Close()
	var readMe = fmt.Sprintf("# [%s](%s)\n## Description\n## Solution\n", *titleFlag, *link)
	f, err = os.Create(pySolutionDir + "README.md")
	if err != nil {
		log.Println("create " + "readme" + " failure." + err.Error())
	} else {
		f.Write([]byte(readMe))
	}
	f.Close()
	f, err = os.Create(goSolutionDir + "README.md")
	if err != nil {
		log.Println("create " + "readme" + " failure." + err.Error())
	} else {
		f.Write([]byte(readMe))
	}
	f.Close()
	f, err = os.OpenFile("../../README.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("open " + "readme" + " failure." + err.Error())
	} else {
		var linkPrefix = "https://github.com/LeviZhang1993/Leetcode-Solutions/tree/main/leetcode/"
		f.Write([]byte(fmt.Sprintf("\n|%05d|%s|[Go](%s%s/go-solution)|[Python](%s%s/python-solution)|holder|",
			num, desc, linkPrefix, wholeTitle, linkPrefix, wholeTitle)))
	}
	f.Close()
	sortReadMeQuestions()
}
