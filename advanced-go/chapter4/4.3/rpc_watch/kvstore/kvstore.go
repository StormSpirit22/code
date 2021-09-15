package kvstore

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// KVStoreService 一个简单的内存KV数据库
type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

// Set 不需要返回，所以 reply 用一个空结构
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	// 如果要修改某个 key，将对该 key 执行过滤器里的所有函数
	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

// Watch 方法用于在某个限定的时间内，观察当有key变化时将key作为返回值返回。如果超过时间后依然没有key被修改，则返回超时的错误。
// Watch的实现中，用唯一的id表示每个Watch调用，然后根据id将自身对应的过滤器函数注册到p.filter列表。
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	fmt.Println("start watching, id: ", id)
	p.mu.Lock()
	// 在 p.filter 里注册该 watch 函数的 id，该函数作用就是将 key 放入到缓存的 channel 中供后面 keyChanged 返回。
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	for {
		select {
		case <-time.After(time.Duration(timeoutSecond) * time.Second):
			fmt.Println(id, "id timeout")
			return fmt.Errorf("timeout")
		case key := <-ch:
			fmt.Println(id, key)
			*keyChanged = key
			return nil
		}
	}


	return nil
}