package datactl

import "testing"

/*  Command line: go test datactl  */
func TestCheckLost(t *testing.T)  {
	long := []string{"a", "ab", "abc", "abcd", "abcde", "abcdef"}
	short := []string{"ab", "abc", "abcd"}
	lost := CheckLost(long, short)
	result := []string{"a",  "abcde", "abcdef"}
	if len(lost) != len(result) {
		t.Error("Check lenth CheckLost() failed. Got", lost, "Expected ",`{"a",  "abcde", "abcdef"}`)
		t.Error(lost,result)
	}
	for i := 0; i < len(result); i++ {
			if result[i] != lost[i] {
				t.Error("Check array CheckLost() failed. Got", lost[i], "Expected ",result[i])
				t.Error(lost,result)
			}
	}
}

/*  Command line: go test -test.bench datactl  */
/* Pprof: go test -bench=".*" -cpuprofile=cpu.prof -c */
/* Pprof: go tool pprof *.test cpu.prof */
func BenchmarkCheckLost(b *testing.B)  {
	b.StopTimer()
	// Do something.... It not in benchmark time
	b.StartTimer()
	long := []string{"a", "ab", "abc", "abcd", "abcde", "abcdef"}
	short := []string{"ab", "abc", "abcd"}
	//var result []string
	for i := 0; i < b.N; i++ {
		_ = CheckLost(long, short)
	}
}

func TestDeletesame(t *testing.T) {
	baselist := []string{"1","2","3","4","5","3","5"}
	newlist := Deletesame(baselist)
	result := []string{"1","2","4","3","5"}
	if len(newlist) != len(result) {
		t.Error("Deletesame() lenth chaeck failed. Got", newlist, "Expected ",result)
	}
	for i := 0; i < len(result); i++ {
			if result[i] != newlist[i] {
				t.Error("Deletesame() array check failed. Got", newlist[i], "Expected ",result[i])
				t.Error(newlist,result)
			}
	}

}