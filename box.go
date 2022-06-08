package golang_united_school_homework

import "fmt"

const ErrInd = "Index went out of the range"
// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	//panic("implement me")
	if b.shapesCapacity > len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return fmt.Errorf("Out of the shapesCapacity range")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	//panic("implement me")
	if i >= len(b.shapes) {
		return nil, fmt.Errorf(ErrInd)
	} else {
		return b.shapes[i], nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	//panic("implement me")
	if i >= len(b.shapes) {
		return nil, fmt.Errorf(ErrInd)
	} else {
		s := b.shapes[i]
		copy(b.shapes[i:], b.shapes[i+1:])
		//b.shapes[len(b.shapes)-1] = nil
		b.shapes = b.shapes[:len(b.shapes)-1]
		return s, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//panic("implement me")
	if i >= len(b.shapes) {
		return nil, fmt.Errorf(ErrInd)
	} else {
		s := b.shapes[i]
		b.shapes[i] = shape
		return s, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	//panic("implement me")
	var p float64
	for i := 0; i < len(b.shapes); i++ {
		p += b.shapes[i].CalcPerimeter()
	}
	return p
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	//panic("implement me")
	var s float64
	for i := 0; i < len(b.shapes); i++ {
		s += b.shapes[i].CalcArea()
	}
	return s
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	//panic("implement me")
	n := 0
	for i := 0; i < len(b.shapes); i++ {
		switch b.shapes[i].(type) {
		case Circle, *Circle:
			copy(b.shapes[i:], b.shapes[i+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			n++
			i--
		}
	}
	if n==0 {
		return fmt.Errorf("There are no circles in the list")
	} else {
		return nil
	}
}


