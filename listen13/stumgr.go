package main

import (
	"fmt"
	"os"
)

type Student struct {
	Username string
	Score float32
	Grade string
	Sex int
}

type StudentMgr struct{
	AllStudents []*Student
}

var StuMgr = new(StudentMgr)

func NewStudent(Username string, Score float32, Grade string, Sex int) (stu *Student) {
	stu = &Student{
		Username: Username,
		Score: Score,
		Grade: Grade,
		Sex: Sex,
	}
	return
}

func createStudent() *Student {
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
	return stu
}

func (p *StudentMgr) AddStudent(stu *Student) (err error) {
	for _, value := range p.AllStudents{
		if value.Username == stu.Username{
			err = fmt.Errorf("Student is already exists!\n")
			return
		}
	}
	p.AllStudents = append(p.AllStudents, stu)	
	return
}

func (p *StudentMgr) ShowStudentList() {
	for i := 0; i< len(p.AllStudents); i++ {
		fmt.Printf("%#v\n", p.AllStudents[i])
	}
}

func (p *StudentMgr) ModifyStudent(stu *Student) (err error) {
	for index, value := range p.AllStudents{
		if value.Username == stu.Username{
			p.AllStudents[index] = stu
			fmt.Println("Modify Sucess!")
			return
		}
	}
	err = fmt.Errorf("Student is not exists!\n")
	return
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
	var	sel int
	for {
		fmt.Println(Tips)
		fmt.Printf("Please Input Your Select: ")
		fmt.Scanf("%d\n", &sel)

		switch sel {
		case 1:
			stu := createStudent()
			err := StuMgr.AddStudent(stu)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			stu := createStudent()
			err := StuMgr.ModifyStudent(stu)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			StuMgr.ShowStudentList()
		case 0:
			os.Exit(0)
		}	
	}
	
}


func main() {
	Run()
}
