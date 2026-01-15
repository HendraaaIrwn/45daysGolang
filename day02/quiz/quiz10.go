package main

import "fmt"

func displayTableString(matrix [9][11]string) {
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("%3s", matrix[row][col])
		}
		fmt.Println()
	}
}

func tableInit(matrix [9][11]string) [9][11]string {
	for row := range matrix {
		for col := range matrix[row] {
			if row == 0 && col == 0 {
				matrix[row][col] = "         "
			} else if col == 0 {
				matrix[row][col] = fmt.Sprintf("Student-%d", row-1)
			} else if row == 0 {
				matrix[row][col] = fmt.Sprintf("%d", col-1)
			} else {
				matrix[row][col] = " "
			}
		}
	}
	return matrix
}

func assignValue(matrix [9][11]string) [9][11]string {
	students := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"}, // Student-0 => row 1
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"}, // Student-1 => row 2
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"}, // Student-2 => row 3
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"}, // Student-3 => row 4
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"}, // Student-4 => row 5
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"}, // Student-5 => row 6
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"}, // Student-6 => row 7
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"}, // Student-7 => row 8
	}

	for sIdx, answers := range students {
		row := sIdx + 1 // row 1..8
		for i, v := range answers {
			matrix[row][i+1] = v // col 1..10
		}
	}

	return matrix
}


func countRightAnswer(studentAnswers []string, answerKey []string) int {
	count := 0
	for i := range studentAnswers {
		if studentAnswers[i] == answerKey[i] {
			count++
		}
	}
	return count
}

func getStudentAnswers(matrix [9][11]string, studentIndex int) []string {
	answers := []string{}
	for col := 1; col < len(matrix[studentIndex]); col++ {
		answers = append(answers, matrix[studentIndex][col])
	}
	return answers
}

func main() {
	var matrix [9][11]string
	var answerKey = []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	matrix = tableInit(matrix)
	matrix = assignValue(matrix)

	displayTableString(matrix)

	fmt.Println("\njawaban benar masing-masing siswa:")
	for studentRow := 1; studentRow <= 8; studentRow++ {
		studentName := matrix[studentRow][0] // "Student-x"
		score := countRightAnswer(getStudentAnswers(matrix, studentRow), answerKey)
		fmt.Printf("Jawaban %s yang benar : %d\n", studentName, score)
	}
}
