package plot

import (
	"bytes"
	"fmt"
	"sort"
	"math"
)

type BoxAndWhisker struct {
	LeftWhisker  float64 // smallest x that x >= Q1 - 1.5IQR
	Left         float64 // == Q1
	Mid          float64 // == Q2
	Right        float64 // == Q3
	RightWhisker float64 // largest x that x <= Q3 + 1.5IQR
}

type NormalizedBoxAndWhisker struct {
	LeftWhisker  int // smallest x that x >= Q1 - 1.5IQR
	Left         int // == Q1
	Mid          int // == Q2
	Right        int // == Q3
	RightWhisker int // largest x that x <= Q3 + 1.5IQR
}

type Quartiles struct {
	First, Second, Third float64
}

func (box BoxAndWhisker) normalize(maxWidth int) NormalizedBoxAndWhisker {
	// ಠ_ಠ  this is ridiculous but I don't know a better way to do this max
	max := math.Max(box.LeftWhisker, math.Max(box.Left, math.Max(box.Mid, math.Max(box.Right, box.RightWhisker))))
	normalize := func(val float64) int  {
		return int(val / max * float64(maxWidth))
	}
	return NormalizedBoxAndWhisker{
		normalize(box.LeftWhisker),
		normalize(box.Left),
		normalize(box.Mid),
		normalize(box.Right),
		normalize(box.RightWhisker),
	}
}

func (box NormalizedBoxAndWhisker) validate() {
	if !(box.LeftWhisker <= box.Left &&
		box.Left <= box.Mid &&
		box.Mid <= box.Right &&
		box.Right <= box.RightWhisker) {
		panic(fmt.Sprintf("Box points are not in expected order: %v, %v, %v, %v, %v",
			box.LeftWhisker, box.Left, box.Mid, box.Right, box.RightWhisker))
	}
}

func calculateBoxAndWhisker(values []float64) BoxAndWhisker {
	if len(values) == 0 {
		return BoxAndWhisker{}
	}

	// sort a copy of input values
	sorted := make([]float64, len(values))
	copy(sorted, values)
	sort.Float64s(sorted)

	quartiles := quartilesOfSorted(sorted)
	box := BoxAndWhisker{Left: quartiles.First, Mid: quartiles.Second, Right: quartiles.Third}

	oneHalfIQR := (quartiles.Third - quartiles.First) * 1.5
	for i := 0; i < len(sorted); i++ {
		if sorted[i] >= quartiles.First-oneHalfIQR {
			box.LeftWhisker = sorted[i]
			break
		}
	}
	for i := len(sorted) - 1; i >= 0; i-- {
		if sorted[i] <= quartiles.Third+oneHalfIQR {
			box.RightWhisker = sorted[i]
			break
		}
	}
	return box
}

func quartilesOfSorted(sortedValues []float64) Quartiles {
	if len(sortedValues) == 0 {
		return Quartiles{}
	}

	// using method 2 of computing quartiles from https://en.wikipedia.org/wiki/Quartile
	// use the median to divide the ordered data set into two halves.
	median := medianOfSorted(sortedValues)
	mid := int(len(sortedValues) / 2)
	var firstHalf, secondHalf []float64
	if len(sortedValues)%2 == 0 {
		firstHalf = sortedValues[:mid]
	} else {
		// including the median value in the first half if odd number of values
		firstHalf = sortedValues[:mid+1]
	}
	secondHalf = sortedValues[mid:]
	q1 := medianOfSorted(firstHalf)
	q3 := medianOfSorted(secondHalf)
	return Quartiles{q1, median, q3}
}

func medianOfSorted(sortedValues []float64) float64 {
	if len(sortedValues) == 0 {
		return 0
	}

	mid := int(len(sortedValues) / 2)
	if len(sortedValues)%2 == 0 {
		return (sortedValues[mid-1] + sortedValues[mid]) / 2
	} else {
		return sortedValues[mid]
	}
}

func drawBox(box NormalizedBoxAndWhisker) string {
	box.validate()
	buf := new(bytes.Buffer)
	drawBoxTop(box, buf)
	drawBoxMid(box, buf)
	drawBoxBottom(box, buf)
	return buf.String()
}

func drawBoxTop(box NormalizedBoxAndWhisker, buf *bytes.Buffer) {
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

func drawBoxMid(box NormalizedBoxAndWhisker, buf *bytes.Buffer) {
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

func drawBoxBottom(box NormalizedBoxAndWhisker, buf *bytes.Buffer) {
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
