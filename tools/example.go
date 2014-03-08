package main

import (
	"flag"
	"fmt"
	"github.com/huichen/sego"
)

var (
	text = flag.String("text", "lambgini", "要分词的文本")
)

func main() {
	flag.Parse()

	var seg sego.Segmenter
	seg.LoadDictionary("../data/dictionary.txt")

	segments := seg.Segment([]byte(*text))
	fmt.Println(sego.SegmentsToString(segments, true))
}
