package session

import (
	"github.com/nobypass/fds-bot/internal/bot/model"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strconv"
)

func (c *FDSConnection) Verify(sp opentracing.Span, input *model.VerifyRequest) (*model.VerifyResponse, error) {
	req, err := c.newRequest(http.MethodPost, "/discord/verify", input, sp)
	if err != nil {
		return nil, err
	}

	return do[model.VerifyResponse](req)
}

func (c *FDSConnection) Daily(sp opentracing.Span, id string) (*model.MemberResponse, error) {
	req, err := c.newRequest(http.MethodPatch, "/discord/daily/"+id, nil, sp)
	if err != nil {
		return nil, err
	}

	return do[model.MemberResponse](req)
}

func (c *FDSConnection) Login(sp opentracing.Span, pwd string) (*model.LoginResponse, error) {
	req, err := c.newRequest(http.MethodPost, "/discord/bot-login", &model.LoginRequest{Pwd: pwd}, sp)
	if err != nil {
		return nil, err
	}

	return do[model.LoginResponse](req)
}

func (c *FDSConnection) Leaderboard(sp opentracing.Span, page int) (*model.LeaderboardResponse, error) {
	req, err := c.newRequest(http.MethodGet, "/discord/leaderboard/"+strconv.Itoa(page), nil, sp)
	if err != nil {
		return nil, err
	}

	return do[model.LeaderboardResponse](req)
}

func (c *FDSConnection) Member(sp opentracing.Span, id string) (*model.MemberResponse, error) {
	req, err := c.newRequest(http.MethodGet, "/discord/member/"+id, nil, sp)
	if err != nil {
		return nil, err
	}

	return do[model.MemberResponse](req)
}

func (c *FDSConnection) Revoke(sp opentracing.Span, id string) (*model.MemberResponse, error) {
	req, err := c.newRequest(http.MethodDelete, "/discord/revoke/"+id, nil, sp)
	if err != nil {
		return nil, err
	}

	return do[model.MemberResponse](req)
}
