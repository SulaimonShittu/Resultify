package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Shittu",
			LastName:  "Zami",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 85},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 99},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 98},
			},
		},
		Student{
			ID:        2,
			FirstName: "Adebayo",
			LastName:  "Salami",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 72},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 88},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 91},
			},
		},
		Student{
			ID:        3,
			FirstName: "Folake",
			LastName:  "Adebola",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 90},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 95},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 85},
			},
		},
		Student{
			ID:        4,
			FirstName: "Ikenna",
			LastName:  "Eze",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 81},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 92},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 89},
			},
		},
		Student{
			ID:        5,
			FirstName: "Ngozi",
			LastName:  "Okeke",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 76},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 86},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 80},
			},
		},
		Student{
			ID:        6,
			FirstName: "Adamu",
			LastName:  "Suleiman",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 84},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 90},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 87},
			},
		},
		Student{
			ID:        7,
			FirstName: "Fatima",
			LastName:  "Abubakar",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 95},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 97},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 94},
			},
		},
		Student{
			ID:        8,
			FirstName: "Bello",
			LastName:  "Musa",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 83},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 89},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 86},
			},
		},
		Student{
			ID:        9,
			FirstName: "Uchenna",
			LastName:  "Obi",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 79},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 93},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 88},
			},
		},
		Student{
			ID:        10,
			FirstName: "Aisha",
			LastName:  "Mohammed",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 77},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 84},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 79},
			},
		},
		Student{
			ID:        11,
			FirstName: "Bimpe",
			LastName:  "Olowo",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 89},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 94},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 90},
			},
		},
		Student{
			ID:        12,
			FirstName: "Chinedu",
			LastName:  "Okafor",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 85},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 88},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 82},
			},
		},
		Student{
			ID:        13,
			FirstName: "Deborah",
			LastName:  "Ojo",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 78},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 80},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 77},
			},
		},
		Student{
			ID:        14,
			FirstName: "Emeka",
			LastName:  "Nwosu",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 88},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 91},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 87},
			},
		},
		Student{
			ID:        15,
			FirstName: "Gbemi",
			LastName:  "Adegoke",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 82},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 85},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 80},
			},
		},
		Student{
			ID:        16,
			FirstName: "Isaac",
			LastName:  "Omole",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 91},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 96},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 92},
			},
		},
		Student{
			ID:        17,
			FirstName: "Kehinde",
			LastName:  "Balogun",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 73},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 82},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 79},
			},
		},
		Student{
			ID:        18,
			FirstName: "Oluwatobi",
			LastName:  "Fashola",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 87},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 91},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 89},
			},
		},
		Student{
			ID:        19,
			FirstName: "Rasheed",
			LastName:  "Afolabi",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 92},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 93},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 95},
			},
		},
		Student{
			ID:        20,
			FirstName: "Taiwo",
			LastName:  "Babatunde",
			Grades: []Grade{
				Grade{Title: "Quiz 1", Type: GradeQuiz, Score: 74},
				Grade{Title: "Week 1 Homework", Type: GradeHomework, Score: 83},
				Grade{Title: "Quiz 2", Type: GradeQuiz, Score: 78},
			},
		},
	}
}
