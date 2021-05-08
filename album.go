package main

type Photo struct {
	URI               string `json:"uri"`
	CreationTimestamp int    `json:"creation_timestamp"`
	MediaMetadata     struct {
		PhotoMetadata struct {
			ExifData []struct {
				UploadIP string `json:"upload_ip"`
			} `json:"exif_data"`
		} `json:"photo_metadata"`
	} `json:"media_metadata"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type CoverPhoto struct {
	URI               string `json:"uri"`
	CreationTimestamp int    `json:"creation_timestamp"`
	MediaMetadata     struct {
		PhotoMetadata struct {
			ExifData []struct {
				UploadIP string `json:"upload_ip"`
			} `json:"exif_data"`
		} `json:"photo_metadata"`
	} `json:"media_metadata"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Album struct {
	Name                  string     `json:"name"`
	Photos                []Photo    `json:"photos"`
	CoverPhoto            CoverPhoto `json:"cover_photo"`
	LastModifiedTimestamp int        `json:"last_modified_timestamp"`
	Description           string     `json:"description"`
}
