package repo

import (
	"TinyURL/entity"
	"TinyURL/logic"

	"sync"
	"time"
)

/**
 * 类似于 Service 层
 */

// MemoryRepo 存数据，一个计数器墨子自增id，一把锁保证并发安全
type MemoryRepo struct {
	mux    sync.RWMutex
	data   map[uint64]*entity.URLMapping
	nextID uint64
}

// NewMemoryRepo 构造函数
func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		data: make(map[uint64]*entity.URLMapping),
	}
}

// Save 数据库层的操作，拿到长链URL，存入数据库并返回ID
func (r *MemoryRepo) Save(longURL string) (*entity.URLMapping, error) {
	// 1. 加锁（写锁）
	r.mux.Lock()
	defer r.mux.Unlock()
	// 2. nextID + 1
	r.nextID++
	// 3. 构造 URLMapping 对象
	now := time.Now()
	mapping := &entity.URLMapping{
		ID:        r.nextID,
		LongURL:   longURL,
		CreatedAt: now,
	}

	// 4. 存进 data map
	r.data[mapping.ID] = mapping
	// 5. 返回
	return mapping, nil
}

// FindByID 数据库层已经获得到了 id，直接查，返回长链
func (r *MemoryRepo) FindByID(id uint64) (*entity.URLMapping, error) {
	// 1. 加读锁 (RLock)
	r.mux.RLock()
	defer r.mux.RUnlock()

	// 2. 从 r.data 里查，判断 key 是否存在
	value, ok := r.data[id]

	if !ok {
		// key 不存在
		return nil, logic.ErrKeyNotFound
	}

	// 3. 存在就返回 mapping，不存在就返回 (nil, ErrKeyNotFound)
	return value, nil
}
