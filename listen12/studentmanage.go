package main

import (
	"fmt"
	"os"
)

var	sel int
var	Students []*Student

type Student struct {
	Username string
	Score float32
	Grade string
	Sex int
}

func NewStudent(Username string, Score float32, Grade string, Sex int) *Student {
	stu := &Student{
		Username: Username,
		Score: Score,
		Grade: Grade,
		Sex: Sex,
	}
	return stu
}

func AddStudent() {
	var (
		Username string
		Score float32
		Grade string
		Sex int
	)
	fmt.Printf("Username:")
	fmt.Scanf("%s\n", &Username)

	fmt.Printf("Score:")
	fmt.Scanf("%f\n", &Score)

	fmt.Printf("Grade:")
	fmt.Scanf("%s\n", &Grade)

	fmt.Printf("Sex:")
	fmt.Scanf("%d\n", &Sex)

	stu := NewStudent(Username, Score, Grade, Sex)
	Students = append(Students, stu)	
}

func ShowStudentList() {
	// fmt.Printf("%#v", Students)
	for i := 0; i< len(Students); i++ {
		fmt.Printf("%#v\n", Students[i])
	}
}

func ModifyStudent() {
	// fmt.Println("Modify")
	var (
		Username string
		Score float32
		Grade string
		Sex int
	)
	fmt.Printf("Username:")
	fmt.Scanf("%s\n", &Username)

	fmt.Printf("Score:")
	fmt.Scanf("%f\n", &Score)

	fmt.Printf("Grade:")
	fmt.Scanf("%s\n", &Grade)

	fmt.Printf("Sex:")
	fmt.Scanf("%d\n", &Sex)

	for i := 0; i< len(Students); i++ {
		if Students[i].Username == Username{
			Students[i].Username = Username
			Students[i].Score = Score
			Students[i].Grade = Grade
			Students[i].Sex = Sex
			return
		}
	}
}


func Run(){
	Tips := `
	##########################
	# 1.Add Student          #
	# 2.Modify Student       #
	# 3.Show Student List    #
	#                        #
	# 0.Exit                 #
	##########################
	`
	for {
		fmt.Println(Tips)
		fmt.Printf("Please Input Your Select: ")
		fmt.Scanf("%d\n", &sel)

		switch sel {
		case 1:
			AddStudent()
		case 2:
			ModifyStudent()
		case 3:
			ShowStudentList()
		case 0:
			os.Exit(0)
		}	
	}
	
}


func main() {
	Run()
}
