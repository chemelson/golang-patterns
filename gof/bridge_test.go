package gof

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &MacComputer{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &WindowsComputer{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

// Abstraction
type IComputer interface {
	Print()
	SetPrinter(IPrinter)
}

// Refined abstractions
type MacComputer struct {
	printer IPrinter
}

func (m *MacComputer) Print() {
	fmt.Println("Print request for Mac")
	m.printer.PrintFile()
}

func (m *MacComputer) SetPrinter(p IPrinter) {
	m.printer = p
}

type WindowsComputer struct {
	printer IPrinter
}

func (w *WindowsComputer) Print() {
	fmt.Println("Print request for Windows")
	w.printer.PrintFile()
}

func (w *WindowsComputer) SetPrinter(p IPrinter) {
	w.printer = p
}

// Implementation
type IPrinter interface {
	PrintFile()
}

// Refined implementations
type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type HP struct {
}

func (hp *HP) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
