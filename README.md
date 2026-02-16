# One elevator - TTK4145

Control of single elevator for TTK4145 Real-Time Programming at NTNU


## Overview

This package contains an implementation for controlling a single elevator as a finite state machine,
using message passing between separate threads

The hardware polling, timer instance and fsm logic all run on separate threads,
using message passing from the hardware and the timer to the fsm loop to signal events for triggering a state transition.

Using message passing allows all fsm helper functions to be pure, since they can calculate the state transitions,
signal the outputted actions on message channels and return the transitioned states to the main fsm loop.

The action outputs are handled in main, which in turn performs the actions by messaging the timer and elevator packages,
which actually do the work of setting the timers and executing the elevator commands.

Structuring the program this way allows all the behavioural logic to be entirely contained within the fsm package,
only needing the other packages for interacting with the outside world,
and thus maintains a clean concept for what a computer program should do


## Dependencies

This project requires the following dependencies to be installed on the host machine:
- `elevatorserver` (for the physical elevator hardware) or [`simelevatorserver`](https://github.com/TTK4145/Simulator-v2)
 needs to be installed and in the path of the root user
- `golang` >= 1.25


## Deployment

To deploy the elevator you first run the elevatorserver (or simulator):

`elevatorserver`

and then run the elevator program:

`go run main.go`

Communication to the both the server and the simulator is over TCP,
and this program is set up to use the default port options.
