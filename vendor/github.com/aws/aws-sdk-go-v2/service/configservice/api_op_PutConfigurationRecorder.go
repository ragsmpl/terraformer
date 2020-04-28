// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// The input for the PutConfigurationRecorder action.
type PutConfigurationRecorderInput struct {
	_ struct{} `type:"structure"`

	// The configuration recorder object that records each configuration change
	// made to the resources.
	//
	// ConfigurationRecorder is a required field
	ConfigurationRecorder *ConfigurationRecorder `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutConfigurationRecorderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationRecorderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationRecorderInput"}

	if s.ConfigurationRecorder == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationRecorder"))
	}
	if s.ConfigurationRecorder != nil {
		if err := s.ConfigurationRecorder.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationRecorder", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutConfigurationRecorderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationRecorderOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutConfigurationRecorder = "PutConfigurationRecorder"

// PutConfigurationRecorderRequest returns a request value for making API operation for
// AWS Config.
//
// Creates a new configuration recorder to record the selected resource configurations.
//
// You can use this action to change the role roleARN or the recordingGroup
// of an existing recorder. To change the role, call the action on the existing
// configuration recorder and specify a role.
//
// Currently, you can specify only one configuration recorder per region in
// your account.
//
// If ConfigurationRecorder does not have the recordingGroup parameter specified,
// the default is to record all supported resource types.
//
//    // Example sending a request using PutConfigurationRecorderRequest.
//    req := client.PutConfigurationRecorderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutConfigurationRecorder
func (c *Client) PutConfigurationRecorderRequest(input *PutConfigurationRecorderInput) PutConfigurationRecorderRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationRecorder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutConfigurationRecorderInput{}
	}

	req := c.newRequest(op, input, &PutConfigurationRecorderOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutConfigurationRecorderRequest{Request: req, Input: input, Copy: c.PutConfigurationRecorderRequest}
}

// PutConfigurationRecorderRequest is the request type for the
// PutConfigurationRecorder API operation.
type PutConfigurationRecorderRequest struct {
	*aws.Request
	Input *PutConfigurationRecorderInput
	Copy  func(*PutConfigurationRecorderInput) PutConfigurationRecorderRequest
}

// Send marshals and sends the PutConfigurationRecorder API request.
func (r PutConfigurationRecorderRequest) Send(ctx context.Context) (*PutConfigurationRecorderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationRecorderResponse{
		PutConfigurationRecorderOutput: r.Request.Data.(*PutConfigurationRecorderOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationRecorderResponse is the response type for the
// PutConfigurationRecorder API operation.
type PutConfigurationRecorderResponse struct {
	*PutConfigurationRecorderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationRecorder request.
func (r *PutConfigurationRecorderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}