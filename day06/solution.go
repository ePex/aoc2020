package main

func CalculateYesCount(batch []string) (answerCount int) {
	var groupAnswers []string
	for idx, answers := range batch {
		groupAnswers = append(groupAnswers, answers)
		if answers == "" || idx+1 == len(batch) {
			answerCount += countGroup(groupAnswers)
			groupAnswers = []string{}
		}
	}

	return answerCount
}

func countGroup(groupAnswers []string) int {
	answerMap := make(map[string]int)
	for _, personAnswers := range groupAnswers {
		for _, answer := range personAnswers {
			answerMap[string(answer)] = 1
		}
	}

	return len(answerMap)
}
