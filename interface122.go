package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"strconv"
	"math/rand"
	"time"
	//"net/http/cookiejar"
	//"net/url"
	//url2 "net/url"
	"os"
	url2 "net/url"
)

var userAgent=[]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r=rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomUserAgent()map[string][]string  {
	return map[string][]string{"User-Agent":[]string{userAgent[r.Intn(len(userAgent))]},
		"Accept":{"application/json, text/javascript, */*; q=0.01"},
		"Accept-Encoding":{"gzip, deflate"},
		"Content-Type": {"application/json;charset=UTF-8"},
		"Transfer-Encoding": {"chunked"},
		"Connection": {"keep-alive"},
		"X-Frame-Options": {"SAMEORIGIN"},
		"Accept-Language":{"zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7"},
		"Host":{"sd.122.gov.cn"},
		"Origin":{"http://sd.122.gov.cn"},
		"Referer":{"http://sd.122.gov.cn/views/inquiry.html?q=j"},
		"X-Requested-With":{"XMLHttpRequest"},
	}
}
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

func GetCurtime() string{
	curtime := time.Now()
	curtime=time.Now().UTC()
	return curtime.Format(RFC1123)[:26]+"GMT"

}
var client=&http.Client{}

func saveImg(url string,headers map[string][]string) map[string][]string {
	resp,err:=client.Get(url)
	//var temp_cookies =resp.Cookies()
	for i:=range headers{
		resp.Header.Set(i,headers[i][0])
	}
	if err!=nil{
		fmt.Println("请求出错!!")
	}
	defer resp.Body.Close()
	f,err:=os.Create("./code.jpg")
	if err!=nil{
		fmt.Println("文件创建出错!!")
	}
	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("文件写入出错!!")
	}
	f.Write(b)
	f.Close()
	resp.Header.Del("Date")
	return resp.Header
}

func post(url string,headers map[string][]string,code string)  {
	//resp,err:=http.Get(url)
	var postdata=url2.Values{}
	//postdata["hpzl"]=[]string{"02"}
	//postdata["hphm1b"]=[]string{"AQV901"}
	//postdata["hphm"]=[]string{"鲁AQV901"}
	//postdata["fdjh"]=[]string{"220548"}
	//postdata["qm"]=[]string{"wf"}
	//postdata["page"]=[]string{"1"}
	//postdata["captcha"]=[]string{code}
	postdata.Add("hpzl","02")
	postdata.Add("hphm1b","AQV901")
	postdata.Add("hphm","鲁AQV901")
	postdata.Add("fdjh","220548")
	postdata.Add("qm","wf")
	postdata.Add("page","1")
	postdata["captcha"]=[]string{code}
	resp,err:=client.PostForm(url,postdata)
	fmt.Println(resp.Header)

	//for i:=range headers{
	//	resp.Header.Set(i,headers[i][0])
	//}
	//resp.Header.Set("Data",GetCurtime())
	resp.Header.Add("Referer","http://sd.122.gov.cn/views/inquiry.html?q=j")
	resp.Header.Add("User-Agent",userAgent[r.Intn(len(userAgent))])
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	//fmt.Println(resp.Header)
	if err!=nil{
		fmt.Println("post error")
	}
	defer resp.Body.Close()
	n,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(n))
}

type Cp struct {
	address string
	url string
	cp string
	sftjb int
}

