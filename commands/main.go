package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {

	financialStatementReaderHtml()

	// time.Sleep(5)

	// financialStatement()

	// balanceSheet()

	// companyStockList()

	// queryCompanyImportantMessageDailyList()

	// queryCompanyImportantMessageList()

	// companyImportantMessage()
}

type financialStatementJson struct {
	Name             string
	Value            string
	Percentage       string
	CountSpace       int
	OriginIndex      int
	ParentGroupIndex int
}

func financialStatementReaderHtml() {
	var res io.Reader
	res, _ = os.Open("commands/financialStatementHtml.html")

	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		log.Fatal(err)
	}

	var countSpace int
	var parentIndex int
	financialStatementList := []*financialStatementJson{}
	doc.Find("table.hasBorder tr").Each(func(i int, docSecondary *goquery.Selection) {
		if i < 4 {
			return
		}

		if countSpace == 0 {
			parentIndex = i - 1 // record on the previous index
		}

		countSpace = 0
		financialStatement := &financialStatementJson{}
		financialStatement.OriginIndex = i
		docSecondary.Find("td").Each(func(j int, s *goquery.Selection) {

			value := s.Text()

			if len(value) == 0 {
				return
			}

			// 會計科目名稱
			if j == 0 {
				for _, v := range value {
					// 12288 is space from ASCII
					if v != 12288 {
						break
					}
					countSpace++
				}

				if countSpace > 0 {
					financialStatement.ParentGroupIndex = parentIndex
				}

				financialStatement.Name = strings.TrimSpace(value)
				financialStatement.CountSpace = countSpace
			} else if j == 1 {
				financialStatement.Value = strings.TrimSpace(value)
			} else if j == 2 {
				financialStatement.Percentage = strings.TrimSpace(value)
			}
		})

		financialStatementList = append(financialStatementList, financialStatement)
	})

	for _, fs := range financialStatementList {
		fmt.Println(fs)
	}
}

// 綜合損益表
func financialStatement() {
	url := "https://mops.twse.com.tw/mops/web/ajax_t164sb04"
	method := "POST"

	payload := strings.NewReader("encodeURIComponent=1&step=1&firstin=1&off=1&queryName=co_id&inpuType=co_id&TYPEK=all&isnew=false&co_id=2330&year=110&season=2")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "jcsession=jHttpSession@2e3e0afd")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table.hasBorder tr").Each(func(i int, docSecondary *goquery.Selection) {
		if i < 5 {
			return
		}

		docSecondary.Find("td").Each(func(j int, s *goquery.Selection) {

			if j == 0 || j == 1 || j == 2 {
				value := strings.Trim(s.Text(), "	") // Remove tab
				fmt.Println(strings.Trim(value, " "))

				// value := s.Text()

				// if len(value) > 0 {
				// var count int
				// if j == 0 {

				// if value == "　　推銷費用" {
				// 	// for i := 0; i < len(value); i++ {
				// 	// 	fmt.Println(value[i], i)
				// 	// }
				// 	for k, v := range value {
				// 		fmt.Println(v, k, string(v))
				// 	}
				// }

				// if value == "　營業費用合計" {
				// 	fmt.Println("-----")
				// 	// for i := 0; i < len(value); i++ {
				// 	// 	fmt.Println(value[i], i)
				// 	// }
				// 	for k, v := range value {
				// 		fmt.Println(v, k, string(v))
				// 	}
				// }

				// count = countLeadingSpaces(value)
				// match, _ := regexp.MatchString(`^　`, s.Text())

				// fmt.Println(s.Text(), "regex : ", match, "count : ", count)
				// }

				// }
			}
		})
	})
}

func countLeadingSpaces(line string) int {
	// len(line) - len(strings.TrimLeft(line, "	"))
	fmt.Println(len(line), len(strings.TrimLeft(line, " ")), len(strings.Trim(line, " ")))
	return len(line) - len(strings.Trim(line, "	")) // Remove tab
}

