package provider
import (
	"../ledger"
)
type Initializer struct {
	para1    int
	para2    int
	para3    string
	para4    string
}

type struct theProvider {
}
var v Initializer;

func (p *theProvider)Init(int){}
func (p *theProvider)Init(i *Initializer){
	v = *i
}
func (p *theProvider)Open(){}
func (p *theProvider)Read() int{}
func (p *theProvider)Write(int){}
func (p *theProvider)Close(){}
