// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package translation

const (
	// DefaultTranslationRulesYaml defines default translation rules that will be applied to metrics if
	// config.SendCompatibleMetrics set to true and config.TranslationRules not specified explicitly.
	// Keep it in YAML format to be able to easily copy and paste it in config if modifications needed.
	DefaultTranslationRulesYaml = `
translation_rules:

- action: rename_dimension_keys
  mapping:

    # dimensions
    container.image.name: container_image
    container.name: container_spec_name
    k8s.cluster.name: kubernetes_cluster
    k8s.daemonset.name: kubernetes_name
    k8s.daemonset.uid: kubernetes_uid
    k8s.deployment.name: kubernetes_name
    k8s.deployment.uid: kubernetes_uid
    k8s.hpa.name: kubernetes_name
    k8s.hpa.uid: kubernetes_uid
    k8s.namespace.name: kubernetes_namespace
    k8s.node.name: kubernetes_node
    k8s.node.uid: kubernetes_node_uid
    k8s.pod.name: kubernetes_pod_name
    k8s.pod.uid: kubernetes_pod_uid
    k8s.replicaset.name: kubernetes_name
    k8s.replicaset.uid: kubernetes_uid
    k8s.replicationcontroller.name: kubernetes_name
    k8s.replicationcontroller.uid: kubernetes_uid
    k8s.resourcequota.name: quota_name
    k8s.resourcequota.uid: kubernetes_uid
    k8s.statefulset.name: kubernetes_name
    k8s.statefulset.uid: kubernetes_uid

    # properties
    cronjob_uid: cronJob_uid
    cronjob: cronJob
    daemonset_uid: daemonSet_uid
    daemonset: daemonSet
    k8s.workload.kind: kubernetes_workload
    k8s.workload.name: kubernetes_workload_name
    replicaset_uid: replicaSet_uid
    replicaset: replicaSet
    statefulset_uid: statefulSet_uid
    statefulset: statefulSet

- action: rename_metrics
  mapping:

    # kubeletstats receiver metrics
    container.cpu.time: container_cpu_utilization
    container.filesystem.available: container_fs_available_bytes
    container.filesystem.capacity: container_fs_capacity_bytes
    container.filesystem.usage: container_fs_usage_bytes
    container.memory.available: container_memory_available_bytes
    container.memory.rss: container_memory_rss_bytes
    container.memory.usage: container_memory_usage_bytes
    k8s.node.cpu.utilization: cpu.utilization

    # k8s cluster receiver metrics
    k8s/container/cpu/limit: kubernetes.container_cpu_limit
    k8s/container/cpu/request: kubernetes.container_cpu_request
    k8s/container/memory/limit: kubernetes.container_memory_limit
    k8s/container/memory/request: kubernetes.container_memory_request
    k8s/container/ready: kubernetes.container_ready
    k8s/container/restarts: kubernetes.container_restart_count
    k8s/cronjob/active_jobs: kubernetes.cronjob.active
    k8s/daemon_set/current_scheduled_nodes: kubernetes.daemon_set.current_scheduled
    k8s/daemon_set/desired_scheduled_nodes: kubernetes.daemon_set.desired_scheduled
    k8s/daemon_set/misscheduled_nodes: kubernetes.daemon_set.misscheduled
    k8s/daemon_set/ready_nodes: kubernetes.daemon_set.ready
    k8s/deployment/available: kubernetes.deployment.available
    k8s/deployment/desired: kubernetes.deployment.desired
    k8s/job/active_pods: kubernetes.job.active
    k8s/job/desired_successful_pods: kubernetes.job.completions
    k8s/job/failed_pods: kubernetes.job.failed
    k8s/job/max_parallel_pods: kubernetes.job.parallelism
    k8s/job/successful_pods: kubernetes.job.succeeded
    k8s/hpa/current_replicas: kubernetes.hpa.status.current_replicas
    k8s/hpa/desired_replicas: kubernetes.hpa.status.desired_replicas
    k8s/hpa/max_replicas: kubernetes.hpa.spec.max_replicas
    k8s/hpa/min_replicas: kubernetes.hpa.spec.min_replicas
    k8s/namespace/phase: kubernetes.namespace_phase
    k8s/node/condition_memory_pressure: kubernetes.node_memory_pressure
    k8s/node/condition_network_unavailable: kubernetes.node_network_unavailable
    k8s/node/condition_out_of_disk: kubernetes.node_out_of_disk
    k8s/node/condition_p_i_d_pressure: kubernetes.node_p_i_d_pressure
    k8s/node/condition_ready: kubernetes.node_ready
    k8s/pod/phase: kubernetes.pod_phase
    k8s/replica_set/available: kubernetes.replica_set.available
    k8s/replica_set/desired: kubernetes.replica_set.desired
    k8s/replication_controller/available: kubernetes.replication_controller.available
    k8s/replication_controller/desired: kubernetes.replication_controller.desired
    k8s/resource_quota/hard_limt: kubernetes.resource_quota_hard
    k8s/resource_quota/used: kubernetes.resource_quota_used
    k8s/stateful_set/current_pods: kubernetes.stateful_set.current
    k8s/stateful_set/desired_pods: kubernetes.stateful_set.desired
    k8s/stateful_set/ready_pods: kubernetes.stateful_set.ready
    k8s/stateful_set/updated_pods: kubernetes.stateful_set.updated

# container network metrics
- action: split_metric
  metric_name: k8s.pod.network.io
  dimension_key: direction
  mapping: 
    receive: pod_network_receive_bytes_total
    transmit: pod_network_transmit_bytes_total
- action: split_metric
  metric_name: k8s.pod.network.errors
  dimension_key: direction
  mapping: 
    receive: pod_network_receive_errors_total
    transmit: pod_network_transmit_errors_total

# convert cpu metrics
- action: split_metric
  metric_name: system.cpu.time
  dimension_key: state
  mapping:
    idle: cpu.idle
    interrupt: cpu.interrupt
    system: cpu.system
    user: cpu.user
    steal: cpu.steal
    wait: cpu.wait
    softirq: cpu.softirq
    nice: cpu.nice
- action: multiply_float
  scale_factors_float:
    container_cpu_utilization: 100
    cpu.idle: 100
    cpu.interrupt: 100
    cpu.system: 100
    cpu.user: 100
    cpu.steal: 100
    cpu.wait: 100
    cpu.softirq: 100
    cpu.nice: 100
- action: convert_values
  types_mapping:
    container_cpu_utilization: int
    cpu.idle: int
    cpu.interrupt: int
    cpu.system: int
    cpu.user: int
    cpu.steal: int
    cpu.wait: int
    cpu.softirq: int
    cpu.nice: int

# compute cpu.num_processors
- action: copy_metrics
  mapping:
    cpu.idle: cpu.num_processors
- action: aggregate_metric
  metric_name: cpu.num_processors
  aggregation_method: count
  dimensions:
  - host
  - kubernetes_node
  - kubernetes_cluster

# compute memory.total
- action: copy_metrics
  mapping:
    system.memory.usage: memory.total
  dimension_key: state
  dimension_values:
    buffered: true
    cached: true
    free: true
    used: true
- action: aggregate_metric
  metric_name: memory.total
  aggregation_method: sum
  dimensions:
  - host
  - kubernetes_node
  - kubernetes_cluster

# convert memory metrics
- action: split_metric
  metric_name: system.memory.usage
  dimension_key: state
  mapping:
    buffered: memory.buffered
    cached: memory.cached
    free: memory.free
    slab_reclaimable: memory.slab_recl
    slab_unreclaimable: memory.slab_unrecl
    used: memory.used

# convert filesystem metrics
- action: split_metric
  metric_name: system.filesystem.usage
  dimension_key: state
  mapping:
    free: df_complex.free
    reserved: df_complex.reserved
    used: df_complex.used
- action: split_metric
  metric_name: system.filesystem.inodes.usage
  dimension_key: state
  mapping:
    free: df_inodes.free
    used: df_inodes.used
`
)
