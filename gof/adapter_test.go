package gof

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	client := Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer")
	com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into Mac machine")
}

type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB Connector is plugged into Windows machine.")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}
