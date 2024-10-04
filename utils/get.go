package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func Get_test(url string) (string, bool) {
	_, err := http.Get(url)
	if err != nil {
		return err.Error(), false
	} else {
		return "成功", true
	}
}

func Get(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("=> 请求错误 %d %s", res.StatusCode, res.Status)
	} else {
		fmt.Println("成功")
		decode(res.Body)
	}

}

func decode(body io.ReadCloser) {
	fmt.Print("==> 解析返回结果... ")

	cart, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	name := cart.Find(".d-sm-none .bottom-10").Text()

	paint := cart.Find("div[class=col-12] span[class=ng-binding]").Text()

	var author string

	cart.Find(".d-md-block:nth-of-type(4)").Each(func(i int, s *goquery.Selection) {
		s.Find("span").Remove()
		result := s.Text()
		author = strings.Join(strings.FieldsFunc(result, unicode.IsSpace), "")
	})

	fmt.Println("成功")
	fmt.Println(LINE)
	fmt.Println("名称:", strings.Join(strings.FieldsFunc(name, unicode.IsSpace), ""))
	fmt.Println("作者:", author)
	fmt.Println("获取到的画数:", paint)

}
func Get_paint(start, end int, url string) {
	fmt.Println(LINE)
	for i := start; i < end+1; i++ {
		log.Println("Get: " + url + "/chapter-" + strconv.Itoa(i))
		res, err := http.Get(url + "/chapter-" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		} else {
			dir := url[26:] + "-" + strconv.Itoa(i)
			log.Println("mkdir: " + dir)
			os.MkdirAll(dir, 0777)
			decode_paint(res.Body, dir)
		}
		defer res.Body.Close()

	}
}

func decode_paint(body io.ReadCloser, dir string) {
	cart_paint, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	cart_paint.Find(".ng-scope img").Each(func(i int, s *goquery.Selection) {
		url := s.AttrOr("src", "default")
		get_paint(url, dir)
	})

}

func get_paint(url, dir string) {
	log.Println("Get:", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	file_name := path.Base(url)
	log.Println("Save:", file_name)
	out, err := os.Create("./" + dir + "/" + file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err_copy := io.Copy(out, res.Body)
	if err_copy != nil {
		log.Fatal(err_copy)
	}
}
