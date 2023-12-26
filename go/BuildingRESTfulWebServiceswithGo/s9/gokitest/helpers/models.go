package helpers

type EncryptRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type EncryptResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

type DecryptRequest struct {
	Cipher string `json:"cipher"`
	Key    string `json:"key"`
}

type DecryptResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}
