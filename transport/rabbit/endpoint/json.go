package endpoint

import "github.com/vinceprignano/bunny/transport/rabbit"

type JsonEndpoint struct {
	EndpointName string
	Handler      func(*rabbit.RabbitRequest) ([]byte, error)
}

func (j *JsonEndpoint) Name() string {
	return j.EndpointName
}

func (j *JsonEndpoint) HandleRequest(req *rabbit.RabbitRequest) ([]byte, error) {
	return j.Handler(req)
}
