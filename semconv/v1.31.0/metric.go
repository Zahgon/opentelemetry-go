package semconv

const (
	AzureCosmosDBClientActiveInstanceCountName        = "azure.cosmosdb.client.active_instance.count"
	AzureCosmosDBClientActiveInstanceCountUnit        = "{instance}"
	AzureCosmosDBClientActiveInstanceCountDescription = "Number of active client instances"

	AzureCosmosDBClientOperationRequestChargeName        = "azure.cosmosdb.client.operation.request_charge"
	AzureCosmosDBClientOperationRequestChargeUnit        = "{request_unit}"
	AzureCosmosDBClientOperationRequestChargeDescription = "[Request units](https://learn.microsoft.com/azure/cosmos-db/request-units) consumed by the operation"

	CICDPipelineRunActiveName        = "cicd.pipeline.run.active"
	CICDPipelineRunActiveUnit        = "{run}"
	CICDPipelineRunActiveDescription = "The number of pipeline runs currently active in the system by state."

	CICDPipelineRunDurationName        = "cicd.pipeline.run.duration"
	CICDPipelineRunDurationUnit        = "s"
	CICDPipelineRunDurationDescription = "Duration of a pipeline run grouped by pipeline, state and result."

	CICDPipelineRunErrorsName        = "cicd.pipeline.run.errors"
	CICDPipelineRunErrorsUnit        = "{error}"
	CICDPipelineRunErrorsDescription = "The number of errors encountered in pipeline runs (eg. compile, test failures)."

	CICDSystemErrorsName        = "cicd.system.errors"
	CICDSystemErrorsUnit        = "{error}"
	CICDSystemErrorsDescription = "The number of errors in a component of the CICD system (eg. controller, scheduler, agent)."

	CICDWorkerCountName        = "cicd.worker.count"
	CICDWorkerCountUnit        = "{count}"
	CICDWorkerCountDescription = "The number of workers on the CICD system by state."

	ContainerCPUTimeName        = "container.cpu.time"
	ContainerCPUTimeUnit        = "s"
	ContainerCPUTimeDescription = "Total CPU time consumed"

	ContainerCPUUsageName        = "container.cpu.usage"
	ContainerCPUUsageUnit        = "{cpu}"
	ContainerCPUUsageDescription = "Container's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"

	ContainerDiskIoName        = "container.disk.io"
	ContainerDiskIoUnit        = "By"
	ContainerDiskIoDescription = "Disk bytes for the container."

	ContainerMemoryUsageName        = "container.memory.usage"
	ContainerMemoryUsageUnit        = "By"
	ContainerMemoryUsageDescription = "Memory usage of the container."

	ContainerNetworkIoName        = "container.network.io"
	ContainerNetworkIoUnit        = "By"
	ContainerNetworkIoDescription = "Network bytes for the container."

	ContainerUptimeName        = "container.uptime"
	ContainerUptimeUnit        = "s"
	ContainerUptimeDescription = "The time the container has been running"

	CPUFrequencyName        = "cpu.frequency"
	CPUFrequencyUnit        = "Hz"
	CPUFrequencyDescription = "Operating frequency of the logical CPU in Hertz."

	CPUTimeName        = "cpu.time"
	CPUTimeUnit        = "s"
	CPUTimeDescription = "Seconds each logical CPU spent on each mode"

	CPUUtilizationName        = "cpu.utilization"
	CPUUtilizationUnit        = "1"
	CPUUtilizationDescription = "For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time."

	DBClientConnectionCountName        = "db.client.connection.count"
	DBClientConnectionCountUnit        = "{connection}"
	DBClientConnectionCountDescription = "The number of connections that are currently in state described by the `state` attribute"

	DBClientConnectionCreateTimeName        = "db.client.connection.create_time"
	DBClientConnectionCreateTimeUnit        = "s"
	DBClientConnectionCreateTimeDescription = "The time it took to create a new connection"

	DBClientConnectionIdleMaxName        = "db.client.connection.idle.max"
	DBClientConnectionIdleMaxUnit        = "{connection}"
	DBClientConnectionIdleMaxDescription = "The maximum number of idle open connections allowed"

	DBClientConnectionIdleMinName        = "db.client.connection.idle.min"
	DBClientConnectionIdleMinUnit        = "{connection}"
	DBClientConnectionIdleMinDescription = "The minimum number of idle open connections allowed"

	DBClientConnectionMaxName        = "db.client.connection.max"
	DBClientConnectionMaxUnit        = "{connection}"
	DBClientConnectionMaxDescription = "The maximum number of open connections allowed"

	DBClientConnectionPendingRequestsName        = "db.client.connection.pending_requests"
	DBClientConnectionPendingRequestsUnit        = "{request}"
	DBClientConnectionPendingRequestsDescription = "The number of current pending requests for an open connection"

	DBClientConnectionTimeoutsName        = "db.client.connection.timeouts"
	DBClientConnectionTimeoutsUnit        = "{timeout}"
	DBClientConnectionTimeoutsDescription = "The number of connection timeouts that have occurred trying to obtain a connection from the pool"

	DBClientConnectionUseTimeName        = "db.client.connection.use_time"
	DBClientConnectionUseTimeUnit        = "s"
	DBClientConnectionUseTimeDescription = "The time between borrowing a connection and returning it to the pool"

	DBClientConnectionWaitTimeName        = "db.client.connection.wait_time"
	DBClientConnectionWaitTimeUnit        = "s"
	DBClientConnectionWaitTimeDescription = "The time it took to obtain an open connection from the pool"

	DBClientOperationDurationName        = "db.client.operation.duration"
	DBClientOperationDurationUnit        = "s"
	DBClientOperationDurationDescription = "Duration of database client operations."

	DBClientResponseReturnedRowsName        = "db.client.response.returned_rows"
	DBClientResponseReturnedRowsUnit        = "{row}"
	DBClientResponseReturnedRowsDescription = "The actual number of records returned by the database operation."

	DNSLookupDurationName        = "dns.lookup.duration"
	DNSLookupDurationUnit        = "s"
	DNSLookupDurationDescription = "Measures the time taken to perform a DNS lookup."

	FaaSColdstartsName        = "faas.coldstarts"
	FaaSColdstartsUnit        = "{coldstart}"
	FaaSColdstartsDescription = "Number of invocation cold starts"

	FaaSCPUUsageName        = "faas.cpu_usage"
	FaaSCPUUsageUnit        = "s"
	FaaSCPUUsageDescription = "Distribution of CPU usage per invocation"

	FaaSErrorsName        = "faas.errors"
	FaaSErrorsUnit        = "{error}"
	FaaSErrorsDescription = "Number of invocation errors"

	FaaSInitDurationName        = "faas.init_duration"
	FaaSInitDurationUnit        = "s"
	FaaSInitDurationDescription = "Measures the duration of the function's initialization, such as a cold start"

	FaaSInvocationsName        = "faas.invocations"
	FaaSInvocationsUnit        = "{invocation}"
	FaaSInvocationsDescription = "Number of successful invocations"

	FaaSInvokeDurationName        = "faas.invoke_duration"
	FaaSInvokeDurationUnit        = "s"
	FaaSInvokeDurationDescription = "Measures the duration of the function's logic execution"

	FaaSMemUsageName        = "faas.mem_usage"
	FaaSMemUsageUnit        = "By"
	FaaSMemUsageDescription = "Distribution of max memory usage per invocation"

	FaaSNetIoName        = "faas.net_io"
	FaaSNetIoUnit        = "By"
	FaaSNetIoDescription = "Distribution of net I/O usage per invocation"

	FaaSTimeoutsName        = "faas.timeouts"
	FaaSTimeoutsUnit        = "{timeout}"
	FaaSTimeoutsDescription = "Number of invocation timeouts"

	GenAIClientOperationDurationName        = "gen_ai.client.operation.duration"
	GenAIClientOperationDurationUnit        = "s"
	GenAIClientOperationDurationDescription = "GenAI operation duration"

	GenAIClientTokenUsageName        = "gen_ai.client.token.usage"
	GenAIClientTokenUsageUnit        = "{token}"
	GenAIClientTokenUsageDescription = "Measures number of input and output tokens used"

	GenAIServerRequestDurationName        = "gen_ai.server.request.duration"
	GenAIServerRequestDurationUnit        = "s"
	GenAIServerRequestDurationDescription = "Generative AI server request duration such as time-to-last byte or last output token"

	GenAIServerTimePerOutputTokenName        = "gen_ai.server.time_per_output_token"
	GenAIServerTimePerOutputTokenUnit        = "s"
	GenAIServerTimePerOutputTokenDescription = "Time per output token generated after the first token for successful responses"

	GenAIServerTimeToFirstTokenName        = "gen_ai.server.time_to_first_token"
	GenAIServerTimeToFirstTokenUnit        = "s"
	GenAIServerTimeToFirstTokenDescription = "Time to generate first token for successful responses"

	GoConfigGogcName        = "go.config.gogc"
	GoConfigGogcUnit        = "%"
	GoConfigGogcDescription = "Heap size target percentage configured by the user, otherwise 100."

	GoGoroutineCountName        = "go.goroutine.count"
	GoGoroutineCountUnit        = "{goroutine}"
	GoGoroutineCountDescription = "Count of live goroutines."

	GoMemoryAllocatedName        = "go.memory.allocated"
	GoMemoryAllocatedUnit        = "By"
	GoMemoryAllocatedDescription = "Memory allocated to the heap by the application."

	GoMemoryAllocationsName        = "go.memory.allocations"
	GoMemoryAllocationsUnit        = "{allocation}"
	GoMemoryAllocationsDescription = "Count of allocations to the heap by the application."

	GoMemoryGCGoalName        = "go.memory.gc.goal"
	GoMemoryGCGoalUnit        = "By"
	GoMemoryGCGoalDescription = "Heap size target for the end of the GC cycle."

	GoMemoryLimitName        = "go.memory.limit"
	GoMemoryLimitUnit        = "By"
	GoMemoryLimitDescription = "Go runtime memory limit configured by the user, if a limit exists."

	GoMemoryUsedName        = "go.memory.used"
	GoMemoryUsedUnit        = "By"
	GoMemoryUsedDescription = "Memory used by the Go runtime."

	GoProcessorLimitName        = "go.processor.limit"
	GoProcessorLimitUnit        = "{thread}"
	GoProcessorLimitDescription = "The number of OS threads that can execute user-level Go code simultaneously."

	GoScheduleDurationName        = "go.schedule.duration"
	GoScheduleDurationUnit        = "s"
	GoScheduleDurationDescription = "The time goroutines have spent in the scheduler in a runnable state before actually running."

	HTTPClientActiveRequestsName        = "http.client.active_requests"
	HTTPClientActiveRequestsUnit        = "{request}"
	HTTPClientActiveRequestsDescription = "Number of active HTTP requests."

	HTTPClientConnectionDurationName        = "http.client.connection.duration"
	HTTPClientConnectionDurationUnit        = "s"
	HTTPClientConnectionDurationDescription = "The duration of the successfully established outbound HTTP connections."

	HTTPClientOpenConnectionsName        = "http.client.open_connections"
	HTTPClientOpenConnectionsUnit        = "{connection}"
	HTTPClientOpenConnectionsDescription = "Number of outbound HTTP connections that are currently active or idle on the client."

	HTTPClientRequestBodySizeName        = "http.client.request.body.size"
	HTTPClientRequestBodySizeUnit        = "By"
	HTTPClientRequestBodySizeDescription = "Size of HTTP client request bodies."

	HTTPClientRequestDurationName        = "http.client.request.duration"
	HTTPClientRequestDurationUnit        = "s"
	HTTPClientRequestDurationDescription = "Duration of HTTP client requests."

	HTTPClientResponseBodySizeName        = "http.client.response.body.size"
	HTTPClientResponseBodySizeUnit        = "By"
	HTTPClientResponseBodySizeDescription = "Size of HTTP client response bodies."

	HTTPServerActiveRequestsName        = "http.server.active_requests"
	HTTPServerActiveRequestsUnit        = "{request}"
	HTTPServerActiveRequestsDescription = "Number of active HTTP server requests."

	HTTPServerRequestBodySizeName        = "http.server.request.body.size"
	HTTPServerRequestBodySizeUnit        = "By"
	HTTPServerRequestBodySizeDescription = "Size of HTTP server request bodies."

	HTTPServerRequestDurationName        = "http.server.request.duration"
	HTTPServerRequestDurationUnit        = "s"
	HTTPServerRequestDurationDescription = "Duration of HTTP server requests."

	HTTPServerResponseBodySizeName        = "http.server.response.body.size"
	HTTPServerResponseBodySizeUnit        = "By"
	HTTPServerResponseBodySizeDescription = "Size of HTTP server response bodies."

	HwEnergyName        = "hw.energy"
	HwEnergyUnit        = "J"
	HwEnergyDescription = "Energy consumed by the component"

	HwErrorsName        = "hw.errors"
	HwErrorsUnit        = "{error}"
	HwErrorsDescription = "Number of errors encountered by the component"

	HwHostAmbientTemperatureName        = "hw.host.ambient_temperature"
	HwHostAmbientTemperatureUnit        = "Cel"
	HwHostAmbientTemperatureDescription = "Ambient (external) temperature of the physical host"

	HwHostEnergyName        = "hw.host.energy"
	HwHostEnergyUnit        = "J"
	HwHostEnergyDescription = "Total energy consumed by the entire physical host, in joules"

	HwHostHeatingMarginName        = "hw.host.heating_margin"
	HwHostHeatingMarginUnit        = "Cel"
	HwHostHeatingMarginDescription = "By how many degrees Celsius the temperature of the physical host can be increased, before reaching a warning threshold on one of the internal sensors"

	HwHostPowerName        = "hw.host.power"
	HwHostPowerUnit        = "W"
	HwHostPowerDescription = "Instantaneous power consumed by the entire physical host in Watts (`hw.host.energy` is preferred)"

	HwPowerName        = "hw.power"
	HwPowerUnit        = "W"
	HwPowerDescription = "Instantaneous power consumed by the component"

	HwStatusName        = "hw.status"
	HwStatusUnit        = "1"
	HwStatusDescription = "Operational status: `1` (true) or `0` (false) for each of the possible states"

	K8SCronJobActiveJobsName        = "k8s.cronjob.active_jobs"
	K8SCronJobActiveJobsUnit        = "{job}"
	K8SCronJobActiveJobsDescription = "The number of actively running jobs for a cronjob"

	K8SDaemonSetCurrentScheduledNodesName        = "k8s.daemonset.current_scheduled_nodes"
	K8SDaemonSetCurrentScheduledNodesUnit        = "{node}"
	K8SDaemonSetCurrentScheduledNodesDescription = "Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod"

	K8SDaemonSetDesiredScheduledNodesName        = "k8s.daemonset.desired_scheduled_nodes"
	K8SDaemonSetDesiredScheduledNodesUnit        = "{node}"
	K8SDaemonSetDesiredScheduledNodesDescription = "Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)"

	K8SDaemonSetMisscheduledNodesName        = "k8s.daemonset.misscheduled_nodes"
	K8SDaemonSetMisscheduledNodesUnit        = "{node}"
	K8SDaemonSetMisscheduledNodesDescription = "Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod"

	K8SDaemonSetReadyNodesName        = "k8s.daemonset.ready_nodes"
	K8SDaemonSetReadyNodesUnit        = "{node}"
	K8SDaemonSetReadyNodesDescription = "Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready"

	K8SDeploymentAvailablePodsName        = "k8s.deployment.available_pods"
	K8SDeploymentAvailablePodsUnit        = "{pod}"
	K8SDeploymentAvailablePodsDescription = "Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment"

	K8SDeploymentDesiredPodsName        = "k8s.deployment.desired_pods"
	K8SDeploymentDesiredPodsUnit        = "{pod}"
	K8SDeploymentDesiredPodsDescription = "Number of desired replica pods in this deployment"

	K8SHpaCurrentPodsName        = "k8s.hpa.current_pods"
	K8SHpaCurrentPodsUnit        = "{pod}"
	K8SHpaCurrentPodsDescription = "Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler"

	K8SHpaDesiredPodsName        = "k8s.hpa.desired_pods"
	K8SHpaDesiredPodsUnit        = "{pod}"
	K8SHpaDesiredPodsDescription = "Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler"

	K8SHpaMaxPodsName        = "k8s.hpa.max_pods"
	K8SHpaMaxPodsUnit        = "{pod}"
	K8SHpaMaxPodsDescription = "The upper limit for the number of replica pods to which the autoscaler can scale up"

	K8SHpaMinPodsName        = "k8s.hpa.min_pods"
	K8SHpaMinPodsUnit        = "{pod}"
	K8SHpaMinPodsDescription = "The lower limit for the number of replica pods to which the autoscaler can scale down"

	K8SJobActivePodsName        = "k8s.job.active_pods"
	K8SJobActivePodsUnit        = "{pod}"
	K8SJobActivePodsDescription = "The number of pending and actively running pods for a job"

	K8SJobDesiredSuccessfulPodsName        = "k8s.job.desired_successful_pods"
	K8SJobDesiredSuccessfulPodsUnit        = "{pod}"
	K8SJobDesiredSuccessfulPodsDescription = "The desired number of successfully finished pods the job should be run with"

	K8SJobFailedPodsName        = "k8s.job.failed_pods"
	K8SJobFailedPodsUnit        = "{pod}"
	K8SJobFailedPodsDescription = "The number of pods which reached phase Failed for a job"

	K8SJobMaxParallelPodsName        = "k8s.job.max_parallel_pods"
	K8SJobMaxParallelPodsUnit        = "{pod}"
	K8SJobMaxParallelPodsDescription = "The max desired number of pods the job should run at any given time"

	K8SJobSuccessfulPodsName        = "k8s.job.successful_pods"
	K8SJobSuccessfulPodsUnit        = "{pod}"
	K8SJobSuccessfulPodsDescription = "The number of pods which reached phase Succeeded for a job"

	K8SNamespacePhaseName        = "k8s.namespace.phase"
	K8SNamespacePhaseUnit        = "{namespace}"
	K8SNamespacePhaseDescription = "Describes number of K8s namespaces that are currently in a given phase."

	K8SNodeCPUTimeName        = "k8s.node.cpu.time"
	K8SNodeCPUTimeUnit        = "s"
	K8SNodeCPUTimeDescription = "Total CPU time consumed"

	K8SNodeCPUUsageName        = "k8s.node.cpu.usage"
	K8SNodeCPUUsageUnit        = "{cpu}"
	K8SNodeCPUUsageDescription = "Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"

	K8SNodeMemoryUsageName        = "k8s.node.memory.usage"
	K8SNodeMemoryUsageUnit        = "By"
	K8SNodeMemoryUsageDescription = "Memory usage of the Node"

	K8SNodeNetworkErrorsName        = "k8s.node.network.errors"
	K8SNodeNetworkErrorsUnit        = "{error}"
	K8SNodeNetworkErrorsDescription = "Node network errors"

	K8SNodeNetworkIoName        = "k8s.node.network.io"
	K8SNodeNetworkIoUnit        = "By"
	K8SNodeNetworkIoDescription = "Network bytes for the Node"

	K8SNodeUptimeName        = "k8s.node.uptime"
	K8SNodeUptimeUnit        = "s"
	K8SNodeUptimeDescription = "The time the Node has been running"

	K8SPodCPUTimeName        = "k8s.pod.cpu.time"
	K8SPodCPUTimeUnit        = "s"
	K8SPodCPUTimeDescription = "Total CPU time consumed"

	K8SPodCPUUsageName        = "k8s.pod.cpu.usage"
	K8SPodCPUUsageUnit        = "{cpu}"
	K8SPodCPUUsageDescription = "Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"

	K8SPodMemoryUsageName        = "k8s.pod.memory.usage"
	K8SPodMemoryUsageUnit        = "By"
	K8SPodMemoryUsageDescription = "Memory usage of the Pod"

	K8SPodNetworkErrorsName        = "k8s.pod.network.errors"
	K8SPodNetworkErrorsUnit        = "{error}"
	K8SPodNetworkErrorsDescription = "Pod network errors"

	K8SPodNetworkIoName        = "k8s.pod.network.io"
	K8SPodNetworkIoUnit        = "By"
	K8SPodNetworkIoDescription = "Network bytes for the Pod"

	K8SPodUptimeName        = "k8s.pod.uptime"
	K8SPodUptimeUnit        = "s"
	K8SPodUptimeDescription = "The time the Pod has been running"

	K8SReplicaSetAvailablePodsName        = "k8s.replicaset.available_pods"
	K8SReplicaSetAvailablePodsUnit        = "{pod}"
	K8SReplicaSetAvailablePodsDescription = "Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset"

	K8SReplicaSetDesiredPodsName        = "k8s.replicaset.desired_pods"
	K8SReplicaSetDesiredPodsUnit        = "{pod}"
	K8SReplicaSetDesiredPodsDescription = "Number of desired replica pods in this replicaset"

	K8SReplicationControllerAvailablePodsName        = "k8s.replicationcontroller.available_pods"
	K8SReplicationControllerAvailablePodsUnit        = "{pod}"
	K8SReplicationControllerAvailablePodsDescription = "Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller"

	K8SReplicationControllerDesiredPodsName        = "k8s.replicationcontroller.desired_pods"
	K8SReplicationControllerDesiredPodsUnit        = "{pod}"
	K8SReplicationControllerDesiredPodsDescription = "Number of desired replica pods in this replication controller"

	K8SStatefulSetCurrentPodsName        = "k8s.statefulset.current_pods"
	K8SStatefulSetCurrentPodsUnit        = "{pod}"
	K8SStatefulSetCurrentPodsDescription = "The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision"

	K8SStatefulSetDesiredPodsName        = "k8s.statefulset.desired_pods"
	K8SStatefulSetDesiredPodsUnit        = "{pod}"
	K8SStatefulSetDesiredPodsDescription = "Number of desired replica pods in this statefulset"

	K8SStatefulSetReadyPodsName        = "k8s.statefulset.ready_pods"
	K8SStatefulSetReadyPodsUnit        = "{pod}"
	K8SStatefulSetReadyPodsDescription = "The number of replica pods created for this statefulset with a Ready Condition"

	K8SStatefulSetUpdatedPodsName        = "k8s.statefulset.updated_pods"
	K8SStatefulSetUpdatedPodsUnit        = "{pod}"
	K8SStatefulSetUpdatedPodsDescription = "Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision"

	KestrelActiveConnectionsName        = "kestrel.active_connections"
	KestrelActiveConnectionsUnit        = "{connection}"
	KestrelActiveConnectionsDescription = "Number of connections that are currently active on the server."

	KestrelActiveTLSHandshakesName        = "kestrel.active_tls_handshakes"
	KestrelActiveTLSHandshakesUnit        = "{handshake}"
	KestrelActiveTLSHandshakesDescription = "Number of TLS handshakes that are currently in progress on the server."

	KestrelConnectionDurationName        = "kestrel.connection.duration"
	KestrelConnectionDurationUnit        = "s"
	KestrelConnectionDurationDescription = "The duration of connections on the server."

	KestrelQueuedConnectionsName        = "kestrel.queued_connections"
	KestrelQueuedConnectionsUnit        = "{connection}"
	KestrelQueuedConnectionsDescription = "Number of connections that are currently queued and are waiting to start."

	KestrelQueuedRequestsName        = "kestrel.queued_requests"
	KestrelQueuedRequestsUnit        = "{request}"
	KestrelQueuedRequestsDescription = "Number of HTTP requests on multiplexed connections (HTTP/2 and HTTP/3) that are currently queued and are waiting to start."

	KestrelRejectedConnectionsName        = "kestrel.rejected_connections"
	KestrelRejectedConnectionsUnit        = "{connection}"
	KestrelRejectedConnectionsDescription = "Number of connections rejected by the server."

	KestrelTLSHandshakeDurationName        = "kestrel.tls_handshake.duration"
	KestrelTLSHandshakeDurationUnit        = "s"
	KestrelTLSHandshakeDurationDescription = "The duration of TLS handshakes on the server."

	KestrelUpgradedConnectionsName        = "kestrel.upgraded_connections"
	KestrelUpgradedConnectionsUnit        = "{connection}"
	KestrelUpgradedConnectionsDescription = "Number of connections that are currently upgraded (WebSockets). ."

	MessagingClientConsumedMessagesName        = "messaging.client.consumed.messages"
	MessagingClientConsumedMessagesUnit        = "{message}"
	MessagingClientConsumedMessagesDescription = "Number of messages that were delivered to the application."

	MessagingClientOperationDurationName        = "messaging.client.operation.duration"
	MessagingClientOperationDurationUnit        = "s"
	MessagingClientOperationDurationDescription = "Duration of messaging operation initiated by a producer or consumer client."

	MessagingClientSentMessagesName        = "messaging.client.sent.messages"
	MessagingClientSentMessagesUnit        = "{message}"
	MessagingClientSentMessagesDescription = "Number of messages producer attempted to send to the broker."

	MessagingProcessDurationName        = "messaging.process.duration"
	MessagingProcessDurationUnit        = "s"
	MessagingProcessDurationDescription = "Duration of processing operation."

	OTelSDKExporterSpanExportedCountName        = "otel.sdk.exporter.span.exported.count"
	OTelSDKExporterSpanExportedCountUnit        = "{span}"
	OTelSDKExporterSpanExportedCountDescription = "The number of spans for which the export has finished, either successful or failed"

	OTelSDKExporterSpanInflightCountName        = "otel.sdk.exporter.span.inflight.count"
	OTelSDKExporterSpanInflightCountUnit        = "{span}"
	OTelSDKExporterSpanInflightCountDescription = "The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)"

	OTelSDKProcessorSpanProcessedCountName        = "otel.sdk.processor.span.processed.count"
	OTelSDKProcessorSpanProcessedCountUnit        = "{span}"
	OTelSDKProcessorSpanProcessedCountDescription = "The number of spans for which the processing has finished, either successful or failed"

	OTelSDKProcessorSpanQueueCapacityName        = "otel.sdk.processor.span.queue.capacity"
	OTelSDKProcessorSpanQueueCapacityUnit        = "{span}"
	OTelSDKProcessorSpanQueueCapacityDescription = "The maximum number of spans the queue of a given instance of an SDK span processor can hold"

	OTelSDKProcessorSpanQueueSizeName        = "otel.sdk.processor.span.queue.size"
	OTelSDKProcessorSpanQueueSizeUnit        = "{span}"
	OTelSDKProcessorSpanQueueSizeDescription = "The number of spans in the queue of a given instance of an SDK span processor"

	OTelSDKSpanEndedCountName        = "otel.sdk.span.ended.count"
	OTelSDKSpanEndedCountUnit        = "{span}"
	OTelSDKSpanEndedCountDescription = "The number of created spans for which the end operation was called"

	OTelSDKSpanLiveCountName        = "otel.sdk.span.live.count"
	OTelSDKSpanLiveCountUnit        = "{span}"
	OTelSDKSpanLiveCountDescription = "The number of created spans for which the end operation has not been called yet"

	ProcessContextSwitchesName        = "process.context_switches"
	ProcessContextSwitchesUnit        = "{context_switch}"
	ProcessContextSwitchesDescription = "Number of times the process has been context switched."

	ProcessCPUTimeName        = "process.cpu.time"
	ProcessCPUTimeUnit        = "s"
	ProcessCPUTimeDescription = "Total CPU seconds broken down by different states."

	ProcessCPUUtilizationName        = "process.cpu.utilization"
	ProcessCPUUtilizationUnit        = "1"
	ProcessCPUUtilizationDescription = "Difference in process.cpu.time since the last measurement, divided by the elapsed time and number of CPUs available to the process."

	ProcessDiskIoName        = "process.disk.io"
	ProcessDiskIoUnit        = "By"
	ProcessDiskIoDescription = "Disk bytes transferred."

	ProcessMemoryUsageName        = "process.memory.usage"
	ProcessMemoryUsageUnit        = "By"
	ProcessMemoryUsageDescription = "The amount of physical memory in use."

	ProcessMemoryVirtualName        = "process.memory.virtual"
	ProcessMemoryVirtualUnit        = "By"
	ProcessMemoryVirtualDescription = "The amount of committed virtual memory."

	ProcessNetworkIoName        = "process.network.io"
	ProcessNetworkIoUnit        = "By"
	ProcessNetworkIoDescription = "Network bytes transferred."

	ProcessOpenFileDescriptorCountName        = "process.open_file_descriptor.count"
	ProcessOpenFileDescriptorCountUnit        = "{file_descriptor}"
	ProcessOpenFileDescriptorCountDescription = "Number of file descriptors in use by the process."

	ProcessPagingFaultsName        = "process.paging.faults"
	ProcessPagingFaultsUnit        = "{fault}"
	ProcessPagingFaultsDescription = "Number of page faults the process has made."

	ProcessThreadCountName        = "process.thread.count"
	ProcessThreadCountUnit        = "{thread}"
	ProcessThreadCountDescription = "Process threads count."

	ProcessUptimeName        = "process.uptime"
	ProcessUptimeUnit        = "s"
	ProcessUptimeDescription = "The time the process has been running."

	RPCClientDurationName        = "rpc.client.duration"
	RPCClientDurationUnit        = "ms"
	RPCClientDurationDescription = "Measures the duration of outbound RPC."

	RPCClientRequestSizeName        = "rpc.client.request.size"
	RPCClientRequestSizeUnit        = "By"
	RPCClientRequestSizeDescription = "Measures the size of RPC request messages (uncompressed)."

	RPCClientRequestsPerRPCName        = "rpc.client.requests_per_rpc"
	RPCClientRequestsPerRPCUnit        = "{count}"
	RPCClientRequestsPerRPCDescription = "Measures the number of messages received per RPC."

	RPCClientResponseSizeName        = "rpc.client.response.size"
	RPCClientResponseSizeUnit        = "By"
	RPCClientResponseSizeDescription = "Measures the size of RPC response messages (uncompressed)."

	RPCClientResponsesPerRPCName        = "rpc.client.responses_per_rpc"
	RPCClientResponsesPerRPCUnit        = "{count}"
	RPCClientResponsesPerRPCDescription = "Measures the number of messages sent per RPC."

	RPCServerDurationName        = "rpc.server.duration"
	RPCServerDurationUnit        = "ms"
	RPCServerDurationDescription = "Measures the duration of inbound RPC."

	RPCServerRequestSizeName        = "rpc.server.request.size"
	RPCServerRequestSizeUnit        = "By"
	RPCServerRequestSizeDescription = "Measures the size of RPC request messages (uncompressed)."

	RPCServerRequestsPerRPCName        = "rpc.server.requests_per_rpc"
	RPCServerRequestsPerRPCUnit        = "{count}"
	RPCServerRequestsPerRPCDescription = "Measures the number of messages received per RPC."

	RPCServerResponseSizeName        = "rpc.server.response.size"
	RPCServerResponseSizeUnit        = "By"
	RPCServerResponseSizeDescription = "Measures the size of RPC response messages (uncompressed)."

	RPCServerResponsesPerRPCName        = "rpc.server.responses_per_rpc"
	RPCServerResponsesPerRPCUnit        = "{count}"
	RPCServerResponsesPerRPCDescription = "Measures the number of messages sent per RPC."

	SignalrServerActiveConnectionsName        = "signalr.server.active_connections"
	SignalrServerActiveConnectionsUnit        = "{connection}"
	SignalrServerActiveConnectionsDescription = "Number of connections that are currently active on the server."

	SignalrServerConnectionDurationName        = "signalr.server.connection.duration"
	SignalrServerConnectionDurationUnit        = "s"
	SignalrServerConnectionDurationDescription = "The duration of connections on the server."

	SystemCPULogicalCountName        = "system.cpu.logical.count"
	SystemCPULogicalCountUnit        = "{cpu}"
	SystemCPULogicalCountDescription = "Reports the number of logical (virtual) processor cores created by the operating system to manage multitasking"

	SystemCPUPhysicalCountName        = "system.cpu.physical.count"
	SystemCPUPhysicalCountUnit        = "{cpu}"
	SystemCPUPhysicalCountDescription = "Reports the number of actual physical processor cores on the hardware"

	SystemDiskIoName = "system.disk.io"
	SystemDiskIoUnit = "By"

	SystemDiskIoTimeName        = "system.disk.io_time"
	SystemDiskIoTimeUnit        = "s"
	SystemDiskIoTimeDescription = "Time disk spent activated"

	SystemDiskLimitName        = "system.disk.limit"
	SystemDiskLimitUnit        = "By"
	SystemDiskLimitDescription = "The total storage capacity of the disk"

	SystemDiskMergedName = "system.disk.merged"
	SystemDiskMergedUnit = "{operation}"

	SystemDiskOperationTimeName        = "system.disk.operation_time"
	SystemDiskOperationTimeUnit        = "s"
	SystemDiskOperationTimeDescription = "Sum of the time each operation took to complete"

	SystemDiskOperationsName = "system.disk.operations"
	SystemDiskOperationsUnit = "{operation}"

	SystemFilesystemLimitName        = "system.filesystem.limit"
	SystemFilesystemLimitUnit        = "By"
	SystemFilesystemLimitDescription = "The total storage capacity of the filesystem"

	SystemFilesystemUsageName        = "system.filesystem.usage"
	SystemFilesystemUsageUnit        = "By"
	SystemFilesystemUsageDescription = "Reports a filesystem's space usage across different states."

	SystemFilesystemUtilizationName = "system.filesystem.utilization"
	SystemFilesystemUtilizationUnit = "1"

	SystemLinuxMemoryAvailableName        = "system.linux.memory.available"
	SystemLinuxMemoryAvailableUnit        = "By"
	SystemLinuxMemoryAvailableDescription = "An estimate of how much memory is available for starting new applications, without causing swapping"

	SystemLinuxMemorySlabUsageName        = "system.linux.memory.slab.usage"
	SystemLinuxMemorySlabUsageUnit        = "By"
	SystemLinuxMemorySlabUsageDescription = "Reports the memory used by the Linux kernel for managing caches of frequently used objects."

	SystemMemoryLimitName        = "system.memory.limit"
	SystemMemoryLimitUnit        = "By"
	SystemMemoryLimitDescription = "Total memory available in the system."

	SystemMemorySharedName        = "system.memory.shared"
	SystemMemorySharedUnit        = "By"
	SystemMemorySharedDescription = "Shared memory used (mostly by tmpfs)."

	SystemMemoryUsageName        = "system.memory.usage"
	SystemMemoryUsageUnit        = "By"
	SystemMemoryUsageDescription = "Reports memory in use by state."

	SystemMemoryUtilizationName = "system.memory.utilization"
	SystemMemoryUtilizationUnit = "1"

	SystemNetworkConnectionsName = "system.network.connections"
	SystemNetworkConnectionsUnit = "{connection}"

	SystemNetworkDroppedName        = "system.network.dropped"
	SystemNetworkDroppedUnit        = "{packet}"
	SystemNetworkDroppedDescription = "Count of packets that are dropped or discarded even though there was no error"

	SystemNetworkErrorsName        = "system.network.errors"
	SystemNetworkErrorsUnit        = "{error}"
	SystemNetworkErrorsDescription = "Count of network errors detected"

	SystemNetworkIoName = "system.network.io"
	SystemNetworkIoUnit = "By"

	SystemNetworkPacketsName = "system.network.packets"
	SystemNetworkPacketsUnit = "{packet}"

	SystemPagingFaultsName = "system.paging.faults"
	SystemPagingFaultsUnit = "{fault}"

	SystemPagingOperationsName = "system.paging.operations"
	SystemPagingOperationsUnit = "{operation}"

	SystemPagingUsageName        = "system.paging.usage"
	SystemPagingUsageUnit        = "By"
	SystemPagingUsageDescription = "Unix swap or windows pagefile usage"

	SystemPagingUtilizationName = "system.paging.utilization"
	SystemPagingUtilizationUnit = "1"

	SystemProcessCountName        = "system.process.count"
	SystemProcessCountUnit        = "{process}"
	SystemProcessCountDescription = "Total number of processes in each state"

	SystemProcessCreatedName        = "system.process.created"
	SystemProcessCreatedUnit        = "{process}"
	SystemProcessCreatedDescription = "Total number of processes created over uptime of the host"

	SystemUptimeName        = "system.uptime"
	SystemUptimeUnit        = "s"
	SystemUptimeDescription = "The time the system has been running"

	VCSChangeCountName        = "vcs.change.count"
	VCSChangeCountUnit        = "{change}"
	VCSChangeCountDescription = "The number of changes (pull requests/merge requests/changelists) in a repository, categorized by their state (e.g. open or merged)"

	VCSChangeDurationName        = "vcs.change.duration"
	VCSChangeDurationUnit        = "s"
	VCSChangeDurationDescription = "The time duration a change (pull request/merge request/changelist) has been in a given state."

	VCSChangeTimeToApprovalName        = "vcs.change.time_to_approval"
	VCSChangeTimeToApprovalUnit        = "s"
	VCSChangeTimeToApprovalDescription = "The amount of time since its creation it took a change (pull request/merge request/changelist) to get the first approval."

	VCSChangeTimeToMergeName        = "vcs.change.time_to_merge"
	VCSChangeTimeToMergeUnit        = "s"
	VCSChangeTimeToMergeDescription = "The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref."

	VCSContributorCountName        = "vcs.contributor.count"
	VCSContributorCountUnit        = "{contributor}"
	VCSContributorCountDescription = "The number of unique contributors to a repository"

	VCSRefCountName        = "vcs.ref.count"
	VCSRefCountUnit        = "{ref}"
	VCSRefCountDescription = "The number of refs of type branch or tag in a repository."

	VCSRefLinesDeltaName        = "vcs.ref.lines_delta"
	VCSRefLinesDeltaUnit        = "{line}"
	VCSRefLinesDeltaDescription = "The number of lines added/removed in a ref (branch) relative to the ref from the `vcs.ref.base.name` attribute."

	VCSRefRevisionsDeltaName        = "vcs.ref.revisions_delta"
	VCSRefRevisionsDeltaUnit        = "{revision}"
	VCSRefRevisionsDeltaDescription = "The number of revisions (commits) a ref (branch) is ahead/behind the branch from the `vcs.ref.base.name` attribute"

	VCSRefTimeName        = "vcs.ref.time"
	VCSRefTimeUnit        = "s"
	VCSRefTimeDescription = "Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`"

	VCSRepositoryCountName        = "vcs.repository.count"
	VCSRepositoryCountUnit        = "{repository}"
	VCSRepositoryCountDescription = "The number of repositories in an organization."
)
