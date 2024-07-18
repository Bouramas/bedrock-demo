# bedrock-demo

A demonstration of using the AWS Bedrock service with Go to list available models and interact with foundation models, starting with Anthropic's Claude 3 Sonnet.

## Overview

This project showcases how to:

1. Connect to AWS Bedrock using Go
2. List all available foundation models
3. Interact with foundation models:
    - Anthropic Claude 3 Sonnet [anthropic.claude-3-sonnet-20240229-v1:0]

The project is designed to be extensible, with plans to add support for more models in the future.

## Prerequisites

- AWS account with Bedrock access - make sure to setup your `.aws/credentials file`
- Go 1.22
- AWS SDK for Go v2

## Installation

Clone this repository: 
   `git clone git@github.com:Bouramas/bedrock-demo.git`

## Usage

Run the main program: `go run main.go` and pass an argument:

    - models - list all models available in AWS Bedrock
    - claude - run a sample with Claude

## ☎️ Get in Touch

I'm always open to discussions, collaborations, and feedback. If you have any questions or just want to connect, feel free to reach out!

- **Email:** gbouramas@gmail.com
- **LinkedIn:** [Giannis Bouramas](https://www.linkedin.com/in/bouramas)