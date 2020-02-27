package dynamicarr

import (
	"errors"
)

// GrowthFactor determines the how much array will resize
const GrowthFactor = 2

// DynamicArr defines the structure of the array object
type DynamicArr struct {
	length   int
	capacity int
	array    []interface{}
}

func (d *DynamicArr) shrinkSize() {

	d.capacity = d.capacity / GrowthFactor
	temArr := make([]interface{}, d.capacity)
	for i := 0; i < d.length; i++ {
		temArr[i] = d.array[i]
	}
	//log.Printf("Shrinked to %d, Length: %d", d.capacity, d.length)
	d.array = temArr
}

func (d *DynamicArr) growSize() {

	if d.length != d.capacity {
		return
	}

	d.capacity = d.capacity * GrowthFactor
	temArr := make([]interface{}, d.capacity)
	for i := 0; i < d.length; i++ {
		temArr[i] = d.array[i]
	}
	//log.Printf("Growed to %d, Length: %d", d.capacity, d.length)
	d.array = temArr
}

// NewDynamicArray returns a DynamicArr instance
func NewDynamicArray() DynamicArr {
	d := DynamicArr{}
	d.length = 0
	d.capacity = 1
	d.array = make([]interface{}, 1)
	return d
}

// GetCapacity returns the length of the capacity
func (d *DynamicArr) GetCapacity() int {
	return d.capacity
}

// GetLength returns the length of the array
func (d *DynamicArr) GetLength() int {
	return d.length
}

// Get returns the element of array on given index
func (d *DynamicArr) Get(index int) (interface{}, error) {
	if index < 0 || index > (d.capacity-1) {
		return nil, errors.New("Index out of range")
	}
	return d.array[index], nil
}

// Set overrides an index of array to given element
func (d *DynamicArr) Set(index int, value interface{}) error {
	if index < 0 {
		return errors.New("Index has to be greater than or equal to zero")
	}

	for index > d.capacity {
		d.growSize()
		d.length = d.capacity
	}

	d.array[index] = value
	d.length++
	return nil
}

// Push an element into end of the array
func (d *DynamicArr) Push(value interface{}) {
	if d.length == d.capacity {
		d.growSize()
	}
	d.array[d.length] = value
	d.length++
}

// Pop removes the last element of the array and returns it
func (d *DynamicArr) Pop() (interface{}, error) {
	if d.length == 0 {
		return nil, errors.New("Empty array")
	}

	for d.capacity/2 > d.length-1 {
		d.shrinkSize()
	}

	tempArr := make([]interface{}, d.capacity)
	for i := 0; i < d.length-1; i++ {
		tempArr[i] = d.array[i]
	}
	val := d.array[d.length-1]
	d.length--
	d.array = tempArr

	return val, nil
}
