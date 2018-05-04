package parse

import (
	"testing"

	"io/ioutil"
)

// 测试城市列表的数据的完整性
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("test.html")
	// 由于网络测试的影响存在问题，所以我们进行一个内容进行文件保存
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("Result Should have %d requests;but had %d ", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("Expected Url #%d : %s;but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("Result Should have %d requests;but had %d ", resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("Expected Url #%d : %s;but was %s", i, city, result.Items[i].(string))
		}
	}
}
