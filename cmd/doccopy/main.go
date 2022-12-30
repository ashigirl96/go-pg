package main

import (
	"fmt"
	"github.com/ashigirl96/go-pg/pkg"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`\w.*`)

func main() {
	pasted, err := pkg.PbPaste()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	texts := strings.Split(pasted, "\n")

	result := ""
	for _, text := range texts {
		result += fmt.Sprintf("%v ", reg.FindString(text))
	}
	err = pkg.PbCopy(fmt.Sprintf("https://www.deepl.com/ja/translator#en/ja/%v", url.PathEscape(result)))
	if err != nil {
		return
	}
}
