package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item //用于存数据
	RequestProcessor Processor //用于Worker rpc,是work()函数
}
//src/crawler/engine/worker.go
type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)

	WorkerChan() chan Request
	Run()
	//ConfigureMasterWorkerChan(chan Request)
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make( chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in,out)
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//要尽快处理
			//log.Printf("Got item #%d: %v",itemCount, item)
			//itemCount++
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false

}

//func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	//in := make(chan Request)
	go func() {
		for {
			// tell scheduler I'm ready
			ready.WorkerReady(in)
			request := <-in
			// 换成call rpc, 这里是分布式的关键//17.8
			result, e := e.RequestProcessor(request)
			//result, e := Worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()

}
