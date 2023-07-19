package goym

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UtilTestSuite struct {
	suite.Suite
	require *require.Assertions
}

func (s *UtilTestSuite) SetupSuite() {
	s.require = s.Require()
}

func (s *UtilTestSuite) TestI2s() {
	var i int = 1234
	var i32 int32 = 15135135
	var i64 int64 = 531135531
	res := iToString(i)
	s.require.Equal(res, "1234")
	res = iToString(i32)
	s.require.Equal(res, "15135135")
	res = iToString(i64)
	s.require.Equal(res, "531135531")
}

func (s *UtilTestSuite) TestGenApiPath() {
	const expected = "https://api.music.yandex.net/users/1234/playlists/list"
	result := genApiPath("users", iToString(1234), "playlists", "list")
	s.require.Equal(expected, result)
}
