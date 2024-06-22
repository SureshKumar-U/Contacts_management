package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// var templ *template.Template

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}

// func ParseTemplates(filepath string) error {

// 	rootDir, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}

// 	templateDir :=

// 	fmt.Println(templateDir)

// 	// Parse templates from the determined directory

// }
