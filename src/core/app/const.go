package app

type CSP string
type ROLE string
type Kind string
type NetworkCni string
type StatusCode int
type Loadbalancer string
type Etcd string
type ServiceType string

const (
	CSP_AWS       CSP = "aws"
	CSP_GCP       CSP = "gcp"
	CSP_AZURE     CSP = "azure"
	CSP_ALIBABA   CSP = "alibaba"
	CSP_TENCENT   CSP = "tencent"
	CSP_OPENSTACK CSP = "openstack"
	CSP_IBM       CSP = "ibm"
	CSP_NCPVPC    CSP = "ncpvpc"
	CSP_NCP       CSP = "ncp"
	CSP_CLOUDIT   CSP = "cloudit"
	CSP_NHNCLOUD  CSP = "nhncloud"

	CONTROL_PLANE ROLE = "control-plane"
	WORKER        ROLE = "worker"

	KIND_STATUS       Kind = "Status"
	KIND_CLUSTER      Kind = "Cluster"
	KIND_CLUSTER_LIST Kind = "ClusterList"
	KIND_NODE         Kind = "Node"
	KIND_NODE_LIST    Kind = "NodeList"

	STATUS_UNKNOWN  = 0
	STATUS_SUCCESS  = 200
	STATUS_NOTFOUND = 404

	NETWORKCNI_KILO    NetworkCni = "kilo"
	NETWORKCNI_CANAL   NetworkCni = "canal"
	NETWORKCNI_FLANNEL NetworkCni = "flannel"
	NETWORKCNI_CALICO  NetworkCni = "calico"

	LB_HAPROXY Loadbalancer = "haproxy"
	LB_NLB     Loadbalancer = "nlb"

	ETCD_LOCAL    Etcd = "local"
	ETCD_EXTERNAL Etcd = "external"

	POD_CIDR       = "10.244.0.0/16"
	SERVICE_CIDR   = "10.96.0.0/12"
	SERVICE_DOMAIN = "cluster.local"

	LABEL_KEY_CSP    = "topology.cloud-barista.github.io/csp"
	LABEL_KEY_REGION = "topology.kubernetes.io/region"
	LABEL_KEY_ZONE   = "topology.kubernetes.io/zone"

	LABEL_KEY_CLUSTER = "kubernetes.io/cluster"

	MCIS_LABEL       = "mcks"
	MCIS_SYSTEMLABEL = "Managed by MCKS"

	ST_MULTI  ServiceType = "multi"
	ST_SINGLE ServiceType = "single"
)

type Status struct {
	Kind    Kind   `json:"kind"`
	Code    int    `json:"code"`
	Message string `json:"message" example:"Any message"`
}

type ClusterReq struct {
	Name         string           `json:"name" example:"cluster-01"`
	ControlPlane []*NodeSetReq    `json:"controlPlane"`
	Worker       []*NodeSetReq    `json:"worker"`
	ServiceType  ServiceType      `json:"serviceType" enums:"multi,single" default:"multi"`
	Config       ClusterConfigReq `json:"config"`
	Label        string           `json:"label"`
	Description  string           `json:"description"`
}

type NodeReq struct {
	ControlPlane []*NodeSetReq `json:"controlPlane"`
	Worker       []*NodeSetReq `json:"worker"`
}

type NodeSetReq struct {
	Connection string `json:"connection" example:"config-aws-ap-northeast-2"`
	Count      int    `json:"count" example:"3"`
	Spec       string `json:"spec" example:"t2.medium"`
	RootDisk   struct {
		Type string `json:"type" example:"default"`
		Size string `json:"size" example:"default"`
	} `json:"rootDisk"`
	Role string `json:"role"`
}

type ClusterConfigReq struct {
	InstallMonAgent string                     `json:"installMonAgent" example:"no"`
	Kubernetes      ClusterConfigKubernetesReq `json:"kubernetes"`
}
type ClusterConfigKubernetesReq struct {
	Version          string     `json:"version" example:"1.23.13"`
	NetworkCni       NetworkCni `json:"networkCni" example:"kilo" enums:"canal,kilo,flannel,calico" default1:"kilo"`
	PodCidr          string     `json:"podCidr" example:"10.244.0.0/16"`
	ServiceCidr      string     `json:"serviceCidr" example:"10.96.0.0/12"`
	ServiceDnsDomain string     `json:"serviceDnsDomain" example:"cluster.local"`
	StorageClass     struct {
		Nfs ClusterStorageClassNfsReq `json:"nfs"`
	} `json:"storageclass"`
	Loadbalancer Loadbalancer `json:"loadbalancer" example:"haproxy" enums:"haproxy,nlb"`
	Etcd         Etcd         `json:"etcd" example:"local" enums:"local,external"`
}

type ClusterStorageClassNfsReq struct {
	Server string `json:"server" example:"163.154.154.89"`
	Path   string `json:"path" example:"/nfs/data"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
