// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compactor

import (
	"fmt"
	"math"
	"strconv"

	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	"github.com/banzaicloud/operator-tools/pkg/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (c *Compactor) deployment() (runtime.Object, reconciler.DesiredState, error) {
	const app = "compactor"
	name := app + "-deployment"
	compactor := c.objectStore.Spec.Compactor.DeepCopy()

	if c.objectStore.Spec.BucketWeb.Enabled {

		var deployment = &appsv1.Deployment{
			ObjectMeta: c.objectMeta(name, &compactor.BaseObject),
			Spec: appsv1.DeploymentSpec{
				Replicas: utils.IntPointer(1),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"app": app},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name:        app,
						Labels:      map[string]string{"app": app},
						Annotations: c.objectStore.Annotations,
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  app,
								Image: fmt.Sprintf("%s:%s", compactor.Image.Repository, compactor.Image.Tag),
								Args: []string{
									"compact",
									"--log.level=info",
									"--http-address=" + compactor.HTTPAddress,
									"--http-grace-period=" + strconv.Itoa(int(math.Floor(compactor.HTTPGracePeriod.Duration.Seconds()))) + "s",
									"--data-dir=" + compactor.DataDir,
									// TODO: get secret file path from secret mount
									"--objstore.config-file=/etc/config/object-store.yaml",
									"--consistency-delay=" + strconv.Itoa(int(math.Floor(compactor.ConsistencyDelay.Duration.Seconds()))) + "s",
									"--retention.resolution-raw=" + strconv.Itoa(int(math.Floor(compactor.RetentionResolutionRaw.Duration.Seconds()))) + "s",
									"--retention.resolution-5m=" + strconv.Itoa(int(math.Floor(compactor.RetentionResolution5m.Duration.Seconds()))) + "s",
									"--retention.resolution-1h=" + strconv.Itoa(int(math.Floor(compactor.RetentionResolution1h.Duration.Seconds()))) + "s",
									"--compact.concurrency=" + strconv.Itoa(compactor.CompactConcurrency),
								},
								Ports: []corev1.ContainerPort{
									{
										Name:          "http",
										ContainerPort: GetPort(compactor.HTTPAddress),
										Protocol:      corev1.ProtocolTCP,
									},
								},
								VolumeMounts: []corev1.VolumeMount{
									{
										Name:      "objectstore-secret",
										ReadOnly:  true,
										MountPath: "/etc/config/",
									},
								},
								Resources:       compactor.Resources,
								ImagePullPolicy: corev1.PullPolicy(compactor.Image.PullPolicy),
							},
						},
						Volumes: []corev1.Volume{
							{
								Name: "objectstore-secret",
								VolumeSource: corev1.VolumeSource{
									Secret: &corev1.SecretVolumeSource{
										SecretName: c.objectStore.Spec.Config.MountFrom.SecretKeyRef.Name,
									},
								},
							},
						},
					},
				},
			},
		}

		if compactor.Wait {
			deployment.Spec.Template.Spec.Containers[0].Args = append(deployment.Spec.Template.Spec.Containers[0].Args, "--wait")
		}
		if compactor.DownsamplingDisable {
			deployment.Spec.Template.Spec.Containers[0].Args = append(deployment.Spec.Template.Spec.Containers[0].Args, "--downsampling.disable")
		}

		return deployment, reconciler.StatePresent, nil
	}

	return &appsv1.Deployment{
		ObjectMeta: c.objectMeta(name, &compactor.BaseObject),
	}, reconciler.StateAbsent, nil
}
