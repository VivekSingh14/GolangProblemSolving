package p2

import (
	"demo1/package4/p1"
	"fmt"
)

type RepoService struct {
	namerepository p1.PersonRepository
}

func (r RepoService) GetNamesByService() { ///now to call this method from main.go
	result := r.namerepository.GetUser(23)
	fmt.Println("Final result ", result)
}
