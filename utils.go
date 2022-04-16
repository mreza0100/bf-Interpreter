package main

func getInsideLoop(commands string, loopStartIdx int) (insideLoop string, endLoopIdx int) {
	depth := 0

Loop:
	for _, r := range commands[loopStartIdx+1:] {
		switch r {
		case loopStart:
			depth++
			insideLoop += string(r)
		case loopEnd:
			if depth <= 0 {
				break Loop
			}
			depth--
			insideLoop += string(r)
		default:
			insideLoop += string(r)
		}
	}

	return insideLoop, loopStartIdx + len(insideLoop) + 1
}
