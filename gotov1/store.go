package gotov1

import (
	"sync"
)

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, preset := s.urls[key]
	if preset {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
	return ""
}

// 结构体中包含map类型的字段，使用前必须先用make初始化。
// 在Go中创建一个结构体实例，一般是通过定义一个前缀为New，能返回该类型已初始化实例的函数（通常是指向实例的指针）
func NewURLStore() *URLStore {
	//锁无需特别指明初始化，这是Go创建结构体实例的惯例
	return &URLStore{urls: make(map[string]string)}
}
