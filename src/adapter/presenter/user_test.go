package presenter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestOutputLoginResult(t *testing.T) {
	/* Arrange */
	expected := "{}\n"
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputLoginResult()

	/* Assert */
	// up.OutputLoginResultがJSONを返すこと
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
	var expected error = nil
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputSignupWithAuth()

	/* Assert */
	// up.OutputLoginResultがJSONを返すこと
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.Equal(t, expected, actual)
	}
}
