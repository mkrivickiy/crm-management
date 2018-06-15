package method

type ImageType struct {
	Code        string   `json:"code"`  //Code for this image type, used in image URLs
	Description string   `json:"desc"`  //Text description of the image type
	Sizes       []string `json:"sizes"` //Available sizes for this image type
}
