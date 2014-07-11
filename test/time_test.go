package test

import(
	"time"

	"testing"
	"github.com/bmizerany/assert"
)

func TestDate(t *testing.T) {

	date := time.Now()     // 2014-07-02 17:35:39.605631353 +0800 CST
	_ = date.Format("MST")  //CST

	_ = date.Format("Jan 2, 2006 at 3:04pm (MST)")  // Jul 2, 2014 at 5:43pm (CST)

	date = date.Add(time.Hour * 8)  // 2014-07-03 01:35:39.605631353 +0800 CST
	_ = date.UTC()  // 2014-07-02 17:35:39.605631353 +0800 UTC

	_ = date.Unix()   // unix时间戳

	t.Log(date)  //CST


	assert.Equal(t, 1, 1)
}
