package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	BrowserBrandsKey = attribute.Key("browser.brands")

	BrowserPlatformKey = attribute.Key("browser.platform")

	BrowserMobileKey = attribute.Key("browser.mobile")

	BrowserLanguageKey = attribute.Key("browser.language")
)

func BrowserBrands(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserPlatform(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserMobile(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserLanguage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CloudProviderKey = attribute.Key("cloud.provider")

	CloudAccountIDKey = attribute.Key("cloud.account.id")

	CloudRegionKey = attribute.Key("cloud.region")

	CloudResourceIDKey = attribute.Key("cloud.resource_id")

	CloudAvailabilityZoneKey = attribute.Key("cloud.availability_zone")

	CloudPlatformKey = attribute.Key("cloud.platform")
)

var (
	CloudProviderAlibabaCloud = CloudProviderKey.String("alibaba_cloud")

	CloudProviderAWS = CloudProviderKey.String("aws")

	CloudProviderAzure = CloudProviderKey.String("azure")

	CloudProviderGCP = CloudProviderKey.String("gcp")

	CloudProviderHeroku = CloudProviderKey.String("heroku")

	CloudProviderIbmCloud = CloudProviderKey.String("ibm_cloud")

	CloudProviderTencentCloud = CloudProviderKey.String("tencent_cloud")
)

var (
	CloudPlatformAlibabaCloudECS = CloudPlatformKey.String("alibaba_cloud_ecs")

	CloudPlatformAlibabaCloudFc = CloudPlatformKey.String("alibaba_cloud_fc")

	CloudPlatformAlibabaCloudOpenshift = CloudPlatformKey.String("alibaba_cloud_openshift")

	CloudPlatformAWSEC2 = CloudPlatformKey.String("aws_ec2")

	CloudPlatformAWSECS = CloudPlatformKey.String("aws_ecs")

	CloudPlatformAWSEKS = CloudPlatformKey.String("aws_eks")

	CloudPlatformAWSLambda = CloudPlatformKey.String("aws_lambda")

	CloudPlatformAWSElasticBeanstalk = CloudPlatformKey.String("aws_elastic_beanstalk")

	CloudPlatformAWSAppRunner = CloudPlatformKey.String("aws_app_runner")

	CloudPlatformAWSOpenshift = CloudPlatformKey.String("aws_openshift")

	CloudPlatformAzureVM = CloudPlatformKey.String("azure_vm")

	CloudPlatformAzureContainerInstances = CloudPlatformKey.String("azure_container_instances")

	CloudPlatformAzureAKS = CloudPlatformKey.String("azure_aks")

	CloudPlatformAzureFunctions = CloudPlatformKey.String("azure_functions")

	CloudPlatformAzureAppService = CloudPlatformKey.String("azure_app_service")

	CloudPlatformAzureOpenshift = CloudPlatformKey.String("azure_openshift")

	CloudPlatformGCPComputeEngine = CloudPlatformKey.String("gcp_compute_engine")

	CloudPlatformGCPCloudRun = CloudPlatformKey.String("gcp_cloud_run")

	CloudPlatformGCPKubernetesEngine = CloudPlatformKey.String("gcp_kubernetes_engine")

	CloudPlatformGCPCloudFunctions = CloudPlatformKey.String("gcp_cloud_functions")

	CloudPlatformGCPAppEngine = CloudPlatformKey.String("gcp_app_engine")

	CloudPlatformGCPOpenshift = CloudPlatformKey.String("gcp_openshift")

	CloudPlatformIbmCloudOpenshift = CloudPlatformKey.String("ibm_cloud_openshift")

	CloudPlatformTencentCloudCvm = CloudPlatformKey.String("tencent_cloud_cvm")

	CloudPlatformTencentCloudEKS = CloudPlatformKey.String("tencent_cloud_eks")

	CloudPlatformTencentCloudScf = CloudPlatformKey.String("tencent_cloud_scf")
)

func CloudAccountID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudRegion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudResourceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudAvailabilityZone(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func AWSECSContainerARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSClusterARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskFamily(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSEKSClusterARNKey = attribute.Key("aws.eks.cluster.arn")
)

func AWSEKSClusterARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSLogGroupNamesKey = attribute.Key("aws.log.group.names")

	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	AWSLogStreamNamesKey = attribute.Key("aws.log.stream.names")

	AWSLogStreamARNsKey = attribute.Key("aws.log.stream.arns")
)

func AWSLogGroupNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogGroupARNs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogStreamNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogStreamARNs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HerokuReleaseCreationTimestampKey = attribute.Key("heroku.release.creation_timestamp")

	HerokuReleaseCommitKey = attribute.Key("heroku.release.commit")

	HerokuAppIDKey = attribute.Key("heroku.app.id")
)

func HerokuReleaseCreationTimestamp(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HerokuReleaseCommit(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HerokuAppID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ContainerNameKey = attribute.Key("container.name")

	ContainerIDKey = attribute.Key("container.id")

	ContainerRuntimeKey = attribute.Key("container.runtime")

	ContainerImageNameKey = attribute.Key("container.image.name")

	ContainerImageTagKey = attribute.Key("container.image.tag")
)

func ContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerRuntime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageTag(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DeploymentEnvironmentKey = attribute.Key("deployment.environment")
)

func DeploymentEnvironment(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DeviceIDKey = attribute.Key("device.id")

	DeviceModelIdentifierKey = attribute.Key("device.model.identifier")

	DeviceModelNameKey = attribute.Key("device.model.name")

	DeviceManufacturerKey = attribute.Key("device.manufacturer")
)

func DeviceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceModelIdentifier(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceModelName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceManufacturer(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	FaaSNameKey = attribute.Key("faas.name")

	FaaSVersionKey = attribute.Key("faas.version")

	FaaSInstanceKey = attribute.Key("faas.instance")

	FaaSMaxMemoryKey = attribute.Key("faas.max_memory")
)

func FaaSName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSInstance(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSMaxMemory(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

	HostArchS390x = HostArchKey.String("s390x")

	HostArchX86 = HostArchKey.String("x86")
)

func HostID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SClusterNameKey = attribute.Key("k8s.cluster.name")
)

func K8SClusterName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SNodeNameKey = attribute.Key("k8s.node.name")

	K8SNodeUIDKey = attribute.Key("k8s.node.uid")
)

func K8SNodeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNodeUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SNamespaceNameKey = attribute.Key("k8s.namespace.name")
)

func K8SNamespaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SPodUIDKey = attribute.Key("k8s.pod.uid")

	K8SPodNameKey = attribute.Key("k8s.pod.name")
)

