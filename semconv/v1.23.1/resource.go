package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	CloudAccountIDKey = attribute.Key("cloud.account.id")

	CloudAvailabilityZoneKey = attribute.Key("cloud.availability_zone")

	CloudPlatformKey = attribute.Key("cloud.platform")

	CloudProviderKey = attribute.Key("cloud.provider")

	CloudRegionKey = attribute.Key("cloud.region")

	CloudResourceIDKey = attribute.Key("cloud.resource_id")
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

	CloudPlatformGCPBareMetalSolution = CloudPlatformKey.String("gcp_bare_metal_solution")

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

var (
	CloudProviderAlibabaCloud = CloudProviderKey.String("alibaba_cloud")

	CloudProviderAWS = CloudProviderKey.String("aws")

	CloudProviderAzure = CloudProviderKey.String("azure")

	CloudProviderGCP = CloudProviderKey.String("gcp")

	CloudProviderHeroku = CloudProviderKey.String("heroku")

	CloudProviderIbmCloud = CloudProviderKey.String("ibm_cloud")

	CloudProviderTencentCloud = CloudProviderKey.String("tencent_cloud")
)

func CloudAccountID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudAvailabilityZone(val string) attribute.KeyValue {
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

const (
	ContainerCommandKey = attribute.Key("container.command")

	ContainerCommandArgsKey = attribute.Key("container.command_args")

	ContainerCommandLineKey = attribute.Key("container.command_line")

	ContainerIDKey = attribute.Key("container.id")

	ContainerImageIDKey = attribute.Key("container.image.id")

	ContainerImageNameKey = attribute.Key("container.image.name")

	ContainerImageRepoDigestsKey = attribute.Key("container.image.repo_digests")

	ContainerImageTagsKey = attribute.Key("container.image.tags")

	ContainerNameKey = attribute.Key("container.name")

	ContainerRuntimeKey = attribute.Key("container.runtime")
)

func ContainerCommand(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerCommandArgs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerCommandLine(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageRepoDigests(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageTags(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerRuntime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OciManifestDigestKey = attribute.Key("oci.manifest.digest")
)

func OciManifestDigest(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AndroidOSAPILevelKey = attribute.Key("android.os.api_level")
)

func AndroidOSAPILevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	BrowserBrandsKey = attribute.Key("browser.brands")

	BrowserLanguageKey = attribute.Key("browser.language")

	BrowserMobileKey = attribute.Key("browser.mobile")

	BrowserPlatformKey = attribute.Key("browser.platform")
)

func BrowserBrands(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserLanguage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserMobile(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserPlatform(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSECSClusterARNKey = attribute.Key("aws.ecs.cluster.arn")

	AWSECSContainerARNKey = attribute.Key("aws.ecs.container.arn")

	AWSECSLaunchtypeKey = attribute.Key("aws.ecs.launchtype")

	AWSECSTaskARNKey = attribute.Key("aws.ecs.task.arn")

	AWSECSTaskFamilyKey = attribute.Key("aws.ecs.task.family")

	AWSECSTaskRevisionKey = attribute.Key("aws.ecs.task.revision")
)

var (
	AWSECSLaunchtypeEC2 = AWSECSLaunchtypeKey.String("ec2")

	AWSECSLaunchtypeFargate = AWSECSLaunchtypeKey.String("fargate")
)

func AWSECSClusterARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSContainerARN(val string) attribute.KeyValue {
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
	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	AWSLogGroupNamesKey = attribute.Key("aws.log.group.names")

	AWSLogStreamARNsKey = attribute.Key("aws.log.stream.arns")

	AWSLogStreamNamesKey = attribute.Key("aws.log.stream.names")
)

func AWSLogGroupARNs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogGroupNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogStreamARNs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogStreamNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	GCPCloudRunJobExecutionKey = attribute.Key("gcp.cloud_run.job.execution")

	GCPCloudRunJobTaskIndexKey = attribute.Key("gcp.cloud_run.job.task_index")
)

func GCPCloudRunJobExecution(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPCloudRunJobTaskIndex(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	GCPGceInstanceHostnameKey = attribute.Key("gcp.gce.instance.hostname")

	GCPGceInstanceNameKey = attribute.Key("gcp.gce.instance.name")
)

func GCPGceInstanceHostname(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPGceInstanceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HerokuAppIDKey = attribute.Key("heroku.app.id")

	HerokuReleaseCommitKey = attribute.Key("heroku.release.commit")

	HerokuReleaseCreationTimestampKey = attribute.Key("heroku.release.creation_timestamp")
)

func HerokuAppID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HerokuReleaseCommit(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HerokuReleaseCreationTimestamp(val string) attribute.KeyValue {
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

	DeviceManufacturerKey = attribute.Key("device.manufacturer")

	DeviceModelIdentifierKey = attribute.Key("device.model.identifier")

	DeviceModelNameKey = attribute.Key("device.model.name")
)

func DeviceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceManufacturer(val string) attribute.KeyValue {
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

const (
	FaaSInstanceKey = attribute.Key("faas.instance")

	FaaSMaxMemoryKey = attribute.Key("faas.max_memory")

	FaaSNameKey = attribute.Key("faas.name")

	FaaSVersionKey = attribute.Key("faas.version")
)

func FaaSInstance(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSMaxMemory(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HostArchKey = attribute.Key("host.arch")

	HostIDKey = attribute.Key("host.id")

	HostImageIDKey = attribute.Key("host.image.id")

	HostImageNameKey = attribute.Key("host.image.name")

	HostImageVersionKey = attribute.Key("host.image.version")

	HostIPKey = attribute.Key("host.ip")

	HostMacKey = attribute.Key("host.mac")

	HostNameKey = attribute.Key("host.name")

	HostTypeKey = attribute.Key("host.type")
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

func HostImageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostIP(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostMac(val ...string) attribute.KeyValue {
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

const (
	HostCPUCacheL2SizeKey = attribute.Key("host.cpu.cache.l2.size")

	HostCPUFamilyKey = attribute.Key("host.cpu.family")

	HostCPUModelIDKey = attribute.Key("host.cpu.model.id")

	HostCPUModelNameKey = attribute.Key("host.cpu.model.name")

	HostCPUSteppingKey = attribute.Key("host.cpu.stepping")

	HostCPUVendorIDKey = attribute.Key("host.cpu.vendor.id")
)

func HostCPUCacheL2Size(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUFamily(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUModelID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUModelName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUStepping(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUVendorID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SClusterNameKey = attribute.Key("k8s.cluster.name")

	K8SClusterUIDKey = attribute.Key("k8s.cluster.uid")
)

func K8SClusterName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SClusterUID(val string) attribute.KeyValue {
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
	K8SPodNameKey = attribute.Key("k8s.pod.name")

	K8SPodUIDKey = attribute.Key("k8s.pod.uid")
)

func K8SPodName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SPodUID(val string) attribute.KeyValue {
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
	K8SReplicaSetNameKey = attribute.Key("k8s.replicaset.name")

	K8SReplicaSetUIDKey = attribute.Key("k8s.replicaset.uid")
)

func K8SReplicaSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicaSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SDeploymentNameKey = attribute.Key("k8s.deployment.name")

	K8SDeploymentUIDKey = attribute.Key("k8s.deployment.uid")
)

func K8SDeploymentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDeploymentUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SStatefulSetNameKey = attribute.Key("k8s.statefulset.name")

	K8SStatefulSetUIDKey = attribute.Key("k8s.statefulset.uid")
)

func K8SStatefulSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStatefulSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SDaemonSetNameKey = attribute.Key("k8s.daemonset.name")

	K8SDaemonSetUIDKey = attribute.Key("k8s.daemonset.uid")
)

func K8SDaemonSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDaemonSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SJobNameKey = attribute.Key("k8s.job.name")

	K8SJobUIDKey = attribute.Key("k8s.job.uid")
)

func K8SJobName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SJobUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	K8SCronJobNameKey = attribute.Key("k8s.cronjob.name")

	K8SCronJobUIDKey = attribute.Key("k8s.cronjob.uid")
)

func K8SCronJobName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SCronJobUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OSBuildIDKey = attribute.Key("os.build_id")

	OSDescriptionKey = attribute.Key("os.description")

	OSNameKey = attribute.Key("os.name")

	OSTypeKey = attribute.Key("os.type")

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

func OSBuildID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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
	ProcessCommandKey = attribute.Key("process.command")

	ProcessCommandArgsKey = attribute.Key("process.command_args")

	ProcessCommandLineKey = attribute.Key("process.command_line")

	ProcessExecutableNameKey = attribute.Key("process.executable.name")

	ProcessExecutablePathKey = attribute.Key("process.executable.path")

	ProcessOwnerKey = attribute.Key("process.owner")

	ProcessParentPIDKey = attribute.Key("process.parent_pid")

	ProcessPIDKey = attribute.Key("process.pid")
)

func ProcessCommand(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommandArgs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommandLine(val string) attribute.KeyValue {
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

func ProcessOwner(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessParentPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ProcessRuntimeDescriptionKey = attribute.Key("process.runtime.description")

	ProcessRuntimeNameKey = attribute.Key("process.runtime.name")

	ProcessRuntimeVersionKey = attribute.Key("process.runtime.version")
)

func ProcessRuntimeDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ServiceNameKey = attribute.Key("service.name")

	ServiceVersionKey = attribute.Key("service.version")
)

func ServiceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ServiceInstanceIDKey = attribute.Key("service.instance.id")

	ServiceNamespaceKey = attribute.Key("service.namespace")
)

func ServiceInstanceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	TelemetrySDKLanguageKey = attribute.Key("telemetry.sdk.language")

	TelemetrySDKNameKey = attribute.Key("telemetry.sdk.name")

	TelemetrySDKVersionKey = attribute.Key("telemetry.sdk.version")
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

	TelemetrySDKLanguageRust = TelemetrySDKLanguageKey.String("rust")

	TelemetrySDKLanguageSwift = TelemetrySDKLanguageKey.String("swift")

	TelemetrySDKLanguageWebjs = TelemetrySDKLanguageKey.String("webjs")
)

func TelemetrySDKName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetrySDKVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	TelemetryDistroNameKey = attribute.Key("telemetry.distro.name")

	TelemetryDistroVersionKey = attribute.Key("telemetry.distro.version")
)

func TelemetryDistroName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetryDistroVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	WebEngineDescriptionKey = attribute.Key("webengine.description")

	WebEngineNameKey = attribute.Key("webengine.name")

	WebEngineVersionKey = attribute.Key("webengine.version")
)

func WebEngineDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineVersion(val string) attribute.KeyValue {
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
