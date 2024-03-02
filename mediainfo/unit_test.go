package mediainfo

import (
	"testing"
)

func TestGetTime(t *testing.T) {
	te, err := GetTime("/Users/zen/Github/RemoveIntroOutro/Brittany rides Trinity futa dick [PuppetMaster].mp4")
	if err != nil {
		t.Error(err)
	}
	t.Log(te)
}
func TestFormat(t *testing.T) {
	str := "69.959s" // 例如：要解析的时间段字符串
	d2t(str)
}
