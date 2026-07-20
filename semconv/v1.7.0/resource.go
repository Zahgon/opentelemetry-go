package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	CloudProviderKey = attribute.Key("cloud.provider")

	CloudAccountIDKey = attribute.Key("cloud.account.id")

	CloudRegionKey = attribute.Key("cloud.region")

	CloudAvailabilityZoneKey = attribute.Key("cloud.availability_zone")

	CloudPlatformKey = attribute.Key("cloud.platform")
)

var (
	CloudProviderAlibabaCloud = CloudProviderKey.String("alibaba_cloud")

	CloudProviderAWS = CloudProviderKey.String("aws")

	CloudProviderAzure = CloudProviderKey.String("azure")

	CloudProviderGCP = CloudProviderKey.String("gcp")
)

var (
	CloudPlatformAlibabaCloudECS = CloudPlatformKey.String("alibaba_cloud_ecs")

	CloudPlatformAlibabaCloudFc = CloudPlatformKey.String("alibaba_cloud_fc")

	CloudPlatformAWSEC2 = CloudPlatformKey.String("aws_ec2")

	CloudPlatformAWSECS = CloudPlatformKey.String("aws_ecs")

	CloudPlatformAWSEKS = CloudPlatformKey.String("aws_eks")

	CloudPlatformAWSLambda = CloudPlatformKey.String("aws_lambda")

	CloudPlatformAWSElasticBeanstalk = CloudPlatformKey.String("aws_elastic_beanstalk")

	CloudPlatformAzureVM = CloudPlatformKey.String("azure_vm")

	CloudPlatformAzureContainerInstances = CloudPlatformKey.String("azure_container_instances")

	CloudPlatformAzureAKS = CloudPlatformKey.String("azure_aks")

	CloudPlatformAzureFunctions = CloudPlatformKey.String("azure_functions")

	CloudPlatformAzureAppService = CloudPlatformKey.String("azure_app_service")

	CloudPlatformGCPComputeEngine = CloudPlatformKey.String("gcp_compute_engine")

	CloudPlatformGCPCloudRun = CloudPlatformKey.String("gcp_cloud_run")

	CloudPlatformGCPKubernetesEngine = CloudPlatformKey.String("gcp_kubernetes_engine")

	CloudPlatformGCPCloudFunctions = CloudPlatformKey.String("gcp_cloud_functions")

	CloudPlatformGCPAppEngine = CloudPlatformKey.String("gcp_app_engine")
)

const (
	AWSECSContainerARNKey = attribute.Key("aws.ecs.container.arn")

	AWSECSClusterARNKey = attribute.Key("aws.ecs.cluster.arn")

	AWSECSLaunchtypeKey = attribute.Key("aws.ecs.launchtype")

	AWSECSTaskARNKey = attribute.Key("aws.ecs.task.arn")

	AWSECSTaskFamilyKey = attribute.Key("aws.ecs.task.family")

	AWSECSTaskRevisionKey = attribute.Key("aws.ecs.task.revision")
)

var (
	AWSECSLaunchtypeEC2 = AWSECSLaunchtypeKey.String("ec2")

	AWSECSLaunchtypeFargate = AWSECSLaunchtypeKey.String("fargate")
)

const (
	AWSEKSClusterARNKey = attribute.Key("aws.eks.cluster.arn")
)

const (
	AWSLogGroupNamesKey = attribute.Key("aws.log.group.names")

	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	AWSLogStreamNamesKey = attribute.Key("aws.log.stream.names")

	AWSLogStreamARNsKey = attribute.Key("aws.log.stream.arns")
)

const (
	ContainerNameKey = attribute.Key("container.name")

	ContainerIDKey = attribute.Key("container.id")

	ContainerRuntimeKey = attribute.Key("container.runtime")

	ContainerImageNameKey = attribute.Key("container.image.name")

	ContainerImageTagKey = attribute.Key("container.image.tag")
)

const (
	DeploymentEnvironmentKey = attribute.Key("deployment.environment")
)

const (
	DeviceIDKey = attribute.Key("device.id")

	DeviceModelIdentifierKey = attribute.Key("device.model.identifier")

	DeviceModelNameKey = attribute.Key("device.model.name")
)

const (
	FaaSNameKey = attribute.Key("faas.name")

	FaaSIDKey = attribute.Key("faas.id")

	FaaSVersionKey = attribute.Key("faas.version")

	FaaSInstanceKey = attribute.Key("faas.instance")

	FaaSMaxMemoryKey = attribute.Key("faas.max_memory")
)

const (
	HostIDKey = attribute.Key("host.id")

	HostNameKey = attribute.Key("host.name")

	HostTypeKey = attribute.Key("host.type")

	HostArchKey = attribute.Key("host.arch")

	HostImageNameKey = attribute.Key("host.image.name")

	HostImageIDKey = attribute.Key("host.image.id")

	HostImageVersionKey = attribute.Key("host.image.version")
)

var (
	HostArchAMD64 = HostArchKey.String("amd64")

	HostArchARM32 = HostArchKey.String("arm32")

	HostArchARM64 = HostArchKey.String("arm64")

	HostArchIA64 = HostArchKey.String("ia64")

	HostArchPPC32 = HostArchKey.String("ppc32")

	HostArchPPC64 = HostArchKey.String("ppc64")

	HostArchX86 = HostArchKey.String("x86")
)

