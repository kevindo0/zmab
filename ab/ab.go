package ab

import (
	"fmt"
	"sync"
	"time"
)

// 压测执行单元
type ABOne struct {
	Path        string `json:"path"`
	Method      string
	Requests    int
	Concurrency int
}

// json文件格式
type ABOnes struct {
	Host string `json:"host"`
	ABS  []ABOne
}

func (b *ABOnes) Init() []ABOne {
	var opt []ABOne
	for _, v := range b.ABS {
		fmt.Println("cell:", v)
		v.Path = fmt.Sprintf("%s%s", b.Host, v.Path)
		opt = append(opt, v)
	}
	return opt
}

func Consumer(res chan Result) {
	for {
		val := <-res
		fmt.Println("con:", val)
	}
}

func StartCell(i int, wg *sync.WaitGroup, v *ABOne, res chan Result) {
	start := time.Now().UnixNano()
	resp, err := Get(v.Method, v.Path)
	end := time.Now().UnixNano()
	time.Sleep(time.Second)
	r := Result{
		StartTime: start,
		EndTime:   end,
		Interval:  end - start,
		Response:  resp,
		Error:     err,
	}
	res <- r
	wg.Done()
}

func Start(s []ABOne) {
	var wg sync.WaitGroup
	res := make(chan Result, 10)
	go func() {
		Consumer(res)
	}()
	for _, v := range s {
		request := v.Requests
		for {
			num := v.Concurrency
			if v.Concurrency > request {
				num = request
			}
			wg.Add(num)
			for i := 0; i < num; i++ {
				go StartCell(i, &wg, &v, res)
			}
			wg.Wait()
			request -= num
			if request <= 0 {
				break
			}
		}
	}
}
