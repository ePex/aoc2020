package main

func CalculateYesCount(batch []string, everyone bool) (answerCount int) {
	var groupAnswers []string
	for idx, answers := range batch {
		groupAnswers = append(groupAnswers, answers)
		if answers == "" || idx+1 == len(batch) {
			answerCount += countGroup(groupAnswers, everyone)
			groupAnswers = []string{}
		}
	}

	return answerCount
}

func countGroup(groupAnswers []string, everyone bool) (answerCount int) {
	// drop last entry that holds the 'group break' (empty line)
	if len(groupAnswers) > 1 && groupAnswers[len(groupAnswers)-1] == "" {
		groupAnswers = groupAnswers[:len(groupAnswers)-1]
	}

	answerMap := make(map[string]int)
	for _, personAnswers := range groupAnswers {
		for _, answer := range personAnswers {
			if everyone {
				answerMap[string(answer)]++
			} else {
				answerMap[string(answer)] = 1
			}
		}
	}

	if everyone {
		for _, count := range answerMap {
			if count == len(groupAnswers) {
				answerCount++
			}
		}
	} else {
		answerCount = len(answerMap)
	}

	return answerCount
}