type addressCp struct{
	ac Cp
}
func main()  {
	var code string
	headers:=GetRandomUserAgent()
	imgurl:="http://sd.122.gov.cn/captcha?nocache=1519907612436"
	head:=saveImg(imgurl,headers)
	fmt.Print("输入验证码:")
	fmt.Scanln(&code)
	var postdata map[string][]string
	postdata=make(map[string][]string)
	postdata["hpzl"]=[]string{"02"}
	postdata["hphm1b"]=[]string{"AQV901"}
	postdata["hphm"]=[]string{"鲁AQV901"}
	postdata["fdjh"]=[]string{"220548"}
	postdata["qm"]=[]string{"wf"}
	postdata["page"]=[]string{"1"}
	postdata["captcha"]=[]string{code}
	//fmt.Println(postdata,headers)
	post("http://sd.122.gov.cn/m/publicquery/vio", head, code)
	//data:=map[string]map[string]string{
	//	"000000": {"address": "\u516c\u5b89\u90e8", "url": "http://gab.122.gov.cn", "cp": "\u90e8O", "sftjb": "1"}, "110000": {"address": "\u5317\u4eac", "url": "http://bj.122.gov.cn", "cp": "\u4eacA", "sftjb": "1"}, "120000": {"address": "\u5929\u6d25", "url": "http://tj.122.gov.cn/", "cp": "\u6d25O", "sftjb": "1"}, "130000": {"address": "\u6cb3\u5317", "url": "http://he.122.gov.cn", "cp": "\u5180O", "sftjb": "1"},"140000": {"address": "\u5c71\u897f", "url": "http://sx.122.gov.cn", "cp": "\u664bO", "sftjb": "1"}, "150000": {"address": "\u5185\u8499\u53e4", "url": "http://nm.122.gov.cn", "cp": "\u8499O", "sftjb": "1"}, "210000": {"address": "\u8fbd\u5b81", "url": "http://ln.122.gov.cn", "cp": "\u8fbdO", "sftjb": "1"}, "220000": {"address": "\u5409\u6797", "url": "http://jl.122.gov.cn", "cp": "\u5409O", "sftjb": "1"}, "230000": {"address": "\u9ed1\u9f99\u6c5f", "url": "http://hl.122.gov.cn", "cp": "\u9ed1O", "sftjb": "1"}, "310000": {"address": "\u4e0a\u6d77", "url": "http://sh.122.gov.cn", "cp": "\u6caaO", "sftjb": "1"}, "320000": {"address": "\u6c5f\u82cf", "url": "http://www.jscd.gov.cn/", "cp": "\u82cfO", "sftjb": "0"}, "320100": {"address": "\u6c5f\u82cf\u5357\u4eac", "url": "http://nkg.122.gov.cn", "cp": "\u82cfA", "sftjb": "1"}, "320200": {"address": "\u6c5f\u82cf\u65e0\u9521", "url": "http://wux.122.gov.cn", "cp": "\u82cfB", "sftjb": "1"}, "320300": {"address": "\u6c5f\u82cf\u5f90\u5dde", "url": "http://xuz.122.gov.cn", "cp": "\u82cfC", "sftjb": "1"}, "320400": {"address": "\u6c5f\u82cf\u5e38\u5dde", "url": "http://czx.122.gov.cn", "cp": "\u82cfD", "sftjb": "1"}, "320500": {"address": "\u6c5f\u82cf\u82cf\u5dde", "url": "http://szv.122.gov.cn", "cp": "\u82cfE", "sftjb": "1"}, "320600": {"address": "\u6c5f\u82cf\u5357\u901a", "url": "http://ntg.122.gov.cn", "cp": "\u82cfF", "sftjb": "1"}, "320700": {"address": "\u6c5f\u82cf\u8fde\u4e91\u6e2f", "url": "http://lyg.122.gov.cn", "cp": "\u82cfG", "sftjb": "1"}, "320800": {"address": "\u6c5f\u82cf\u6dee\u5b89", "url": "http://has.122.gov.cn", "cp": "\u82cfH", "sftjb": "1"}, "320900": {"address": "\u6c5f\u82cf\u76d0\u57ce", "url": "http://ynz.122.gov.cn", "cp": "\u82cfJ", "sftjb": "1"}, "321000": {"address": "\u6c5f\u82cf\u626c\u5dde", "url": "http://yzo.122.gov.cn", "cp": "\u82cfK", "sftjb": "1"}, "321100": {"address": "\u6c5f\u82cf\u9547\u6c5f", "url": "http://zhe.122.gov.cn", "cp": "\u82cfL", "sftjb": "1"}, "321200": {"address": "\u6c5f\u82cf\u6cf0\u5dde", "url": "http://tzs.122.gov.cn", "cp": "\u82cfM", "sftjb": "1"}, "321300": {"address": "\u6c5f\u82cf\u5bbf\u8fc1", "url": "http://suq.122.gov.cn", "cp": "\u82cfN", "sftjb": "1"}, "330000": {"address": "\u6d59\u6c5f", "url": "http://stgs.zjsgat.gov.cn/website/index.aspx", "cp": "\u6d59O", "sftjb": "0"}, "330100": {"address": "\u6d59\u6c5f\u676d\u5dde", "url": "http://hgh.122.gov.cn", "cp": "\u6d59A", "sftjb": "1"}, "330200": {"address": "\u6d59\u6c5f\u5b81\u6ce2", "url": "http://ngb.122.gov.cn", "cp": "\u6d59B", "sftjb": "1"}, "330300": {"address": "\u6d59\u6c5f\u6e29\u5dde", "url": "http://wnz.122.gov.cn", "cp": "\u6d59C", "sftjb": "1"}, "330400": {"address": "\u6d59\u6c5f\u5609\u5174", "url": "http://jix.122.gov.cn", "cp": "\u6d59F", "sftjb": "1"}, "330500": {"address": "\u6d59\u6c5f\u6e56\u5dde", "url": "http://hzh.122.gov.cn", "cp": "\u6d59E", "sftjb": "1"}, "330600": {"address": "\u6d59\u6c5f\u7ecd\u5174", "url": "http://sxg.122.gov.cn", "cp": "\u6d59D", "sftjb": "1"}, "330700": {"address": "\u6d59\u6c5f\u91d1\u534e", "url": "http://jha.122.gov.cn", "cp": "\u6d59G", "sftjb": "1"}, "330800": {"address": "\u6d59\u6c5f\u8862\u5dde", "url": "http://quz.122.gov.cn", "cp": "\u6d59H", "sftjb": "1"}, "330900": {"address": "\u6d59\u6c5f\u821f\u5c71", "url": "http://zos.122.gov.cn", "cp": "\u6d59L", "sftjb": "1"}, "331000": {"address": "\u6d59\u6c5f\u53f0\u5dde", "url": "http://tzz.122.gov.cn", "cp": "\u6d59J", "sftjb": "1"}, "331100": {"address": "\u6d59\u6c5f\u4e3d\u6c34", "url": "http://lss.122.gov.cn", "cp": "\u6d59K", "sftjb": "1"}, "335000": {"address": "\u6d59\u6c5f\u9ad8\u901f", "url": "http://zjgs.122.gov.cn", "cp": "\u6d59Z", "sftjb": "1"}, "340000": {"address": "\u5b89\u5fbd", "url": "http://ah.122.gov.cn", "cp": "\u7696O", "sftjb": "1"}, "350000": {"address": "\u798f\u5efa", "url": "http://fj.122.gov.cn", "cp": "\u95fdO", "sftjb": "1"}, "360000": {"address": "\u6c5f\u897f", "url": "http://jx.122.gov.cn", "cp": "\u8d63O", "sftjb": "1"}, "370000": {"address": "\u5c71\u4e1c", "url": "http://sd.122.gov.cn", "cp": "\u9c81O", "sftjb": "1"}, "410000": {"address": "\u6cb3\u5357", "url": "http://ha.122.gov.cn", "cp": "\u8c6bO", "sftjb": "1"}, "420000": {"address": "\u6e56\u5317", "url": "http://hb.122.gov.cn", "cp": "\u9102O", "sftjb": "1"}, "430000": {"address": "\u6e56\u5357", "url": "http://hn.122.gov.cn", "cp": "\u6e58O", "sftjb": "1"}, "440000": {"address": "\u5e7f\u4e1c", "url": "http://gd.122.gov.cn", "cp": "\u7ca4O", "sftjb": "1"}, "450000": {"address": "\u5e7f\u897f\u58ee\u65cf\u81ea\u6cbb\u533a", "url": "http://gx.122.gov.cn", "cp": "\u6842O", "sftjb": "1"}, "460000": {"address": "\u6d77\u5357", "url": "http://hi.122.gov.cn", "cp": "\u743cO", "sftjb": "1"}, "500000": {"address": "\u91cd\u5e86", "url": "http://cq.122.gov.cn", "cp": "\u6e1dA", "sftjb": "1"}, "510000": {"address": "\u56db\u5ddd", "url": "http://sc.122.gov.cn", "cp": "\u5dddO", "sftjb": "1"}, "520000": {"address": "\u8d35\u5dde", "url": "http://gz.122.gov.cn", "cp": "\u8d35O", "sftjb": "1"}, "530000": {"address": "\u4e91\u5357", "url": "http://yn.122.gov.cn", "cp": "\u4e91O", "sftjb": "1"}, "540000": {"address": "\u897f\u85cf", "url": "http://xz.122.gov.cn", "cp": "\u85cfO", "sftjb": "1"}, "610000": {"address": "\u9655\u897f", "url": "http://sn.122.gov.cn", "cp": "\u9655O", "sftjb": "1"}, "620000": {"address": "\u7518\u8083", "url": "http://gs.122.gov.cn", "cp": "\u7518O", "sftjb": "1"}, "630000": {"address": "\u9752\u6d77", "url": "http://qh.122.gov.cn", "cp": "\u9752O", "sftjb": "1"}, "640000": {"address": "\u5b81\u590f", "url": "http://nx.122.gov.cn", "cp": "\u5b81O", "sftjb": "1"}, "650000": {"address": "\u65b0\u7586", "url": "http://xj.122.gov.cn", "cp": "\u65b0O", "sftjb": "1"}}
	//for i:=range data{
	//	fmt.Println(data[i]["address"])
	//
	//}
}