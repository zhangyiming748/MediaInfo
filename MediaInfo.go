package MediaInfo

import (
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/MediaInfo/log"
	"os"
	"os/exec"
	"strings"
)

type mediainfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                           string `json:"@type"`
			VideoCount                     string `json:"VideoCount,omitempty"`
			AudioCount                     string `json:"AudioCount,omitempty"`
			FileExtension                  string `json:"FileExtension,omitempty"`
			Format                         string `json:"Format"`
			FormatProfile                  string `json:"Format_Profile,omitempty"`
			CodecID                        string `json:"CodecID"`
			CodecIDCompatible              string `json:"CodecID_Compatible,omitempty"`
			FileSize                       string `json:"FileSize,omitempty"`
			Duration                       string `json:"Duration"`
			OverallBitRate                 string `json:"OverallBitRate,omitempty"`
			FrameRate                      string `json:"FrameRate"`
			FrameCount                     string `json:"FrameCount"`
			StreamSize                     string `json:"StreamSize"`
			HeaderSize                     string `json:"HeaderSize,omitempty"`
			DataSize                       string `json:"DataSize,omitempty"`
			FooterSize                     string `json:"FooterSize,omitempty"`
			IsStreamable                   string `json:"IsStreamable,omitempty"`
			Description                    string `json:"Description,omitempty"`
			FileModifiedDate               string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal          string `json:"File_Modified_Date_Local,omitempty"`
			EncodedApplication             string `json:"Encoded_Application,omitempty"`
			StreamOrder                    string `json:"StreamOrder,omitempty"`
			ID                             string `json:"ID,omitempty"`
			FormatLevel                    string `json:"Format_Level,omitempty"`
			FormatSettingsCABAC            string `json:"Format_Settings_CABAC,omitempty"`
			FormatSettingsRefFrames        string `json:"Format_Settings_RefFrames,omitempty"`
			BitRate                        string `json:"BitRate,omitempty"`
			Width                          string `json:"Width,omitempty"`
			Height                         string `json:"Height,omitempty"`
			StoredWidth                    string `json:"Stored_Width,omitempty"`
			SampledWidth                   string `json:"Sampled_Width,omitempty"`
			SampledHeight                  string `json:"Sampled_Height,omitempty"`
			PixelAspectRatio               string `json:"PixelAspectRatio,omitempty"`
			DisplayAspectRatio             string `json:"DisplayAspectRatio,omitempty"`
			Rotation                       string `json:"Rotation,omitempty"`
			FrameRateMode                  string `json:"FrameRate_Mode,omitempty"`
			FrameRateModeOriginal          string `json:"FrameRate_Mode_Original,omitempty"`
			ColorSpace                     string `json:"ColorSpace,omitempty"`
			ChromaSubsampling              string `json:"ChromaSubsampling,omitempty"`
			BitDepth                       string `json:"BitDepth,omitempty"`
			ScanType                       string `json:"ScanType,omitempty"`
			EncodedLibrary                 string `json:"Encoded_Library,omitempty"`
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
				CodecConfigurationBox string `json:"CodecConfigurationBox"`
			} `json:"extra,omitempty"`
			FormatSettingsSBR          string `json:"Format_Settings_SBR,omitempty"`
			FormatAdditionalFeatures   string `json:"Format_AdditionalFeatures,omitempty"`
			SourceDuration             string `json:"Source_Duration,omitempty"`
			BitRateMode                string `json:"BitRate_Mode,omitempty"`
			Channels                   string `json:"Channels,omitempty"`
			ChannelPositions           string `json:"ChannelPositions,omitempty"`
			ChannelLayout              string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame            string `json:"SamplesPerFrame,omitempty"`
			SamplingRate               string `json:"SamplingRate,omitempty"`
			SamplingCount              string `json:"SamplingCount,omitempty"`
			SourceFrameCount           string `json:"Source_FrameCount,omitempty"`
			CompressionMode            string `json:"Compression_Mode,omitempty"`
			StreamSizeProportion       string `json:"StreamSize_Proportion,omitempty"`
			SourceStreamSize           string `json:"Source_StreamSize,omitempty"`
			SourceStreamSizeProportion string `json:"Source_StreamSize_Proportion,omitempty"`
			Default                    string `json:"Default,omitempty"`
			AlternateGroup             string `json:"AlternateGroup,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

func MediaInfo(src, pattern, dst string) {
	old := strings.Join([]string{dst, "Solution.sh"}, "/")
	if err := os.RemoveAll(old); err != nil {
		log.Debug.Println("旧文件并没有被删除")
	} else {
		head := "#!/usr/bin/env bash"
		writeline(old, head)
	}

	files := getFiles(src, pattern)
	for index, file := range files {
		log.Info.Printf("正在判断第%d/%d个文件\n", index+1, len(files))
		mediaInfo_help(src, file, dst)
	}
}
func mediaInfo_help(src, file, dst string) {
	in := strings.Join([]string{src, file}, "/")
	out := ""
	log.Info.Printf("文件名%s\n", in)
	cmd := exec.Command("mediainfo", "--Output=JSON", in)
	log.Debug.Printf("生成的命令是:%s", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Debug.Printf("cmd.StdoutPipe产生的错误:%f", err)
	}
	if err = cmd.Start(); err != nil {
		log.Debug.Printf("cmd.Run产生的错误:%f", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		s := string(tmp)
		s = strings.Replace(s, "\u0000", "", -1)
		//log.Println(s)
		out = strings.Join([]string{out, s}, "")
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Debug.Println("命令执行中有错误产生", err)
	}
	//fmt.Println(out)
	report := getJson(out)
	//fmt.Printf("@ref is %f\n", report.Media.Ref)
	var format string
	for _, f := range report.Media.Track {
		if f.Type == "Video" {
			fmt.Printf("%s\t是\t%s编码\n", file, f.Format)
			format = f.Format
		}
	}
	line := generateSolution(file, format)
	sh := strings.Join([]string{dst, "Solution.sh"}, "/")
	writeline(sh, line)
}

func getJson(s string) mediainfo {
	var mi mediainfo
	if err := json.Unmarshal([]byte(s), &mi); err != nil {
		log.Debug.Printf("err : %s\n", err)
	}
	//fmt.Printf("%+v\n", mi)
	return mi
}

func generateSolution(fname, ftype string) string {
	b := ""
	fname = strings.Join([]string{"\"", fname, "\""}, "")
	if ftype == "HEVC" {
		b = strings.Join([]string{"mv", fname, "h265/"}, " ")
	} else {
		b = strings.Join([]string{"mv", fname, "h264/"}, " ")
	}
	return b
}

func writeline(fname, content string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Debug.Println(err)
	}
	_, err = f.WriteString(content)
	_, _ = f.WriteString("\n")
	if err != nil {
		log.Debug.Println("写文件出错")
	} else {
		//log.Printf("写入%d个字节", n)
	}
}

func getFiles(dir, pattern string) []string {
	files, _ := os.ReadDir(dir)
	var aim []string
	types := strings.Split(pattern, ";") //"wmv;rm"
	for _, f := range files {
		//fmt.Println(f.Name())
		if l := strings.Split(f.Name(), ".")[0]; len(l) != 0 {
			//log.Info.Printf("有效的文件:%v\n", f.Name())
			for _, v := range types {
				if strings.HasSuffix(f.Name(), v) {
					log.Debug.Printf("有效的目标文件:%v\n", f.Name())
					//absPath := strings.Join([]string{dir, f.Name()}, "/")
					//log.Printf("目标文件的绝对路径:%v\n", absPath)
					aim = append(aim, f.Name())
				}
			}
		}
	}
	return aim
}
