/*
MIT License

Copyright (c) 2023 Daniel Abraham

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	workerpb "github.com/daabr/protoc-gen-temporal-go/proto/temporal"
)

func GenerateWorker(g *protogen.GeneratedFile, service *protogen.Service) {
	worker := proto.GetExtension(service.Desc.Options(), workerpb.E_Worker).(*workerpb.Worker)
	if worker == nil || worker.TaskQueue == "" {
		g.Skip()
		return
	}

	g.P("func StartWorker", service.GoName, "(c ", clientPackage.Ident("Client"), ") {")
	g.P(`taskQueue := "`, worker.TaskQueue, `"`)
	g.P("opts := ", workerPackage.Ident("Options"), "{")
	if worker.Options != nil {
		nonDefaultWorkerOptions(g, worker.Options)
	}
	g.P("}")

	g.P("w := ", workerPackage.Ident("New"), "(c, taskQueue, opts)")
	g.P()

	registerWorkerMethods(g, service.Methods)
	g.P()

	g.P("if err := w.Run(", workerPackage.Ident("InterruptCh"), "()); err != nil {")
	g.P(logPackage.Ident("Fatalln"), `("Failed to start Temporal worker:", err)`)
	g.P("}")

	g.P("}")
	g.P()
}

func nonDefaultWorkerOptions(g *protogen.GeneratedFile, o *workerpb.WorkerOptions) {
	options := []struct {
		value  interface{}
		goName string
	}{
		{
			o.MaxConcurrentActivityExecutionSize,
			"MaxConcurrentActivityExecutionSize",
		},
		{
			o.WorkerActivitiesPerSecond,
			"WorkerActivitiesPerSecond",
		},
		{
			o.MaxConcurrentLocalActivityExecutionSize,
			"MaxConcurrentLocalActivityExecutionSize",
		},
		{
			o.WorkerLocalActivitiesPerSecond,
			"WorkerLocalActivitiesPerSecond",
		},
		{
			o.TaskQueueActivitiesPerSecond,
			"TaskQueueActivitiesPerSecond",
		},
		{
			o.MaxConcurrentActivityTaskPollers,
			"MaxConcurrentActivityTaskPollers",
		},
		{
			o.MaxConcurrentWorkflowTaskExecutionSize,
			"MaxConcurrentWorkflowTaskExecutionSize",
		},
		{
			o.MaxConcurrentWorkflowTaskPollers,
			"MaxConcurrentWorkflowTaskPollers",
		},
		{
			o.EnableLoggingInReplay,
			"EnableLoggingInReplay",
		},
		// Deprecated: DisableStickyExecution
		{
			o.StickyScheduleToStartTimeout,
			"StickyScheduleToStartTimeout",
		},
		// TODO: BackgroundActivityContext
		// TODO: WorkflowPanicPolicy
		{
			o.WorkerStopTimeout,
			"WorkerStopTimeout",
		},
		{
			o.EnableSessionWorker,
			"EnableSessionWorker",
		},
		{
			o.MaxConcurrentSessionExecutionSize,
			"MaxConcurrentSessionExecutionSize",
		},
		{
			o.DisableWorkflowWorker,
			"DisableWorkflowWorker",
		},
		{
			o.LocalActivityWorkerOnly,
			"LocalActivityWorkerOnly",
		},
		{
			o.Identity,
			"Identity",
		},
		{
			o.DeadlockDetectionTimeout,
			"DeadlockDetectionTimeout",
		},
		{
			o.MaxHeartbeatThrottleInterval,
			"MaxHeartbeatThrottleInterval",
		},
		{
			o.DefaultHeartbeatThrottleInterval,
			"DefaultHeartbeatThrottleInterval",
		},
		// TODO: Interceptors []WorkerInterceptor
		// TODO: OnFatalError func(error)
		{
			o.DisableEagerActivities,
			"DisableEagerActivities",
		},
		{
			o.MaxConcurrentEagerActivityExecutionSize,
			"MaxConcurrentEagerActivityExecutionSize",
		},
		{
			o.DisableRegistrationAliasing,
			"DisableRegistrationAliasing",
		},
		{
			o.BuildId,
			"BuildID",
		},
		{
			o.UseBuildIdForVersioning,
			"UseBuildIDForVersioning",
		},
	}
	for _, option := range options {
		if v, ok := option.value.(bool); ok && v {
			g.P(option.goName, ": ", v, ",")
			continue
		}
		if v, ok := option.value.(float64); ok && v != 0 {
			g.P(option.goName, ": ", v, ",")
			continue
		}
		if v, ok := option.value.(int32); ok && v != 0 {
			g.P(option.goName, ": ", v, ",")
			continue
		}
		if v, ok := option.value.(string); ok && v != "" {
			g.P(option.goName, `: "`, v, `",`)
			continue
		}
		if v, ok := option.value.(*durationpb.Duration); ok && v != nil {
			s := v.AsDuration().Seconds()
			g.P(option.goName, ": ", timePackage.Ident("Duration"), "(", s, " * float64(time.Second)),")
			continue
		}
	}
}

func registerWorkerMethods(g *protogen.GeneratedFile, methods []*protogen.Method) {
	for _, m := range methods {
		w := proto.GetExtension(m.Desc.Options(), workerpb.E_Workflow).(*workerpb.Workflow)
		if w != nil {
			g.P("w.RegisterWorkflow(", m.GoName, ")")
		} else {
			g.P("w.RegisterActivity(", m.GoName, ")")
		}
	}
}
