package huanwuji

type T interface {
}

type Processor interface {
	Subscriber
	Publisher
}

type Publisher interface {
	subscribe(s Subscriber)
}

type Subscriber interface {
	onSubscribe(s Subscription)
	onError(error interface{})
	onComplete()
}

type Subscription interface {
	request(n int64)
	cancel()
}