package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// from=en
// &to=zh&query=good&simple_means_flag=3&
// sign=262931.57378&token=747080f2902d70a28e1e6a7abf8b85a9&domain=common
type DictRequestHuoshan struct {
	Source          string   `json:"source"`
	Words           []string `json:"words"`
	Source_language string   `json:"source_language"`
	Target_language string   `json:"target_language"`
}

type DictResponseHuoshan struct {
	Details []struct {
		Detail string `json:"detail"`
		Extra  string `json:"extra"`
	} `json:"details"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

type DictResponseHuoshanDetail struct {
	ErrorCode string `json:"errorCode"`
	RequestID string `json:"requestId"`
	Msg       string `json:"msg"`
	Result    []struct {
		Ec struct {
			Basic struct {
				UsPhonetic string `json:"usPhonetic"`
				Phonetic   string `json:"phonetic"`
				Explains   []struct {
					Pos   string `json:"pos"`
					Trans string `json:"trans"`
				} `json:"explains"`
				UkPhonetic string `json:"ukPhonetic"`
			} `json:"basic"`
			Lang   string `json:"lang"`
			IsWord bool   `json:"isWord"`
		} `json:"ec"`
	} `json:"result"`
}

func QueryHuoshan(word string) {
	client := &http.Client{}
	// var data = strings.NewReader(`{"source":"youdao","words":["good"],"source_language":"en","target_language":"zh"}`)

	request := DictRequestHuoshan{Source: "youdao", Words: []string{word},
		Source_language: "en", Target_language: "zh"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/detail/v1/?msToken=&X-Bogus=DFSzswVLQDaZV29HSZDVTF9WX7rt&_signature=_02B4Z6wo000015abjAwAAIDDFpl2T6aV5NOWm4iAAIZqy9CRlkrxo1WeQXhAWv3TlXN3oxFJ67.2AmpMxIgghK5tAMlxCG73eY.RRoAiBUtrrhglcNNe0mwhBQGx0foUGrs7m5Sk47SZ0rTl9e", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16738539773293788; i18next=zh-CN; s_v_web_id=verify_lcyhfv7a_FEea6qLt_0mSo_44Uk_95OR_dUwBoorfdnuD; ttcid=23117e9fbac34d65bf06fd184eba348917; tt_scid=X3Pcx7Uc4OWo-POd-vZkYJt9Sn9GzTKfqYRp3o.W3ORdMsyjAFt7b7Q3F80Fo71866e9")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/?category=&home_language=zh&source_language=detect&target_language=zh&text=good")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var dictResponse DictResponseHuoshan
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	var dictResponseDetail DictResponseHuoshanDetail
	err = json.Unmarshal([]byte(dictResponse.Details[0].Detail), &dictResponseDetail)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("====================Result of \"" + word + "\" in Huoshan====================")
	fmt.Println("UK:", "["+dictResponseDetail.Result[0].Ec.Basic.UkPhonetic+"]",
		"US:", "["+dictResponseDetail.Result[0].Ec.Basic.UsPhonetic+"]")
	for _, item := range dictResponseDetail.Result[0].Ec.Basic.Explains {
		fmt.Println(item.Pos, strings.Split(item.Trans, "；")[0])
	}
}
