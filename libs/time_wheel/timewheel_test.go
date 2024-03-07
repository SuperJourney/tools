package timewheel

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeWheel_multiColock(t *testing.T) {
	timeWheel := NewTimeWheel(time.Second, 10)

	circle, pos := timeWheel.multiColock(10 * time.Second)
	assert.Equal(t, circle, int64(1))
	assert.Equal(t, pos, int64(0))

	for i := 0; i < 3; i++ {
		timeWheel.tick()
	}

	circle, pos = timeWheel.multiColock(10 * time.Second)
	assert.Equal(t, circle, int64(1))
	assert.Equal(t, pos, int64(3))

}

var i = 0

func TestTimeWheel_Tick(t *testing.T) {
	timeWheel := NewTimeWheel(time.Second, 10)
	var lock sync.Locker = &sync.Mutex{}

	add := func(taskID string) {
		lock.Lock()
		defer lock.Unlock()
		t.Log("任务执行, 任务ID:", taskID)
		i = i + 1
	}

	addStop := func(taskID string) {
		lock.Lock()
		defer lock.Unlock()
		t.Log("任务取消未执行, 任务ID:", taskID)
	}

	timeWheel.AddTask(10*time.Second, "1", add, addStop)
	timeWheel.AddTask(10*time.Second, "2", add, addStop)
	timeWheel.StopTask("2")
	timeWheel.AddTask(10*time.Second, "3", add, addStop)

	for i := 0; i < 9; i++ {
		timeWheel.tick()
	}

	assert.Equal(t, 0, i)
	timeWheel.tick()
	timeWheel.StopTask("1")
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, i, 2)
	timeWheel.tick()
	assert.Equal(t, i, 2)
	timeWheel.AddTask(12*time.Second, "4", add, addStop)
	for i := 0; i < 10; i++ {
		timeWheel.tick()
	}
	assert.Equal(t, i, 2)
	time.Sleep(10 * time.Millisecond)
	timeWheel.AddTask(10*time.Second, "5", add, addStop)
	timeWheel.clear()
}
