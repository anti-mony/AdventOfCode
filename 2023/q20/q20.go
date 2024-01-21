package main

import (
	"fmt"
	"log"
	"strings"

	"advent.of.code/list"
	"advent.of.code/util"
)

func main() {

	config, modules, err := parseInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lowP, highP := solveP1(config, modules, 1000)
	fmt.Printf("P1 answer is %d\n", lowP*highP)

	fmt.Printf("P2 answer is %d\n", solveP2(config, modules))
}

func solveP1(config ModuleConfiguration, modules Modules, nButtonPresses int) (int, int) {
	lowPulses, highPulses := 0, 0

	for i := 0; i < nButtonPresses; i++ {
		loopLowPulse, loopHighPulse := pressButton(config, modules)
		lowPulses += loopLowPulse
		highPulses += loopHighPulse
	}

	return lowPulses, highPulses
}

func solveP2(config ModuleConfiguration, modules Modules) int {

	// Worth a shot I guess lol, obviously doesn't work
	i := 0
	for !pressButtonP2(config, modules, "rx", 0) {
		i++
	}
	return i

	// Looking at the input a conjunction module feeds into rx
	// so all the thigns that feeds into that conjunction module
	// need to be 1s and then only it'll send out a 0
	// Maybe there's a cyclic behaviour to those input to the conjunction
	// module and then maybe the LCM of those cycles is the answer.
	// idk
}

func pressButton(config ModuleConfiguration, modules Modules) (int, int) {
	lowPulses, highPulses := 0, 0

	q := list.NewQueue()
	q.Push(Pulse{To: "broadcaster", From: "button", Value: 0})

	for q.Len() > 0 {
		pulse := q.Pop().(Pulse)

		if pulse.Value == 0 {
			lowPulses++
		} else {
			highPulses++
		}

		moduleTo, moduleExists := modules[pulse.To]
		if moduleExists {
			emittedPulse := moduleTo.Emit(pulse.From, pulse.Value)
			if emittedPulse != nil {
				for _, r := range config[pulse.To] {
					q.Push(Pulse{From: pulse.To, To: r, Value: *emittedPulse})
				}
			}
		}
	}

	return lowPulses, highPulses
}

func pressButtonP2(config ModuleConfiguration, modules Modules, outputModuleName string, pulseValue int) bool {

	q := list.NewQueue()
	q.Push(Pulse{To: "broadcaster", From: "button", Value: 0})

	for q.Len() > 0 {
		pulse := q.Pop().(Pulse)

		if pulse.To == outputModuleName && pulse.Value == pulseValue {
			return true
		}

		moduleTo, moduleExists := modules[pulse.To]
		if moduleExists {
			emittedPulse := moduleTo.Emit(pulse.From, pulse.Value)
			if emittedPulse != nil {
				for _, r := range config[pulse.To] {
					q.Push(Pulse{From: pulse.To, To: r, Value: *emittedPulse})
				}
			}
		}
	}

	return false
}

type Pulse struct {
	To, From string
	Value    int
}

func (p Pulse) String() string {
	return fmt.Sprintf("%s --%d-> %s \t", p.From, p.Value, p.To)
}

type Modules map[string]*Module
type ModuleConfiguration map[string][]string

type ModuleType int

const (
	ModuleTypeUnknown ModuleType = iota
	ModuleTypeFlipFlop
	ModuleTypeConjunction
	ModuleTypeBroadCaster
)

type Module struct {
	Type   ModuleType
	Memory map[string]int
	On     bool
}

func (m *Module) Emit(from string, pulse int) *int {

	if m.Type == ModuleTypeUnknown {
		return nil
	}

	if m.Type == ModuleTypeBroadCaster {
		return &pulse
	}

	if m.Type == ModuleTypeFlipFlop {
		if pulse == 1 {
			return nil
		}
		newPulse := 0
		if !m.On {
			m.On = true
			newPulse = 1
		} else {
			m.On = false
		}
		return &newPulse
	}

	// if m.Type == Conjunction
	m.Memory[from] = pulse

	pulses := 0
	for _, pulse := range m.Memory {
		pulses += pulse
	}

	newPulse := 1
	if pulses == len(m.Memory) {
		newPulse = 0
	}

	return &newPulse
}

func (m *Module) String() string {
	return fmt.Sprintf("Type: %d | Memory: %v | On %v", m.Type, m.Memory, m.On)
}

func parseInput(filename string) (ModuleConfiguration, Modules, error) {
	lines, err := util.GetFileAsListOfStrings(filename)
	if err != nil {
		return nil, nil, err
	}

	modules := make(Modules)
	config := make(ModuleConfiguration)

	for _, line := range lines {
		ioSplit := strings.Split(line, "->")
		producer := strings.TrimSpace(ioSplit[0])

		receivers := strings.Split(ioSplit[1], ",")
		for i, v := range receivers {
			receivers[i] = strings.TrimSpace(v)
		}

		newModule, name := makeModule(producer)
		if _, in := modules[name]; !in {
			modules[name] = &newModule
		}

		if rcvrs, in := config[name]; in {
			config[name] = append(rcvrs, receivers...)
		} else {
			config[name] = receivers
		}
	}

	for producer, receivers := range config {
		for _, receiver := range receivers {
			if m, ok := modules[receiver]; ok {
				if m.Type == ModuleTypeConjunction {
					modules[receiver].Memory[producer] = 0
				}
			}
		}
	}

	return config, modules, nil
}

func makeModule(in string) (Module, string) {
	if in == "broadcaster" {
		return Module{Type: ModuleTypeBroadCaster}, "broadcaster"
	}

	if string(in[0]) == "%" {
		return Module{Type: ModuleTypeFlipFlop}, in[1:]
	}

	return Module{Type: ModuleTypeConjunction, Memory: make(map[string]int)}, in[1:]
}

func PrintConfig(c ModuleConfiguration, m Modules) {

	fmt.Println("\n#### Modules ####")

	for name, module := range m {
		fmt.Printf("%12s ---> %v \n", name, module)
	}

	fmt.Println("\n#### Configuration ####")

	for input, output := range c {
		fmt.Printf("%12s ---> %v \n", input, output)
	}

}
