package main

// func getInsideLoop(commands string, loopStartIdx int) (insideLoop string, end int) {
// 	depth := 0

// Loop:
// 	for _, r := range commands[loopStartIdx+1:] {
// 		switch r {
// 		case loopEnter:
// 			depth++
// 			insideLoop += string(r)
// 		case loopExit:
// 			if depth <= 0 {
// 				break Loop
// 			}
// 			depth--
// 			insideLoop += string(r)
// 		default:
// 			insideLoop += string(r)
// 		}
// 	}

// 	return insideLoop, loopStartIdx + len(insideLoop)
// }

func getLoopEnd(commands string, loopStartIdx int) (end int) {
	depth := 0
	loopLen := 0

Loop:
	for _, r := range commands[loopStartIdx+1:] {
		switch r {
		case loopEnter:
			depth++
			loopLen++
		case loopExit:
			if depth <= 0 {
				break Loop
			}
			depth--
			loopLen++
		default:
			loopLen++
		}
	}

	return loopStartIdx + loopLen + 1
}

// func trim(instructions string) string {
// 	re := regexp.MustCompile(`\r?\n| |[a-zA-Z0-9]`)
// 	return re.ReplaceAllString(instructions, "")
// }
