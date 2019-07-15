package main

import (
	"fmt"

	"github.com/wang/ledger"
)

type user struct {
	name  string
	email string
}

func (u *user) change_email_by_pointer(e string) {
	u.email = e
	fmt.Printf("in change email by pointer,name: %s, email: %s\n", u.name, u.email)
}
func (u user) change_email_by_value(e string) {
	u.email = e
	fmt.Printf("in change email by value,name: %s, email: %s\n", u.name, u.email)
}

/*func main() {
	bill := &user{
		name:  "bill",
		email: "bill@yahoo.com.cn",
	}
	mark := &user{"Mark", "mark@yahoo.com.cn"}
	fmt.Printf("name: %s, email: %s\n", bill.name, bill.email)
	fmt.Printf("name: %s, email: %s\n", mark.name, mark.email)
	bill.change_email_by_pointer("bill@189.cn")
	fmt.Printf("name: %s, email: %s\n", bill.name, bill.email)
	mark.change_email_by_value("mark@189.cn")
	fmt.Printf("name: %s, email is %s\n", mark.name, mark.email)

	attrsToIndex := []ledger.IndexableAttr{
		ledger.IndexableAttrBlockHash,
		ledger.IndexableAttrBlockNum,
		ledger.IndexableAttrTxID,
		ledger.IndexableAttrBlockNumTranNum,
		ledger.IndexableAttrBlockTxID,
		ledger.IndexableAttrTxValidationCode,
	}
	fmt.Printf("%s", attrsToIndex)
	//indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}

}
*/
