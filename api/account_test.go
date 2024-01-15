package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/rensawamo/grpc-api/db/mock"
	db "github.com/rensawamo/grpc-api/db/sqlc"
	"github.com/rensawamo/grpc-api/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	ctrl := gomock.NewController(t) // これは mockgenから作成される関数
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	// スタブの作成
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// スタート
	server := NewServer(store)
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}

// ランダムでアカウントを作成
func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
