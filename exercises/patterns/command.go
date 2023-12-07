package patterns

import "fmt"

// Button : invoker
type Button struct {
	command Command
}

// Command : command interface
type Command interface {
	execute()
}

func (b *Button) press() {
	b.command.execute()
}

// Device : receiver interface
type Device interface {
	on()
	off()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Tv concrete receiver
type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}
func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}

// Client code

func CommandPattern() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}
	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

}
