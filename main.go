package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Comment struct {
	Author string
	Body   string
	Score  int
	Id     string
}

func (c Comment) String() string {
	return fmt.Sprintf("author: %s - %d points\n%s\n\n", c.Author, c.Score, c.Body)
}

// Jesus fuck, reddit
type Payload []struct {
	Data struct {
		Children []struct {
			Data Comment
		}
	}
}

// Get shit by url
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

	for _, post := range *p {
		for _, comment := range post.Data.Children {
			comments = append(comments, comment.Data)
		}
	}

	return comments, nil
}

func main() {

	p, err := Get("https://www.reddit.com/r/MMA/comments/523s3d/official_ufc_203_miocic_vs_overeem/.json?sort=new&limit=500")
	if err != nil {
		log.Fatal(err)
	}
	for _, comment := range p {
		fmt.Println(comment)
	}
}
