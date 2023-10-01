package models

type Automaton struct {
	Status bool
	States []string
}

func Validation(text string) *Automaton {

	var states []string
	statesMaxLength := 10

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
				state = "error"
			}
		case "q1":
			if entry >= 'W' && entry <= 'Z' {
				state = "q3"
			} else {
				state = "error"
			}
		case "q2":
			if entry >= 'A' && entry <= 'C' {
				state = "q3"
			} else {
				state = "error"
			}
		case "q3":
			if entry >= 'A' && entry <= 'Z' {
				state = "q4"
			} else {
				state = "error"
			}
		case "q4":
			if entry == '-' {
				state = "q5"
			} else {
				state = "error"
			}
		case "q5":
			if entry == '0' {
				state = "q6"
			} else if entry >= '1' && entry <= '9' {
				state = "q11"
			} else {
				state = "error"
			}
		case "q6":
			if entry == '0' {
				state = "q7"
			} else if entry >= '1' && entry <= '9' {
				state = "q9"
			} else {
				state = "error"
			}
		case "q7":
			if entry >= '1' && entry <= '9' {
				state = "q8"
			} else {
				state = "error"
			}
		case "q8", "q10", "q13":
			if entry == '-' {
				state = "q14"
			} else {
				state = "error"
			}
		case "q9":
			if entry >= '0' && entry <= '9' {
				state = "q10"
			} else {
				state = "error"
			}
		case "q11":
			if entry >= '0' && entry <= '9' {
				state = "q12"
			} else {
				state = "error"
			}
		case "q12":
			if entry >= '0' && entry <= '9' {
				state = "q13"
			} else {
				state = "error"
			}
		case "q14":
			if entry >= 'A' && entry <= 'Z' {
				state = "q15"
			} else {
				state = "error"
			}
		case "error":
			return &Automaton{
				Status: false,
				States: states,
			}
		}
		states = append(states, state)
		if state == "q15" {
			break
		}
	}

	if len(text) > statesMaxLength {
		states = append(states, "La cadena supera lo valido")
		return &Automaton{
			Status: false,
			States: states,
		}
	}

	if state == "q15" {
		return &Automaton{
			Status: true,
			States: states,
		}
	} else {
		return &Automaton{
			Status: false,
			States: states,
		}
	}

}
