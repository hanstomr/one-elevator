# One elevator - TTK4145

Control of single elevator for TTK4145 Real-Time Programming at NTNU


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
