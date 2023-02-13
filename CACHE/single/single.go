package single

import "sync"

type call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

type Group struct{
	mu sync.Mutex
	callMap map[string]*call
}

func (g *Group)Execute(key string,fn func()(interface{},error))(interface{},error){
	g.mu.Lock()
	if g.callMap==nil{
		g.callMap=make(map[string]*call)
	}
	if call,ok:=g.callMap[key];ok{
		g.mu.Unlock()
		call.wg.Wait()
		return call.val,call.err
	}
	call:=new(call)
	call.wg.Add(1)
	g.callMap[key]=call
	g.mu.Unlock()
	
	call.val,call.err=fn()
	call.wg.Done()

	g.mu.Lock()
	delete(g.callMap,key)
	g.mu.Unlock()

	return call.val,call.err
}