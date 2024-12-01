package presenter

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseSetCookie(setCookie string) map[string]string {
	attributes := make(map[string]string)
	parts := strings.Split(setCookie, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		// "key=value"形式の属性を解析
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			attributes[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		} else {
			attributes[strings.TrimSpace(keyValue[0])] = ""
		}
	}
	return attributes
}

func TestOutputCreateResult(t *testing.T) {
	/* Arrange */
	expected := "{}\n"
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputCreateResult()

	/* Assert */
	// up.OutputCreateResultがJSONを返すこと
	if assert.NoError(t, actual) {
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestOutputUpdateResult(t *testing.T) {
	/* Arrange */
	expected := "{}\n"
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputCreateResult()

	/* Assert */
	if assert.NoError(t, actual) {
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestOutputLoginResult(t *testing.T) {
	/* Arrange */
	expected := "{\"userId\":\"id_1\"}\n"
	c, rec := newRouter()
	up := &UserPresenter{c: c}
	userId := "id_1"

	/* Act */
	actual := up.OutputLoginResult(userId)

	/* Assert */
	if assert.NoError(t, actual) {
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestOutputAuthUrl(t *testing.T) {
	/* Arrange */
	url := "https://www.google.com"
	expected := url
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputAuthUrl(url)

	/* Assert */
	// up.OutputAuthUrlがJSONを返すこと
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusFound, rec.Code)
		// リダイレクト先のURLが正しいこと
		assert.Equal(t, expected, rec.HeaderMap["Location"][0])
	}
}

func TestOutputSignupWithAuth(t *testing.T) {
	/* Arrange */
	token := "test_token"
	requestPath := "/editUser"

	var expected error = nil
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputSignupWithAuth(token)

	/* Assert */
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Contains(t, rec.HeaderMap["Location"][0], requestPath)
	assert.Equal(t, expected, actual)

	// レスポンスヘッダーからSet-Cookieを取得
	setCookie := rec.Header().Get("Set-Cookie")
	cookieAttributes := parseSetCookie(setCookie)
	assert.Equal(t, token, cookieAttributes["auth_token"])
}

func TestOutputAlreadySignedup(t *testing.T) {
	/* Arrange */
	var expected error = nil
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputAlreadySignedup()

	/* Assert */
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.Equal(t, expected, actual)
	}
}

func TestOutputHasEmailInRequestBody(t *testing.T) {
	/* Arrange */
	expected := "{\"error\":\"Email is included in Request Body\"}\n"
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputHasEmailInRequestBody()

	/* Assert */
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	if assert.NoError(t, actual) {
		assert.Equal(t, expected, rec.Body.String())
	}
}
