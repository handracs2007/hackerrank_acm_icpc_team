// https://www.hackerrank.com/challenges/acm-icpc-team/problem

package main

import "fmt"

func acmTeam(topic []string) []int {
	var result = make([]int, 2)
	var topicLength = len(topic[0])
	var topicCounts = make([]int, topicLength)
	var maxTopicCount = -1

	for i := 0; i < len(topic); i++ {
		for j := i + 1; j < len(topic); j++ {
			var topicCount = 0
			var topic1 = topic[i]
			var topic2 = topic[j]

			for k := 0; k < topicLength; k++ {
				if topic1[k] == '1' || topic2[k] == '1' {
					topicCount++
				}
			}

			if topicCount > maxTopicCount {
				maxTopicCount = topicCount
			}

			topicCounts[topicCount-1] ++
		}
	}

	result[0] = maxTopicCount
	result[1] = topicCounts[maxTopicCount-1]

	return result
}

func main() {
	fmt.Println(acmTeam([]string{"10101", "11100", "11010", "00101"}))
}
