package mediainfo

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Media struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                          string `json:"@type"`
			Count                         string `json:"Count"`
			StreamCount                   string `json:"StreamCount"`
			StreamKind                    string `json:"StreamKind"`
			StreamKindString              string `json:"StreamKind_String"`
			StreamKindID                  string `json:"StreamKindID"`
			VideoCount                    string `json:"VideoCount,omitempty"`
			AudioCount                    string `json:"AudioCount,omitempty"`
			VideoFormatList               string `json:"Video_Format_List,omitempty"`
			VideoFormatWithHintList       string `json:"Video_Format_WithHint_List,omitempty"`
			VideoCodecList                string `json:"Video_Codec_List,omitempty"`
			AudioFormatList               string `json:"Audio_Format_List,omitempty"`
			AudioFormatWithHintList       string `json:"Audio_Format_WithHint_List,omitempty"`
			AudioCodecList                string `json:"Audio_Codec_List,omitempty"`
			AudioChannelsTotal            string `json:"Audio_Channels_Total,omitempty"`
			CompleteName                  string `json:"CompleteName,omitempty"`
			FolderName                    string `json:"FolderName,omitempty"`
			FileNameExtension             string `json:"FileNameExtension,omitempty"`
			FileName                      string `json:"FileName,omitempty"`
			FileExtension                 string `json:"FileExtension,omitempty"`
			Format                        string `json:"Format"`
			FormatString                  string `json:"Format_String"`
			FormatExtensions              string `json:"Format_Extensions,omitempty"`
			FormatCommercial              string `json:"Format_Commercial"`
			FormatProfile                 string `json:"Format_Profile,omitempty"`
			InternetMediaType             string `json:"InternetMediaType,omitempty"`
			CodecID                       string `json:"CodecID"`
			CodecIDString                 string `json:"CodecID_String,omitempty"`
			CodecIDUrl                    string `json:"CodecID_Url,omitempty"`
			CodecIDCompatible             string `json:"CodecID_Compatible,omitempty"`
			FileSize                      string `json:"FileSize,omitempty"`
			FileSizeString                string `json:"FileSize_String,omitempty"`
			FileSizeString1               string `json:"FileSize_String1,omitempty"`
			FileSizeString2               string `json:"FileSize_String2,omitempty"`
			FileSizeString3               string `json:"FileSize_String3,omitempty"`
			FileSizeString4               string `json:"FileSize_String4,omitempty"`
			Duration                      string `json:"Duration"`
			DurationString                string `json:"Duration_String"`
			DurationString1               string `json:"Duration_String1"`
			DurationString2               string `json:"Duration_String2"`
			DurationString3               string `json:"Duration_String3"`
			DurationString4               string `json:"Duration_String4,omitempty"`
			DurationString5               string `json:"Duration_String5"`
			OverallBitRateMode            string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRateModeString      string `json:"OverallBitRate_Mode_String,omitempty"`
			OverallBitRate                string `json:"OverallBitRate,omitempty"`
			OverallBitRateString          string `json:"OverallBitRate_String,omitempty"`
			FrameRate                     string `json:"FrameRate"`
			FrameRateString               string `json:"FrameRate_String"`
			FrameCount                    string `json:"FrameCount"`
			StreamSize                    string `json:"StreamSize"`
			StreamSizeString              string `json:"StreamSize_String"`
			StreamSizeString1             string `json:"StreamSize_String1"`
			StreamSizeString2             string `json:"StreamSize_String2"`
			StreamSizeString3             string `json:"StreamSize_String3"`
			StreamSizeString4             string `json:"StreamSize_String4"`
			StreamSizeString5             string `json:"StreamSize_String5"`
			StreamSizeProportion          string `json:"StreamSize_Proportion"`
			HeaderSize                    string `json:"HeaderSize,omitempty"`
			DataSize                      string `json:"DataSize,omitempty"`
			FooterSize                    string `json:"FooterSize,omitempty"`
			IsStreamable                  string `json:"IsStreamable,omitempty"`
			FileModifiedDate              string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal         string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication            string `json:"Encoded_Application,omitempty"`
			EncodedApplicationString      string `json:"Encoded_Application_String,omitempty"`
			StreamOrder                   string `json:"StreamOrder,omitempty"`
			ID                            string `json:"ID,omitempty"`
			IDString                      string `json:"ID_String,omitempty"`
			FormatInfo                    string `json:"Format_Info,omitempty"`
			FormatUrl                     string `json:"Format_Url,omitempty"`
			FormatLevel                   string `json:"Format_Level,omitempty"`
			FormatSettings                string `json:"Format_Settings,omitempty"`
			FormatSettingsCABAC           string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsCABACString     string `json:"Format_Settings_CABAC_String,omitempty"`
			FormatSettingsRefFrames       string `json:"Format_Settings_RefFrames,omitempty"`
			FormatSettingsRefFramesString string `json:"Format_Settings_RefFrames_String,omitempty"`
			FormatSettingsGOP             string `json:"Format_Settings_GOP,omitempty"`
			CodecIDInfo                   string `json:"CodecID_Info,omitempty"`
			BitRateMode                   string `json:"BitRate_Mode,omitempty"`
			BitRateModeString             string `json:"BitRate_Mode_String,omitempty"`
			BitRate                       string `json:"BitRate,omitempty"`
			BitRateString                 string `json:"BitRate_String,omitempty"`
			Width                         string `json:"Width,omitempty"`
			WidthString                   string `json:"Width_String,omitempty"`
			Height                        string `json:"Height,omitempty"`
			HeightString                  string `json:"Height_String,omitempty"`
			StoredHeight                  string `json:"Stored_Height,omitempty"`
			SampledWidth                  string `json:"Sampled_Width,omitempty"`
			SampledHeight                 string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio              string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio            string `json:"DisplayAspectRatio,omitempty"`
			DisplayAspectRatioString      string `json:"DisplayAspectRatio_String,omitempty"`
			Rotation                      string `json:"Rotation,omitempty"`
			FrameRateMode                 string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeString           string `json:"FrameRate_Mode_String,omitempty"`
			FrameRateNum                  string `json:"FrameRate_Num,omitempty"`
			FrameRateDen                  string `json:"FrameRate_Den,omitempty"`
			ColorSpace                    string `json:"ColorSpace,omitempty"`
			ChromaSubsampling             string `json:"ChromaSubsampling,omitempty"`
			ChromaSubsamplingString       string `json:"ChromaSubsampling_String,omitempty"`
			BitDepth                      string `json:"BitDepth,omitempty"`
			BitDepthString                string `json:"BitDepth_String,omitempty"`
			ScanType                      string `json:"ScanType,omitempty"`
			ScanTypeString                string `json:"ScanType_String,omitempty"`
			BitsPixelFrame                string `json:"BitsPixel_Frame,omitempty"`
			BufferSize                    string `json:"BufferSize,omitempty"`
			Extra                         struct {
				CodecConfigurationBox string `json:"CodecConfigurationBox"`
			} `json:"extra,omitempty"`
			FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
			DurationLastFrame        string `json:"Duration_LastFrame,omitempty"`
			DurationLastFrameString  string `json:"Duration_LastFrame_String,omitempty"`
			DurationLastFrameString1 string `json:"Duration_LastFrame_String1,omitempty"`
			DurationLastFrameString2 string `json:"Duration_LastFrame_String2,omitempty"`
			DurationLastFrameString3 string `json:"Duration_LastFrame_String3,omitempty"`
			DurationLastFrameString5 string `json:"Duration_LastFrame_String5,omitempty"`
			Channels                 string `json:"Channels,omitempty"`
			ChannelsString           string `json:"Channels_String,omitempty"`
			ChannelPositions         string `json:"ChannelPositions,omitempty"`
			ChannelPositionsString2  string `json:"ChannelPositions_String2,omitempty"`
			ChannelLayout            string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
			SamplingRate             string `json:"SamplingRate,omitempty"`
			SamplingRateString       string `json:"SamplingRate_String,omitempty"`
			SamplingCount            string `json:"SamplingCount,omitempty"`
			CompressionMode          string `json:"Compression_Mode,omitempty"`
			CompressionModeString    string `json:"Compression_Mode_String,omitempty"`
			Default                  string `json:"Default,omitempty"`
			DefaultString            string `json:"Default_String,omitempty"`
			AlternateGroup           string `json:"AlternateGroup,omitempty"`
			AlternateGroupString     string `json:"AlternateGroup_String,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

/*
获取视频时长
*/
func GetTime(s string) (string, error) {
	cmd := exec.Command("mediainfo", "--Output=JSON", "--full", s)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	out := string(output)
	//fmt.Println(out)
	var m Media
	err = json.Unmarshal(output, &m)
	if err != nil {
		return "", err
	}
	for _, v := range m.Media.Track {
		if v.DurationString3 != "" {
			//fmt.Println(v.DurationString3)
			return v.DurationString3, nil
		}
	}
	return out, nil
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

/*
str1-str2
*/
func TimeSub(timeStr1, timeStr2 string) string {
	layout := "15:04:05.000"
	t1, _ := time.Parse(layout, timeStr1)
	t2, _ := time.Parse(layout, timeStr2)

	duration := t1.Sub(t2)
	hms := fmt.Sprint(duration)
	hms = strings.Replace(hms, "h", ":", 1)
	hms = strings.Replace(hms, "m", ":", 1)
	hms = strings.Replace(hms, "s", "", 1)
	return hms
}

type MI struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                           string `json:"@type"`
			Count                          string `json:"Count"`
			StreamCount                    string `json:"StreamCount"`
			StreamKind                     string `json:"StreamKind"`
			StreamKindString               string `json:"StreamKind_String"`
			StreamKindID                   string `json:"StreamKindID"`
			VideoCount                     string `json:"VideoCount,omitempty"`
			AudioCount                     string `json:"AudioCount,omitempty"`
			VideoFormatList                string `json:"Video_Format_List,omitempty"`
			VideoFormatWithHintList        string `json:"Video_Format_WithHint_List,omitempty"`
			VideoCodecList                 string `json:"Video_Codec_List,omitempty"`
			AudioFormatList                string `json:"Audio_Format_List,omitempty"`
			AudioFormatWithHintList        string `json:"Audio_Format_WithHint_List,omitempty"`
			AudioCodecList                 string `json:"Audio_Codec_List,omitempty"`
			CompleteName                   string `json:"CompleteName,omitempty"`
			FolderName                     string `json:"FolderName,omitempty"`
			FileNameExtension              string `json:"FileNameExtension,omitempty"`
			FileName                       string `json:"FileName,omitempty"`
			FileExtension                  string `json:"FileExtension,omitempty"`
			Format                         string `json:"Format"`
			FormatString                   string `json:"Format_String"`
			FormatExtensions               string `json:"Format_Extensions,omitempty"`
			FormatCommercial               string `json:"Format_Commercial"`
			FormatProfile                  string `json:"Format_Profile,omitempty"`
			InternetMediaType              string `json:"InternetMediaType,omitempty"`
			CodecID                        string `json:"CodecID"`
			CodecIDString                  string `json:"CodecID_String,omitempty"`
			CodecIDUrl                     string `json:"CodecID_Url,omitempty"`
			CodecIDCompatible              string `json:"CodecID_Compatible,omitempty"`
			FileSize                       string `json:"FileSize,omitempty"`
			FileSizeString                 string `json:"FileSize_String,omitempty"`
			FileSizeString1                string `json:"FileSize_String1,omitempty"`
			FileSizeString2                string `json:"FileSize_String2,omitempty"`
			FileSizeString3                string `json:"FileSize_String3,omitempty"`
			FileSizeString4                string `json:"FileSize_String4,omitempty"`
			Duration                       string `json:"Duration"`
			DurationString                 string `json:"Duration_String"`
			DurationString1                string `json:"Duration_String1"`
			DurationString2                string `json:"Duration_String2"`
			DurationString3                string `json:"Duration_String3"`
			DurationString4                string `json:"Duration_String4,omitempty"`
			DurationString5                string `json:"Duration_String5"`
			OverallBitRateMode             string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRateModeString       string `json:"OverallBitRate_Mode_String,omitempty"`
			OverallBitRate                 string `json:"OverallBitRate,omitempty"`
			OverallBitRateString           string `json:"OverallBitRate_String,omitempty"`
			FrameRate                      string `json:"FrameRate"`
			FrameRateString                string `json:"FrameRate_String"`
			FrameCount                     string `json:"FrameCount"`
			StreamSize                     string `json:"StreamSize"`
			StreamSizeString               string `json:"StreamSize_String"`
			StreamSizeString1              string `json:"StreamSize_String1"`
			StreamSizeString2              string `json:"StreamSize_String2"`
			StreamSizeString3              string `json:"StreamSize_String3"`
			StreamSizeString4              string `json:"StreamSize_String4"`
			StreamSizeString5              string `json:"StreamSize_String5"`
			StreamSizeProportion           string `json:"StreamSize_Proportion"`
			HeaderSize                     string `json:"HeaderSize,omitempty"`
			DataSize                       string `json:"DataSize,omitempty"`
			FooterSize                     string `json:"FooterSize,omitempty"`
			IsStreamable                   string `json:"IsStreamable,omitempty"`
			FileModifiedDate               string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal          string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication             string `json:"Encoded_Application,omitempty"`
			EncodedApplicationString       string `json:"Encoded_Application_String,omitempty"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			IDString                       string `json:"ID_String,omitempty"`
			FormatInfo                     string `json:"Format_Info,omitempty"`
			FormatUrl                      string `json:"Format_Url,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatSettings                 string `json:"Format_Settings,omitempty"`
			FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsCABACString      string `json:"Format_Settings_CABAC_String,omitempty"`
			FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
			FormatSettingsRefFramesString  string `json:"Format_Settings_RefFrames_String,omitempty"`
			CodecIDInfo                    string `json:"CodecID_Info,omitempty"`
			SourceDuration                 string `json:"Source_Duration,omitempty"`
			SourceDurationString           string `json:"Source_Duration_String,omitempty"`
			SourceDurationString1          string `json:"Source_Duration_String1,omitempty"`
			SourceDurationString2          string `json:"Source_Duration_String2,omitempty"`
			SourceDurationString3          string `json:"Source_Duration_String3,omitempty"`
			SourceDurationString4          string `json:"Source_Duration_String4,omitempty"`
			SourceDurationString5          string `json:"Source_Duration_String5,omitempty"`
			SourceDurationLastFrame        string `json:"Source_Duration_LastFrame,omitempty"`
			SourceDurationLastFrameString3 string `json:"Source_Duration_LastFrame_String3,omitempty"`
			SourceDurationLastFrameString4 string `json:"Source_Duration_LastFrame_String4,omitempty"`
			SourceDurationLastFrameString5 string `json:"Source_Duration_LastFrame_String5,omitempty"`
			BitRate                        string `json:"BitRate,omitempty"`
			BitRateString                  string `json:"BitRate_String,omitempty"`
			BitRateMaximum                 string `json:"BitRate_Maximum,omitempty"`
			BitRateMaximumString           string `json:"BitRate_Maximum_String,omitempty"`
			Width                          string `json:"Width,omitempty"`
			WidthString                    string `json:"Width_String,omitempty"`
			Height                         string `json:"Height,omitempty"`
			HeightString                   string `json:"Height_String,omitempty"`
			StoredHeight                   string `json:"Stored_Height,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			DisplayAspectRatioString       string `json:"DisplayAspectRatio_String,omitempty"`
			Rotation                       string `json:"Rotation,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeString            string `json:"FrameRate_Mode_String,omitempty"`
			FrameRateModeOriginal          string `json:"FrameRate_Mode_Original,omitempty"`
			FrameRateNum                   string `json:"FrameRate_Num,omitempty"`
			FrameRateDen                   string `json:"FrameRate_Den,omitempty"`
			SourceFrameCount               string `json:"Source_FrameCount,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			ChromaSubsamplingString        string `json:"ChromaSubsampling_String,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			BitDepthString                 string `json:"BitDepth_String,omitempty"`
			ScanType                       string `json:"ScanType,omitempty"`
			ScanTypeString                 string `json:"ScanType_String,omitempty"`
			BitsPixelFrame                 string `json:"BitsPixel_Frame,omitempty"`
			SourceStreamSize               string `json:"Source_StreamSize,omitempty"`
			SourceStreamSizeString         string `json:"Source_StreamSize_String,omitempty"`
			SourceStreamSizeString1        string `json:"Source_StreamSize_String1,omitempty"`
			SourceStreamSizeString2        string `json:"Source_StreamSize_String2,omitempty"`
			SourceStreamSizeString3        string `json:"Source_StreamSize_String3,omitempty"`
			SourceStreamSizeString4        string `json:"Source_StreamSize_String4,omitempty"`
			SourceStreamSizeString5        string `json:"Source_StreamSize_String5,omitempty"`
			SourceStreamSizeProportion     string `json:"Source_StreamSize_Proportion,omitempty"`
			EncodedLibrary                 string `json:"Encoded_Library,omitempty"`
			EncodedLibraryString           string `json:"Encoded_Library_String,omitempty"`
			EncodedLibraryName             string `json:"Encoded_Library_Name,omitempty"`
			EncodedLibraryVersion          string `json:"Encoded_Library_Version,omitempty"`
			EncodedLibrarySettings         string `json:"Encoded_Library_Settings,omitempty"`
			ColourDescriptionPresent       string `json:"colour_description_present,omitempty"`
			ColourDescriptionPresentSource string `json:"colour_description_present_Source,omitempty"`
			ColourRange                    string `json:"colour_range,omitempty"`
			ColourRangeSource              string `json:"colour_range_Source,omitempty"`
			ColourPrimaries                string `json:"colour_primaries,omitempty"`
			ColourPrimariesSource          string `json:"colour_primaries_Source,omitempty"`
			TransferCharacteristics        string `json:"transfer_characteristics,omitempty"`
			TransferCharacteristicsSource  string `json:"transfer_characteristics_Source,omitempty"`
			MatrixCoefficients             string `json:"matrix_coefficients,omitempty"`
			MatrixCoefficientsSource       string `json:"matrix_coefficients_Source,omitempty"`
			Extra                          struct {
				SourceDelay           string `json:"Source_Delay"`
				SourceDelaySource     string `json:"Source_Delay_Source"`
				CodecConfigurationBox string `json:"CodecConfigurationBox"`
			} `json:"extra,omitempty"`
			FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
			BitRateMode              string `json:"BitRate_Mode,omitempty"`
			BitRateModeString        string `json:"BitRate_Mode_String,omitempty"`
			Channels                 string `json:"Channels,omitempty"`
			ChannelsString           string `json:"Channels_String,omitempty"`
			ChannelPositions         string `json:"ChannelPositions,omitempty"`
			ChannelPositionsString2  string `json:"ChannelPositions_String2,omitempty"`
			ChannelLayout            string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
			SamplingRate             string `json:"SamplingRate,omitempty"`
			SamplingRateString       string `json:"SamplingRate_String,omitempty"`
			SamplingCount            string `json:"SamplingCount,omitempty"`
			CompressionMode          string `json:"Compression_Mode,omitempty"`
			CompressionModeString    string `json:"Compression_Mode_String,omitempty"`
			Default                  string `json:"Default,omitempty"`
			DefaultString            string `json:"Default_String,omitempty"`
			AlternateGroup           string `json:"AlternateGroup,omitempty"`
			AlternateGroupString     string `json:"AlternateGroup_String,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

/*
mediainfo --Output=JSON --Full
*/
func GetCodec(fp string) (codec, width, height string) {
	cmd := exec.Command("mediainfo", "--Output=JSON", "--Full", fp)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	file, _ := os.OpenFile("exam.json", os.O_WRONLY|os.O_CREATE, 0666)
	file.WriteString(string(output))
	var mi MI
	err = json.Unmarshal(output, &mi)
	if err != nil {
		return
	}
	for _, v := range mi.Media.Track {
		if v.Type == "Video" {
			if v.Format == "HEVC" {
				codec = "HEVC"
			}
			if v.Format == "VP9" {
				codec = "VP9"
			}
			if v.Format == "AVC" {
				codec = "AVC"
			}
			width = v.Width
			height = v.Height
		}
	}
	return codec, width, height
}
func GetCrfForVP9(width, height string) (crf string) {
	width_int, _ := strconv.Atoi(width)
	height_int, _ := strconv.Atoi(height)
	stand := 0
	if width_int > height_int {
		stand = height_int
	} else if height_int > width_int {
		stand = width_int
	} else {
		stand = width_int
	}
	if stand >= 2160 {
		crf = "15"
	} else if stand >= 1440 {
		crf = "24"
	} else if stand >= 1080 {
		crf = "31"
	} else if stand >= 720 {
		crf = "32"
	} else if stand >= 480 {
		crf = "33"
	} else if stand >= 135 {
		crf = "34"
	} else {
		crf = "35"
	}
	return crf
}
