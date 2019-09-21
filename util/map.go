package util

import "sync"

// MergeMap 合并两个map[int]int
func MergeMap(args ...map[int]int) map[int]int {
	mergedMap := map[int]int{}
	for _, m := range args {
		for k, v := range m {
			if _, exists := mergedMap[k]; exists {
				mergedMap[k] += v
			} else {
				mergedMap[k] = v
			}
		}
	}
	return mergedMap
}

// GetMapMinValue 获取map[int]int中最小值, 并返回所有key的集合
// value约定是大于0的
func GetMapMinValue(m map[int]int) (int, []int) {
	// value到keys的对应关系
	keys := []int{}
	// 最小值
	minValue := -1
	for k, v := range m {
		if minValue == -1 || v < minValue {
			minValue = v
			keys = []int{k}
		} else if v == minValue {
			keys = append(keys, k)
		}
	}
	return minValue, keys
}

// GetMapMaxValue 获取map[int]int中最大值, 并返回所有key的集合
// value约定是大于0的
func GetMapMaxValue(m map[int]int) (int, []int) {
	// value到keys的对应关系
	keys := []int{}
	// 最大值
	maxValue := -1
	for k, v := range m {
		if v > maxValue {
			maxValue = v
			keys = []int{k}
		} else if v == maxValue {
			keys = append(keys, k)
		}
	}
	return maxValue, keys
}

// GetMapValues 获取map[int]int结构的所有value
// 返回结果去重
func GetMapValues(m map[int]int) []int {
	values := make([]int, 0)
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

//========================================================option

type Map struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

func (m *Map) init() {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}
}

func (m *Map) UnsafeGet(key interface{}) interface{} {
	if m.m == nil {
		return nil
	} else {
		return m.m[key]
	}
}

func (m *Map) Get(key interface{}) interface{} {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeGet(key)
}

func (m *Map) UnsafeSet(key interface{}, value interface{}) {
	m.init()
	m.m[key] = value
}

func (m *Map) Set(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeSet(key, value)
}

func (m *Map) TestAndSet(key interface{}, value interface{}) interface{} {
	m.Lock()
	defer m.Unlock()

	m.init()

	if v, ok := m.m[key]; ok {
		return v
	} else {
		m.m[key] = value
		return nil
	}
}

func (m *Map) UnsafeDel(key interface{}) {
	m.init()
	delete(m.m, key)
}

func (m *Map) Del(key interface{}) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeDel(key)
}

func (m *Map) UnsafeLen() int {
	if m.m == nil {
		return 0
	} else {
		return len(m.m)
	}
}

func (m *Map) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeLen()
}

func (m *Map) UnsafeRange(f func(interface{}, interface{})) {
	if m.m == nil {
		return
	}
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *Map) RLockRange(f func(interface{}, interface{})) {
	m.RLock()
	defer m.RUnlock()
	m.UnsafeRange(f)
}

func (m *Map) LockRange(f func(interface{}, interface{})) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeRange(f)
}

//========================================================option

//=======================================================saftmap

type safeMap struct {
	m map[string]string
	l *sync.RWMutex
}

func (s *safeMap) Set(key string, value string) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[key] = value
}

func (s *safeMap) Get(key string) string {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.m[key]
}

func newSafeMap() *safeMap {
	return &safeMap{l: new(sync.RWMutex), m: make(map[string]string)}
}

//=======================================================saftmap
