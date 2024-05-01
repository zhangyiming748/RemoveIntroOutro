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
func TestTimeSub(t *testing.T) {
	timeStr1 := "00:45:49.960"
	timeStr2 := "00:43:48.860"
	ret := TimeSub(timeStr1, timeStr2)
	t.Log(ret)
}

// go test -v -run  TestGetCodec
func TestGetCodec(t *testing.T) {
	codec, width, height := GetCodec("/mnt/e/video/Straplez/title1/Sexy lesbians in black nylon catsuits fuck with a big dildo [64d662238e2a4].mp4")
	t.Logf("codec: %v, width: %v, height: %v", codec, width, height)
}
