package qweather

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/zhangyiming748/pretty"
	"github.com/zhangyiming748/qweather/constant"
)

func init() {
	log.SetFlags(2 | 16)
}

func TestNow(t *testing.T) {
	if rep, err := NowApi("101011000", constant.HOST, constant.APIKEY); err != nil {
		t.Error(err)
	} else {
		pretty.P(rep)
		if report,err :=os.OpenFile("report.md", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err!= nil {
			t.Error(err)	
		}else{
			defer report.Close()
			pretty.P(rep, report)
			report.WriteString(fmt.Sprintln("|数据观测时间|温度|体感温度|天气状况图标代码|天气状况描述|风向360角度|风向|风力等级|风速|相对湿度|降水量|大气压强|能见度|云量|露点温度|"))
			report.WriteString(fmt.Sprintln("|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|"))
			report.WriteString(fmt.Sprintf("|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|", rep.Now.ObsTime, rep.Now.Temp, rep.Now.FeelsLike, rep.Now.Icon, rep.Now.Text, rep.Now.Wind360, rep.Now.WindDir, rep.Now.WindScale, rep.Now.WindSpeed, rep.Now.Humidity, rep.Now.Precip, rep.Now.Pressure, rep.Now.Vis, rep.Now.Cloud, rep.Now.Dew))
		}
	}

}
