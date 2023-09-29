package main

type Automaton struct {
	status bool
	states []string
}

func automaton(text string) *Automaton {

	var states []string

	state := "q0"
	states = append(states, state)

	for _, entry := range text {
		switch state {
		case "q0":
			if entry == 'Y' {
				state = "q1"
			} else if entry == 'Z' {
				state = "q2"
			} else {
				break
			}
		case "q1":
			if entry >= 'W' && entry <= 'Z' {
				state = "q3"
			} else {
				break
			}
		case "q2":
			if entry >= 'A' && entry <= 'C' {
				state = "q3"
			} else {
				break
			}
		case "q3":
			if entry >= 'A' && entry <= 'Z' {
				state = "q4"
			} else {
				break
			}
		case "q4":
			if entry == '-' {
				state = "q5"
			} else {
				break
			}
		case "q5":
			if entry == '0' {
				state = "q6"
			} else if entry >= '1' && entry <= '9' {
				state = "q11"
			} else {
				break
			}
		case "q6":
			if entry == '0' {
				state = "q7"
			} else if entry >= '1' && entry <= '9' {
				state = "q9"
			} else {
				break
			}
		case "q7":
			if entry >= '1' && entry <= '9' {
				state = "q8"
			} else {
				break
			}
		case "q8", "q10", "q13":
			if entry == '-' {
				state = "q14"
			} else {
				break
			}
		case "q9":
			if entry >= '0' && entry <= '9' {
				state = "q10"
			} else {
				break
			}
		case "q11":
			if entry >= '0' && entry <= '9' {
				state = "q12"
			} else {
				break
			}
		case "q12":
			if entry >= '0' && entry <= '9' {
				state = "q13"
			} else {
				break
			}
		case "q14":
			if entry >= 'A' && entry <= 'Z' {
				state = "q15"
			} else {
				break
			}
		}
		states = append(states, state)
	}

	if state == "q15" {
		return &Automaton{
			status: true,
			states: states,
		}
	} else {
		return &Automaton{
			status: false,
			states: states,
		}
	}

}
