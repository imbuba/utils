package time

import (
	"reflect"
	"testing"
	"time"
)

func TestDaysBetween(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "same day", args: args{from: time.Now(), to: time.Now()}, want: 0},
		{name: "before day", args: args{from: time.Now(), to: time.Now().Add(-24 * time.Hour)}, want: -1},
		{name: "after day", args: args{from: time.Now(), to: time.Now().Add(24 * time.Hour)}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaysBetween(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("DaysBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayID(t *testing.T) {
	type args struct {
		from time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first day", args: args{from: time.Unix(300, 0)}, want: 0},
		{name: "next day", args: args{from: time.Unix(86400, 0)}, want: 1},
		{name: "next day + 10 secs", args: args{from: time.Unix(86410, 0)}, want: 1},
		{name: "almost next day", args: args{from: time.Unix(2*86400-1, 999)}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayID(tt.args.from); got != tt.want {
				t.Errorf("DayID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	type args struct {
		from     time.Time
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "first day", args: args{from: time.Date(2000, time.January, 1, 12, 34, 56, 789, loc), location: time.UTC}, want: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfDay(tt.args.from, tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
