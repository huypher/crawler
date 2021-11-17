package voz

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/huypher/crawler/internal/cache"
)

var (
	ctx = context.Background()
)

func Test_executor_Do(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		url  string
	}{
		{"test", "https://voz.vn/t/bh-media-phan-phao-chuyen-vtv-noi-minh-nhan-vo-ban-quyen-quoc-ca.425411/"},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cache, _, err := cache.NewCache(&cache.Config{
				Addr: "127.0.0.1:6379",
				Pass: "",
				DB:   0,
			})
			require.Nil(t, err)

			e := &executor{
				cache: cache,
			}

			got, err := e.Do(ctx, tc.url)
			require.Nil(t, err)
			require.NotEqual(t, "", got)
		})
	}
}
