package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// url = 'http://www.someserver.com/cgi-bin/register.cgi'

// values = {'action':'parse','section':0,'prop':'text','format':'json','page':topic}

// data = urllib.parse.urlencode(values)
// data = data.encode('ascii')  # data should be bytes
// req = urllib.request.Request(url, data)
// with urllib.request.urlopen(req) as response:
// 	page = response.read()
// 	y = json.loads(page)
// 	content = y['parse']['text']['*']
// 	return content.count(topic)
type Data struct {
	//Number  int    `json:"number"`
	//Message string `json:"message"`
	Status string `json:"status"`
	//Data   []string `json:"data"`
	//Response []string `json:"response"`
	// data     struct {
	// 	docs string `json:docs`
	// }
}

func main() {
	//url := "http://api.plos.org/search?q=title:DNA"
	url := "http://dummy.restapiexample.com/api/v1/employees"
	//url := "http://api.open-notify.org/astros.json"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data Data
	json.Unmarshal(body, &data)
	//fmt.Println(string(body))
	fmt.Println(data)
}
