package test

import (
	"binginx.com/brush/config"
	"binginx.com/brush/internal/clients"
	"binginx.com/brush/internal/service"
	"binginx.com/brush/model"
	service2 "binginx.com/brush/service"
	"testing"
)

const (
	_token = "eyJhbGciOiJSUzI1NiIsInR5cCI6ImF0K2p3dCJ9.eyJuYmYiOjE2NjkwOTcxMzQsImV4cCI6MTY2OTEwMDczNCwiaXNzIjoiaHR0cHM6Ly9pZHMuY2hpbmEtaW52LmNuIiwiYXVkIjpbImVjb3N5cyIsImVjb3N5c3JjIiwic2VuZGNvZGUiLCJzdmMiXSwiY2xpZW50X2lkIjoiZWNvc3lzIiwic3ViIjoibGliaW5Ac2RqaWN0ZWMuY29tIiwiYXV0aF90aW1lIjoxNjY5MDg2ODAwLCJpZHAiOiJsb2NhbCIsIm5hbWUiOiLmnY7lvawiLCJlbWFpbCI6ImxpYmluQHNkamljdGVjLmNvbSIsImRlcGFydG1lbnQiOiLkuLTml7bpg6jpl6giLCJjb21wYW55Ijoi5bu65oqV5pWw5o2u56eR5oqA77yI5bGx5Lic77yJ5pyJ6ZmQ5YWs5Y-4IiwiZGl2aXNpb24iOiIiLCJlbXBsb3llZWlkIjoiMTIzODk1IiwiZW1haWxfdmVyaWZpZWQiOiJ0cnVlIiwianRpIjoiNzZERUNDOTMzOEZBMDY4RTM2RERDMjE2RERBNEI0NkIiLCJpYXQiOjE2NjkwODY4MDAsInNjb3BlIjpbImVjb3N5cyIsImVjb3N5c3JjIiwib3BlbmlkIiwic2VuZGNvZGUiLCJzdmMiLCJvZmZsaW5lX2FjY2VzcyJdLCJhbXIiOlsiU21zQ29kZSJdfQ.hlxllOnWvlvxx_KOozvu2i8U7IXn1yvNGFgPgXuJ0QiXcrmfmpEGpWXByxPmj7MbddzpYcrBRpOU7yv_RqLhKNFo8vKFWEQdGyNi-BhdXS63G_HkbcX9dhHCYnM_bKuzXk7AQoQyymOVzR7zHRDFM-FuQSx-uouwPJ_jSUbx1uVa1XondJJDpBZQcomlyzAEDGCYaK6pPDTnX-SFxJ1k-JHfo2_OJfGlRQa5u3lw2M8f9Rko1nO3BSJ79kIFqpuS0IU571iGcLbvQ5ZKgXDgnBr6YtKf5RoQg98kaGRKQfGnwYbMmMSAjkex-4xG66fRd_XlwglAWgNJQlCBfGMQsg"
)

func Test_CreateUser(t *testing.T) {
	config.MustInit()
	clients.MustInit()
	user := &model.User{
		Name:  "test",
		Token: service2.Token,
		Score: 5566,
	}
	err1 := service.CreateUser(user)
	if err1 != nil {
		t.Log(err1)
	}
	getUser, err2 := service.GetUserByToken(user.Token)
	if err2 == nil {
		t.Log(getUser)
	}
}

func Test_DB(t *testing.T) {
	config.MustInit()
	clients.MustInit()

}
