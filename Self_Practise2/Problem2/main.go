package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []float64
}

func main() {
	fmt.Println("Problem 2")
	classroom := make(map[string]Student)

	classroom["Student123"] = Student{
		Name:   "Gagan",
		Grades: []float64{90, 80, 78, 90, 93},
	}

	classroom["Student1231"] = Student{
		Name:   "Shona",
		Grades: []float64{94, 88, 92, 97, 93},
	}

	classroom["Student12351"] = Student{
		Name:   "Some1",
		Grades: []float64{50, 70, 54, 50, 43},
	}

	var std string
	fmt.Println("\nEnter studentId")
	fmt.Scan(&std)

	result, grades := GetStudentAverage(classroom, std)
	fmt.Println(result, grades)
}

func CalculateAverage(grades []float64) (float64, error) {
	if grades == nil {
		return 0.0, fmt.Errorf("no Grades Recorded")
	} else {
		var average, sum float64
		for i := 1; i < len(grades); i++ {
			sum += grades[i]
		}
		average = sum / float64(len(grades))
		return average, nil
	}
}

func GetStudentAverage(classroom map[string]Student, studentID string) (float64, error) {
	_, present := classroom[studentID]
	if !present {
		return 0.0, fmt.Errorf("student not found")
	} else {
		result, grades := CalculateAverage(classroom[studentID].Grades)
		return result, grades
	}
}
