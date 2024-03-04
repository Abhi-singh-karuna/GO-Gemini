package main

import (
	"fmt"
	gemini "github.com/Limit-LAB/go-gemini"
	"github.com/Limit-LAB/go-gemini/models"
)

func main() {

	var yourQuestion string
	fmt.Println("Please enter your question those you ask with assistance :  ")
	
	fmt.Scanln(&yourQuestion)
	cli := gemini.NewClient("AIzaSyD9G87j3-_kkhPW4ROB5VHLHetLx0HF6EQ")
	rst, err := cli.GenerateContent(models.GeminiPro,
		models.NewGenerateContentRequest(
			models.NewContent(models.RoleUser, models.NewTextPart(yourQuestion)),
		),
	)
	if err != nil {
		panic(err)
	}

	x := rst.Candidates[0].Content.Parts[0].Text
	fmt.Println(&rst.Candidates[0].Content.Parts[0])
	// fmt.Println(&rst.Candidates[0].Content.Parts[0].Text)

	fmt.Println(*x)



}