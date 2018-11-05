/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/projectriff/riff/pkg/core"
)

/*

The following declarations determine the docker images and kubernetes configuration files that are used by all packages
in this repository. To preserve this consistency, the declarations are shared between the main packages. This ensures
that other packages do not import the declarations directly but must rely on the values being injected and passed as
parameters.

This file is defined in the root of this repository and symbolically linked from other main packages such as that in
./distro.

*/

var (
	buildpacks = map[string]string{
		"java": "projectriff/buildpack",
	}
	invokers = map[string]string{
		"jar":     "https://github.com/projectriff/java-function-invoker/raw/v0.1.1/java-invoker.yaml",
		"command": "https://github.com/projectriff/command-function-invoker/raw/v0.0.7/command-invoker.yaml",
		"node":    "https://github.com/projectriff/node-function-invoker/raw/v0.0.8/node-invoker.yaml",
	}
	manifests = map[string]*core.Manifest{
		"latest": {
			ManifestVersion: "0.1",
			Istio: []string{
				"https://storage.googleapis.com/knative-releases/serving/latest/istio.yaml",
			},
			Knative: []string{
				"https://storage.googleapis.com/knative-releases/build/latest/release.yaml",
				"https://storage.googleapis.com/knative-releases/serving/latest/serving.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/latest/release.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/latest/release-clusterbus-stub.yaml",
			},
			Namespace: []string{
				"https://storage.googleapis.com/riff-releases/latest/riff-build.yaml",
				"https://storage.googleapis.com/riff-releases/riff-cnb-buildtemplate-0.1.0.yaml",
			},
		},
		"stable": {
			ManifestVersion: "0.1",
			Istio: []string{
				"https://storage.googleapis.com/knative-releases/serving/previous/v0.2.0/istio.yaml",
			},
			Knative: []string{
				"https://storage.googleapis.com/knative-releases/build/previous/v0.2.0/release.yaml",
				"https://storage.googleapis.com/knative-releases/serving/previous/v20181101-v0.2.0-11-g877523d/serving.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20181031-a2f9417/release.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20181031-a2f9417/release-clusterbus-stub.yaml",
			},
			Namespace: []string{
				"https://storage.googleapis.com/riff-releases/previous/riff-build/riff-build-0.1.0.yaml",
				"https://storage.googleapis.com/riff-releases/riff-cnb-buildtemplate-0.1.0.yaml",
			},
		},
		"v0.1.3": {
			ManifestVersion: "0.1",
			Istio: []string{
				"https://storage.googleapis.com/knative-releases/serving/previous/v20180921-69811e7/istio.yaml",
			},
			Knative: []string{
				"https://storage.googleapis.com/knative-releases/serving/previous/v20180921-69811e7/release-no-mon.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20180921-01f95cb/release.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20180921-01f95cb/release-clusterbus-stub.yaml",
			},
			Namespace: []string{
				"https://storage.googleapis.com/riff-releases/previous/riff-build/riff-build-0.1.0.yaml",
				"https://storage.googleapis.com/riff-releases/riff-cnb-buildtemplate-0.1.0.yaml",
			},
		},
		"v0.1.2": {
			ManifestVersion: "0.1",
			Istio: []string{
				"https://storage.googleapis.com/knative-releases/serving/previous/v20180828-7c20145/istio.yaml",
			},
			Knative: []string{
				"https://storage.googleapis.com/knative-releases/serving/previous/v20180828-7c20145/release-no-mon.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20180830-5d35af5/release.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20180830-5d35af5/release-clusterbus-stub.yaml",
			},
			Namespace: []string{
				"https://storage.googleapis.com/riff-releases/previous/riff-build/riff-build-0.1.0.yaml",
			},
		},
		"v0.1.1": {
			ManifestVersion: "0.1",
			Istio: []string{
				"https://storage.googleapis.com/riff-releases/istio/istio-1.0.0-riff-crds.yaml",
				"https://storage.googleapis.com/riff-releases/istio/istio-1.0.0-riff-main.yaml",
			},
			Knative: []string{
				"https://storage.googleapis.com/knative-releases/serving/previous/v20180809-6b01d8e/release-no-mon.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20180809-34ab480/release.yaml",
				"https://storage.googleapis.com/knative-releases/eventing/previous/v20180809-34ab480/release-clusterbus-stub.yaml",
			},
			Namespace: []string{
				"https://storage.googleapis.com/riff-releases/previous/riff-build/riff-build-0.1.0.yaml",
			},
		},
		"v0.1.0": {
			ManifestVersion: "0.1",
			Istio: []string{
				"https://storage.googleapis.com/riff-releases/istio-riff-0.1.0.yaml",
			},
			Knative: []string{
				"https://storage.googleapis.com/riff-releases/release-no-mon-riff-0.1.0.yaml",
				"https://storage.googleapis.com/riff-releases/release-eventing-riff-0.1.0.yaml",
				"https://storage.googleapis.com/riff-releases/release-eventing-clusterbus-stub-riff-0.1.0.yaml",
			},
			Namespace: []string{
				"https://storage.googleapis.com/riff-releases/previous/riff-build/riff-build-0.1.0.yaml",
			},
		},
	}
)
