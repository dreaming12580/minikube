/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package ktmpl

import "text/template"

// V1Beta1 is kubeadm config template for Kubernetes v1.13+
var V1Beta1 = template.Must(template.New("configTmpl-v1beta1").Funcs(template.FuncMap{
	"printMapInOrder": printMapInOrder,
}).Parse(`apiVersion: kubeadm.k8s.io/v1beta1
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: {{.AdvertiseAddress}}
  bindPort: {{.APIServerPort}}
bootstrapTokens:
  - groups:
      - system:bootstrappers:kubeadm:default-node-token
    ttl: 24h0m0s
    usages:
      - signing
      - authentication
nodeRegistration:
  criSocket: {{if .CRISocket}}{{.CRISocket}}{{else}}/var/run/dockershim.sock{{end}}
  name: {{.NodeName}}
  taints: []
---
apiVersion: kubeadm.k8s.io/v1beta1
kind: ClusterConfiguration
{{ if .ImageRepository}}imageRepository: {{.ImageRepository}}
{{end}}{{range .ExtraArgs}}{{.Component}}:
  extraArgs:
{{- range $i, $val := printMapInOrder .Options ": " }}
    {{$val}}
{{- end}}
{{end -}}
{{if .FeatureArgs}}featureGates:
{{range $i, $val := .FeatureArgs}}{{$i}}: {{$val}}
{{end -}}{{end -}}
certificatesDir: {{.CertDir}}
clusterName: kubernetes
controlPlaneEndpoint: localhost:{{.APIServerPort}}
dns:
  type: CoreDNS
etcd:
  local:
    dataDir: {{.EtcdDataDir}}
kubernetesVersion: {{.KubernetesVersion}}
networking:
  dnsDomain: {{if .DNSDomain}}{{.DNSDomain}}{{else}}cluster.local{{end}}{{if ne .PodSubnet ""}}
  podSubnet: {{.PodSubnet}}{{end}}
  serviceSubnet: {{.ServiceCIDR}}
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
imageGCHighThresholdPercent: 100
evictionHard:
  nodefs.available: "0%"
  nodefs.inodesFree: "0%"
  imagefs.available: "0%"
`))
