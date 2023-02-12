package goym

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	require *require.Assertions
}

func (s *ServiceTestSuite) SetupSuite() {
	s.require = s.Require()
}

func (s ServiceTestSuite) TestI2s() {
	var i int = 1234
	var i32 int32 = 15135135
	var i64 int64 = 531135531
	res := i2s(i)
	s.require.Equal(res, "1234")
	res = i2s(i32)
	s.require.Equal(res, "15135135")
	res = i2s(i64)
	s.require.Equal(res, "531135531")
}

func (s ServiceTestSuite) TestGenApiPath() {
	const expected = "https://api.music.yandex.net/users/1234/playlists/list"
	result := genApiPath([]string{"users", i2s(1234), "playlists", "list"})
	s.require.Equal(expected, result)
}
