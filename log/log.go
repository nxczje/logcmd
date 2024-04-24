package log

import (
	"fmt"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

type LogData struct {
	Level string
	Func  string
	Data  string
}

func Log(data LogData) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	data.Func = color.YellowString(data.Func)
	switch data.Level {
	case "info":
		log.Info(fmt.Sprintf("| %-20s | [+] %s", data.Func, data.Data))
	case "error":
		log.Error(fmt.Sprintf("| %-20s | [x] %s", data.Func, data.Data))
	case "warn":
		log.Warn(fmt.Sprintf("| %-20s | [-] %s", data.Func, data.Data))
	case "done":
		log.Info(fmt.Sprintf("| %-20s | [✓] %s", data.Func, data.Data))
	}
}

func Banner() {
	banner := `
	Nxczje	                |     |
                                \\_V_//
                                \/=|=\/
                                 [=v=]
                               __\___/_____
                              /..[  _____  ]
                             /_  [ [  N /] ]
                            /../.[ [ N /@] ]
                           <-->[_[ [N /@/] ]
                          /../ [.[ [ /@/ ] ]
     _________________]\ /__/  [_[ [/@/ C] ]
    <_________________>>0---]  [=\ \@/ C / /
       ___      ___   ]/000o   /__\ \ C / /
          \    /              /....\ \_/ /
       ....\||/....           [___/=\___/
      .    .  .    .          [...] [...]
     .      ..      .         [___/ \___]
     .    0 .. 0    .         <---> <--->
  /\/\.    .  .    ./\/\      [..]   [..]
 / / / .../|  |\... \ \ \    _[__]   [__]_
/ / /       \/       \ \ \  [____>   <____]
	`
	color.HiMagenta(banner)
}
