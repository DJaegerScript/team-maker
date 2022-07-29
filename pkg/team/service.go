package team

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"mime/multipart"
	"team-maker/utils"
)

type Service interface {
	GetAllTeams() (err error, status int, res interface{})
	GetTeam(id int) (err error, status int, res interface{})
	CreateTeam(name string, file *multipart.FileHeader) (err error, status int, res interface{})
	UpdateTeam(id int, name string) (err error, status int, res interface{})
	DeleteTeam(id int) (err error, status int, res interface{})
}

type service struct {
	Repo Repo
}

func NewService(repo Repo) *service {
	svc := &service{
		Repo: repo,
	}

	return svc
}

func (s *service) GetAllTeams() (err error, status int, res interface{}) {
	status = fiber.StatusOK

	err, res = s.Repo.GetAllTeams()
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func (s *service) GetTeam(id int) (err error, status int, res interface{}) {
	status = fiber.StatusOK

	err, res = s.Repo.GetTeam(id)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func (s *service) CreateTeam(name string, file *multipart.FileHeader) (err error, status int, res interface{}) {
	status = fiber.StatusCreated

	type user struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"LastName"`
	}

	type data struct {
		Users []user `json:"users"`
	}

	var resData data

	err, status, body := utils.HttpClient(fiber.MethodGet, "https://dummyjson.com/users?limit=100")

	err = json.Unmarshal(body, &resData)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	var members []string

	for i := 0; i < 4; i++ {
		member := resData.Users[rand.Intn(len(resData.Users))]
		members = append(members, fmt.Sprintf("%s %s", member.FirstName, member.LastName))
	}

	err, imageUrl := utils.S3Client(file, name)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	team := Team{
		Name:    name,
		Members: members,
		Image:   imageUrl,
	}

	err = s.Repo.CreateTeam(team)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func (s *service) UpdateTeam(id int, name string) (err error, status int, res interface{}) {
	status = fiber.StatusOK

	team := Team{
		ID:   id,
		Name: name,
	}

	err = s.Repo.UpdateTeam(team)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func (s *service) DeleteTeam(id int) (err error, status int, res interface{}) {
	status = fiber.StatusOK

	err = s.Repo.DeleteTeam(id)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}
