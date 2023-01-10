package goym

import (
	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AlbumTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *AlbumTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s AlbumTestSuite) getAlbumId() int64 {
	res, err := s.cl.Search("crystal castles iii", 0, schema.SearchTypeAlbum, false)
	s.require.Nil(err)
	s.require.NotNil(res)
	s.require.NotEmpty(res.Albums.Results)
	return res.Albums.Results[0].ID
}

func (s AlbumTestSuite) getAlbumIds() []int64 {
	res, err := s.cl.Search("moby", 0, schema.SearchTypeAlbum, false)
	s.require.Nil(err)
	s.require.NotNil(res)
	s.require.NotEmpty(res.Albums.Results)
	var ids = []int64{}
	for i, al := range res.Albums.Results {
		ids = append(ids, al.ID)
		if i == 5 {
			break
		}
	}
	return ids
}

func (s AlbumTestSuite) TestGetAlbumById() {
	// without tracks
	var id = s.getAlbumId()
	data, err := s.cl.GetAlbumById(id, false)
	s.require.Nil(err)
	s.require.Positive(data.ID)

	// with tracks
	data, err = s.cl.GetAlbumById(231541, true)
	s.require.Nil(err)
	s.require.Positive(data.ID)
	s.require.NotEmpty(data.Volumes)
}

func (s AlbumTestSuite) TestGetAlbumsByIds() {
	var ids = s.getAlbumIds()
	albums, err := s.cl.GetAlbumsByIds(ids)
	s.require.Nil(err)
	s.require.NotEmpty(albums)
	s.require.Positive(albums[0].ID)
}

func (s AlbumTestSuite) TestLikeUnlikeAlbum() {
	res, err := s.cl.Search("mujuice downshifting", 0, schema.SearchTypeAlbum, false)
	s.require.Nil(err)
	s.require.NotNil(res)
	s.require.NotEmpty(res.Albums.Results)
	var al = res.Albums.Results[0]

	// like
	err = s.cl.LikeAlbum(al)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeAlbum(al)
	s.require.Nil(err)
}
