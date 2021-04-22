// this is wrapper for github.com/jinzhu/now to set fake date
package utils

import (
	"time"

	"github.com/jinzhu/now"
)

type Time interface {
	Now() time.Time
	BeginningOfDay() time.Time
	EndOfDay() time.Time
}

type FakeTimeSource struct {
	now      time.Time
	Location *time.Location
}

func NewFakeTimeSource(now time.Time, l *time.Location) Time {
	return &FakeTimeSource{now, l}
}

func (t *FakeTimeSource) BeginningOfDay() time.Time {
	return now.New(t.now).BeginningOfDay()
}

func (t *FakeTimeSource) EndOfDay() time.Time {
	return now.New(t.now).EndOfDay()
}

func (t *FakeTimeSource) Now() time.Time {
	if t.Location != nil {
		t.now = t.now.In(t.Location)
	}
	return t.now
}

type TimeSource struct {
	Location *time.Location
}

func NewTimeSource(l *time.Location) Time {
	return &TimeSource{
		Location: l,
	}
}

func (t *TimeSource) BeginningOfDay() time.Time {
	return now.New(t.Now()).BeginningOfDay()
}

func (t *TimeSource) EndOfDay() time.Time {
	return now.New(t.Now()).EndOfDay()
}

func (t *TimeSource) Now() time.Time {
	tm := time.Now()
	if t.Location != nil {
		// This is important to construct date without location.
		// Date with location different from default passed to ORM
		// will be converted back to default
		l := "2006-01-02 15:04:05"
		tm, _ = time.Parse(l, tm.In(t.Location).Format(l))
	}
	return tm
}
