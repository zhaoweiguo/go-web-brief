package test

import(
//	"fmt"
	"encoding/json"

	"testing"
	"github.com/bmizerany/assert"

)

func TestUnmarshal(t *testing.T) {
	var jsonBlob = []byte(`[
        {"Name": "number 0","Order": "order 0"},
        {"Name": "number 1","Order": "order 1"}
    ]`)
	type Animal struct {
		Name  string     // 注意1: 字段要大写
		Order string     // 注意2: 字段名要与json串中的key相同
	}

	jsonRtn := make([]Animal, 2)

	jsonRtn[0] = Animal{ "number 0", "order 0"}
	jsonRtn[1] = Animal{"number 1", "order 1"}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)

t.Log(animals)
	assert.Equal(t, err, nil)
	assert.Equal(t, animals, jsonRtn)
}

func TestUnmarshal2(t *testing.T) {
	var jsonBlog = []byte(`{
            "date":"20140709",
            "news":[
               {
                 "title":"工资不高花钱却不少：得看花在哪儿，兴许是好事",
                 "url":"http:\/\/news-at.zhihu.com\/api\/1.2\/news\/4021498"
               },
               {
                  "title":"音量增加一倍，什么跟着变化了",
                  "url":"http:\/\/news-at.zhihu.com\/api\/1.2\/news\/4022651"
               }
            ]
         }`)

	type Item struct {
		Title 	string
		Url 	string
	}

	type News struct {
		Date 	string
		News 	[]*Item
	}

	var jsonNews News
	err := json.Unmarshal(jsonBlog, &jsonNews)

	assert.Equal(t, err, nil)
	t.Log(jsonNews.News[0])
}