// 資產負債表
func balanceSheet() {
	url := "https://mops.twse.com.tw/mops/web/ajax_t164sb03"
	method := "POST"

	payload := strings.NewReader("encodeURIComponent=1&step=1&firstin=1&off=1&queryName=co_id&inpuType=co_id&TYPEK=all&isnew=false&co_id=2330&year=110&season=2")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "jcsession=jHttpSession@2e3e0afd")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table.hasBorder tr").Each(func(i int, docSecondary *goquery.Selection) {
		if i < 4 {
			return
		}

		docSecondary.Find("td").Each(func(j int, s *goquery.Selection) {

			if j == 0 || j == 1 || j == 2 {
				// fmt.Println(strings.Replace(s.Text(), " ", "", -1))

				// fmt.Println(compressStr(s.Text()))
				// fmt.Println(deleteTailBlank(s.Text()))

				// strconv.Unquote()
				value := strings.Trim(s.Text(), "　　") // Remove tab
				fmt.Println(strings.Trim(value, " "))
			}
		})
	})
}

func deleteTailBlank(str string) string {
	spaceNum := 0
	for i := len(str) - 1; i >= 0; i-- { // 去除字符串尾部的所有空格
		if str[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	return str[:len(str)-spaceNum]
}

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

type requestInfo struct {
	action     string `form:"action"`
	seqNo      int    `form:"seq_no"`
	spokeTime  int    `form:"spoke_time"`
	spokeDate  int    `form:"spoke_date"`
	coId       int    `form:"co_id"`
	typeConfig string `form:"TYPEK"`
	year       int    `form:"year"`
}

type Data struct {
	number string
	title  string
	heat   string
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func decodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	// tr := simplifiedchinese.GB18030.NewDecoder()
	tr := simplifiedchinese.GBK.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

func companyStockList() {

	res, err := http.Get("https://isin.twse.com.tw/isin/class_main.jsp?owncode=&stockname=&isincode=&market=1&issuetype=1&industry_code=&Page=1&c")
	if err != nil {
		// handle error
	}
	defer res.Body.Close()

	// utfBody, err := iconv.NewReader(res.Body, "MS950", "utf-8")
	// if err != nil {
	// 	// handler error
	// }

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		// handler error
	}

	doc.Find("table.h4 tr").Each(func(i int, docSecondarys *goquery.Selection) {
		if i == 3 {
			// fmt.Println(docSecondarys)
			docSecondarys.Find("td").Each(func(j int, s *goquery.Selection) {

				//converts a  string from UTF-8 to gbk encoding.
				// fmt.Println(s.Text())
				gbkTitle, _ := decodeToGBK(s.Text())
				fmt.Println(gbkTitle)

				// fmt.Println(s.Text())
			})
		}
	})

	// url := "https://isin.twse.com.tw/isin/C_public.jsp?strMode=2"
	// method := "GET"

	// client := &http.Client{}
	// req, err := http.NewRequest(method, url, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// doc.Find("table.h4 tr").Each(func(i int, docSecondarys *goquery.Selection) {
	// 	if i == 3 {
	// 		// fmt.Println(docSecondarys)
	// 		docSecondarys.Find("td").Each(func(j int, s *goquery.Selection) {

	// 			//converts a  string from UTF-8 to gbk encoding.
	// 			fmt.Println(ConvertToString(s.Text(), "gbk", "utf-8"))

	// 			// fmt.Println(s.Text())
	// 		})
	// 	}
	// })

	// 爬取微博熱搜網頁
	// res, err := http.Get("https://s.weibo.com/top/summary")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode != 200 {
	// 	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// }
	// //將html生成goquery的Document
	// dom, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// var data []Data
	// // 篩選class為td-01的元素
	// dom.Find(".td-01").Each(func(i int, selection *goquery.Selection) {
	// 	data = append(data, Data{number: selection.Text()})
	// })
	// // 篩選class為td-02的元素下的a元素
	// dom.Find(".td-02>a").Each(func(i int, selection *goquery.Selection) {
	// 	data[i].title = selection.Text()
	// })
	// // 篩選class為td-02的元素下的span元素
	// dom.Find(".td-02>span").Each(func(i int, selection *goquery.Selection) {
	// 	data[i].heat = selection.Text()
	// })
	// fmt.Println(data)
}

func queryCompanyImportantMessageDailyList() {
	url := "https://mops.twse.com.tw/mops/web/ajax_t05sr01_1"
	method := "POST"

	payload := strings.NewReader("encodeURIComponent=1&TYPEK=all&step=0")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "jcsession=jHttpSession@79b80ceb")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(doc.Find("table.hasBorder tr input").Length())

	// [
	// 	document.fm_t05sr01_1.SEQ_NO.value='1';
	// 	document.fm_t05sr01_1.SPOKE_TIME.value='70002';
	// 	document.fm_t05sr01_1.SPOKE_DATE.value='20210825';
	// 	document.fm_t05sr01_1.COMPANY_ID.value='5392';
	// 	document.fm_t05sr01_1.skey.value='5392202108251';
	// 	openWindow(this.form ,'');
	// ]

	// document.t05st01_fm.action='ajax_t05st01';
	// document.t05st01_fm.seq_no.value='2';
	// document.t05st01_fm.spoke_time.value='134241';
	// document.t05st01_fm.spoke_date.value='20210825';
	// document.t05st01_fm.co_id.value='6725';
	// document.t05st01_fm.TYPEK.value='pub';
	// openWindow(this.form ,'')

	doc.Find("table.hasBorder tr input").Each(func(i int, docSecondarys *goquery.Selection) {
		onclickValue := docSecondarys.Get(0).Attr[2].Val
		dataInfo := strings.Split(onclickValue, "document.fm_t05sr01_1.")

		var trimVal string
		request := &requestInfo{}
		for _, val := range dataInfo {
			request.year = 110
			request.action = "ajax_t05st01"
			request.typeConfig = "sii"

			switch {
			case strings.Contains(val, "SEQ_NO"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "SEQ_NO.value='", "", 1)
				request.seqNo, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "SPOKE_TIME"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "SPOKE_TIME.value='", "", 1)
				request.spokeTime, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "SPOKE_DATE"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "SPOKE_DATE.value='", "", 1)
				request.spokeDate, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "COMPANY_ID"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "COMPANY_ID.value='", "", 1)
				request.coId, _ = strconv.Atoi(trimVal)
			default:
			}
		}

		fmt.Println(request)
		fmt.Println("-------------")
	})

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))
}

