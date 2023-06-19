//
//MIT License
//
//Copyright (c) 2023 Daniel Abraham
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

// Code generated by protoc-gen-temporal-go. DO NOT EDIT.
// versions:
// - protoc-gen-temporal-go v0.0.0
// - protoc                 v4.23.2
// source: activity_with_empty_options.proto

package workflows

import (
	context "context"
	client "go.temporal.io/sdk/client"
	worker "go.temporal.io/sdk/worker"
	workflow "go.temporal.io/sdk/workflow"
	log "log"
)

func StartWorkerActivityWithEmptyOptions(c client.Client) {
	taskQueue := "my-task-queue"
	opts := worker.Options{}
	w := worker.New(c, taskQueue, opts)

	w.RegisterActivity(Foo)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Failed to start Temporal worker:", err)
	}
}

type ActivityWithEmptyOptionsTemporalClient interface {
	// Foo activity.
	Foo(ctx context.Context, in *FooInput) (*FooOutput, error)
}

type activityWithEmptyOptionsTemporalClient struct {
	t client.Client
}

func NewActivityWithEmptyOptionsTemporalClient(c client.Client) *ActivityWithEmptyOptionsTemporalClient {
	return &activityWithEmptyOptionsTemporalClient{c}
}

// Foo activity.
//
// This method starts the activity with pre-configured options, and returns a
// Future to interact with it until completion. For more information, see
// https://docs.temporal.io/dev-guide/go/foundations#activity-execution.
func (c *activityWithEmptyOptionsTemporalClient) StartActivityActivityWithEmptyOptionsFoo(ctx workflow.Context, in *FooInput) workflow.Future {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{})
	return workflow.ExecuteActivity(ctx, c.Foo, in)
}

// Foo activity.
//
// This method executes the activity with pre-configured options, blocks until
// completion, and returns the output/error results. For more information, see
// https://docs.temporal.io/dev-guide/go/foundations#activity-execution.
func (c *activityWithEmptyOptionsTemporalClient) ExecuteActivityActivityWithEmptyOptionsFoo(ctx workflow.Context, in *FooInput) (*FooOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{})
	var out *FooOutput
	err := workflow.ExecuteActivity(ctx, c.Foo, in).Get(ctx, &out)
	return out, err
}

// Foo activity.
//
// This method starts the activity (locally) with pre-configured options, and
// returns a Future to interact with it until completion. For more information,
// see https://docs.temporal.io/dev-guide/go/foundations#activity-execution
// and https://docs.temporal.io/activities#local-activity.
func (c *activityWithEmptyOptionsTemporalClient) StartLocalActivityActivityWithEmptyOptionsFoo(ctx workflow.Context, in *FooInput) workflow.Future {
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{})
	return workflow.ExecuteActivity(ctx, c.Foo, in)
}

// Foo activity.
//
// This method executes the activity (locally) with pre-configured options,
// blocks until completion, and returns the output/error. For more information,
// see https://docs.temporal.io/dev-guide/go/foundations#activity-execution
// and https://docs.temporal.io/activities#local-activity.
func (c *activityWithEmptyOptionsTemporalClient) ExecuteLocalActivityActivityWithEmptyOptionsFoo(ctx workflow.Context, in *FooInput) (*FooOutput, error) {
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{})
	var out *FooOutput
	err := workflow.ExecuteLocalActivity(ctx, c.Foo, in).Get(ctx, &out)
	return out, err
}
