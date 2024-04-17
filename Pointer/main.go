package main

import "fmt"

// Device struct defines a model and a flag indicating if it's malfunctioning.
type Device struct {
    Model            string
    IsMalfunctioning bool
}

// NewDevice is a constructor that initializes a Device struct and returns a pointer to it.
func NewDevice(model string) *Device {
    return &Device{
        Model:            model,
        IsMalfunctioning: false,
    }
}

// MarkMalfunctioning sets the device's malfunctioning status.
func (d *Device) MarkMalfunctioning(status bool) {
    d.IsMalfunctioning = status
}

// Display prints the current state of the device.
func (d *Device) Display() {
    fmt.Printf("Device %s, Malfunctioning: %t\n", d.Model, d.IsMalfunctioning)
}

func main() {
    // Array of pointers to Device structs.
    devices := []*Device{
        NewDevice("DeviceA"),
        NewDevice("DeviceB"),
        NewDevice("DeviceC"),
    }

    // Mark the second device in the list as malfunctioning.
    devices[1].MarkMalfunctioning(true)

    // Iterate over the array and display each device's status.
    for _, device := range devices {
        device.Display()
    }
}
