package main

import (
	"fmt"
	"os"
	"time"
)

type EndToEndId struct {
	operation_type string
	ispb           string
	date           string
	random         string
}

func parse_end_to_end_id(end_to_end_id string) *EndToEndId {
	return &EndToEndId{
		operation_type: end_to_end_id[0:1],
		ispb:           end_to_end_id[1:9],
		date:           end_to_end_id[9:21],
		random:         end_to_end_id[21:],
	}
}

func ent_to_end_id_operation_type_map(operation_type string) string {
	switch operation_type {
	case "E":
		return "Pix Efetuado"
	case "D":
		return "Pix Devolução"
	default:
		return "Unmapped Operation Type"
	}
}

func end_to_end_id_date_map(date string) string {
	t, err := time.Parse("200601021504", date)
	if err != nil {
		return "Error parsing date"
	}

	return t.Format("02/01/2006 15:04")
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <endToEndId>")
		return
	}

	end_to_end_id := args[1]

	if len(end_to_end_id) != 32 {
		fmt.Println("EndToEndId must have 32 characters")
		return
	}

	parsed_end_to_end_id := parse_end_to_end_id(end_to_end_id)

	fmt.Println("Operation Type: ", ent_to_end_id_operation_type_map(parsed_end_to_end_id.operation_type))
	fmt.Println("ISPB: ", parsed_end_to_end_id.ispb)
	fmt.Println("Date: ", end_to_end_id_date_map(parsed_end_to_end_id.date))
	fmt.Println("Random: ", parsed_end_to_end_id.random)
}
