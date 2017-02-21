package plot

import (
	"bytes"
	"fmt"
)

type BoxAndWhisker struct {
	LeftWhisker  int
	Left         int
	Mid          int
	Right        int
	RightWhisker int
}

func (box BoxAndWhisker) validate() {
	if !(box.LeftWhisker <= box.Left &&
		box.Left <= box.Mid &&
		box.Mid <= box.Right &&
		box.Right <= box.RightWhisker) {
		panic(fmt.Sprintf("Box points are not in expected order: %v, %v, %v, %v, %v",
			box.LeftWhisker, box.Left, box.Mid, box.Right, box.RightWhisker))
	}
}

func drawBox(box BoxAndWhisker) string {
	box.validate()
	var buf = new(bytes.Buffer)
	drawBoxTop(buf, box)
	drawBoxMid(buf, box)
	drawBoxBottom(buf, box)
	return buf.String()
}

func drawBoxTop(buf *bytes.Buffer, box BoxAndWhisker) {
	//  ¯\_(ツ)_/¯
	writeRepeat(buf, " ", box.Left)
	switch {
	case box.Left == box.Right:
		buf.WriteString("╥")
	case box.Left == box.Mid:
		buf.WriteString("╓")
		writeRepeat(buf, "─", box.Right-box.Mid-1)
		buf.WriteString("┐")
	case box.Left < box.Mid:
		buf.WriteString("┌")
		writeRepeat(buf, "─", box.Mid-box.Left-1)
		if box.Mid == box.Right {
			buf.WriteString("╖")
		} else {
			buf.WriteString("┰")
			writeRepeat(buf, "─", box.Right-box.Mid-1)
			buf.WriteString("┐")
		}
	}
	buf.WriteString("\n")
}

func drawBoxMid(buf *bytes.Buffer, box BoxAndWhisker) {
	//  ¯\_(ツ)_/¯
	writeRepeat(buf, " ", box.LeftWhisker)
	if box.LeftWhisker < box.Left {
		buf.WriteString("├")
		writeRepeat(buf, "╶", box.Left-box.LeftWhisker-1)
	}

	switch {
	case box.Left == box.Right:
		buf.WriteString("║")
	case box.Left == box.Mid:
		buf.WriteString("╢")
		writeRepeat(buf, " ", box.Right-box.Left-1)
		buf.WriteString("├")
	case box.Left < box.Mid:
		buf.WriteString("┤")
		writeRepeat(buf, " ", box.Mid-box.Left-1)
		if box.Mid == box.Right {
			buf.WriteString("╟")
		} else {
			buf.WriteString("┃")
			writeRepeat(buf, " ", box.Right-box.Mid-1)
			buf.WriteString("├")
		}
	}

	if box.RightWhisker > box.Right {
		writeRepeat(buf, "╶", box.RightWhisker-box.Right-1)
		buf.WriteString("┤")
	}
	buf.WriteString("\n")
}

func drawBoxBottom(buf *bytes.Buffer, box BoxAndWhisker) {
	//  ¯\_(ツ)_/¯
	writeRepeat(buf, " ", box.Left)
	switch {
	case box.Left == box.Right:
		buf.WriteString("╨")
	case box.Left == box.Mid:
		buf.WriteString("╙")
		writeRepeat(buf, "─", box.Right-box.Mid-1)
		buf.WriteString("┘")
	case box.Left < box.Mid:
		buf.WriteString("└")
		writeRepeat(buf, "─", box.Mid-box.Left-1)
		if box.Mid == box.Right {
			buf.WriteString("╜")
		} else {
			buf.WriteString("┸")
			writeRepeat(buf, "─", box.Right-box.Mid-1)
			buf.WriteString("┘")
		}
	}
	buf.WriteString("\n")
}

func writeRepeat(buf *bytes.Buffer, s string, count int) {
	for i := 0; i < count; i++ {
		buf.WriteString(s)
	}
}
