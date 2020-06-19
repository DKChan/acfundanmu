package acfundanmu

import (
	"context"
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

// ass文件的Script Info
const scriptInfo = `[Script Info]
Title: %s
ScriptType: v4.00+
Collisions: Normal
PlayResX: %d
PlayResY: %d

`

// ass文件的V4+ Styles
const sytles = `[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Danmu,Microsoft YaHei,%d,&H00FFFFFF,&H00FFFFFF,&H00000000,&H00000000,0,0,0,0,100,100,0,0,1,1,0,2,20,20,2,0

`

// ass文件的Events
const events = `[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`

// 弹幕字幕
const dialogue = `Dialogue: 0,%s,%s,Danmu,%d,20,20,2,,{\move(%d,%d,%d,%d)}%s
`

// ass文件里弹幕的出现和消失时间
const startEnd = `%d:%02d:%02d.%02d`

// 弹幕持续时间，单位为纳秒
const duration = int64(10 * time.Second)

// 弹幕在视频里出现和消失的时间，单位为纳秒
type danmuTime int64

// SubConfig 是字幕的详细设置
type SubConfig struct {
	Title     string // 字幕标题
	PlayResX  int    // 视频分辨率
	PlayResY  int    // 视频分辨率
	FontSize  int    // 字体大小
	StartTime int64  // 直播录播开始的时间，单位为纳秒
}

// dTime就是计算弹幕碰撞需要的数据
type dTime struct {
	appear    int64 // 弹幕出现的时间
	emerge    int64 // 整个弹幕完全出现在视频右边的时间
	disappear int64 // 弹幕消失的时间
}

// 将指定时间转换为ass字幕特定格式
func (d danmuTime) String() string {
	if d < 0 {
		return fmt.Sprintf(startEnd, 0, 0, 0, 0)
	}
	t := time.Unix(0, int64(d)).UTC()
	return fmt.Sprintf(
		startEnd,
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1e7,
	)
}

// WriteASS 将ass字幕写入到file里，s为字幕的设置，ctx用来结束写入ass字幕
func (q *Queue) WriteASS(ctx context.Context, s SubConfig, file string) {
	info := fmt.Sprintf(scriptInfo, s.Title, s.PlayResX, s.PlayResY)
	style := fmt.Sprintf(sytles, s.FontSize)

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	checkErr(err)
	defer f.Close()

	_, err = f.WriteString(info)
	checkErr(err)
	_, err = f.WriteString(style)
	checkErr(err)
	_, err = f.WriteString(events)
	checkErr(err)

	// lastTime存放每一行最后的弹幕的dTime
	lastTime := make([]dTime, queueLen)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			comments := q.GetDanmu()
			if comments == nil {
				return
			}
			for _, c := range comments {
				length := utf8.RuneCountInString(c.Content) * s.FontSize
				// leftTime就是弹幕运动到视频左边的时间
				leftTime := c.SendTime*1e6 - s.StartTime + (int64(s.PlayResX)*duration)/int64(s.PlayResX+length)
				dt := dTime{
					appear:    c.SendTime*1e6 - s.StartTime,
					emerge:    c.SendTime*1e6 - s.StartTime + (int64(length)*duration)/int64(s.PlayResX+length),
					disappear: c.SendTime*1e6 - s.StartTime + duration}
				for i, t := range lastTime {
					// 防止弹幕发生碰撞重叠
					if dt.appear > t.emerge && leftTime > t.disappear {
						lastTime[i] = dt
						s := fmt.Sprintf(dialogue,
							danmuTime(dt.appear),
							danmuTime(dt.disappear),
							c.UserID,
							s.PlayResX+length/2,
							s.FontSize*(i+1),
							-length/2,
							s.FontSize*(i+1),
							c.Content,
						)
						_, err = f.WriteString(s)
						checkErr(err)
						break
					}
				}
			}
		}
	}
}