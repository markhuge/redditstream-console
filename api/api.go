package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Comment is the data structure we want, burried in Reddit's other bullshit.
type Comment struct {
	Author string
	Body   string
	Score  int
	Id     string
}

// Implement stringer to format output
func (c Comment) String() string {
	return fmt.Sprintf("author: %s - %d points\n%s\n\n", c.Author, c.Score, c.Body)
}

// Jesus fuck, reddit.
type Payload []struct {
	Data struct {
		Children []struct {
			Data Comment
		}
	}
}

// I'm really creative with these function names
func Get(url string) ([]Comment, error) {
	var comments []Comment
	client := &http.Client{}
	// TODO specify new posts here instead of in URL
	// TODO pagination?
	// TODO cache by id

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Redditstream/1.0")
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	log.Print(res.StatusCode)
	p := new(Payload)
	json.NewDecoder(res.Body).Decode(p)

	// TODO this seems wrong
	for _, post := range *p {
		for _, comment := range post.Data.Children {
			comments = append(comments, comment.Data)
		}
	}

	return comments, nil
}

// This is either gonna get weird or go away
func Print(comments []Comment) {
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
