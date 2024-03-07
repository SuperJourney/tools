package timewheel

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertInOrder(t *testing.T) {
	type args struct {
		l    *list.List
		task *Task
	}

	emptyList := list.New()
	sortList := list.New()
	sortList.PushBack(&Task{
		circle: 2,
	})
	sortList.PushBack(&Task{
		circle: 3,
	})
	sortList.PushBack(&Task{
		circle: 6,
	})

	composertion := func(tt assert.TestingT, i1, i2 interface{}, i3 ...interface{}) bool {
		l := i1.(*list.List)
		want := i2.([]int64)

		for e := l.Front(); e != nil; e = e.Next() {
			if len(want) == 0 {
				continue
			}
			if e.Value.(*Task).circle != want[0] {
				tt.Errorf("want:%v, got:%v", want, e.Value.(int))
			}
			want = want[1:]
		}
		return true
	}
	tests := []struct {
		name       string
		args       args
		want       []int64
		compostion assert.ComparisonAssertionFunc
	}{
		{
			name: "test case 1",
			args: args{
				l: emptyList,
				task: &Task{
					circle: 1,
					job:    []func(string){func(string) {}, func(string) {}},
				},
			},
			want:       []int64{1},
			compostion: composertion,
		},
		{
			name: "test case 1",
			args: args{
				l: sortList,
				task: &Task{
					circle: 1,
				},
			},
			want:       []int64{1, 2, 3, 6},
			compostion: composertion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertInOrder(tt.args.l, tt.args.task)
			tt.compostion(t, tt.args.l, tt.want)
		})
	}
}