func queryCompanyImportantMessageList() {
	url := "https://mops.twse.com.tw/mops/web/ajax_t05st01"
	method := "POST"

	payload := strings.NewReader("encodeURIComponent=1&step=1&firstin=1&off=1&queryName=co_id&inpuType=co_id&TYPEK=all&co_id=2330&year=110")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "jcsession=jHttpSession@6ccbffc6")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// var requestInfo map[string]map[string]interface{}
	// var request requestInfo

	doc.Find(".hasBorder tr input").Last().Each(func(i int, docSecondarys *goquery.Selection) {
		onclickValue := docSecondarys.Get(0).Attr[2].Val
		dataInfo := strings.Split(onclickValue, "document.t05st01_fm.")

		var trimVal string
		request := &requestInfo{}
		for _, val := range dataInfo {
			request.year = 110

			switch {
			case strings.Contains(val, "action"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "action='", "", 1)
				request.action = trimVal
			case strings.Contains(val, "seq_no"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "seq_no.value='", "", 1)
				request.seqNo, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "spoke_time"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "spoke_time.value='", "", 1)
				request.spokeTime, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "spoke_date"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "spoke_date.value='", "", 1)
				request.spokeDate, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "co_id"):
				trimVal = strings.Replace(val, "';", "", 1)
				trimVal = strings.Replace(trimVal, "co_id.value='", "", 1)
				request.coId, _ = strconv.Atoi(trimVal)
			case strings.Contains(val, "TYPEK"):
				trimVal = strings.Replace(val, "';openWindow(this.form ,'');", "", 1)
				trimVal = strings.Replace(trimVal, "TYPEK.value='", "", 1)
				request.typeConfig = trimVal
			default:
			}
		}

		fmt.Println(request)

		// request.companyImportantMessage()

		fmt.Println("-------------")
	})

	// doc.Find(".hasBorder tr").Each(func(i int, docSecondarys *goquery.Selection) {
	// 	// fmt.Println(docSecondarys.Text())
	// 	if i == 2 {
	// 		fmt.Println(docSecondarys.Text())
	// 		// input := docSecondarys.Find("input").Text()
	// 		// fmt.Println(input)
	// 		// fmt.Println("-------------")
	// 	}
	// })
}

