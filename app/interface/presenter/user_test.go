package presenter

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/tayusa/notugly_backend/app/domain"
)

func TestResponseUser(t *testing.T) {
	userPresenter := NewUserPresenter()
	inputUser := domain.User{Id: "1", Name: "John", Sex: 1, Age: 24}

	body, err := userPresenter.ResponseUser(inputUser)
	if err != nil {
		log.Fatalln(err)
	}

	outputUser := domain.User{}
	if err := json.Unmarshal(body, &outputUser); err != nil {
		log.Fatalln(err)
	}

	if inputUser.Id != outputUser.Id ||
		inputUser.Name != outputUser.Name ||
		inputUser.Sex != outputUser.Sex ||
		inputUser.Age != outputUser.Age {
		t.Errorf("input: %s, %s, %d, %d output: %s, %s, %d, %d",
			inputUser.Id, inputUser.Name, inputUser.Sex, inputUser.Age,
			outputUser.Id, outputUser.Name, outputUser.Sex, outputUser.Age)
	}
}
