package v1

type Response struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
}

type Shell struct {
	PodName string `json:"pod_name"`
}