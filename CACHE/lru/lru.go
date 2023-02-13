package lru

import "container/list"

type Cache struct {
	maxBytes int64
	usedBytes   int64
	list     *list.List
	cache map[string]*list.Element
	OnEvicted func(key string,value Value)
}

type entry struct{
	key string
	value Value
}

type Value interface{
  Len() int
}

func New(maxBytes int64,onEvicted func(string,Value))*Cache{
	return &Cache{
		maxBytes: maxBytes,
		list:list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache)Get(key string)(value Value,ok bool){
	if element,ok:=c.cache[key];ok{
		c.list.MoveToFront(element)
		kv:=element.Value.(*entry)
		return kv.value,true
	}
	return
}

func (c *Cache)Add(key string,value Value){
	if element,ok:=c.cache[key];ok{
		c.list.MoveToFront(element)
		kv:=element.Value.(*entry)
		c.usedBytes+=int64(value.Len())-int64(kv.value.Len())
		kv.value=value
	}else{
		element:=c.list.PushFront(&entry{key,value})
		c.cache[key]=element
		c.usedBytes+=int64(value.Len())-int64(value.Len())
	}
	for c.maxBytes!=0&&c.maxBytes<c.usedBytes{
		c.RemoveOldElement()
	}
}

func (c *Cache)RemoveOldElement(){
	element:=c.list.Back()
	if element!=nil{
		c.list.Remove(element)
		kv:=element.Value.(*entry)
		delete(c.cache,kv.key)
		c.usedBytes-=int64(len(kv.key))+int64(kv.value.Len())
		if c.OnEvicted!=nil{
			c.OnEvicted(kv.key,kv.value)
		}
	}
}

func(c *Cache)Len()int{
	return c.list.Len()
}