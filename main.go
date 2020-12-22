package main

import "fmt"

var (
	name                        string
	opc, cN, g1, g2, g3, search int
	exit                        bool
	stdt                        = make([]student, 0)
)

type student struct {
	name                            string
	ctrlNum, grade1, grade2, grade3 int
}

func main() {
	for exit == false {
		fmt.Println("Choose an action")
		fmt.Println("1. Add student")
		fmt.Println("2. Search student")
		fmt.Println("3. Delete student")
		fmt.Println("4. List student")
		fmt.Println("5. Exit")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			add()
		case 2:
			showData(searchStudent())
		case 3:
			deleteStudent(searchStudent())
		case 4:
			listStudents()
		case 5:
			fmt.Println("E X I T . . .")
			exit = true
		default:
			fmt.Println("Option not valid")
		}
	}
}

func add() {
	fmt.Println("Enter control number")
	fmt.Scan(&cN)
	fmt.Println("Enter name")
	fmt.Scan(&name)
	fmt.Println("Enter 1st grade")
	fmt.Scan(&g1)
	fmt.Println("Enter 2nd grade")
	fmt.Scan(&g2)
	fmt.Println("Enter 3rd grade")
	fmt.Scan(&g3)

	stdt = append(stdt, student{ctrlNum: cN, name: name, grade1: g1, grade2: g2, grade3: g3})

}

func searchStudent() int {
	fmt.Println("Enter ctrl. Number to search")
	fmt.Scan(&search)
	for i, v := range stdt {
		if search == v.ctrlNum {
			return i
		}
	}
	return -1
}

func showData(search int) {
	if search >= 0 {
		fmt.Println("Name :", stdt[search].name)
		fmt.Println("Ctrl. Number :", stdt[search].ctrlNum)
		fmt.Println("1st Grade :", stdt[search].grade1)
		fmt.Println("2nd Grade:", stdt[search].grade2)
		fmt.Println("3rd Grade :", stdt[search].grade3)
	} else {
		fmt.Println("Data not found")
	}
}

func deleteStudent(search int) {
	if search >= 0 {
		stdt = append(stdt[:search], stdt[search+1:]...)
		fmt.Println("D E L E T I N G . . .")
	} else {
		fmt.Println("Data not found")
	}
}

func listStudents() {
	fmt.Println(stdt)
}
