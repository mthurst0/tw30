package tinytools_test

import (
	"testing"

	"github.com/mthurst0/tw30/internal/tinytools"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type SliceTestSuite struct {
	suite.Suite
	ginMode string
}

func TestSliceTestSuite(t *testing.T) {
	suite.Run(t, &SliceTestSuite{})
}

func (s *SliceTestSuite) SetupTest() {
	s.ginMode = gin.Mode()
	gin.SetMode(gin.TestMode)
}

func (s *SliceTestSuite) TearDownTest() {
	gin.SetMode(s.ginMode)
}

func (s *SliceTestSuite) TestRemoveFromSlice() {
	t1 := []string{"a", "b", "c", "d", "e"}
	r1 := tinytools.RemoveFromSlice(t1, "c")
	s.Equal([]string{"a", "b", "d", "e"}, r1)
	t2 := []string{"a", "b", "c", "d", "e"}
	r2 := tinytools.RemoveFromSlice(t2, "z")
	s.Equal([]string{"a", "b", "c", "d", "e"}, r2)
}

func (s *SliceTestSuite) TestRemoveFromSliceNil() {
	r1 := tinytools.RemoveFromSlice(nil, "z")
	s.Nil(r1)
	r2 := tinytools.RemoveFromSlice(nil, "")
	s.Nil(r2)
	t3 := []string{"a", "b", "c", "d", "e"}
	r3 := tinytools.RemoveFromSlice(t3, "")
	s.Equal([]string{"a", "b", "c", "d", "e"}, r3)
}

func (s *SliceTestSuite) TestRemoveAllFromSlice() {
	t1 := []string{"a", "b", "c", "d", "e"}
	r1 := tinytools.RemoveAllFromSlice(t1, []string{"c", "e"})
	s.Equal([]string{"a", "b", "d"}, r1)
	t2 := []string{"a", "b", "c", "d", "e"}
	r2 := tinytools.RemoveAllFromSlice(t2, []string{"z"})
	s.Equal([]string{"a", "b", "c", "d", "e"}, r2)
}

func (s *SliceTestSuite) TestRemoveAllFromSliceNil() {
	r1 := tinytools.RemoveAllFromSlice(nil, []string{"z"})
	s.Nil(r1)
	r2 := tinytools.RemoveAllFromSlice(nil, nil)
	s.Nil(r2)
	t3 := []string{"a", "b", "c", "d", "e"}
	r3 := tinytools.RemoveAllFromSlice(t3, nil)
	s.Equal([]string{"a", "b", "c", "d", "e"}, r3)
}

func (s *SliceTestSuite) TestUniqueSlice() {
	t1 := []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"}
	r1 := tinytools.UniqueSlice(t1)
	s.Equal([]string{"a", "b", "c", "d", "e"}, r1)
	t2 := []string{"a", "b", "c", "d", "e"}
	r2 := tinytools.UniqueSlice(t2)
	s.Equal([]string{"a", "b", "c", "d", "e"}, r2)
}

func (s *SliceTestSuite) TestUniqueSliceNil() {
	r1 := tinytools.UniqueSlice(nil)
	s.Nil(r1)
	var t2 []string
	r2 := tinytools.UniqueSlice(t2)
	s.Equal([]string(nil), r2)
}
