package xb

import (
	"fmt"
	. "github.com/gherlein/xbevents"
)

/*
#include <linux/joystick.h>
#include <fcntl.h>
#include <unistd.h>
int jsOpen(char* f) {
  return open(f, O_RDONLY);
}
int jsClose(int fd) {
  return close(fd);
}
void jsRead(int js, struct js_event* e) {
  ssize_t r=read(js, e, sizeof(struct js_event));
}
*/
import "C"

var (
	event    xbc_event
	device   string = "/dev/input/js0"
	deadzone int16  = 200
	debug    bool   = false
	fd       int    = -1
	x        int16  = 0
	y        int16  = 0
)

const (
	PRESS    uint8 = 1 // EventType
	JOYSTICK uint8 = 2 // EventType
)

const (
	A      uint8 = 0
	B      uint8 = 1
	X      uint8 = 2
	Y      uint8 = 3
	BACK   uint8 = 9
	START  uint8 = 8
	LTOP   uint8 = 4
	LBOT   uint8 = 6
	RTOP   uint8 = 5
	RBOT   uint8 = 7
	RSTICK uint8 = 12
	LSTICK uint8 = 11
	GUIDE  uint8 = 10
	PADL   uint8 = 13
	PADR   uint8 = 14
	PADU   uint8 = 15
	PADD   uint8 = 16
)

type xbc_event struct {
	Time      uint32
	Value     int16
	EventType uint8
	Number    uint8
}

func Open() int {
	i := C.jsOpen(C.CString(device))
	// need some error handling here
	return int(i)
}

func xbclose() int {
	return int(C.jsClose(C.int(fd)))
}

func xbGetEvent() int {
	var e C.struct_js_event

	if fd == -1 {
		fd = Open()
		// put some error handling here
	}
	C.jsRead(C.int(fd), &e)

	event.Time = uint32(e.time)
	event.Value = int16(e.value)
	event.EventType = uint8(e._type)
	event.Number = uint8(e.number)
	return 1
}

func DEBUG(a ...interface{}) {
	if debug {
		fmt.Println(a)
	}
}

func DebugModeOn() {
	debug = true
}

func DebugModeOff() {
	debug = false
}

func GetEvent() *XBevent {
	r := xbGetEvent()
	var xbe XBevent
	var valid bool = false
	if r > 0 {
		if event.EventType == PRESS {
			valid = true
			switch event.Number {
			case A:
				if event.Value == 1 {
					xbe.Code = A_DOWN
					xbe.Name = "A_DOWN"
				} else {
					xbe.Code = A_UP
					xbe.Name = "A_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}
			case B:
				if event.Value == 1 {
					xbe.Code = B_DOWN
					xbe.Name = "B_DOWN"
				} else {
					xbe.Code = B_UP
					xbe.Name = "B_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}
			case X:
				if event.Value == 1 {
					xbe.Code = X_DOWN
					xbe.Name = "X_DOWN"
				} else {
					xbe.Code = X_UP
					xbe.Name = "X_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case Y:
				if event.Value == 1 {
					xbe.Code = Y_DOWN
					xbe.Name = "Y_DOWN"
				} else {
					xbe.Code = Y_UP
					xbe.Name = "Y_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case BACK:
				if event.Value == 1 {
					xbe.Code = BACK_DOWN
					xbe.Name = "BACK_DOWN"
				} else {
					xbe.Code = BACK_UP
					xbe.Name = "BACK_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case START:
				if event.Value == 1 {
					xbe.Code = START_DOWN
					xbe.Name = "START_DOWN"
				} else {
					xbe.Code = START_UP
					xbe.Name = "START_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case LTOP:
				if event.Value == 1 {
					xbe.Code = LTOP_DOWN
					xbe.Name = "LTOP_DOWN"
				} else {
					xbe.Code = LTOP_UP
					xbe.Name = "LTOP_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case LBOT:
				if event.Value == 1 {
					xbe.Code = LBOT_DOWN
					xbe.Name = "LBOT_DOWN"
				} else {
					xbe.Code = LBOT_UP
					xbe.Name = "LBOT_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case RTOP:
				if event.Value == 1 {
					xbe.Code = RTOP_DOWN
					xbe.Name = "RTOP_DOWN"
				} else {
					xbe.Code = RTOP_UP
					xbe.Name = "RTOP_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case RBOT:
				if event.Value == 1 {
					xbe.Code = RBOT_DOWN
					xbe.Name = "RBOT_DOWN"
				} else {
					xbe.Code = RBOT_UP
					xbe.Name = "RBOT_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case RSTICK:
				if event.Value == 1 {
					xbe.Code = RJOY_DOWN
					xbe.Name = "RJOY_DOWN"
				} else {
					xbe.Code = RJOY_UP
					xbe.Name = "RJOY_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case LSTICK:
				if event.Value == 1 {
					xbe.Code = LJOY_DOWN
					xbe.Name = "LJOY_DOWN"
				} else {
					xbe.Code = LJOY_UP
					xbe.Name = "LJOY_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case GUIDE:
				if event.Value == 1 {
					xbe.Code = GUIDE_DOWN
					xbe.Name = "GUIDE_DOWN"
				} else {
					xbe.Code = GUIDE_UP
					xbe.Name = "GUIDE_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case PADL:
				if event.Value == 1 {
					xbe.Code = PADL_DOWN
					xbe.Name = "PADL_DOWN"
				} else {
					xbe.Code = PADL_UP
					xbe.Name = "PADL_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case PADR:
				if event.Value == 1 {
					xbe.Code = PADR_DOWN
					xbe.Name = "PADR_DOWN"
				} else {
					xbe.Code = PADR_UP
					xbe.Name = "PADR_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case PADU:
				if event.Value == 1 {
					xbe.Code = PADU_DOWN
					xbe.Name = "PADU_DOWN"
				} else {
					xbe.Code = PADU_UP
					xbe.Name = "PADU_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}

			case PADD:
				if event.Value == 1 {
					xbe.Code = PADD_DOWN
					xbe.Name = "PADD_DOWN"
				} else {
					xbe.Code = PADD_UP
					xbe.Name = "PADD_UP"
				}
				if debug {
					fmt.Println(xbe.Name)
				}
			}

		}
		if event.EventType == JOYSTICK {
			valid = true
			switch event.Number {
			case LJOYX:
				xbe.Code = LJOYX
				xbe.Name = "LJOYX"
				x = event.Value
				xbe.X = x
				xbe.Y = y

				if debug {
					fmt.Printf("%s - x: %d    y: %d\n", xbe.Name, xbe.X, xbe.Y)
				}
			case LJOYY:
				xbe.Code = LJOYY
				xbe.Name = "LJOYY"
				y = event.Value
				xbe.X = x
				xbe.Y = y

				if debug {
					fmt.Printf("%s - x: %d    y: %d\n", xbe.Name, xbe.X, xbe.Y)
				}

			case RJOYX:
				xbe.Code = RJOYX
				xbe.Name = "RJOYX"
				x = event.Value
				xbe.X = x
				xbe.Y = y

				if debug {
					fmt.Printf("%s - x: %d    y: %d\n", xbe.Name, xbe.X, xbe.Y)
				}
			case RJOYY:
				xbe.Code = RJOYY
				xbe.Name = "RJOYY"
				y = event.Value
				xbe.X = x
				xbe.Y = y

				if debug {
					fmt.Printf("%s - x: %d    y: %d\n", xbe.Name, xbe.X, xbe.Y)
				}

			}

		}

	}
	if valid == true {
		return &xbe
	}
	return nil
}
