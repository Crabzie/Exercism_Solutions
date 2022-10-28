package elon

import "fmt"

// Drive moves the car is it has enough juice
func (c *Car) Drive() {
	if c.batteryDrain <= c.battery {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance shows the odometer
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery shows how much juice is remaining.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish calculates if the car has enough juice to complete the track
func (c Car) CanFinish(trackDistance int) bool {
	totalBatteryDrained := (c.batteryDrain * trackDistance) / c.speed
	return totalBatteryDrained <= c.battery
}
