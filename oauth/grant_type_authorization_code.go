package oauth

import (
	"errors"
	"net/http"

	"github.com/artpar/go-oauth2-server/response"
	"github.com/artpar/go-oauth2-server/util"
)

var (
	// ErrInvalidRedirectURI ...
	ErrInvalidRedirectURI = errors.New("Invalid redirect URI")
)

func (s *Service) authorizationCodeGrant(w http.ResponseWriter, r *http.Request, client *Client) {
	// Fetch the authorization code
	authorizationCode, err := s.getValidAuthorizationCode(
		r.Form.Get("code"), // authorization code
		client,             // client
	)
	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect URI must match if it was used to obtain the authorization code
	if util.StringOrNull(r.Form.Get("redirect_uri")) != authorizationCode.RedirectURI {
		response.Error(w, ErrInvalidRedirectURI.Error(), http.StatusBadRequest)
		return
	}

	// Log in the user
	accessToken, refreshToken, err := s.Login(
		authorizationCode.Client,
		authorizationCode.User,
		authorizationCode.Scope,
	)
	if err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the authorization code
	s.db.Unscoped().Delete(&authorizationCode)

	// Write the JSON access token to the response
	accessTokenRespone := &AccessTokenResponse{
		ID:           accessToken.ID,
		AccessToken:  accessToken.Token,
		ExpiresIn:    s.cnf.Oauth.AccessTokenLifetime,
		TokenType:    TokenType,
		Scope:        accessToken.Scope,
		RefreshToken: refreshToken.Token,
	}
	if accessToken.User != nil {
		accessTokenRespone.UserID = accessToken.User.ID
	}
	response.WriteJSON(w, accessTokenRespone, 200)
}
