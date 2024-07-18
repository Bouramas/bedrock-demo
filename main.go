package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"bedrock-demo/claude"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument:")
		fmt.Println("models - list all models available in AWS Bedrock")
		fmt.Println("claude - run a sample with Claude")
		return
	}

	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	param := os.Args[1]
	if param == "models" {
		listModels(&sdkConfig)
	} else if param == "claude" {
		runClaude(&sdkConfig)
	} else {
		fmt.Println("unknown param: ", param)
	}

}

func runClaude(sdkConfig *aws.Config) {

	modelId := "anthropic.claude-3-sonnet-20240229-v1:0"
	anthropicVersion := "bedrock-2023-05-31"
	prompt := "What is the capital of Greece?"

	request := claude.BedrockRequest{
		AnthropicVersion: anthropicVersion,
		MaxTokens:        100,
		Messages: []claude.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 1,
	}

	body, err := json.Marshal(request)
	if err != nil {
		log.Panicln("Couldn't marshal the request: ", err)
	}

	client := bedrockruntime.NewFromConfig(*sdkConfig)

	modelOutput, err := client.InvokeModel(context.Background(), &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(modelId),
		ContentType: aws.String("application/json"),
		Body:        body,
	})

	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "no such host") {
			fmt.Printf("Error: The Bedrock service is not available in the selected region. Please double-check the service availability for your region at https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/.\n")
		} else if strings.Contains(errMsg, "Could not resolve the foundation model") {
			fmt.Printf("Error: Could not resolve the foundation model from model identifier: \"%v\". Please verify that the requested model exists and is accessible within the specified region.\n", modelId)
		} else {
			fmt.Printf("Error: Couldn't invoke Anthropic Claude. Here's why: %v\n", err)
		}
		os.Exit(1)
	}

	var response claude.BedrockResponse

	err = json.Unmarshal(modelOutput.Body, &response)

	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}
	fmt.Println("Prompt:\n", prompt)
	fmt.Printf("Response from Anthropic Claude:  %v \n", response)
	fmt.Printf("Response from Anthropic Claude:  %v \n", response.Content[0].Text)
}

func listModels(sdkConfig *aws.Config) {

	bedrockClient := bedrock.NewFromConfig(*sdkConfig)
	result, err := bedrockClient.ListFoundationModels(context.TODO(), &bedrock.ListFoundationModelsInput{})
	if err != nil {
		fmt.Printf("Couldn't list foundation models. Here's why: %v\n", err)
		return
	}
	if len(result.ModelSummaries) == 0 {
		fmt.Println("There are no foundation models.")
	}
	for _, modelSummary := range result.ModelSummaries {
		fmt.Println(*modelSummary.ModelId)
	}
}
