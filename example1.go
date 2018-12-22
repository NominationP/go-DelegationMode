/**
这个是 Go 语中的委托和接口多态的编程方式，其实是面向对象和原型编程的综合玩法
https://time.geekbang.org/column/article/2748
 */

package main

import (
	"fmt"
)

type Widget struct {
	X, Y int
}

/**
Label
 */
type Label struct {
	Widget
	Text string
	X    int
}

func (label Label) Paint() {
	//
	fmt.Printf("[%p] - Label.Paint(%q)\n",
		&label, label.Text)
}

/**
Button
 */
type Button struct {
	Label
}

func NewButton(x, y int, text string) Button {
	return Button{Label{Widget{x, y}, text, x}}
}

func (button Button) Paint() {
	fmt.Printf("[%p] - Button.Paint(%q)\n",
		&button, button.Text)
}

func (button Button) Click() {
	fmt.Printf("[%p] - Button.Click()\n", &button)
}

/**
ListBox
 */
type ListBox struct {
	Widget
	Texts []string
	Index int
}

func (listBox ListBox) Paint() {
	fmt.Printf("[%p] - ListBox.Paint(%q)\n",
		&listBox, listBox.Texts)
}

func (listBox ListBox) Click() {
	fmt.Printf("[%p] - ListBox.Click()\n", &listBox)
}

/**
interface
 */
type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func main() {

	/**
	example1 output
	 */

	label := Label{Widget{10, 10}, "State", 100}
	//fmt.Printf("X=%d, Y=%d, Text=%s Widget.X=%d\n",
	//	label.X, label.Y, label.Text,
	//	label.Widget.X)
	//fmt.Println()
	//fmt.Printf("%+v\n%v\n", label, label)
	//label.Paint()

	/**
	example2 output
	 */
	button1 := Button{Label{Widget{10, 70}, "ok", 10}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 10},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	fmt.Println()
	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	fmt.Println()
	for _, widget := range []interface{}{label, listBox, button1, button2} {
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
	}
}
