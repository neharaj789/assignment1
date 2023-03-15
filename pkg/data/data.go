package data

import (
	"assignment1.1/pkg/video"
)

var Videoplat = []*video.Video{
	&video.Video{Id: "1", Title: "first video", Viewcount: 0},
	&video.Video{Id: "2", Title: "Second video", Viewcount: 0},
	&video.Video{Id: "3", Title: "Third video", Viewcount: 0},
	&video.Video{Id: "4", Title: "Fourth video", Viewcount: 0},
}

/* func Myfunc() {

	Videoplat = append(Videoplat, video.Video{Id: "1", Title: "first video", Viewcount: 0})
	Videoplat = append(Videoplat, video.Video{Id: "2", Title: "Second video", Viewcount: 0})
	Videoplat = append(Videoplat, video.Video{Id: "3", Title: "Third video", Viewcount: 0})
	Videoplat = append(Videoplat, video.Video{Id: "4", Title: "Fourth video", Viewcount: 0})
}
*/
