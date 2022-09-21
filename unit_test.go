package MediaInfo

import "testing"

func TestUnit(t *testing.T) {
	src := "/Volumes/T7/slacking/Telegram/flexible/FlexiLady/h264"
	dst := "/Users/zen/Github/MediaInfo"
	pattern := "mp4"
	MediaInfo(src, pattern, dst)
}
