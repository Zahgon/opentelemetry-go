package semconv

const (
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

	DBClientConnectionsCreateTimeName        = "db.client.connections.create_time"
	DBClientConnectionsCreateTimeUnit        = "ms"
	DBClientConnectionsCreateTimeDescription = "Deprecated, use `db.client.connection.create_time` instead. Note: the unit also changed from `ms` to `s`."

	DBClientConnectionsIdleMaxName        = "db.client.connections.idle.max"
	DBClientConnectionsIdleMaxUnit        = "{connection}"
	DBClientConnectionsIdleMaxDescription = "Deprecated, use `db.client.connection.idle.max` instead."

	DBClientConnectionsIdleMinName        = "db.client.connections.idle.min"
	DBClientConnectionsIdleMinUnit        = "{connection}"
	DBClientConnectionsIdleMinDescription = "Deprecated, use `db.client.connection.idle.min` instead."

	DBClientConnectionsMaxName        = "db.client.connections.max"
	DBClientConnectionsMaxUnit        = "{connection}"
	DBClientConnectionsMaxDescription = "Deprecated, use `db.client.connection.max` instead."

	DBClientConnectionsPendingRequestsName        = "db.client.connections.pending_requests"
	DBClientConnectionsPendingRequestsUnit        = "{request}"
	DBClientConnectionsPendingRequestsDescription = "Deprecated, use `db.client.connection.pending_requests` instead."

	DBClientConnectionsTimeoutsName        = "db.client.connections.timeouts"
	DBClientConnectionsTimeoutsUnit        = "{timeout}"
	DBClientConnectionsTimeoutsDescription = "Deprecated, use `db.client.connection.timeouts` instead."

	DBClientConnectionsUsageName        = "db.client.connections.usage"
	DBClientConnectionsUsageUnit        = "{connection}"
	DBClientConnectionsUsageDescription = "Deprecated, use `db.client.connection.count` instead."

	DBClientConnectionsUseTimeName        = "db.client.connections.use_time"
	DBClientConnectionsUseTimeUnit        = "ms"
	DBClientConnectionsUseTimeDescription = "Deprecated, use `db.client.connection.use_time` instead. Note: the unit also changed from `ms` to `s`."

	DBClientConnectionsWaitTimeName        = "db.client.connections.wait_time"
	DBClientConnectionsWaitTimeUnit        = "ms"
	DBClientConnectionsWaitTimeDescription = "Deprecated, use `db.client.connection.wait_time` instead. Note: the unit also changed from `ms` to `s`."

	DBClientOperationDurationName        = "db.client.operation.duration"
	DBClientOperationDurationUnit        = "s"
	DBClientOperationDurationDescription = "Duration of database client operations."

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

	HwPowerName        = "hw.power"
	HwPowerUnit        = "W"
	HwPowerDescription = "Instantaneous power consumed by the component"

	HwStatusName        = "hw.status"
	HwStatusUnit        = "1"
	HwStatusDescription = "Operational status: `1` (true) or `0` (false) for each of the possible states"

	K8SNodeCPUTimeName        = "k8s.node.cpu.time"
	K8SNodeCPUTimeUnit        = "s"
	K8SNodeCPUTimeDescription = "Total CPU time consumed"

	K8SNodeCPUUsageName        = "k8s.node.cpu.usage"
	K8SNodeCPUUsageUnit        = "{cpu}"
	K8SNodeCPUUsageDescription = "Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"

	K8SNodeMemoryUsageName        = "k8s.node.memory.usage"
	K8SNodeMemoryUsageUnit        = "By"
	K8SNodeMemoryUsageDescription = "Memory usage of the Node"

	K8SPodCPUTimeName        = "k8s.pod.cpu.time"
	K8SPodCPUTimeUnit        = "s"
	K8SPodCPUTimeDescription = "Total CPU time consumed"

	K8SPodCPUUsageName        = "k8s.pod.cpu.usage"
	K8SPodCPUUsageUnit        = "{cpu}"
	K8SPodCPUUsageDescription = "Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs"

	K8SPodMemoryUsageName        = "k8s.pod.memory.usage"
	K8SPodMemoryUsageUnit        = "By"
	K8SPodMemoryUsageDescription = "Memory usage of the Pod"

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

	MessagingClientPublishedMessagesName        = "messaging.client.published.messages"
	MessagingClientPublishedMessagesUnit        = "{message}"
	MessagingClientPublishedMessagesDescription = "Deprecated. Use `messaging.client.sent.messages` instead."

	MessagingClientSentMessagesName        = "messaging.client.sent.messages"
	MessagingClientSentMessagesUnit        = "{message}"
	MessagingClientSentMessagesDescription = "Number of messages producer attempted to send to the broker."

	MessagingProcessDurationName        = "messaging.process.duration"
	MessagingProcessDurationUnit        = "s"
	MessagingProcessDurationDescription = "Duration of processing operation."

	MessagingProcessMessagesName        = "messaging.process.messages"
	MessagingProcessMessagesUnit        = "{message}"
	MessagingProcessMessagesDescription = "Deprecated. Use `messaging.client.consumed.messages` instead."

	MessagingPublishDurationName        = "messaging.publish.duration"
	MessagingPublishDurationUnit        = "s"
	MessagingPublishDurationDescription = "Deprecated. Use `messaging.client.operation.duration` instead."

	MessagingPublishMessagesName        = "messaging.publish.messages"
	MessagingPublishMessagesUnit        = "{message}"
	MessagingPublishMessagesDescription = "Deprecated. Use `messaging.client.produced.messages` instead."

	MessagingReceiveDurationName        = "messaging.receive.duration"
	MessagingReceiveDurationUnit        = "s"
	MessagingReceiveDurationDescription = "Deprecated. Use `messaging.client.operation.duration` instead."

	MessagingReceiveMessagesName        = "messaging.receive.messages"
	MessagingReceiveMessagesUnit        = "{message}"
	MessagingReceiveMessagesDescription = "Deprecated. Use `messaging.client.consumed.messages` instead."

	ProcessContextSwitchesName        = "process.context_switches"
	ProcessContextSwitchesUnit        = "{count}"
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
	ProcessOpenFileDescriptorCountUnit        = "{count}"
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

	SystemCPUFrequencyName        = "system.cpu.frequency"
	SystemCPUFrequencyUnit        = "{Hz}"
	SystemCPUFrequencyDescription = "Reports the current frequency of the CPU in Hz"

	SystemCPULogicalCountName        = "system.cpu.logical.count"
	SystemCPULogicalCountUnit        = "{cpu}"
	SystemCPULogicalCountDescription = "Reports the number of logical (virtual) processor cores created by the operating system to manage multitasking"

	SystemCPUPhysicalCountName        = "system.cpu.physical.count"
	SystemCPUPhysicalCountUnit        = "{cpu}"
	SystemCPUPhysicalCountDescription = "Reports the number of actual physical processor cores on the hardware"

	SystemCPUTimeName        = "system.cpu.time"
	SystemCPUTimeUnit        = "s"
	SystemCPUTimeDescription = "Seconds each logical CPU spent on each mode"

	SystemCPUUtilizationName        = "system.cpu.utilization"
	SystemCPUUtilizationUnit        = "1"
	SystemCPUUtilizationDescription = "Difference in system.cpu.time since the last measurement, divided by the elapsed time and number of logical CPUs"

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
)
