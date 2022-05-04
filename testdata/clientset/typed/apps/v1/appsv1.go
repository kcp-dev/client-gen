//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code auto-generated. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	appsapiv1 "github.com/kcp-dev/client-gen/testdata/pkg/apis/apps/v1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type WrappedAppsV1 struct {
	Cluster  logicalcluster.LogicalCluster
	Delegate appsv1.AppsV1Interface
}

func (w *WrappedAppsV1) RESTClient() rest.Interface {
	return w.Delegate.RESTClient()
}

func (w *WrappedAppsV1) Deployments(namespace string) appsv1.DeploymentInterface {
	return &wrappedDeployment{
		cluster:  w.Cluster,
		delegate: w.Delegate.Deployments(namespace),
	}
}

type wrappedDeployment struct {
	cluster  logicalcluster.LogicalCluster
	delegate appsv1.DeploymentInterface
}

func (w *wrappedDeployment) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

func (w *wrappedDeployment) Create(ctx context.Context, deployment *appsapiv1.Deployment, opts metav1.CreateOptions) (*appsapiv1.Deployment, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, deployment, opts)
}

func (w *wrappedDeployment) Update(ctx context.Context, deployment *appsapiv1.Deployment, opts metav1.UpdateOptions) (*appsapiv1.Deployment, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, deployment, opts)
}

func (w *wrappedDeployment) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

func (w *wrappedDeployment) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

func (w *wrappedDeployment) Get(ctx context.Context, name string, opts metav1.GetOptions) (*appsapiv1.Deployment, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

func (w *wrappedDeployment) List(ctx context.Context, opts metav1.ListOptions) (*appsapiv1.DeploymentList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

func (w *wrappedDeployment) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

func (w *wrappedDeployment) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *appsapiv1.Deployment, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}
