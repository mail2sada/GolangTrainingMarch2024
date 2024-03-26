package main

import (
	"encoding/xml"
	"fmt"
)

type Marks struct {
	S1    int     `xml:"Maths"`
	S2    int     `xml:"Science"`
	S3    int     `xml:"ComputerScience"`
	Total int     `xml:"Total"`
	Avg   float64 `xml:"Avergae"`
}

type Student struct {
	RollNo int    `xml:"roll_no"`
	Name   string `xml:"name"`
	Class  string `xml:"class"`
	M      Marks  `xml:"marks"`
}

func (m *Marks) SetMarks(s1, s2, s3 int) {
	m.S1 = s1
	m.S2 = s2
	m.S3 = s3
	m.Total = m.S1 + m.S2 + m.S3
	m.Avg = float64(m.Total) / 3
}

func main() {
	fmt.Println("Xml emcoding...")
	std := new(Student)
	std.RollNo = 1
	std.Name = "Alok"
	std.Class = "10B"
	std.M.SetMarks(99, 100, 97)

	data, err := xml.Marshal(std)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	data, err = xml.MarshalIndent(std, "\n", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	//fmt.Println(std)
}