func (ri *requestInfo) companyImportantMessage() {
	url := fmt.Sprintf("https://mops.twse.com.tw/mops/web/%s", ri.action)
	method := "POST"

	formDate := fmt.Sprintf("seq_no=%d&spoke_time=%d&spoke_date=%d&co_id=%d&TYPEK=%s&year=%d&month=all&e_month=all&step=2&off=1&firstin=true&b_date=&e_date=&type=",
		ri.seqNo,
		ri.spokeTime,
		ri.spokeDate,
		ri.coId,
		ri.typeConfig,
		ri.year,
	)
	payload := strings.NewReader(formDate)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "jcsession=jHttpSession@6ccbffc6")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table.noBorder .compName").Each(func(i int, s *goquery.Selection) {
		title := s.Find("b").Text()
		fmt.Printf("%s\n", title)
		fmt.Println("-------------")
	})

	var key string
	var value string

	doc.Find("table.hasBorder tr").Each(func(i int, docSecondary *goquery.Selection) {

		docSecondary.Find("td").Each(func(j int, s *goquery.Selection) {
			j++

			if s.HasClass("tblHead") {
				key = s.ToggleClass("tblHead").Text()
			} else if s.HasClass("odd") {
				value = s.ToggleClass("odd").Text()
			}

			if j%2 == 0 {
				fmt.Println(key, value)
				fmt.Println("-------------")
			}
		})
	})
}

func companyImportantMessage() {
	url := "https://mops.twse.com.tw/mops/web/ajax_t05st01"
	method := "POST"

	payload := strings.NewReader("seq_no=1&spoke_time=170309&spoke_date=20210816&co_id=2330&TYPEK=sii&year=110&month=all&e_month=all&step=2&off=1&firstin=true&b_date=&e_date=&type=")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "jcsession=jHttpSession@6ccbffc6")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table.noBorder .compName").Each(func(i int, s *goquery.Selection) {
		title := s.Find("b").Text()
		fmt.Printf("%s\n", title)
		fmt.Println("-------------")
	})

	var key string
	var value string

	doc.Find("table.hasBorder tr").Each(func(i int, docSecondary *goquery.Selection) {

		docSecondary.Find("td").Each(func(j int, s *goquery.Selection) {
			j++

			if s.HasClass("tblHead") {
				key = s.ToggleClass("tblHead").Text()
			} else if s.HasClass("odd") {
				value = s.ToggleClass("odd").Text()
			}

			if j%2 == 0 {
				fmt.Println(key, value)
				fmt.Println("-------------")
			}
		})

		// docSecondary.Find(".tblHead").Each(func(j int, s *goquery.Selection) {
		// 	title := s.Text()
		// 	fmt.Println(title)
		// })

		// docSecondary.Find(".odd").Each(func(j int, s *goquery.Selection) {
		// 	content := s.Text()
		// 	fmt.Println(content)
		// })
		// fmt.Println("-------------")
	})
}
