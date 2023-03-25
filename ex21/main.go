package main

import "fmt"

type IPhone4s struct {
}

type IPhone14Pro struct {
}

type SimCardConnector interface {
	NanoSim()
}

type IPhone4sWithNanoSimAdapter struct {
	iPhone4s *IPhone4s
}

func (i4s *IPhone4s) MicroSim() {
	fmt.Println("The iPhone 4s runs on the MicroSim standard.")
}

func (i14Pro *IPhone14Pro) NanoSim() {
	fmt.Println("The iPhone 4s runs on the NanoSim standard.")
}

func (adapter *IPhone4sWithNanoSimAdapter) NanoSim() {
	fmt.Println("Inserted NanoSim into NanoSimAdapter (MicroSim)")
	adapter.iPhone4s.MicroSim()
}

func InsertNanoSim(phone SimCardConnector) {
	phone.NanoSim()
}

func main() {
	iphone4s := &IPhone4s{}
	iphone14Pro := &IPhone14Pro{}
	adapter := &IPhone4sWithNanoSimAdapter{
		iPhone4s: iphone4s,
	}

	InsertNanoSim(iphone14Pro)
	InsertNanoSim(adapter)

}
