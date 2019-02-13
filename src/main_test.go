package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/tayusa/notugly_backend/configs"
	"github.com/tayusa/notugly_backend/internal/domain"
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/router"
	"github.com/tayusa/notugly_backend/internal/infrastructure/repository/dummy"
	"github.com/tayusa/notugly_backend/internal/registry"
	"github.com/tayusa/notugly_backend/pkg/ctx"
)

const (
	dummyUserId = "A1"
	imageName   = "test.png"
)

var (
	testRouter *httprouter.Router
)

func dummyAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		next(
			w,
			r.WithContext(
				ctx.SetUserId(r.Context(), dummyUserId)),
			p)
	}
}

func setUp() {
	configs.LoadConfig(configPath)

	interactor := registry.NewDummyInteractor()
	handler := interactor.NewAppHandler()
	testRouter = router.NewRouter(handler, dummyAuth)
}

func toReader(model interface{}) (io.Reader, error) {
	var err error
	var output []byte

	switch value := model.(type) {
	case domain.User:
		output, err = json.Marshal(&value)
	case domain.Coordinate:
		output, err = json.Marshal(&value)
	default:
		err = fmt.Errorf("Unexpected type")
	}
	if err != nil {
		return bytes.NewReader([]byte{}), err
	}
	return bytes.NewReader(output), nil
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	dummyUsers, err := dummy.GetUsers(dummy.GET)
	if err != nil {
		log.Fatalln(err)
	}

	for _, dummyUser := range dummyUsers {
		request, err := http.NewRequest("GET", "/users/"+dummyUser.Id, nil)
		if err != nil {
			log.Fatalln(err)
		}

		writer := httptest.NewRecorder()
		testRouter.ServeHTTP(writer, request)

		if writer.Code != 200 {
			t.Errorf("Response code is %d", writer.Code)
		}

		var user domain.User
		json.Unmarshal(writer.Body.Bytes(), &user)
		if user.Id != dummyUser.Id ||
			user.Name != dummyUser.Name ||
			user.Sex != dummyUser.Sex ||
			user.Age != dummyUser.Age {

			t.Errorf("Cannot retrieve JSON user")
		}
	}
}

func TestPostUser(t *testing.T) {
	t.Parallel()

	dummyUsers, err := dummy.GetUsers(dummy.POST)
	if err != nil {
		log.Fatalln(err)
	}

	for _, dummyUser := range dummyUsers {
		body, err := toReader(dummyUser)
		if err != nil {
			log.Fatalln(err)
		}

		request, err := http.NewRequest("POST", "/users/me", body)
		if err != nil {
			log.Fatalln(err)
		}

		writer := httptest.NewRecorder()
		testRouter.ServeHTTP(writer, request)

		if writer.Code != 200 {
			t.Errorf("Response code is %d", writer.Code)
		}
	}
}

func TestGetCoordinates(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip()
	}

	dummyCoordinates, err := dummy.GetCoordinates(dummy.GET)
	if err != nil {
		log.Fatalln(err)
	}

	coordinates := make([]domain.Coordinate, 0)
	for _, dummyCoordinate := range dummyCoordinates {
		if dummyCoordinate.UserId == dummyUserId {
			coordinates = append(coordinates, dummyCoordinate)
		}
	}

	for idx, coordinate := range coordinates {
		request, err := http.NewRequest("GET",
			"/users/"+coordinate.UserId+"/coordinates", nil)
		if err != nil {
			log.Fatalln(err)
		}

		writer := httptest.NewRecorder()
		testRouter.ServeHTTP(writer, request)

		if writer.Code != 200 {
			t.Errorf("Response code is %d", writer.Code)
		}

		var output []domain.Coordinate
		json.Unmarshal(writer.Body.Bytes(), &output)
		if output[idx].Id != coordinate.Id ||
			output[idx].ImageName != coordinate.ImageName ||
			output[idx].UserId != coordinate.UserId ||
			output[idx].CreatedAt != coordinate.CreatedAt ||
			output[idx].Favorites != coordinate.Favorites ||
			output[idx].IsFavorited != coordinate.IsFavorited {

			t.Errorf("Cannot retrieve JSON user")
		}
	}
}

func imageToBase64(path string) (string, error) {
	file, err := os.Open(path)
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func TestPostCoordinate(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip()
	}

	dummyCoordinates, err := dummy.GetCoordinates(dummy.POST)
	if err != nil {
		log.Fatalln(err)
	}

	image, err := imageToBase64(filepath.Join("test/images", imageName))
	if err != nil {
		log.Fatalln(err)
	}

	for _, dummyCoordinate := range dummyCoordinates {
		dummyCoordinate.Image = image
		body, err := toReader(dummyCoordinate)
		if err != nil {
			log.Fatalln(err)
		}

		request, err := http.NewRequest("POST", "/coordinates", body)
		if err != nil {
			log.Fatalln(err)
		}

		writer := httptest.NewRecorder()
		testRouter.ServeHTTP(writer, request)

		if writer.Code != 200 {
			t.Errorf("Response code is %d", writer.Code)
		}
	}
}
