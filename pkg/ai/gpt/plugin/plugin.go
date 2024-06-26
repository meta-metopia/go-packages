package plugin

import "github.com/meta-metopia/go-packages/pkg/ai/gpt/dto"

type ConvertedOutputAction string

// ContinueOutputAction will call the next plugin
var ContinueOutputAction ConvertedOutputAction = "continue"

// TerminateOutputAction will terminate the output
var TerminateOutputAction ConvertedOutputAction = "terminate"

type ConvertedResponse struct {
	Action       ConvertedOutputAction
	Message      *dto.Message
	AddToHistory bool
}

type Interface interface {
	// Name returns the name of the plugin
	Name() string
	// Description returns the description of the plugin
	Description() string
	// ConvertInput converts the input to a string that can be used by the LLM Model
	ConvertInput(input any) (any, error)
	// ConvertOutput converts the output from the LLM Model to the expected output.
	// Return nil will have 0 impact on the output.
	// Return a non-nil message will replace the output with the returned message or add it to the output depending on the `action`.
	ConvertOutput(response dto.Message) (*ConvertedResponse, error)
}

type Client struct {
}

func (c *Client) ConvertInput(input any) (any, error) {
	return nil, nil
}

func (c *Client) ConvertOutput(response dto.Message) (*ConvertedResponse, error) {
	return nil, nil
}
