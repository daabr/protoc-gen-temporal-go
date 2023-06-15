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
// source: 5_worker_with_empty_options.proto

package worker

import (
	client "go.temporal.io/sdk/client"
	worker "go.temporal.io/sdk/worker"
	log "log"
)

func StartWorkerWorkerWithEmptyOptions(c client.Client) {
	taskQueue := "my-task-queue"
	opts := worker.Options{}
	w := worker.New(c, taskQueue, opts)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Failed to start Temporal worker:", err)
	}
}
