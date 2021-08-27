package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name         string
	IdenticalNum string
	EnteringYear string
	Field        string
}

type Professor struct {
	Name         string
	IdenticalNum string
	Field        string
}

type Class struct {
	Name      string
	ClassId   string
	Field     string
	Professor *Professor
	Students  []Student
}

var students []Student
var classes []Class
var professors []Professor

func IsRegisteredStudent(id string) *Student {
	for _, student := range students {
		if student.IdenticalNum == id {
			return &student
		}
	}
	return nil
}

func IsRegisteredProfessor(id string) *Professor {
	for _, pro := range professors {
		if pro.IdenticalNum == id {
			return &pro
		}
	}
	return nil
}

func IsRegisteredClass(id string) *Class {
	for _, class := range classes {
		if class.ClassId == id {
			return &class
		}
	}
	return nil
}

func RegisterStudent() bool {
	var name, id, enteringYear, field string
	fmt.Scan(&name)
	fmt.Scan(&id)
	fmt.Scan(&enteringYear)
	fmt.Scan(&field)

	if IsRegisteredStudent(id) != nil {
		fmt.Println("this identical number previously registered")
		return false
	}

	students = append(students, Student{
		Name:         name,
		IdenticalNum: id,
		EnteringYear: enteringYear,
		Field:        field,
	})

	fmt.Printf("welcome to golestan")
	return true
}

func RegisterProfessor() bool {
	var name, id, field string

	fmt.Scan(&name)
	fmt.Scan(&id)
	fmt.Scan(&field)

	if IsRegisteredProfessor(id) != nil {
		fmt.Println("this identical number previously registered")
		return false
	}

	professors = append(professors, Professor{
		Name:         name,
		IdenticalNum: id,
		Field:        field,
	})

	fmt.Println("welcome to golestan")
	return true
}

func MakeClass() bool {
	var name, classId, field string

	fmt.Scan(&name)
	fmt.Scan(&classId)
	fmt.Scan(&field)

	if IsRegisteredClass(classId) != nil {
		fmt.Println("this class id previously used")
		return false
	}

	classes = append(classes, Class{
		Name:    name,
		ClassId: classId,
		Field:   field,
	})

	fmt.Printf("class added successfully")
	return true
}

func AddStudent() bool {
	var studentId, classId string

	fmt.Scan(&studentId)
	fmt.Scan(&classId)

	var student *Student
	if student = IsRegisteredStudent(studentId); student == nil {
		fmt.Println("invalid student")
		return false
	}

	var class *Class
	if class = IsRegisteredClass(classId); class == nil {
		fmt.Println("invalid class ")
		return false
	}

	if class.Field != student.Field {
		fmt.Println("student field is not match")
		return false
	}

	for _, s := range class.Students {
		if s.IdenticalNum == studentId {
			fmt.Println("student is already registered")
			return false
		}
	}

	class.Students = append(class.Students, *student)
	fmt.Println("student added successfully to the class")
	return true
}

func AddProfessor() bool {
	var professorId, classId string

	fmt.Scan(&professorId)
	fmt.Scan(&classId)

	var professor *Professor
	if professor = IsRegisteredProfessor(professorId); professor == nil {
		fmt.Println("invalid professor ")
		return false
	}

	var class *Class
	if class = IsRegisteredClass(classId); class == nil {
		fmt.Println("invalid class ")
		return false
	}

	if class.Field != professor.Field {
		fmt.Println("professor field is not match")
		return false
	}

	if class.Professor != nil {
		fmt.Println("this class has a professor")
		return false
	}

	class.Professor = professor
	fmt.Println("professor added successfully to the class")
	return true
}
func StudentStatus() bool {
	var studentId string

	fmt.Scan(&studentId)

	student := IsRegisteredStudent(studentId)
	if student == nil {
		fmt.Println("invalid student")
		return false
	}

	text := fmt.Sprintf("%s %s %s", student.Name, student.EnteringYear, student.Field)
	var classNames []string
	for _, class := range classes {
		for _, s := range class.Students {
			if s.IdenticalNum == student.IdenticalNum {
				classNames = append(classNames, class.Name)
				break
			}
		}
	}
	if len(classNames) > 0 {
		text += strings.Join(classNames, " ")
	}
	fmt.Printf(text)
	return true
}
func ProfessorStatus() bool {
	var professorId string

	fmt.Scan(&professorId)

	professor := IsRegisteredProfessor(professorId)
	if professor == nil {
		fmt.Println("invalid professor")
		return false
	}

	text := fmt.Sprintf("%s %s", professor.Name, professor.Field)
	var classNames []string
	for _, class := range classes {
		if class.Professor.IdenticalNum == professorId {
			classNames = append(classNames, class.Name)
		}
	}
	if len(classNames) > 0 {
		text += strings.Join(classNames, " ")
	}
	fmt.Printf(text)
	return true
}

func ClassStatus() bool {
	var classId string

	fmt.Scan(&classId)

	class := IsRegisteredClass(classId)
	if class == nil {
		fmt.Println("invalid class")
		return false
	}

	teacher := "None"
	if class.Professor != nil {
		teacher = class.Professor.Name
	}
	text := fmt.Sprintf("%s %s", class.Name, teacher)

	var studentNames []string
	for _, student := range class.Students {
		studentNames = append(studentNames, student.Name)
	}
	if len(studentNames) > 0 {
		text += strings.Join(studentNames, " ")
	}
	fmt.Printf(text)
	return true
}

func main() {

	m := map[string]func() bool{
		"register_student":   RegisterStudent,
		"register_professor": RegisterProfessor,
		"make_class":         MakeClass,
		"add_student":        AddStudent,
		"add_professor":      AddProfessor,
		"student_status":     StudentStatus,
		"professor_status":   ProfessorStatus,
		"class_status":       ClassStatus,

		// TODO: Phase 3
		// "set_final_mark": ,
		// "mark_student": ,
		// "mark_list": ,
		// "average_mark_professor": ,
		// "average_mark_student": ,
		// "top_mark": ,
		// "top_student": ,
	}

	for true {
		var action string
		fmt.Scan(&action)

		if action == "end" {
			return
		}

		if fn, ok := m[action]; ok {
			fn()
			continue
		}
	}

	return
}
