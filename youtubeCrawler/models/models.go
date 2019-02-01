package structs

/* VideoData struct to hold data about video link
Id - PK
JobId - Id of the job that fetched and parsed data
Title - title of the video
Link - part of the link: eg /watch?v=sWcXBRTGrWo
*/
type VideoData struct {
	Id    int
	JobId int
	Title string
	Link  string
}

type Job struct {
	Id              int
	Name            string
	FirstVideoLink  string
	FirstVideoTitle string
	NumIterations   uint
	Progress        uint
	Finished        bool
}
