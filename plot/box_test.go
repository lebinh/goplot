package plot

import (
	"testing"
)

//  ¯\_(ツ)_/¯
var box1 = `
╥
║
╨
`[1:]

var box2 = `
   ┌───┰──┐
├╶╶┤   ┃  ├╶╶╶┤
   └───┸──┘
`[1:]

var box3 = `
    ╓───┐
  ├╶╢   ├╶┤
    ╙───┘
`[1:]

var box4 = `
   ┌──╖
├╶╶┤  ╟╶╶╶┤
   └──╜
`[1:]

func TestDrawBox(t *testing.T) {
	tests := []struct {
		box  BoxAndWhisker
		want string
	}{
		{BoxAndWhisker{0, 0, 0, 0, 0}, box1},
		{BoxAndWhisker{0, 3, 7, 10, 14}, box2},
		{BoxAndWhisker{2, 4, 4, 8, 10}, box3},
		{BoxAndWhisker{0, 3, 6, 6, 10}, box4},
	}
	for _, test := range tests {
		if got := drawBox(test.box); got != test.want {
			t.Errorf("drawBox(%v):\n%v, expected:\n%v", test.box, got, test.want)
		}
	}
}
