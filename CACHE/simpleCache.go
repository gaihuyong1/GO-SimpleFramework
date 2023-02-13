package cache

import (
	"fmt"
	"sync"
)

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers PeerPick
}

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu sync.Mutex
	groups=make(map[string]*Group)
)

func NewGroup(name string,cacheBytes int64,getter Getter)*Group{
	if getter==nil{
		panic("getter can't be nil")
	}
	mu.Lock()
	defer mu.Unlock()
	group:=&Group{
		name: name,
		getter:getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name]=group

	return group
}

func GetGroup(name string)*Group{
	mu.Lock()
	group:=groups[name]
	mu.Unlock()
	return group
}

func (g *Group)Get(name string)(ByteView,error){
	if name==""{
		return ByteView{},fmt.Errorf("name can't be nil")
	}
	if v,ok:=g.mainCache.get(name);ok{
		return v,nil
	}

	return g.load(name)
}

func (g *Group) load(key string) (value ByteView, err error) {
	if g.peers!=nil{
		if peer,ok:=g.peers.PickPeer(key);ok{
			if value,err:=g.getFromPeer(peer,key);err==nil{
				return value,nil
			}
		}
	}
	return g.getLocal(key)
}

func (g *Group)getLocal(name string)(ByteView,error){
	bytes,err:=g.getter.Get(name)
	if err!=nil{
		return ByteView{},err
	}
	value:=ByteView{b:cloneBytes(bytes)}
	g.populateCache(name,value)
	return value,nil
}

func (g *Group)populateCache(name string,value ByteView){
	g.mainCache.add(name,value)
}

func (g *Group)RegisterPeers(peers PeerPick){
	if g.peers!=nil{
		panic("already have peers")
	}
	g.peers=peers
}

func (g *Group)getFromPeer(peer PeerGetter,key string)(ByteView,error){
	bytes,err:=peer.Get(g.name,key)
	if err!=nil{
		return ByteView{},err
	}

	return ByteView{b:bytes},nil
}