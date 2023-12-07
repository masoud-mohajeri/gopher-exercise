package patterns

import "fmt"

type houseBuilder struct {
	doorType string
	floor    int8
}

func (h *houseBuilder) setFloor(m int8) *houseBuilder {
	h.floor = m
	return h
}

func (h *houseBuilder) setDoorType(t string) *houseBuilder {
	h.doorType = t
	return h
}

func (h *houseBuilder) build() houseBuilder {

	return *h
}

func Builder() {

	house := houseBuilder{}
	house.setDoorType("Good Door!!")
	house.setFloor(22)
	completedHouse := house.build()
	fmt.Println("completedHouse", completedHouse)

}
