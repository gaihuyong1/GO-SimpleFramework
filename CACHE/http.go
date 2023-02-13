package cache

import (
	"SimpleCache/consistenthash"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const defaultPath = "/cache/"
const defaultReplicas=50

type HTTPPool struct {
	self     string
	basePath string
	mu  sync.Mutex
	peers *consistenthash.Map
	httpGetters map[string]*httpGetter
}

type httpGetter struct{
	baseUrl string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultPath,
	}
}

func (pool *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", pool.self, fmt.Sprintf(format, v...))
}

func (pool *HTTPPool)ServerHttp(w http.ResponseWriter,r *http.Request){
	if !strings.HasPrefix(r.URL.Path,pool.basePath){
		panic("url error")
	}
	pool.Log("%s %s", r.Method, r.URL.Path)
	parts := strings.SplitN(r.URL.Path[len(pool.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "error request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}

func (pool *HTTPPool)Set(peers ...string){
	pool.mu.Lock()
	defer pool.mu.Unlock()

	pool.peers=consistenthash.New(defaultReplicas,nil)
	pool.peers.Add(peers...)
	pool.httpGetters=make(map[string]*httpGetter,len(peers))
	for _,peer:=range peers{
		pool.httpGetters[peer]=&httpGetter{baseUrl: peer+pool.basePath}
	}
}

func (pool *HTTPPool)PickPeer(key string)(PeerGetter,bool){
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if peer:=pool.peers.Get(key);peer!=""&&peer!=pool.self{
		pool.Log("Pick peer %s",peer)
		return pool.httpGetters[peer],true
	}

	return nil,false
}

func(getter *httpGetter)Get(group string,key string)([]byte,error){
	url:=fmt.Sprintf("%v%v/%v",getter.baseUrl,url.QueryEscape(group),url.QueryEscape(key))

	response,err:=http.Get(url)
	if err!=nil{
		return nil,err
	}

	defer response.Body.Close()

	if response.StatusCode!=http.StatusOK{
		return nil,fmt.Errorf("server error: %v",response.Status)
	}

	bytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		return nil,fmt.Errorf("reading response error: %v",err)
	}
	return bytes,nil
}

var _ PeerGetter=(*httpGetter)(nil)