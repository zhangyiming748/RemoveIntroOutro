package mediainfo

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type MediaInfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                    string `json:"@type"`
			VideoCount              string `json:"VideoCount,omitempty"`
			AudioCount              string `json:"AudioCount,omitempty"`
			FileExtension           string `json:"FileExtension,omitempty"`
			Format                  string `json:"Format"`
			FormatProfile           string `json:"Format_Profile,omitempty"`
			CodecID                 string `json:"CodecID"`
			CodecIDCompatible       string `json:"CodecID_Compatible,omitempty"`
			FileSize                string `json:"FileSize,omitempty"`
			Duration                string `json:"Duration"`
			OverallBitRateMode      string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRate          string `json:"OverallBitRate,omitempty"`
			FrameRate               string `json:"FrameRate"`
			FrameCount              string `json:"FrameCount"`
			StreamSize              string `json:"StreamSize"`
			HeaderSize              string `json:"HeaderSize,omitempty"`
			DataSize                string `json:"DataSize,omitempty"`
			FooterSize              string `json:"FooterSize,omitempty"`
			IsStreamable            string `json:"IsStreamable,omitempty"`
			FileModifiedDate        string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal   string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication      string `json:"Encoded_Application,omitempty"`
			StreamOrder             string `json:"StreamOrder,omitempty"`
			ID                      string `json:"ID,omitempty"`
			FormatLevel             string `json:"Format_Level,omitempty"`
			FormatSettingsCABAC     string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsRefFrames string `json:"Format_Settings_RefFrames,omitempty"`
			FormatSettingsGOP       string `json:"Format_Settings_GOP,omitempty"`
			BitRateMode             string `json:"BitRate_Mode,omitempty"`
			BitRate                 string `json:"BitRate,omitempty"`
			Width                   string `json:"Width,omitempty"`
			Height                  string `json:"Height,omitempty"`
			StoredHeight            string `json:"Stored_Height,omitempty"`
			SampledWidth            string `json:"Sampled_Width,omitempty"`
			SampledHeight           string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio        string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio      string `json:"DisplayAspectRatio,omitempty"`
			Rotation                string `json:"Rotation,omitempty"`
			FrameRateMode           string `json:"FrameRate_Mode,omitempty"`
			FrameRateNum            string `json:"FrameRate_Num,omitempty"`
			FrameRateDen            string `json:"FrameRate_Den,omitempty"`
			ColorSpace              string `json:"ColorSpace,omitempty"`
			ChromaSubsampling       string `json:"ChromaSubsampling,omitempty"`
			BitDepth                string `json:"BitDepth,omitempty"`
			ScanType                string `json:"ScanType,omitempty"`
			BufferSize              string `json:"BufferSize,omitempty"`
			Extra                   struct {
				CodecConfigurationBox string `json:"CodecConfigurationBox"`
			} `json:"extra,omitempty"`
			FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
			DurationLastFrame        string `json:"Duration_LastFrame,omitempty"`
			Channels                 string `json:"Channels,omitempty"`
			ChannelPositions         string `json:"ChannelPositions,omitempty"`
			ChannelLayout            string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
			SamplingRate             string `json:"SamplingRate,omitempty"`
			SamplingCount            string `json:"SamplingCount,omitempty"`
			CompressionMode          string `json:"Compression_Mode,omitempty"`
			Default                  string `json:"Default,omitempty"`
			AlternateGroup           string `json:"AlternateGroup,omitempty"`
		} `json:"track"`
	} `json:"media"`
}
type General struct {
	Type                  string `json:"@type"`
	VideoCount            string `json:"VideoCount"`
	AudioCount            string `json:"AudioCount"`
	FileExtension         string `json:"FileExtension"`
	Format                string `json:"Format"`
	FormatProfile         string `json:"Format_Profile"`
	CodecID               string `json:"CodecID"`
	CodecIDCompatible     string `json:"CodecID_Compatible"`
	FileSize              string `json:"FileSize"`
	Duration              string `json:"Duration"`
	OverallBitRateMode    string `json:"OverallBitRate_Mode"`
	OverallBitRate        string `json:"OverallBitRate"`
	FrameRate             string `json:"FrameRate"`
	FrameCount            string `json:"FrameCount"`
	StreamSize            string `json:"StreamSize"`
	HeaderSize            string `json:"HeaderSize"`
	DataSize              string `json:"DataSize"`
	FooterSize            string `json:"FooterSize"`
	IsStreamable          string `json:"IsStreamable"`
	FileModifiedDate      string `json:"File_Modified_Date"`
	FileModifiedDateLocal string `json:"File_Modified_Date_Local"`
	EncodedApplication    string `json:"Encoded_Application"`
}
type Video struct {
	Type                    string `json:"@type"`
	StreamOrder             string `json:"StreamOrder"`
	ID                      string `json:"ID"`
	Format                  string `json:"Format"`
	FormatProfile           string `json:"Format_Profile"`
	FormatLevel             string `json:"Format_Level"`
	FormatSettingsCABAC     string `json:"Format_Settings_CABAC"`
	FormatSettingsRefFrames string `json:"Format_Settings_RefFrames"`
	FormatSettingsGOP       string `json:"Format_Settings_GOP"`
	CodecID                 string `json:"CodecID"`
	Duration                string `json:"Duration"`
	BitRateMode             string `json:"BitRate_Mode"`
	BitRate                 string `json:"BitRate"`
	Width                   string `json:"Width"`
	Height                  string `json:"Height"`
	StoredHeight            string `json:"Stored_Height"`
	SampledWidth            string `json:"Sampled_Width"`
	SampledHeight           string `json:"Sampled_Height"`
	PixelAspectRatio        string `json:"PixelAspectRatio"`
	DisplayAspectRatio      string `json:"DisplayAspectRatio"`
	Rotation                string `json:"Rotation"`
	FrameRateMode           string `json:"FrameRate_Mode"`
	FrameRate               string `json:"FrameRate"`
	FrameRateNum            string `json:"FrameRate_Num"`
	FrameRateDen            string `json:"FrameRate_Den"`
	FrameCount              string `json:"FrameCount"`
	ColorSpace              string `json:"ColorSpace"`
	ChromaSubsampling       string `json:"ChromaSubsampling"`
	BitDepth                string `json:"BitDepth"`
	ScanType                string `json:"ScanType"`
	StreamSize              string `json:"StreamSize"`
	BufferSize              string `json:"BufferSize"`
	Extra                   struct {
		CodecConfigurationBox string `json:"CodecConfigurationBox"`
	} `json:"extra"`
}
type Audio struct {
	Type                     string `json:"@type"`
	StreamOrder              string `json:"StreamOrder"`
	ID                       string `json:"ID"`
	Format                   string `json:"Format"`
	FormatAdditionalFeatures string `json:"Format_AdditionalFeatures"`
	CodecID                  string `json:"CodecID"`
	Duration                 string `json:"Duration"`
	DurationLastFrame        string `json:"Duration_LastFrame"`
	BitRateMode              string `json:"BitRate_Mode"`
	BitRate                  string `json:"BitRate"`
	Channels                 string `json:"Channels"`
	ChannelPositions         string `json:"ChannelPositions"`
	ChannelLayout            string `json:"ChannelLayout"`
	SamplesPerFrame          string `json:"SamplesPerFrame"`
	SamplingRate             string `json:"SamplingRate"`
	SamplingCount            string `json:"SamplingCount"`
	FrameRate                string `json:"FrameRate"`
	FrameCount               string `json:"FrameCount"`
	CompressionMode          string `json:"Compression_Mode"`
	StreamSize               string `json:"StreamSize"`
	Default                  string `json:"Default"`
	AlternateGroup           string `json:"AlternateGroup"`
}

/*
获取视频时长
*/
func GetTime(s string) (string, error) {
	cmd := exec.Command("mediainfo", "--Output=JSON", s)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	var media MediaInfo
	err = json.Unmarshal(output, &media)
	if err != nil {
		return "", err
	}
	for _, t := range media.Media.Track {
		if t.Type == "General" {
			//fmt.Println("General")
		} else if t.Type == "Video" {
			//fmt.Printf("%v is %T\n", t.Duration, t.Duration)
			fulltime, _ := d2t(t.Duration)
			fmt.Printf("视频全长%v\n", fulltime)
		} else if t.Type == "Audio" {
			//fmt.Println("Audio")
		}
	}
	return "", nil
}

func d2t(s string) (string, error) {
	// 例如：要解析的时间段字符串 69.959
	s = strings.Join([]string{s, "s"}, "")
	duration, err := time.ParseDuration(s)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	mill := strings.Replace(strings.Split(s, ".")[1], "s", "", 1)
	//fmt.Println(mill)
	milliseconds, _ := strconv.Atoi(mill)
	formattedTime := fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
	//fmt.Println(formattedTime)
	return formattedTime, nil
}
