package parse

import (
	"io/ioutil"
	"testing"

	"baseLearn/15.爬虫实现--单任务版-并发版/crawler(暂时无用了)/model"
)

// 测试网址:http://album.zhenai.com/u/105904218
func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "huijiao")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element;but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:        31,
		Height:     162,
		Weight:     51,
		Income:     "5001-8000元",
		Gender:     "女",
		Name:       "huijiao",
		Xinzuo:     "天蝎座",
		Occupation: "人事主管",
		Marriage:   "离异",
		House:      "已购房",
		Hukou:      "河南南阳",
		Education:  "大专",
		Car:        "未购车",
	}
	if profile != expected {
		t.Errorf("Expected %v;But was %v", expected, profile)
	}
}
