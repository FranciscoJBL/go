package video

import "io/ioutil"

// Video : A video struct.
type Video struct {
	Title       string
	Media       []byte
	Description string
}

// Upload : Upload a video.
func (video *Video) Upload() error {
	filename := video.Title + ".txt"
	return ioutil.WriteFile(filename, video.Media, 0600)
}

// RequestVideo : Request a Video By title.
func RequestVideo(title string) (*Video, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Video{Title: title, Media: body}, nil
}
