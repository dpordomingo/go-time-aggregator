package aggregator

import (
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestNewTimeAggregator(c *C) {
	_, err := NewTimeAggregator(Hour, Year)
	c.Assert(err, Equals, InvalidOrderError)
}

func (s *UtilsSuite) TestTimeAggregator_Add_YearHour(c *C) {
	a, _ := NewTimeAggregator(Year, Hour)
	a.Add(date2014November, 15)
	a.Add(date2015November, 10)
	a.Add(date2015December, 10)

	c.Assert(a.Values, HasLen, 2)
	c.Assert(a.Get(date2015January), Equals, int64(20))
	c.Assert(a.Get(date2014February), Equals, int64(15))
}

func (s *UtilsSuite) TestTimeAggregator_Add_MonthHour(c *C) {
	a, _ := NewTimeAggregator(Month, Hour)
	a.Add(date2014November, 15)
	a.Add(date2015November, 10)
	a.Add(date2015December, 10)

	c.Assert(a.Values, HasLen, 2)
	c.Assert(a.Get(date2013November), Equals, int64(25))
	c.Assert(a.Get(date2013December), Equals, int64(10))
}

func (s *UtilsSuite) TestTimeAggregator_Add_YearMonthHour(c *C) {
	a, _ := NewTimeAggregator(Year, Month, Hour)
	a.Add(date2014November, 10)
	a.Add(date2015November, 10)
	a.Add(date2015December, 10)
	a.Add(time.Date(2015, time.November, 1, 23, 1, 1, 0, time.UTC), 40)

	c.Assert(a.Values, HasLen, 3)
	c.Assert(a.Get(date2015November), Equals, int64(50))
}

func (s *UtilsSuite) TestTimeAggregator_Add_Only(c *C) {
	a, _ := NewTimeAggregator(Hour)
	a.Add(date2014November, 10)
	a.Add(date2015November, 10)
	a.Add(date2015December, 10)

	h21 := time.Date(2015, time.November, 1, 21, 1, 1, 0, time.UTC)
	a.Add(h21, 40)

	c.Assert(a.Values, HasLen, 1)
	c.Assert(a.Get(date2015November), Equals, int64(30))
	c.Assert(a.Get(h21), Equals, int64(40))
}

func (s *UtilsSuite) TestTimeAggregator_MarshalAndUnmarshal(c *C) {
	d := time.Now()

	a, _ := NewTimeAggregator(Year, Hour)
	a.Add(d, 10)
	a.Add(d, 10)

	c.Assert(a.Values, HasLen, 1)
	c.Assert(a.Get(d), Equals, int64(20))

	m := a.Marshal()

	b := &TimeAggregator{}
	err := b.Unmarshal(m)

	c.Assert(err, IsNil)
	c.Assert(a.flags, Equals, b.flags)
	c.Assert(b.Values, HasLen, 1)
	c.Assert(b.Get(d), Equals, int64(20))
}

var date2013December = time.Date(2013, time.December, 12, 23, 59, 59, 0, time.UTC)
var date2013November = time.Date(2013, time.November, 12, 23, 59, 59, 0, time.UTC)
var date2014February = time.Date(2014, time.February, 12, 23, 59, 59, 0, time.UTC)
var date2014November = time.Date(2014, time.November, 12, 23, 59, 59, 0, time.UTC)
var date2015January = time.Date(2015, time.January, 12, 23, 59, 59, 0, time.UTC)
var date2015December = time.Date(2015, time.December, 12, 23, 59, 59, 0, time.UTC)
var date2015November = time.Date(2015, time.November, 12, 23, 59, 59, 0, time.UTC)
