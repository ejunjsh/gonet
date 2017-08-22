package netgo

import "fmt"

type Pipeline struct {
	head *HandlerContext
	tail *HandlerContext
	chl *channel
}

func (p *Pipeline) fireNextRead(data interface{}){
	p.head.FireRead(data)
}

func (p *Pipeline) fireNextConnected(){
	p.head.FireConnected()
}



func (p *Pipeline) AddLast(handler Handler) *Pipeline{
    prev:=p.tail.prev
	newH:=newHandlerContext(p,handler)
	newH.prev=prev
	newH.next=p.tail
	prev.next=newH
	p.tail.prev=newH
	return p
}

func (p *Pipeline) AddFirst(handler Handler) *Pipeline{
	next:=p.head.next
	newH:=newHandlerContext(p,handler)
	newH.prev=p.head
	newH.next=next
	p.head.next=newH
	next.prev=newH
	return p
}

type headHandler struct {

}

func (h *headHandler) Read(c *HandlerContext,data interface{}) error{
	c.FireRead(data)
	return nil
}

func (h *headHandler) Connected(c *HandlerContext) error{
	c.FireConnected()
	return  nil
}

func (h *headHandler) ErrorCaught(c *HandlerContext,err error){

}

func (h *headHandler) Write(c *HandlerContext,data interface{}) error{
	b,ok:=data.([]byte)
	if ok{
		c.p.chl.Write(b)
		fmt.Println("write 2")
	}
	return nil
}

type tailHandler struct {

}

func (t *tailHandler) ErrorCaught(c *HandlerContext,err error){

}



func newPipeline() *Pipeline{
     p:=&Pipeline{}
	p.tail=&HandlerContext{p,nil,nil,&tailHandler{}}
	p.head=&HandlerContext{p,nil,nil,&headHandler{}}
	p.head.next=p.tail
	p.tail.prev=p.head
	return p
}



