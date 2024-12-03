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
	id := "id_1"
	requestPath := "/editUser"
	queryParameter := "id=" + id

	var expected error = nil
	c, rec := newRouter()
	up := &UserPresenter{c: c}

	/* Act */
	actual := up.OutputSignupWithAuth(id)

	/* Assert */
	// up.OutputLoginResultがJSONを返すこと
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.Contains(t, rec.HeaderMap["Location"][0], requestPath)
		assert.Contains(t, rec.HeaderMap["Location"][0], queryParameter)
		assert.Equal(t, expected, actual)
	}
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
