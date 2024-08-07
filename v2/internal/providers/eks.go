package providers

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/url"
	"os"
	"strings"
)

const EnvVarSkipEKSHostnameCheck = "STRATUS_SKIP_EKS_HOSTNAME_CHECK"

type EKSProvider struct {
	awsProvider         *AWSProvider
	k8sProvider         *K8sProvider
	UniqueCorrelationId uuid.UUID // unique value injected in the user-agent, to differentiate Stratus Red Team executions
}

func NewEKSProvider(uuid uuid.UUID) *EKSProvider {
	return &EKSProvider{
		awsProvider:         NewAWSProvider(uuid),
		k8sProvider:         NewK8sProvider(uuid),
		UniqueCorrelationId: uuid,
	}
}

func (m *EKSProvider) GetAWSConnection() aws.Config {
	return m.awsProvider.GetConnection()
}

func (m *EKSProvider) GetK8sClient() *kubernetes.Clientset {
	return m.k8sProvider.GetClient()
}

func (m *EKSProvider) IsAuthenticatedAgainstEKS() bool {
	// Check if we're properly authenticated against AWS
	if !m.awsProvider.IsAuthenticatedAgainstAWS() {
		return false
	}

	// Check if our current K8s context is a valid EKS cluster
	if os.Getenv(EnvVarSkipEKSHostnameCheck) == "1" {
		return true
	}
	apiServerUrl := m.k8sProvider.GetRestConfig().Host
	parsedAPIServerUrl, err := url.Parse(apiServerUrl)
	if err != nil {
		log.Fatalf("unable to parse API server URL %s: %v", apiServerUrl, err)
	}

	return strings.HasSuffix(parsedAPIServerUrl.Host, ".eks.amazonaws.com")
}

// taken from https://github.com/DataDog/managed-kubernetes-auditing-toolkit/blob/main/internal/utils/kubernetes.go#L59C1-L74C2
func (m *EKSProvider) GetEKSClusterName() string {
	// Most of (if not all) the time, the KubeConfig file generated by "aws eks update-kubeconfig" will have an
	// ExecProvider section that runs "aws eks get-token <...> --cluster-name foo"
	// We parse it and extract the cluster name from there

	execProvider := m.k8sProvider.GetRestConfig().ExecProvider
	if execProvider == nil || execProvider.Command != "aws" {
		return ""
	}
	for i, arg := range execProvider.Args {
		if arg == "--cluster-name" && i+1 < len(execProvider.Args) {
			return execProvider.Args[i+1]
		}
	}
	return ""
}
