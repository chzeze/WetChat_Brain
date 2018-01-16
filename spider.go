package wechat_brain

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"github.com/coreos/goproxy"
	"time"
)

var (
	_spider = newSpider()
)

type spider struct {
	proxy *goproxy.ProxyHttpServer
}

func Run(port string) {
	_spider.Init()
	_spider.Run(port)
}

func Close() {
	memoryDb.Close()
}

func newSpider() *spider {
	sp := &spider{}
	sp.proxy = goproxy.NewProxyHttpServer()
	sp.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	return sp
}

func (s *spider) Run(port string) {
	log.Println("server will at port:" + port)
	log.Fatal(http.ListenAndServe(":"+port, s.proxy))
}

func (s *spider) Init() {
	requestHandleFunc := func(request *http.Request, ctx *goproxy.ProxyCtx) (req *http.Request, resp *http.Response) {
		req = request
		if ctx.Req.URL.Path == `/question/fight/findQuiz` {
			bs, _ := ioutil.ReadAll(req.Body)
			req.Body = ioutil.NopCloser(bytes.NewReader(bs))
			handleQuestionReq(bs)
		} else if ctx.Req.URL.Path == `/question/fight/choose` {
			bs, _ := ioutil.ReadAll(req.Body)
			req.Body = ioutil.NopCloser(bytes.NewReader(bs))
			handleChooseReq(bs)
		} else if ctx.Req.URL.Hostname() == `abc.com` {
			resp = new(http.Response)
			resp.StatusCode = 200
			resp.Header = make(http.Header)
			resp.Header.Add("Content-Disposition", "attachment; filename=ca.crt")
			resp.Header.Add("Content-Type", "application/octet-stream")
			resp.Body = ioutil.NopCloser(bytes.NewReader(goproxy.CA_CERT))
		}
		return
	}
	responseHandleFunc := func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		if resp == nil {
			return resp
		}
		if ctx.Req.URL.Path == `/question/fight/findQuiz` {
			bs, _ := ioutil.ReadAll(resp.Body)
			bsNew,ans := handleQuestionResp(bs)
			
			resp.Body = ioutil.NopCloser(bytes.NewReader(bsNew))
			go loop(ans)
			

		} else if ctx.Req.URL.Path == `/question/fight/choose` {
			bs, _ := ioutil.ReadAll(resp.Body)
			resp.Body = ioutil.NopCloser(bytes.NewReader(bs))
			go handleChooseResponse(bs)
		}else if ctx.Req.URL.Path == `/question/fight/fightResult` {
			go continueFight()//继续挑战
		}else{
			log.Printf(ctx.Req.URL.Path)
			//
		}
		return resp
	}
	s.proxy.OnResponse().DoFunc(responseHandleFunc)
	s.proxy.OnRequest().DoFunc(requestHandleFunc)
}

func continueFight() {
	log.Printf("继续挑战")
	time.Sleep(time.Millisecond * 5000)
	var err error
	touchX, touchY := strconv.Itoa(550), strconv.Itoa(1400)
	log.Printf("开始点击")
	_, err = exec.Command("adb","shell", "input", "swipe", touchX, touchY,touchX, touchY).Output()
	if err != nil {
		log.Fatal("模拟点击失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
	}

	time.Sleep(time.Millisecond * 3000)
	touchX, touchY = strconv.Itoa(550), strconv.Itoa(1600)
	log.Printf("开始点击")
	_, err = exec.Command("adb","shell", "input", "swipe", touchX, touchY,touchX, touchY).Output()
	if err != nil {
		log.Fatal("模拟点击失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
	}
}

func loop(ans int) {
	if ans==0{
		ans=1
	}
	log.Printf("正确答案选项：%d \n",ans)
	time.Sleep(time.Millisecond * 3000)
	var err error
	touchX, touchY := strconv.Itoa(550), strconv.Itoa(800+200*ans)
	log.Printf("开始点击")
	_, err = exec.Command("adb","shell", "input", "swipe", touchX, touchY,touchX, touchY).Output()
	if err != nil {
		log.Fatal("模拟点击失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
	}
}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}
