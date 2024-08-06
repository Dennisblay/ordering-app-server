package api

import (
	"fmt"
	mockdb "github.com/Dennisblay/ordering-app-server/internal/database/mock"
	"github.com/Dennisblay/ordering-app-server/internal/database/models"
	"github.com/Dennisblay/ordering-app-server/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserByIDAPI(t *testing.T) {
	user := RandomUser()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		GetUserById(gomock.Any(), gomock.Eq(user.ID)).
		Times(1).
		Return(user, nil)

	// Start test server and send request
	server, err := NewServer(store)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%d", user.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// Check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func RandomUser() db.GetUserByIdRow {

	return db.GetUserByIdRow{
		ID:        int32(util.RandomInt(1, 1000)),
		FirstName: util.RandomUser(),
		LastName:  util.RandomUser(),
		Email:     util.RandomEmail(),
		Phone:     util.RandomPhone(),
		Address:   util.RandomAddress(),
	}
}