func K8SPodUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SPodName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SContainerNameKey = attribute.Key("k8s.container.name")

	K8SContainerRestartCountKey = attribute.Key("k8s.container.restart_count")
)

func K8SContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SContainerRestartCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SReplicaSetUIDKey = attribute.Key("k8s.replicaset.uid")

	K8SReplicaSetNameKey = attribute.Key("k8s.replicaset.name")
)

func K8SReplicaSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicaSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SDeploymentUIDKey = attribute.Key("k8s.deployment.uid")

	K8SDeploymentNameKey = attribute.Key("k8s.deployment.name")
)

func K8SDeploymentUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDeploymentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SStatefulSetUIDKey = attribute.Key("k8s.statefulset.uid")

	K8SStatefulSetNameKey = attribute.Key("k8s.statefulset.name")
)

func K8SStatefulSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStatefulSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SDaemonSetUIDKey = attribute.Key("k8s.daemonset.uid")

	K8SDaemonSetNameKey = attribute.Key("k8s.daemonset.name")
)

func K8SDaemonSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDaemonSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SJobUIDKey = attribute.Key("k8s.job.uid")

	K8SJobNameKey = attribute.Key("k8s.job.name")
)

func K8SJobUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SJobName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SCronJobUIDKey = attribute.Key("k8s.cronjob.uid")

	K8SCronJobNameKey = attribute.Key("k8s.cronjob.name")
)

func K8SCronJobUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SCronJobName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func OSDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OSName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OSVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ProcessPIDKey = attribute.Key("process.pid")

	ProcessParentPIDKey = attribute.Key("process.parent_pid")

	ProcessExecutableNameKey = attribute.Key("process.executable.name")

	ProcessExecutablePathKey = attribute.Key("process.executable.path")

	ProcessCommandKey = attribute.Key("process.command")

	ProcessCommandLineKey = attribute.Key("process.command_line")

	ProcessCommandArgsKey = attribute.Key("process.command_args")

	ProcessOwnerKey = attribute.Key("process.owner")
)

func ProcessPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessParentPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutableName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutablePath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommand(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommandLine(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommandArgs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessOwner(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ProcessRuntimeNameKey = attribute.Key("process.runtime.name")

	ProcessRuntimeVersionKey = attribute.Key("process.runtime.version")

	ProcessRuntimeDescriptionKey = attribute.Key("process.runtime.description")
)

func ProcessRuntimeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ServiceNameKey = attribute.Key("service.name")

	ServiceNamespaceKey = attribute.Key("service.namespace")

	ServiceInstanceIDKey = attribute.Key("service.instance.id")

	ServiceVersionKey = attribute.Key("service.version")
)

func ServiceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceInstanceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

	TelemetrySDKLanguageSwift = TelemetrySDKLanguageKey.String("swift")
)

func TelemetrySDKName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetrySDKVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetryAutoVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	WebEngineNameKey = attribute.Key("webengine.name")

	WebEngineVersionKey = attribute.Key("webengine.version")

	WebEngineDescriptionKey = attribute.Key("webengine.description")
)

func WebEngineName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OTelScopeNameKey = attribute.Key("otel.scope.name")

	OTelScopeVersionKey = attribute.Key("otel.scope.version")
)

func OTelScopeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelScopeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OTelLibraryNameKey = attribute.Key("otel.library.name")

	OTelLibraryVersionKey = attribute.Key("otel.library.version")
)

func OTelLibraryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelLibraryVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