const (
	K8SClusterNameKey = attribute.Key("k8s.cluster.name")
)

const (
	K8SNodeNameKey = attribute.Key("k8s.node.name")

	K8SNodeUIDKey = attribute.Key("k8s.node.uid")
)

const (
	K8SNamespaceNameKey = attribute.Key("k8s.namespace.name")
)

const (
	K8SPodUIDKey = attribute.Key("k8s.pod.uid")

	K8SPodNameKey = attribute.Key("k8s.pod.name")
)

const (
	K8SContainerNameKey = attribute.Key("k8s.container.name")
)

const (
	K8SReplicaSetUIDKey = attribute.Key("k8s.replicaset.uid")

	K8SReplicaSetNameKey = attribute.Key("k8s.replicaset.name")
)

const (
	K8SDeploymentUIDKey = attribute.Key("k8s.deployment.uid")

	K8SDeploymentNameKey = attribute.Key("k8s.deployment.name")
)

const (
	K8SStatefulSetUIDKey = attribute.Key("k8s.statefulset.uid")

	K8SStatefulSetNameKey = attribute.Key("k8s.statefulset.name")
)

const (
	K8SDaemonSetUIDKey = attribute.Key("k8s.daemonset.uid")

	K8SDaemonSetNameKey = attribute.Key("k8s.daemonset.name")
)

const (
	K8SJobUIDKey = attribute.Key("k8s.job.uid")

	K8SJobNameKey = attribute.Key("k8s.job.name")
)

const (
	K8SCronJobUIDKey = attribute.Key("k8s.cronjob.uid")

	K8SCronJobNameKey = attribute.Key("k8s.cronjob.name")
)

const (
	OSTypeKey = attribute.Key("os.type")

	OSDescriptionKey = attribute.Key("os.description")

	OSNameKey = attribute.Key("os.name")

	OSVersionKey = attribute.Key("os.version")
)

var (
	OSTypeWindows = OSTypeKey.String("windows")

	OSTypeLinux = OSTypeKey.String("linux")

	OSTypeDarwin = OSTypeKey.String("darwin")

	OSTypeFreeBSD = OSTypeKey.String("freebsd")

	OSTypeNetBSD = OSTypeKey.String("netbsd")

	OSTypeOpenBSD = OSTypeKey.String("openbsd")

	OSTypeDragonflyBSD = OSTypeKey.String("dragonflybsd")

	OSTypeHPUX = OSTypeKey.String("hpux")

	OSTypeAIX = OSTypeKey.String("aix")

	OSTypeSolaris = OSTypeKey.String("solaris")

	OSTypeZOS = OSTypeKey.String("z_os")
)

const (
	ProcessPIDKey = attribute.Key("process.pid")

	ProcessExecutableNameKey = attribute.Key("process.executable.name")

	ProcessExecutablePathKey = attribute.Key("process.executable.path")

	ProcessCommandKey = attribute.Key("process.command")

	ProcessCommandLineKey = attribute.Key("process.command_line")

	ProcessCommandArgsKey = attribute.Key("process.command_args")

	ProcessOwnerKey = attribute.Key("process.owner")
)

const (
	ProcessRuntimeNameKey = attribute.Key("process.runtime.name")

	ProcessRuntimeVersionKey = attribute.Key("process.runtime.version")

	ProcessRuntimeDescriptionKey = attribute.Key("process.runtime.description")
)

const (
	ServiceNameKey = attribute.Key("service.name")

	ServiceNamespaceKey = attribute.Key("service.namespace")

	ServiceInstanceIDKey = attribute.Key("service.instance.id")

	ServiceVersionKey = attribute.Key("service.version")
)

const (
	TelemetrySDKNameKey = attribute.Key("telemetry.sdk.name")

	TelemetrySDKLanguageKey = attribute.Key("telemetry.sdk.language")

	TelemetrySDKVersionKey = attribute.Key("telemetry.sdk.version")

	TelemetryAutoVersionKey = attribute.Key("telemetry.auto.version")
)

var (
	TelemetrySDKLanguageCPP = TelemetrySDKLanguageKey.String("cpp")

	TelemetrySDKLanguageDotnet = TelemetrySDKLanguageKey.String("dotnet")

	TelemetrySDKLanguageErlang = TelemetrySDKLanguageKey.String("erlang")

	TelemetrySDKLanguageGo = TelemetrySDKLanguageKey.String("go")

	TelemetrySDKLanguageJava = TelemetrySDKLanguageKey.String("java")

	TelemetrySDKLanguageNodejs = TelemetrySDKLanguageKey.String("nodejs")

	TelemetrySDKLanguagePHP = TelemetrySDKLanguageKey.String("php")

	TelemetrySDKLanguagePython = TelemetrySDKLanguageKey.String("python")

	TelemetrySDKLanguageRuby = TelemetrySDKLanguageKey.String("ruby")

	TelemetrySDKLanguageWebjs = TelemetrySDKLanguageKey.String("webjs")
)

const (
	WebEngineNameKey = attribute.Key("webengine.name")

	WebEngineVersionKey = attribute.Key("webengine.version")

	WebEngineDescriptionKey = attribute.Key("webengine.description")
)
