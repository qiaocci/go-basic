package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

type Door struct {
	Name string
	FSM  *fsm.FSM
}

func NewDoor(name string) *Door {
	d := &Door{
		Name: name,
	}
	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "开", Src: []string{"closed"}, Dst: "open"},
			{Name: "关", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(event *fsm.Event) {
				fmt.Printf("the door ： %s is %s\n", d.Name, event.Dst)
			},
		},
	)
	return d

}
func main() {
	door := NewDoor("宣武门")
	err := door.FSM.Event("开")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = door.FSM.Event("关")
	if err != nil {
		fmt.Println(err)
		return
	}
}
