package main

import (
	"elevatorControl/elevator"
	"elevatorControl/fsm"
	"elevatorControl/timer"
	"fmt"
)

func main() {
	fmt.Println("Starting Elevator...")

	// - - - - - - Channels - - - - - - - - -

	// Input message channels for events in finite state machine
	requestEvent := make(chan elevator.ButtonEvent)
	floorEvent := make(chan int)
	doorTimeout := make(chan bool)

	// Output message channels for performing actions on elevator hardware
	setFloorIndicator := make(chan int)
	setLights := make(chan [elevator.N_FLOORS][elevator.N_BUTTONS]bool)
	changeMotorDirection := make(chan elevator.MotorDirection)
	openDoor := make(chan bool)
	closeDoor := make(chan bool)
	keepDoorOpen := make(chan bool)

	// Output message channel for performing action on timer instance
	resetDoorTimer := make(chan bool)

	// - - - - - - Initializing - - - - - - -

	startFloor := elevator.HardwareInit()

	// - - - - - - Deploying - - - - - - -

	go timer.DoorTimer(resetDoorTimer, doorTimeout)
	go elevator.PollButtons(requestEvent)
	go elevator.PollFloorSensor(floorEvent)

	// Finite state machine transition logic
	go fsm.FSM(startFloor,
		requestEvent, floorEvent,
		doorTimeout, setFloorIndicator,
		setLights, changeMotorDirection,
		openDoor, closeDoor, keepDoorOpen)

	// Finite state machine action handling
	for {
		select {
		case newFloor := <-setFloorIndicator:
			elevator.FloorIndicator(newFloor)

		case requestList := <-setLights:
			elevator.SetAllLights(requestList)

		case dir := <-changeMotorDirection:
			elevator.SetMotorDirection(dir)

		case <-openDoor:
			elevator.DoorLight(true)
			resetDoorTimer <- true

		case <-closeDoor:
			elevator.DoorLight(false)

		case <-keepDoorOpen:
			resetDoorTimer <- true
		}
	}
}
