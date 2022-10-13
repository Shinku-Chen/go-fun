package tree

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/x-funs/go-fun"
)

func TestTire_FindAll(t *testing.T) {
	tire := new(Tire)
	tire.Add("挖土豆").Add("土豆").Add("上海").Add("上海帮忙").Add("to").Add("to box").Add("watch")

	text := "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。to box, into box, sport watch, watching tv"
	all := tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true, HasGroup: true})
	fmt.Printf("%+v\n", all)
}

func BenchmarkTire_FindAll(b *testing.B) {
	tire := new(Tire)
	tire.Add("挖土豆").Add("土豆").Add("上海").Add("上海帮忙").Add("to").Add("to box").Add("watch")

	wordPath := "./word.txt"
	if fun.IsExist(wordPath) {
		buf, _ := ioutil.ReadFile(wordPath)
		for _, word := range strings.Split(fun.String(buf), fun.LF) {
			if word != "" {
				tire.Add(word)
			}
		}
	}

	text := "去上海帮忙挖土豆，土豆地瓜哪里挖，一挖一麻袋。to box, into box, sport watch, watching tv"

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tire.FindAll(text, Opt{Limit: -1, Greed: true, Density: true, HasGroup: true})
	}
	b.StopTimer()
	b.ReportAllocs()
}
