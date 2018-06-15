package method

type ListingFile struct {
	ListingFileID int    `json:"listing_file_id"` //The numeric ID of the listing file.
	ListingID     int    `json:"listing_id"`      //The numeric ID of the listing the file belongs to.
	Rank          int    `json:"rank"`            //Display order.
	Filename      string `json:"filename"`        //The file's displayable name.
	Filesize      string `json:"filesize"`        //The file's size in a displayable format.
	SizeBytes     int    `json:"size_bytes"`      //The file's size in raw bytes.
	Filetype      string `json:"filetype"`        //The file's mimetype.
	CreateDate    int    `json:"create_date"`     //The time when this file was uploaded, in Epoch seconds.
}
