package session

import (
	"github.com/nobypass/fds-bot/internal/bot/models"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strconv"
)

func (c *FDSConnection) Verify(sp opentracing.Span, input *models.VerifyRequest) (*models.VerifyResponse, error) {
	req, err := c.newRequest(http.MethodPost, "/discord/verify", input, sp)
	if err != nil {
		return nil, err
	}

	return do[models.VerifyResponse](req)
}

func (c *FDSConnection) Daily(sp opentracing.Span, id string) (*models.MemberResponse, error) {
	req, err := c.newRequest(http.MethodPost, "/discord/daily/"+id, nil, sp)
	if err != nil {
		return nil, err
	}

	return do[models.MemberResponse](req)
}

func (c *FDSConnection) Login(sp opentracing.Span, pwd string) (*models.LoginResponse, error) {
	req, err := c.newRequest(http.MethodPost, "/discord/bot-login", &models.LoginRequest{Pwd: pwd}, sp)
	if err != nil {
		return nil, err
	}

	return do[models.LoginResponse](req)
}

func (c *FDSConnection) Leaderboard(sp opentracing.Span, page int) (*models.LeaderboardResponse, error) {
	req, err := c.newRequest(http.MethodGet, "/discord/leaderboard/"+strconv.Itoa(page), nil, sp)
	if err != nil {
		return nil, err
	}

	return do[models.LeaderboardResponse](req)
}

func (c *FDSConnection) Member(sp opentracing.Span, id string) (*models.MemberResponse, error) {
	req, err := c.newRequest(http.MethodGet, "/discord/member/"+id, nil, sp)
	if err != nil {
		return nil, err
	}

	return do[models.MemberResponse](req)
}
