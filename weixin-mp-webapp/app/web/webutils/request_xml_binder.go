package webutils

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type RequestBodyBinder struct {
	data []byte
}

func (inst *RequestBodyBinder) Init(c *gin.Context) error {
	return inst.readBodyData(c)
}

func (inst *RequestBodyBinder) readBodyData(c *gin.Context) error {

	body := c.Request.Body
	if body == nil {
		return fmt.Errorf("no body")
	}

	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	inst.data = data
	return nil
}

func (inst *RequestBodyBinder) BindXML(v any) error {
	return xml.Unmarshal(inst.data, v)
}
