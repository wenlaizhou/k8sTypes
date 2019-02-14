package k8sTypes

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	NamespaceNodeLease string = "kube-node-lease"
)

type MicroTime struct {
	time.Time `protobuf:"-"`
}

type Volume struct {
	Name         string `json:"name" protobuf:"bytes,1,opt,name=name"`
	VolumeSource `json:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}

type VolumeSource struct {
	HostPath              *HostPathVolumeSource              `json:"hostPath,omitempty" protobuf:"bytes,1,opt,name=hostPath"`
	EmptyDir              *EmptyDirVolumeSource              `json:"emptyDir,omitempty" protobuf:"bytes,2,opt,name=emptyDir"`
	GCEPersistentDisk     *GCEPersistentDiskVolumeSource     `json:"gcePersistentDisk,omitempty" protobuf:"bytes,3,opt,name=gcePersistentDisk"`
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSource  `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,4,opt,name=awsElasticBlockStore"`
	GitRepo               *GitRepoVolumeSource               `json:"gitRepo,omitempty" protobuf:"bytes,5,opt,name=gitRepo"`
	Secret                *SecretVolumeSource                `json:"secret,omitempty" protobuf:"bytes,6,opt,name=secret"`
	NFS                   *NFSVolumeSource                   `json:"nfs,omitempty" protobuf:"bytes,7,opt,name=nfs"`
	ISCSI                 *ISCSIVolumeSource                 `json:"iscsi,omitempty" protobuf:"bytes,8,opt,name=iscsi"`
	Glusterfs             *GlusterfsVolumeSource             `json:"glusterfs,omitempty" protobuf:"bytes,9,opt,name=glusterfs"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty" protobuf:"bytes,10,opt,name=persistentVolumeClaim"`
	RBD                   *RBDVolumeSource                   `json:"rbd,omitempty" protobuf:"bytes,11,opt,name=rbd"`
	FlexVolume            *FlexVolumeSource                  `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume"`
	Cinder                *CinderVolumeSource                `json:"cinder,omitempty" protobuf:"bytes,13,opt,name=cinder"`
	CephFS                *CephFSVolumeSource                `json:"cephfs,omitempty" protobuf:"bytes,14,opt,name=cephfs"`
	Flocker               *FlockerVolumeSource               `json:"flocker,omitempty" protobuf:"bytes,15,opt,name=flocker"`
	DownwardAPI           *DownwardAPIVolumeSource           `json:"downwardAPI,omitempty" protobuf:"bytes,16,opt,name=downwardAPI"`
	FC                    *FCVolumeSource                    `json:"fc,omitempty" protobuf:"bytes,17,opt,name=fc"`
	AzureFile             *AzureFileVolumeSource             `json:"azureFile,omitempty" protobuf:"bytes,18,opt,name=azureFile"`
	ConfigMap             *ConfigMapVolumeSource             `json:"configMap,omitempty" protobuf:"bytes,19,opt,name=configMap"`
	VsphereVolume         *VsphereVirtualDiskVolumeSource    `json:"vsphereVolume,omitempty" protobuf:"bytes,20,opt,name=vsphereVolume"`
	Quobyte               *QuobyteVolumeSource               `json:"quobyte,omitempty" protobuf:"bytes,21,opt,name=quobyte"`
	AzureDisk             *AzureDiskVolumeSource             `json:"azureDisk,omitempty" protobuf:"bytes,22,opt,name=azureDisk"`
	PhotonPersistentDisk  *PhotonPersistentDiskVolumeSource  `json:"photonPersistentDisk,omitempty" protobuf:"bytes,23,opt,name=photonPersistentDisk"`
	Projected             *ProjectedVolumeSource             `json:"projected,omitempty" protobuf:"bytes,26,opt,name=projected"`
	PortworxVolume        *PortworxVolumeSource              `json:"portworxVolume,omitempty" protobuf:"bytes,24,opt,name=portworxVolume"`
	ScaleIO               *ScaleIOVolumeSource               `json:"scaleIO,omitempty" protobuf:"bytes,25,opt,name=scaleIO"`
	StorageOS             *StorageOSVolumeSource             `json:"storageos,omitempty" protobuf:"bytes,27,opt,name=storageos"`
}

type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `json:"claimName" protobuf:"bytes,1,opt,name=claimName"`
	ReadOnly  bool   `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
}

type PersistentVolumeSource struct {
	GCEPersistentDisk    *GCEPersistentDiskVolumeSource    `json:"gcePersistentDisk,omitempty" protobuf:"bytes,1,opt,name=gcePersistentDisk"`
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,2,opt,name=awsElasticBlockStore"`
	HostPath             *HostPathVolumeSource             `json:"hostPath,omitempty" protobuf:"bytes,3,opt,name=hostPath"`
	Glusterfs            *GlusterfsPersistentVolumeSource  `json:"glusterfs,omitempty" protobuf:"bytes,4,opt,name=glusterfs"`
	NFS                  *NFSVolumeSource                  `json:"nfs,omitempty" protobuf:"bytes,5,opt,name=nfs"`
	RBD                  *RBDPersistentVolumeSource        `json:"rbd,omitempty" protobuf:"bytes,6,opt,name=rbd"`
	ISCSI                *ISCSIPersistentVolumeSource      `json:"iscsi,omitempty" protobuf:"bytes,7,opt,name=iscsi"`
	Cinder               *CinderPersistentVolumeSource     `json:"cinder,omitempty" protobuf:"bytes,8,opt,name=cinder"`
	CephFS               *CephFSPersistentVolumeSource     `json:"cephfs,omitempty" protobuf:"bytes,9,opt,name=cephfs"`
	FC                   *FCVolumeSource                   `json:"fc,omitempty" protobuf:"bytes,10,opt,name=fc"`
	Flocker              *FlockerVolumeSource              `json:"flocker,omitempty" protobuf:"bytes,11,opt,name=flocker"`
	FlexVolume           *FlexPersistentVolumeSource       `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume"`
	AzureFile            *AzureFilePersistentVolumeSource  `json:"azureFile,omitempty" protobuf:"bytes,13,opt,name=azureFile"`
	VsphereVolume        *VsphereVirtualDiskVolumeSource   `json:"vsphereVolume,omitempty" protobuf:"bytes,14,opt,name=vsphereVolume"`
	Quobyte              *QuobyteVolumeSource              `json:"quobyte,omitempty" protobuf:"bytes,15,opt,name=quobyte"`
	AzureDisk            *AzureDiskVolumeSource            `json:"azureDisk,omitempty" protobuf:"bytes,16,opt,name=azureDisk"`
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty" protobuf:"bytes,17,opt,name=photonPersistentDisk"`
	PortworxVolume       *PortworxVolumeSource             `json:"portworxVolume,omitempty" protobuf:"bytes,18,opt,name=portworxVolume"`
	ScaleIO              *ScaleIOPersistentVolumeSource    `json:"scaleIO,omitempty" protobuf:"bytes,19,opt,name=scaleIO"`
	Local                *LocalVolumeSource                `json:"local,omitempty" protobuf:"bytes,20,opt,name=local"`
	StorageOS            *StorageOSPersistentVolumeSource  `json:"storageos,omitempty" protobuf:"bytes,21,opt,name=storageos"`
	CSI                  *CSIPersistentVolumeSource        `json:"csi,omitempty" protobuf:"bytes,22,opt,name=csi"`
}

const (
	BetaStorageClassAnnotation = "volume.beta.kubernetes.io/storage-class"
	MountOptionAnnotation      = "volume.beta.kubernetes.io/mount-options"
)

type PersistentVolume struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       PersistentVolumeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     PersistentVolumeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PersistentVolumeSpec struct {
	Capacity                      ResourceList `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	PersistentVolumeSource        `json:",inline" protobuf:"bytes,2,opt,name=persistentVolumeSource"`
	AccessModes                   []PersistentVolumeAccessMode  `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	ClaimRef                      *ObjectReference              `json:"claimRef,omitempty" protobuf:"bytes,4,opt,name=claimRef"`
	PersistentVolumeReclaimPolicy PersistentVolumeReclaimPolicy `json:"persistentVolumeReclaimPolicy,omitempty" protobuf:"bytes,5,opt,name=persistentVolumeReclaimPolicy,casttype=PersistentVolumeReclaimPolicy"`
	StorageClassName              string                        `json:"storageClassName,omitempty" protobuf:"bytes,6,opt,name=storageClassName"`
	MountOptions                  []string                      `json:"mountOptions,omitempty" protobuf:"bytes,7,opt,name=mountOptions"`
	VolumeMode                    *PersistentVolumeMode         `json:"volumeMode,omitempty" protobuf:"bytes,8,opt,name=volumeMode,casttype=PersistentVolumeMode"`
	NodeAffinity                  *VolumeNodeAffinity           `json:"nodeAffinity,omitempty" protobuf:"bytes,9,opt,name=nodeAffinity"`
}

type VolumeNodeAffinity struct {
	Required *NodeSelector `json:"required,omitempty" protobuf:"bytes,1,opt,name=required"`
}

type PersistentVolumeReclaimPolicy string

const (
	PersistentVolumeReclaimRecycle PersistentVolumeReclaimPolicy = "Recycle"
	PersistentVolumeReclaimDelete  PersistentVolumeReclaimPolicy = "Delete"
	PersistentVolumeReclaimRetain  PersistentVolumeReclaimPolicy = "Retain"
)

type PersistentVolumeMode string

const (
	PersistentVolumeBlock      PersistentVolumeMode = "Block"
	PersistentVolumeFilesystem PersistentVolumeMode = "Filesystem"
)

type PersistentVolumeStatus struct {
	Phase   PersistentVolumePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PersistentVolumePhase"`
	Message string                `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	Reason  string                `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
}

type PersistentVolumeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []PersistentVolume `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type PersistentVolumeClaim struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       PersistentVolumeClaimSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     PersistentVolumeClaimStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PersistentVolumeClaimList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []PersistentVolumeClaim `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type PersistentVolumeClaimSpec struct {
	AccessModes      []PersistentVolumeAccessMode `json:"accessModes,omitempty" protobuf:"bytes,1,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	Selector         *LabelSelector               `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector"`
	Resources        ResourceRequirements         `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources"`
	VolumeName       string                       `json:"volumeName,omitempty" protobuf:"bytes,3,opt,name=volumeName"`
	StorageClassName *string                      `json:"storageClassName,omitempty" protobuf:"bytes,5,opt,name=storageClassName"`
	VolumeMode       *PersistentVolumeMode        `json:"volumeMode,omitempty" protobuf:"bytes,6,opt,name=volumeMode,casttype=PersistentVolumeMode"`
	DataSource       *TypedLocalObjectReference   `json:"dataSource" protobuf:"bytes,7,opt,name=dataSource"`
}

type PersistentVolumeClaimConditionType string

const (
	PersistentVolumeClaimResizing                PersistentVolumeClaimConditionType = "Resizing"
	PersistentVolumeClaimFileSystemResizePending PersistentVolumeClaimConditionType = "FileSystemResizePending"
)

type PersistentVolumeClaimCondition struct {
	Type               PersistentVolumeClaimConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PersistentVolumeClaimConditionType"`
	Status             ConditionStatus                    `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	LastProbeTime      Time                               `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	LastTransitionTime Time                               `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	Reason             string                             `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	Message            string                             `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type PersistentVolumeClaimStatus struct {
	Phase       PersistentVolumeClaimPhase       `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PersistentVolumeClaimPhase"`
	AccessModes []PersistentVolumeAccessMode     `json:"accessModes,omitempty" protobuf:"bytes,2,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	Capacity    ResourceList                     `json:"capacity,omitempty" protobuf:"bytes,3,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	Conditions  []PersistentVolumeClaimCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,4,rep,name=conditions"`
}

type PersistentVolumeAccessMode string

const (
	ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
	ReadOnlyMany  PersistentVolumeAccessMode = "ReadOnlyMany"
	ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
)

type PersistentVolumePhase string

const (
	VolumePending   PersistentVolumePhase = "Pending"
	VolumeAvailable PersistentVolumePhase = "Available"
	VolumeBound     PersistentVolumePhase = "Bound"
	VolumeReleased  PersistentVolumePhase = "Released"
	VolumeFailed    PersistentVolumePhase = "Failed"
)

type PersistentVolumeClaimPhase string

const (
	ClaimPending PersistentVolumeClaimPhase = "Pending"
	ClaimBound   PersistentVolumeClaimPhase = "Bound"
	ClaimLost    PersistentVolumeClaimPhase = "Lost"
)

type HostPathType string

const (
	HostPathUnset             HostPathType = ""
	HostPathDirectoryOrCreate HostPathType = "DirectoryOrCreate"
	HostPathDirectory         HostPathType = "Directory"
	HostPathFileOrCreate      HostPathType = "FileOrCreate"
	HostPathFile              HostPathType = "File"
	HostPathSocket            HostPathType = "Socket"
	HostPathCharDev           HostPathType = "CharDevice"
	HostPathBlockDev          HostPathType = "BlockDevice"
)

type HostPathVolumeSource struct {
	Path string        `json:"path" protobuf:"bytes,1,opt,name=path"`
	Type *HostPathType `json:"type,omitempty" protobuf:"bytes,2,opt,name=type"`
}

type EmptyDirVolumeSource struct {
	Medium    StorageMedium `json:"medium,omitempty" protobuf:"bytes,1,opt,name=medium,casttype=StorageMedium"`
	SizeLimit string        `json:"sizeLimit,omitempty" protobuf:"bytes,2,opt,name=sizeLimit"`
}

type GlusterfsVolumeSource struct {
	EndpointsName string `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints"`
	Path          string `json:"path" protobuf:"bytes,2,opt,name=path"`
	ReadOnly      bool   `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
}

type GlusterfsPersistentVolumeSource struct {
	EndpointsName      string  `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints"`
	Path               string  `json:"path" protobuf:"bytes,2,opt,name=path"`
	ReadOnly           bool    `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
	EndpointsNamespace *string `json:"endpointsNamespace,omitempty" protobuf:"bytes,4,opt,name=endpointsNamespace"`
}

type RBDVolumeSource struct {
	CephMonitors []string              `json:"monitors" protobuf:"bytes,1,rep,name=monitors"`
	RBDImage     string                `json:"image" protobuf:"bytes,2,opt,name=image"`
	FSType       string                `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`
	RBDPool      string                `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool"`
	RadosUser    string                `json:"user,omitempty" protobuf:"bytes,5,opt,name=user"`
	Keyring      string                `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring"`
	SecretRef    *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef"`
	ReadOnly     bool                  `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly"`
}

type RBDPersistentVolumeSource struct {
	CephMonitors []string         `json:"monitors" protobuf:"bytes,1,rep,name=monitors"`
	RBDImage     string           `json:"image" protobuf:"bytes,2,opt,name=image"`
	FSType       string           `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`
	RBDPool      string           `json:"pool,omitempty" protobuf:"bytes,4,opt,name=pool"`
	RadosUser    string           `json:"user,omitempty" protobuf:"bytes,5,opt,name=user"`
	Keyring      string           `json:"keyring,omitempty" protobuf:"bytes,6,opt,name=keyring"`
	SecretRef    *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,7,opt,name=secretRef"`
	ReadOnly     bool             `json:"readOnly,omitempty" protobuf:"varint,8,opt,name=readOnly"`
}

type CinderVolumeSource struct {
	VolumeID  string                `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"`
	FSType    string                `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	ReadOnly  bool                  `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,4,opt,name=secretRef"`
}

type CinderPersistentVolumeSource struct {
	VolumeID  string           `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"`
	FSType    string           `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	ReadOnly  bool             `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
	SecretRef *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,4,opt,name=secretRef"`
}

type CephFSVolumeSource struct {
	Monitors   []string              `json:"monitors" protobuf:"bytes,1,rep,name=monitors"`
	Path       string                `json:"path,omitempty" protobuf:"bytes,2,opt,name=path"`
	User       string                `json:"user,omitempty" protobuf:"bytes,3,opt,name=user"`
	SecretFile string                `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile"`
	SecretRef  *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
	ReadOnly   bool                  `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"`
}

type SecretReference struct {
	Name      string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
}

type CephFSPersistentVolumeSource struct {
	Monitors   []string         `json:"monitors" protobuf:"bytes,1,rep,name=monitors"`
	Path       string           `json:"path,omitempty" protobuf:"bytes,2,opt,name=path"`
	User       string           `json:"user,omitempty" protobuf:"bytes,3,opt,name=user"`
	SecretFile string           `json:"secretFile,omitempty" protobuf:"bytes,4,opt,name=secretFile"`
	SecretRef  *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
	ReadOnly   bool             `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"`
}

type FlockerVolumeSource struct {
	DatasetName string `json:"datasetName,omitempty" protobuf:"bytes,1,opt,name=datasetName"`
	DatasetUUID string `json:"datasetUUID,omitempty" protobuf:"bytes,2,opt,name=datasetUUID"`
}

type StorageMedium string

const (
	StorageMediumDefault   StorageMedium = ""
	StorageMediumMemory    StorageMedium = "Memory"
	StorageMediumHugePages StorageMedium = "HugePages"
)

type Protocol string

const (
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolSCTP Protocol = "SCTP"
)

type GCEPersistentDiskVolumeSource struct {
	PDName    string `json:"pdName" protobuf:"bytes,1,opt,name=pdName"`
	FSType    string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	Partition int32  `json:"partition,omitempty" protobuf:"varint,3,opt,name=partition"`
	ReadOnly  bool   `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
}

type QuobyteVolumeSource struct {
	Registry string `json:"registry" protobuf:"bytes,1,opt,name=registry"`
	Volume   string `json:"volume" protobuf:"bytes,2,opt,name=volume"`
	ReadOnly bool   `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
	User     string `json:"user,omitempty" protobuf:"bytes,4,opt,name=user"`
	Group    string `json:"group,omitempty" protobuf:"bytes,5,opt,name=group"`
}

type FlexPersistentVolumeSource struct {
	Driver    string            `json:"driver" protobuf:"bytes,1,opt,name=driver"`
	FSType    string            `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	SecretRef *SecretReference  `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
	ReadOnly  bool              `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
	Options   map[string]string `json:"options,omitempty" protobuf:"bytes,5,rep,name=options"`
}

type FlexVolumeSource struct {
	Driver    string                `json:"driver" protobuf:"bytes,1,opt,name=driver"`
	FSType    string                `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
	ReadOnly  bool                  `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
	Options   map[string]string     `json:"options,omitempty" protobuf:"bytes,5,rep,name=options"`
}

type AWSElasticBlockStoreVolumeSource struct {
	VolumeID  string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"`
	FSType    string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	Partition int32  `json:"partition,omitempty" protobuf:"varint,3,opt,name=partition"`
	ReadOnly  bool   `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
}

type GitRepoVolumeSource struct {
	Repository string `json:"repository" protobuf:"bytes,1,opt,name=repository"`
	Revision   string `json:"revision,omitempty" protobuf:"bytes,2,opt,name=revision"`
	Directory  string `json:"directory,omitempty" protobuf:"bytes,3,opt,name=directory"`
}

type SecretVolumeSource struct {
	SecretName  string      `json:"secretName,omitempty" protobuf:"bytes,1,opt,name=secretName"`
	Items       []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	DefaultMode *int32      `json:"defaultMode,omitempty" protobuf:"bytes,3,opt,name=defaultMode"`
	Optional    *bool       `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

const (
	SecretVolumeSourceDefaultMode int32 = 0644
)

type SecretProjection struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Items                []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	Optional             *bool       `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

type NFSVolumeSource struct {
	Server   string `json:"server" protobuf:"bytes,1,opt,name=server"`
	Path     string `json:"path" protobuf:"bytes,2,opt,name=path"`
	ReadOnly bool   `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
}

type ISCSIVolumeSource struct {
	TargetPortal      string                `json:"targetPortal" protobuf:"bytes,1,opt,name=targetPortal"`
	IQN               string                `json:"iqn" protobuf:"bytes,2,opt,name=iqn"`
	Lun               int32                 `json:"lun" protobuf:"varint,3,opt,name=lun"`
	ISCSIInterface    string                `json:"iscsiInterface,omitempty" protobuf:"bytes,4,opt,name=iscsiInterface"`
	FSType            string                `json:"fsType,omitempty" protobuf:"bytes,5,opt,name=fsType"`
	ReadOnly          bool                  `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"`
	Portals           []string              `json:"portals,omitempty" protobuf:"bytes,7,opt,name=portals"`
	DiscoveryCHAPAuth bool                  `json:"chapAuthDiscovery,omitempty" protobuf:"varint,8,opt,name=chapAuthDiscovery"`
	SessionCHAPAuth   bool                  `json:"chapAuthSession,omitempty" protobuf:"varint,11,opt,name=chapAuthSession"`
	SecretRef         *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,10,opt,name=secretRef"`
	InitiatorName     *string               `json:"initiatorName,omitempty" protobuf:"bytes,12,opt,name=initiatorName"`
}

type ISCSIPersistentVolumeSource struct {
	TargetPortal      string           `json:"targetPortal" protobuf:"bytes,1,opt,name=targetPortal"`
	IQN               string           `json:"iqn" protobuf:"bytes,2,opt,name=iqn"`
	Lun               int32            `json:"lun" protobuf:"varint,3,opt,name=lun"`
	ISCSIInterface    string           `json:"iscsiInterface,omitempty" protobuf:"bytes,4,opt,name=iscsiInterface"`
	FSType            string           `json:"fsType,omitempty" protobuf:"bytes,5,opt,name=fsType"`
	ReadOnly          bool             `json:"readOnly,omitempty" protobuf:"varint,6,opt,name=readOnly"`
	Portals           []string         `json:"portals,omitempty" protobuf:"bytes,7,opt,name=portals"`
	DiscoveryCHAPAuth bool             `json:"chapAuthDiscovery,omitempty" protobuf:"varint,8,opt,name=chapAuthDiscovery"`
	SessionCHAPAuth   bool             `json:"chapAuthSession,omitempty" protobuf:"varint,11,opt,name=chapAuthSession"`
	SecretRef         *SecretReference `json:"secretRef,omitempty" protobuf:"bytes,10,opt,name=secretRef"`
	InitiatorName     *string          `json:"initiatorName,omitempty" protobuf:"bytes,12,opt,name=initiatorName"`
}

type FCVolumeSource struct {
	TargetWWNs []string `json:"targetWWNs,omitempty" protobuf:"bytes,1,rep,name=targetWWNs"`
	Lun        *int32   `json:"lun,omitempty" protobuf:"varint,2,opt,name=lun"`
	FSType     string   `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`
	ReadOnly   bool     `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
	WWIDs      []string `json:"wwids,omitempty" protobuf:"bytes,5,rep,name=wwids"`
}

type AzureFileVolumeSource struct {
	SecretName string `json:"secretName" protobuf:"bytes,1,opt,name=secretName"`
	ShareName  string `json:"shareName" protobuf:"bytes,2,opt,name=shareName"`
	ReadOnly   bool   `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
}

type AzureFilePersistentVolumeSource struct {
	SecretName      string  `json:"secretName" protobuf:"bytes,1,opt,name=secretName"`
	ShareName       string  `json:"shareName" protobuf:"bytes,2,opt,name=shareName"`
	ReadOnly        bool    `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
	SecretNamespace *string `json:"secretNamespace" protobuf:"bytes,4,opt,name=secretNamespace"`
}

type VsphereVirtualDiskVolumeSource struct {
	VolumePath        string `json:"volumePath" protobuf:"bytes,1,opt,name=volumePath"`
	FSType            string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	StoragePolicyName string `json:"storagePolicyName,omitempty" protobuf:"bytes,3,opt,name=storagePolicyName"`
	StoragePolicyID   string `json:"storagePolicyID,omitempty" protobuf:"bytes,4,opt,name=storagePolicyID"`
}

type PhotonPersistentDiskVolumeSource struct {
	PdID   string `json:"pdID" protobuf:"bytes,1,opt,name=pdID"`
	FSType string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
}

type AzureDataDiskCachingMode string
type AzureDataDiskKind string

const (
	AzureDataDiskCachingNone      AzureDataDiskCachingMode = "None"
	AzureDataDiskCachingReadOnly  AzureDataDiskCachingMode = "ReadOnly"
	AzureDataDiskCachingReadWrite AzureDataDiskCachingMode = "ReadWrite"
	AzureSharedBlobDisk           AzureDataDiskKind        = "Shared"
	AzureDedicatedBlobDisk        AzureDataDiskKind        = "Dedicated"
	AzureManagedDisk              AzureDataDiskKind        = "Managed"
)

type AzureDiskVolumeSource struct {
	DiskName    string                    `json:"diskName" protobuf:"bytes,1,opt,name=diskName"`
	DataDiskURI string                    `json:"diskURI" protobuf:"bytes,2,opt,name=diskURI"`
	CachingMode *AzureDataDiskCachingMode `json:"cachingMode,omitempty" protobuf:"bytes,3,opt,name=cachingMode,casttype=AzureDataDiskCachingMode"`
	FSType      *string                   `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType"`
	ReadOnly    *bool                     `json:"readOnly,omitempty" protobuf:"varint,5,opt,name=readOnly"`
	Kind        *AzureDataDiskKind        `json:"kind,omitempty" protobuf:"bytes,6,opt,name=kind,casttype=AzureDataDiskKind"`
}

type PortworxVolumeSource struct {
	VolumeID string `json:"volumeID" protobuf:"bytes,1,opt,name=volumeID"`
	FSType   string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	ReadOnly bool   `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
}

type ScaleIOVolumeSource struct {
	Gateway          string                `json:"gateway" protobuf:"bytes,1,opt,name=gateway"`
	System           string                `json:"system" protobuf:"bytes,2,opt,name=system"`
	SecretRef        *LocalObjectReference `json:"secretRef" protobuf:"bytes,3,opt,name=secretRef"`
	SSLEnabled       bool                  `json:"sslEnabled,omitempty" protobuf:"varint,4,opt,name=sslEnabled"`
	ProtectionDomain string                `json:"protectionDomain,omitempty" protobuf:"bytes,5,opt,name=protectionDomain"`
	StoragePool      string                `json:"storagePool,omitempty" protobuf:"bytes,6,opt,name=storagePool"`
	StorageMode      string                `json:"storageMode,omitempty" protobuf:"bytes,7,opt,name=storageMode"`
	VolumeName       string                `json:"volumeName,omitempty" protobuf:"bytes,8,opt,name=volumeName"`
	FSType           string                `json:"fsType,omitempty" protobuf:"bytes,9,opt,name=fsType"`
	ReadOnly         bool                  `json:"readOnly,omitempty" protobuf:"varint,10,opt,name=readOnly"`
}

type ScaleIOPersistentVolumeSource struct {
	Gateway          string           `json:"gateway" protobuf:"bytes,1,opt,name=gateway"`
	System           string           `json:"system" protobuf:"bytes,2,opt,name=system"`
	SecretRef        *SecretReference `json:"secretRef" protobuf:"bytes,3,opt,name=secretRef"`
	SSLEnabled       bool             `json:"sslEnabled,omitempty" protobuf:"varint,4,opt,name=sslEnabled"`
	ProtectionDomain string           `json:"protectionDomain,omitempty" protobuf:"bytes,5,opt,name=protectionDomain"`
	StoragePool      string           `json:"storagePool,omitempty" protobuf:"bytes,6,opt,name=storagePool"`
	StorageMode      string           `json:"storageMode,omitempty" protobuf:"bytes,7,opt,name=storageMode"`
	VolumeName       string           `json:"volumeName,omitempty" protobuf:"bytes,8,opt,name=volumeName"`
	FSType           string           `json:"fsType,omitempty" protobuf:"bytes,9,opt,name=fsType"`
	ReadOnly         bool             `json:"readOnly,omitempty" protobuf:"varint,10,opt,name=readOnly"`
}

type StorageOSVolumeSource struct {
	VolumeName      string                `json:"volumeName,omitempty" protobuf:"bytes,1,opt,name=volumeName"`
	VolumeNamespace string                `json:"volumeNamespace,omitempty" protobuf:"bytes,2,opt,name=volumeNamespace"`
	FSType          string                `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`
	ReadOnly        bool                  `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
	SecretRef       *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
}

type StorageOSPersistentVolumeSource struct {
	VolumeName      string           `json:"volumeName,omitempty" protobuf:"bytes,1,opt,name=volumeName"`
	VolumeNamespace string           `json:"volumeNamespace,omitempty" protobuf:"bytes,2,opt,name=volumeNamespace"`
	FSType          string           `json:"fsType,omitempty" protobuf:"bytes,3,opt,name=fsType"`
	ReadOnly        bool             `json:"readOnly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
	SecretRef       *ObjectReference `json:"secretRef,omitempty" protobuf:"bytes,5,opt,name=secretRef"`
}

type ConfigMapVolumeSource struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Items                []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	DefaultMode          *int32      `json:"defaultMode,omitempty" protobuf:"varint,3,opt,name=defaultMode"`
	Optional             *bool       `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

const (
	ConfigMapVolumeSourceDefaultMode int32 = 0644
)

type ConfigMapProjection struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Items                []KeyToPath `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	Optional             *bool       `json:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

type ServiceAccountTokenProjection struct {
	Audience          string `json:"audience,omitempty" protobuf:"bytes,1,rep,name=audience"`
	ExpirationSeconds *int64 `json:"expirationSeconds,omitempty" protobuf:"varint,2,opt,name=expirationSeconds"`
	Path              string `json:"path" protobuf:"bytes,3,opt,name=path"`
}

type ProjectedVolumeSource struct {
	Sources     []VolumeProjection `json:"sources" protobuf:"bytes,1,rep,name=sources"`
	DefaultMode *int32             `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode"`
}

type VolumeProjection struct {
	Secret              *SecretProjection              `json:"secret,omitempty" protobuf:"bytes,1,opt,name=secret"`
	DownwardAPI         *DownwardAPIProjection         `json:"downwardAPI,omitempty" protobuf:"bytes,2,opt,name=downwardAPI"`
	ConfigMap           *ConfigMapProjection           `json:"configMap,omitempty" protobuf:"bytes,3,opt,name=configMap"`
	ServiceAccountToken *ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty" protobuf:"bytes,4,opt,name=serviceAccountToken"`
}

const (
	ProjectedVolumeSourceDefaultMode int32 = 0644
)

type KeyToPath struct {
	Key  string `json:"key" protobuf:"bytes,1,opt,name=key"`
	Path string `json:"path" protobuf:"bytes,2,opt,name=path"`
	Mode *int32 `json:"mode,omitempty" protobuf:"varint,3,opt,name=mode"`
}

type LocalVolumeSource struct {
	Path   string  `json:"path" protobuf:"bytes,1,opt,name=path"`
	FSType *string `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
}

type CSIPersistentVolumeSource struct {
	Driver                     string            `json:"driver" protobuf:"bytes,1,opt,name=driver"`
	VolumeHandle               string            `json:"volumeHandle" protobuf:"bytes,2,opt,name=volumeHandle"`
	ReadOnly                   bool              `json:"readOnly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
	FSType                     string            `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType"`
	VolumeAttributes           map[string]string `json:"volumeAttributes,omitempty" protobuf:"bytes,5,rep,name=volumeAttributes"`
	ControllerPublishSecretRef *SecretReference  `json:"controllerPublishSecretRef,omitempty" protobuf:"bytes,6,opt,name=controllerPublishSecretRef"`
	NodeStageSecretRef         *SecretReference  `json:"nodeStageSecretRef,omitempty" protobuf:"bytes,7,opt,name=nodeStageSecretRef"`
	NodePublishSecretRef       *SecretReference  `json:"nodePublishSecretRef,omitempty" protobuf:"bytes,8,opt,name=nodePublishSecretRef"`
}

type ContainerPort struct {
	Name          string   `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	HostPort      int32    `json:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort"`
	ContainerPort int32    `json:"containerPort" protobuf:"varint,3,opt,name=containerPort"`
	Protocol      Protocol `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol"`
	HostIP        string   `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP"`
}

type VolumeMount struct {
	Name             string                `json:"name" protobuf:"bytes,1,opt,name=name"`
	ReadOnly         bool                  `json:"readOnly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
	MountPath        string                `json:"mountPath" protobuf:"bytes,3,opt,name=mountPath"`
	SubPath          string                `json:"subPath,omitempty" protobuf:"bytes,4,opt,name=subPath"`
	MountPropagation *MountPropagationMode `json:"mountPropagation,omitempty" protobuf:"bytes,5,opt,name=mountPropagation,casttype=MountPropagationMode"`
}

type MountPropagationMode string

const (
	MountPropagationNone            MountPropagationMode = "None"
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	MountPropagationBidirectional   MountPropagationMode = "Bidirectional"
)

type VolumeDevice struct {
	Name       string `json:"name" protobuf:"bytes,1,opt,name=name"`
	DevicePath string `json:"devicePath" protobuf:"bytes,2,opt,name=devicePath"`
}

type EnvVar struct {
	Name      string        `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value     string        `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty" protobuf:"bytes,3,opt,name=valueFrom"`
}

type EnvVarSource struct {
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty" protobuf:"bytes,1,opt,name=fieldRef"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,2,opt,name=resourceFieldRef"`
	ConfigMapKeyRef  *ConfigMapKeySelector  `json:"configMapKeyRef,omitempty" protobuf:"bytes,3,opt,name=configMapKeyRef"`
	SecretKeyRef     *SecretKeySelector     `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

type ObjectFieldSelector struct {
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,1,opt,name=apiVersion"`
	FieldPath  string `json:"fieldPath" protobuf:"bytes,2,opt,name=fieldPath"`
}

type ResourceFieldSelector struct {
	ContainerName string `json:"containerName,omitempty" protobuf:"bytes,1,opt,name=containerName"`
	Resource      string `json:"resource" protobuf:"bytes,2,opt,name=resource"`
	Divisor       string `json:"divisor,omitempty" protobuf:"bytes,3,opt,name=divisor"`
}

type ConfigMapKeySelector struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Key                  string `json:"key" protobuf:"bytes,2,opt,name=key"`
	Optional             *bool  `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional"`
}

type SecretKeySelector struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Key                  string `json:"key" protobuf:"bytes,2,opt,name=key"`
	Optional             *bool  `json:"optional,omitempty" protobuf:"varint,3,opt,name=optional"`
}

type EnvFromSource struct {
	Prefix       string              `json:"prefix,omitempty" protobuf:"bytes,1,opt,name=prefix"`
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty" protobuf:"bytes,2,opt,name=configMapRef"`
	SecretRef    *SecretEnvSource    `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
}

type ConfigMapEnvSource struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Optional             *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional"`
}

type SecretEnvSource struct {
	LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Optional             *bool `json:"optional,omitempty" protobuf:"varint,2,opt,name=optional"`
}

type HTTPHeader struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

type HTTPGetAction struct {
	Path        string       `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	Port        json.Number  `json:"port,Number" protobuf:"bytes,2,opt,name=port"`
	Host        string       `json:"host,omitempty" protobuf:"bytes,3,opt,name=host"`
	Scheme      URIScheme    `json:"scheme,omitempty" protobuf:"bytes,4,opt,name=scheme,casttype=URIScheme"`
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty" protobuf:"bytes,5,rep,name=httpHeaders"`
}

type URIScheme string

const (
	URISchemeHTTP  URIScheme = "HTTP"
	URISchemeHTTPS URIScheme = "HTTPS"
)

type TCPSocketAction struct {
	Port json.Number `json:"port,Number" protobuf:"bytes,1,opt,name=port"`
	Host string      `json:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
}

type ExecAction struct {
	Command []string `json:"command,omitempty" protobuf:"bytes,1,rep,name=command"`
}

type Probe struct {
	Handler             `json:",inline" protobuf:"bytes,1,opt,name=handler"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initialDelaySeconds"`
	TimeoutSeconds      int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeoutSeconds"`
	PeriodSeconds       int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=periodSeconds"`
	SuccessThreshold    int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=successThreshold"`
	FailureThreshold    int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failureThreshold"`
}

type PullPolicy string

const (
	PullAlways       PullPolicy = "Always"
	PullNever        PullPolicy = "Never"
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

type TerminationMessagePolicy string

const (
	TerminationMessageReadFile              TerminationMessagePolicy = "File"
	TerminationMessageFallbackToLogsOnError TerminationMessagePolicy = "FallbackToLogsOnError"
)

type Capability string

type Capabilities struct {
	Add  []Capability `json:"add,omitempty" protobuf:"bytes,1,rep,name=add,casttype=Capability"`
	Drop []Capability `json:"drop,omitempty" protobuf:"bytes,2,rep,name=drop,casttype=Capability"`
}

type ResourceRequirements struct {
	Limits   ResourceList `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName"`
	Requests ResourceList `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName"`
}

const (
	TerminationMessagePathDefault string = "/dev/termination-log"
)

type Container struct {
	Name                     string                   `json:"name" protobuf:"bytes,1,opt,name=name"`
	Image                    string                   `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"`
	Command                  []string                 `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	Args                     []string                 `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	WorkingDir               string                   `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir"`
	Ports                    []ContainerPort          `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`
	EnvFrom                  []EnvFromSource          `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`
	Env                      []EnvVar                 `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
	Resources                ResourceRequirements     `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
	VolumeMounts             []VolumeMount            `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"`
	VolumeDevices            []VolumeDevice           `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"`
	LivenessProbe            *Probe                   `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	ReadinessProbe           *Probe                   `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	Lifecycle                *Lifecycle               `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`
	TerminationMessagePath   string                   `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath"`
	TerminationMessagePolicy TerminationMessagePolicy `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy"`
	ImagePullPolicy          PullPolicy               `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"`
	SecurityContext          *SecurityContext         `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`
	Stdin                    bool                     `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin"`
	StdinOnce                bool                     `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce"`
	TTY                      bool                     `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty"`
}

type Handler struct {
	Exec      *ExecAction      `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	HTTPGet   *HTTPGetAction   `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=httpGet"`
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty" protobuf:"bytes,3,opt,name=tcpSocket"`
}

type Lifecycle struct {
	PostStart *Handler `json:"postStart,omitempty" protobuf:"bytes,1,opt,name=postStart"`
	PreStop   *Handler `json:"preStop,omitempty" protobuf:"bytes,2,opt,name=preStop"`
}

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type ContainerStateWaiting struct {
	Reason  string `json:"reason,omitempty" protobuf:"bytes,1,opt,name=reason"`
	Message string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
}

type ContainerStateRunning struct {
	StartedAt Time `json:"startedAt,omitempty" protobuf:"bytes,1,opt,name=startedAt"`
}

type ContainerStateTerminated struct {
	ExitCode    int32  `json:"exitCode" protobuf:"varint,1,opt,name=exitCode"`
	Signal      int32  `json:"signal,omitempty" protobuf:"varint,2,opt,name=signal"`
	Reason      string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Message     string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
	StartedAt   Time   `json:"startedAt,omitempty" protobuf:"bytes,5,opt,name=startedAt"`
	FinishedAt  Time   `json:"finishedAt,omitempty" protobuf:"bytes,6,opt,name=finishedAt"`
	ContainerID string `json:"containerID,omitempty" protobuf:"bytes,7,opt,name=containerID"`
}

type ContainerState struct {
	Waiting    *ContainerStateWaiting    `json:"waiting,omitempty" protobuf:"bytes,1,opt,name=waiting"`
	Running    *ContainerStateRunning    `json:"running,omitempty" protobuf:"bytes,2,opt,name=running"`
	Terminated *ContainerStateTerminated `json:"terminated,omitempty" protobuf:"bytes,3,opt,name=terminated"`
}

type ContainerStatus struct {
	Name                 string         `json:"name" protobuf:"bytes,1,opt,name=name"`
	State                ContainerState `json:"state,omitempty" protobuf:"bytes,2,opt,name=state"`
	LastTerminationState ContainerState `json:"lastState,omitempty" protobuf:"bytes,3,opt,name=lastState"`
	Ready                bool           `json:"ready" protobuf:"varint,4,opt,name=ready"`
	RestartCount         int32          `json:"restartCount" protobuf:"varint,5,opt,name=restartCount"`
	Image                string         `json:"image" protobuf:"bytes,6,opt,name=image"`
	ImageID              string         `json:"imageID" protobuf:"bytes,7,opt,name=imageID"`
	ContainerID          string         `json:"containerID,omitempty" protobuf:"bytes,8,opt,name=containerID"`
}

type PodPhase string

const (
	PodPending   PodPhase = "Pending"
	PodRunning   PodPhase = "Running"
	PodSucceeded PodPhase = "Succeeded"
	PodFailed    PodPhase = "Failed"
	PodUnknown   PodPhase = "Unknown"
)

type PodConditionType string

const (
	PodScheduled           PodConditionType = "PodScheduled"
	PodReady               PodConditionType = "Ready"
	PodInitialized         PodConditionType = "Initialized"
	PodReasonUnschedulable                  = "Unschedulable"
	ContainersReady        PodConditionType = "ContainersReady"
)

type PodCondition struct {
	Type               PodConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PodConditionType"`
	Status             ConditionStatus  `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	LastProbeTime      Time             `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	LastTransitionTime Time             `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	Reason             string           `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	Message            string           `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)

type DNSPolicy string

const (
	DNSClusterFirstWithHostNet DNSPolicy = "ClusterFirstWithHostNet"
	DNSClusterFirst            DNSPolicy = "ClusterFirst"
	DNSDefault                 DNSPolicy = "Default"
	DNSNone                    DNSPolicy = "None"
)
const (
	DefaultTerminationGracePeriodSeconds = 30
)

type NodeSelector struct {
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms" protobuf:"bytes,1,rep,name=nodeSelectorTerms"`
}

type NodeSelectorTerm struct {
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,1,rep,name=matchExpressions"`
	MatchFields      []NodeSelectorRequirement `json:"matchFields,omitempty" protobuf:"bytes,2,rep,name=matchFields"`
}

type NodeSelectorRequirement struct {
	Key      string               `json:"key" protobuf:"bytes,1,opt,name=key"`
	Operator NodeSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=NodeSelectorOperator"`
	Values   []string             `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type NodeSelectorOperator string

const (
	NodeSelectorOpIn           NodeSelectorOperator = "In"
	NodeSelectorOpNotIn        NodeSelectorOperator = "NotIn"
	NodeSelectorOpExists       NodeSelectorOperator = "Exists"
	NodeSelectorOpDoesNotExist NodeSelectorOperator = "DoesNotExist"
	NodeSelectorOpGt           NodeSelectorOperator = "Gt"
	NodeSelectorOpLt           NodeSelectorOperator = "Lt"
)

type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement `json:"matchLabelExpressions,omitempty" protobuf:"bytes,1,rep,name=matchLabelExpressions"`
}

type TopologySelectorLabelRequirement struct {
	Key    string   `json:"key" protobuf:"bytes,1,opt,name=key"`
	Values []string `json:"values" protobuf:"bytes,2,rep,name=values"`
}

type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"nodeAffinity,omitempty" protobuf:"bytes,1,opt,name=nodeAffinity"`
	PodAffinity     *PodAffinity     `json:"podAffinity,omitempty" protobuf:"bytes,2,opt,name=podAffinity"`
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty" protobuf:"bytes,3,opt,name=podAntiAffinity"`
}

type PodAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"`
}

type PodAntiAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"`
}

type WeightedPodAffinityTerm struct {
	Weight          int32           `json:"weight" protobuf:"varint,1,opt,name=weight"`
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm" protobuf:"bytes,2,opt,name=podAffinityTerm"`
}

type PodAffinityTerm struct {
	LabelSelector *LabelSelector `json:"labelSelector,omitempty" protobuf:"bytes,1,opt,name=labelSelector"`
	Namespaces    []string       `json:"namespaces,omitempty" protobuf:"bytes,2,rep,name=namespaces"`
	TopologyKey   string         `json:"topologyKey" protobuf:"bytes,3,opt,name=topologyKey"`
}

type NodeAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *NodeSelector             `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,1,opt,name=requiredDuringSchedulingIgnoredDuringExecution"`
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty" protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution"`
}

type PreferredSchedulingTerm struct {
	Weight     int32            `json:"weight" protobuf:"varint,1,opt,name=weight"`
	Preference NodeSelectorTerm `json:"preference" protobuf:"bytes,2,opt,name=preference"`
}

type Taint struct {
	Key       string      `json:"key" protobuf:"bytes,1,opt,name=key"`
	Value     string      `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
	Effect    TaintEffect `json:"effect" protobuf:"bytes,3,opt,name=effect,casttype=TaintEffect"`
	TimeAdded *Time       `json:"timeAdded,omitempty" protobuf:"bytes,4,opt,name=timeAdded"`
}

type TaintEffect string

const (
	TaintEffectNoSchedule       TaintEffect = "NoSchedule"
	TaintEffectPreferNoSchedule TaintEffect = "PreferNoSchedule"
	TaintEffectNoExecute        TaintEffect = "NoExecute"
)

type Toleration struct {
	Key               string             `json:"key,omitempty" protobuf:"bytes,1,opt,name=key"`
	Operator          TolerationOperator `json:"operator,omitempty" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator"`
	Value             string             `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`
	Effect            TaintEffect        `json:"effect,omitempty" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect"`
	TolerationSeconds *int64             `json:"tolerationSeconds,omitempty" protobuf:"varint,5,opt,name=tolerationSeconds"`
}

type TolerationOperator string

const (
	TolerationOpExists TolerationOperator = "Exists"
	TolerationOpEqual  TolerationOperator = "Equal"
)

type PodReadinessGate struct {
	ConditionType PodConditionType `json:"conditionType" protobuf:"bytes,1,opt,name=conditionType,casttype=PodConditionType"`
}

type PodSpec struct {
	Volumes                       []Volume               `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
	InitContainers                []Container            `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
	Containers                    []Container            `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
	RestartPolicy                 RestartPolicy          `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy"`
	TerminationGracePeriodSeconds *int64                 `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds"`
	ActiveDeadlineSeconds         *int64                 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=activeDeadlineSeconds"`
	DNSPolicy                     DNSPolicy              `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy"`
	NodeSelector                  map[string]string      `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`
	ServiceAccountName            string                 `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"`
	DeprecatedServiceAccount      string                 `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount"`
	AutomountServiceAccountToken  *bool                  `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken"`
	NodeName                      string                 `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
	HostNetwork                   bool                   `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork"`
	HostPID                       bool                   `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`
	HostIPC                       bool                   `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC"`
	ShareProcessNamespace         *bool                  `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace"`
	SecurityContext               *PodSecurityContext    `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
	ImagePullSecrets              []LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
	Hostname                      string                 `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"`
	Subdomain                     string                 `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"`
	Affinity                      *Affinity              `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
	SchedulerName                 string                 `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"`
	Tolerations                   []Toleration           `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
	HostAliases                   []HostAlias            `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
	PriorityClassName             string                 `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName"`
	Priority                      *int32                 `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority"`
	DNSConfig                     *PodDNSConfig          `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig"`
	ReadinessGates                []PodReadinessGate     `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates"`
	RuntimeClassName              *string                `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
	EnableServiceLinks            *bool                  `json:"enableServiceLinks,omitempty" protobuf:"varint,30,opt,name=enableServiceLinks"`
}

const (
	DefaultEnableServiceLinks = true
)

type HostAlias struct {
	IP        string   `json:"ip,omitempty" protobuf:"bytes,1,opt,name=ip"`
	Hostnames []string `json:"hostnames,omitempty" protobuf:"bytes,2,rep,name=hostnames"`
}

type PodSecurityContext struct {
	SELinuxOptions     *SELinuxOptions `json:"seLinuxOptions,omitempty" protobuf:"bytes,1,opt,name=seLinuxOptions"`
	RunAsUser          *int64          `json:"runAsUser,omitempty" protobuf:"varint,2,opt,name=runAsUser"`
	RunAsGroup         *int64          `json:"runAsGroup,omitempty" protobuf:"varint,6,opt,name=runAsGroup"`
	RunAsNonRoot       *bool           `json:"runAsNonRoot,omitempty" protobuf:"varint,3,opt,name=runAsNonRoot"`
	SupplementalGroups []int64         `json:"supplementalGroups,omitempty" protobuf:"varint,4,rep,name=supplementalGroups"`
	FSGroup            *int64          `json:"fsGroup,omitempty" protobuf:"varint,5,opt,name=fsGroup"`
	Sysctls            []Sysctl        `json:"sysctls,omitempty" protobuf:"bytes,7,rep,name=sysctls"`
}

type PodQOSClass string

const (
	PodQOSGuaranteed PodQOSClass = "Guaranteed"
	PodQOSBurstable  PodQOSClass = "Burstable"
	PodQOSBestEffort PodQOSClass = "BestEffort"
)

type PodDNSConfig struct {
	Nameservers []string             `json:"nameservers,omitempty" protobuf:"bytes,1,rep,name=nameservers"`
	Searches    []string             `json:"searches,omitempty" protobuf:"bytes,2,rep,name=searches"`
	Options     []PodDNSConfigOption `json:"options,omitempty" protobuf:"bytes,3,rep,name=options"`
}

type PodDNSConfigOption struct {
	Name  string  `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Value *string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

type PodStatus struct {
	Phase                 PodPhase          `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=PodPhase"`
	Conditions            []PodCondition    `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
	Message               string            `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	Reason                string            `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	NominatedNodeName     string            `json:"nominatedNodeName,omitempty" protobuf:"bytes,11,opt,name=nominatedNodeName"`
	HostIP                string            `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP"`
	PodIP                 string            `json:"podIP,omitempty" protobuf:"bytes,6,opt,name=podIP"`
	StartTime             *Time             `json:"startTime,omitempty" protobuf:"bytes,7,opt,name=startTime"`
	InitContainerStatuses []ContainerStatus `json:"initContainerStatuses,omitempty" protobuf:"bytes,10,rep,name=initContainerStatuses"`
	ContainerStatuses     []ContainerStatus `json:"containerStatuses,omitempty" protobuf:"bytes,8,rep,name=containerStatuses"`
	QOSClass              PodQOSClass       `json:"qosClass,omitempty" protobuf:"bytes,9,rep,name=qosClass"`
}

type PodStatusResult struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Status     PodStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

type Pod struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       PodSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PodList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Pod `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type PodTemplateSpec struct {
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       PodSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type PodTemplate struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Template   PodTemplateSpec `json:"template,omitempty" protobuf:"bytes,2,opt,name=template"`
}

type PodTemplateList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []PodTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ReplicationControllerSpec struct {
	Replicas        *int32            `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	MinReadySeconds int32             `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds"`
	Selector        map[string]string `json:"selector,omitempty" protobuf:"bytes,2,rep,name=selector"`
	Template        *PodTemplateSpec  `json:"template,omitempty" protobuf:"bytes,3,opt,name=template"`
}

type ReplicationControllerStatus struct {
	Replicas             int32                            `json:"replicas" protobuf:"varint,1,opt,name=replicas"`
	FullyLabeledReplicas int32                            `json:"fullyLabeledReplicas,omitempty" protobuf:"varint,2,opt,name=fullyLabeledReplicas"`
	ReadyReplicas        int32                            `json:"readyReplicas,omitempty" protobuf:"varint,4,opt,name=readyReplicas"`
	AvailableReplicas    int32                            `json:"availableReplicas,omitempty" protobuf:"varint,5,opt,name=availableReplicas"`
	ObservedGeneration   int64                            `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration"`
	Conditions           []ReplicationControllerCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
}

type ReplicationControllerConditionType string

const (
	ReplicationControllerReplicaFailure ReplicationControllerConditionType = "ReplicaFailure"
)

type ReplicationControllerCondition struct {
	Type               ReplicationControllerConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ReplicationControllerConditionType"`
	Status             ConditionStatus                    `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	LastTransitionTime Time                               `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	Reason             string                             `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Message            string                             `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

type ReplicationController struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       ReplicationControllerSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     ReplicationControllerStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ReplicationControllerList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ReplicationController `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ServiceAffinity string

const (
	ServiceAffinityClientIP ServiceAffinity = "ClientIP"
	ServiceAffinityNone     ServiceAffinity = "None"
)
const DefaultClientIPServiceAffinitySeconds int32 = 10800

type SessionAffinityConfig struct {
	ClientIP *ClientIPConfig `json:"clientIP,omitempty" protobuf:"bytes,1,opt,name=clientIP"`
}

type ClientIPConfig struct {
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,1,opt,name=timeoutSeconds"`
}

type ServiceType string

const (
	ServiceTypeClusterIP    ServiceType = "ClusterIP"
	ServiceTypeNodePort     ServiceType = "NodePort"
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeExternalName ServiceType = "ExternalName"
)

type ServiceExternalTrafficPolicyType string

const (
	ServiceExternalTrafficPolicyTypeLocal   ServiceExternalTrafficPolicyType = "Local"
	ServiceExternalTrafficPolicyTypeCluster ServiceExternalTrafficPolicyType = "Cluster"
)

type ServiceStatus struct {
	LoadBalancer LoadBalancerStatus `json:"loadBalancer,omitempty" protobuf:"bytes,1,opt,name=loadBalancer"`
}

type LoadBalancerStatus struct {
	Ingress []LoadBalancerIngress `json:"ingress,omitempty" protobuf:"bytes,1,rep,name=ingress"`
}

type LoadBalancerIngress struct {
	IP       string `json:"ip,omitempty" protobuf:"bytes,1,opt,name=ip"`
	Hostname string `json:"hostname,omitempty" protobuf:"bytes,2,opt,name=hostname"`
}

type ServiceSpec struct {
	Ports                    []ServicePort                    `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
	Selector                 map[string]string                `json:"selector,omitempty" protobuf:"bytes,2,rep,name=selector"`
	ClusterIP                string                           `json:"clusterIP,omitempty" protobuf:"bytes,3,opt,name=clusterIP"`
	Type                     ServiceType                      `json:"type,omitempty" protobuf:"bytes,4,opt,name=type,casttype=ServiceType"`
	ExternalIPs              []string                         `json:"externalIPs,omitempty" protobuf:"bytes,5,rep,name=externalIPs"`
	SessionAffinity          ServiceAffinity                  `json:"sessionAffinity,omitempty" protobuf:"bytes,7,opt,name=sessionAffinity,casttype=ServiceAffinity"`
	LoadBalancerIP           string                           `json:"loadBalancerIP,omitempty" protobuf:"bytes,8,opt,name=loadBalancerIP"`
	LoadBalancerSourceRanges []string                         `json:"loadBalancerSourceRanges,omitempty" protobuf:"bytes,9,opt,name=loadBalancerSourceRanges"`
	ExternalName             string                           `json:"externalName,omitempty" protobuf:"bytes,10,opt,name=externalName"`
	ExternalTrafficPolicy    ServiceExternalTrafficPolicyType `json:"externalTrafficPolicy,omitempty" protobuf:"bytes,11,opt,name=externalTrafficPolicy"`
	HealthCheckNodePort      int32                            `json:"healthCheckNodePort,omitempty" protobuf:"bytes,12,opt,name=healthCheckNodePort"`
	PublishNotReadyAddresses bool                             `json:"publishNotReadyAddresses,omitempty" protobuf:"varint,13,opt,name=publishNotReadyAddresses"`
	SessionAffinityConfig    *SessionAffinityConfig           `json:"sessionAffinityConfig,omitempty" protobuf:"bytes,14,opt,name=sessionAffinityConfig"`
}

type ServicePort struct {
	Name       string      `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Protocol   Protocol    `json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol,casttype=Protocol"`
	Port       int32       `json:"port" protobuf:"varint,3,opt,name=port"`
	TargetPort json.Number `json:"targetPort,Number,omitempty" protobuf:"bytes,4,opt,name=targetPort"`
	NodePort   int32       `json:"nodePort,omitempty" protobuf:"varint,5,opt,name=nodePort"`
}

type Service struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       ServiceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     ServiceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

const (
	ClusterIPNone = "None"
)

type ServiceList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Service `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ServiceAccount struct {
	TypeMeta                     `json:",inline"`
	ObjectMeta                   `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Secrets                      []ObjectReference      `json:"secrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=secrets"`
	ImagePullSecrets             []LocalObjectReference `json:"imagePullSecrets,omitempty" protobuf:"bytes,3,rep,name=imagePullSecrets"`
	AutomountServiceAccountToken *bool                  `json:"automountServiceAccountToken,omitempty" protobuf:"varint,4,opt,name=automountServiceAccountToken"`
}

type ServiceAccountList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ServiceAccount `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Endpoints struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Subsets    []EndpointSubset `json:"subsets,omitempty" protobuf:"bytes,2,rep,name=subsets"`
}

type EndpointSubset struct {
	Addresses         []EndpointAddress `json:"addresses,omitempty" protobuf:"bytes,1,rep,name=addresses"`
	NotReadyAddresses []EndpointAddress `json:"notReadyAddresses,omitempty" protobuf:"bytes,2,rep,name=notReadyAddresses"`
	Ports             []EndpointPort    `json:"ports,omitempty" protobuf:"bytes,3,rep,name=ports"`
}

type EndpointAddress struct {
	IP        string           `json:"ip" protobuf:"bytes,1,opt,name=ip"`
	Hostname  string           `json:"hostname,omitempty" protobuf:"bytes,3,opt,name=hostname"`
	NodeName  *string          `json:"nodeName,omitempty" protobuf:"bytes,4,opt,name=nodeName"`
	TargetRef *ObjectReference `json:"targetRef,omitempty" protobuf:"bytes,2,opt,name=targetRef"`
}

type EndpointPort struct {
	Name     string   `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Port     int32    `json:"port" protobuf:"varint,2,opt,name=port"`
	Protocol Protocol `json:"protocol,omitempty" protobuf:"bytes,3,opt,name=protocol,casttype=Protocol"`
}

type EndpointsList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Endpoints `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type NodeSpec struct {
	PodCIDR             string            `json:"podCIDR,omitempty" protobuf:"bytes,1,opt,name=podCIDR"`
	ProviderID          string            `json:"providerID,omitempty" protobuf:"bytes,3,opt,name=providerID"`
	Unschedulable       bool              `json:"unschedulable,omitempty" protobuf:"varint,4,opt,name=unschedulable"`
	Taints              []Taint           `json:"taints,omitempty" protobuf:"bytes,5,opt,name=taints"`
	ConfigSource        *NodeConfigSource `json:"configSource,omitempty" protobuf:"bytes,6,opt,name=configSource"`
	DoNotUse_ExternalID string            `json:"externalID,omitempty" protobuf:"bytes,2,opt,name=externalID"`
}

type NodeConfigSource struct {
	ConfigMap *ConfigMapNodeConfigSource `json:"configMap,omitempty" protobuf:"bytes,2,opt,name=configMap"`
}

type ConfigMapNodeConfigSource struct {
	Namespace        string `json:"namespace" protobuf:"bytes,1,opt,name=namespace"`
	Name             string `json:"name" protobuf:"bytes,2,opt,name=name"`
	UID              string `json:"uid,omitempty" protobuf:"bytes,3,opt,name=uid"`
	ResourceVersion  string `json:"resourceVersion,omitempty" protobuf:"bytes,4,opt,name=resourceVersion"`
	KubeletConfigKey string `json:"kubeletConfigKey" protobuf:"bytes,5,opt,name=kubeletConfigKey"`
}

type DaemonEndpoint struct {
	/*
		The port tag was not properly in quotes in earlier releases, so it must be
		uppercased for backwards compat (since it was falling back to var name of
		'Port').
	*/
	Port int32 `json:"Port" protobuf:"varint,1,opt,name=Port"`
}

type NodeDaemonEndpoints struct {
	KubeletEndpoint DaemonEndpoint `json:"kubeletEndpoint,omitempty" protobuf:"bytes,1,opt,name=kubeletEndpoint"`
}

type NodeSystemInfo struct {
	MachineID               string `json:"machineID" protobuf:"bytes,1,opt,name=machineID"`
	SystemUUID              string `json:"systemUUID" protobuf:"bytes,2,opt,name=systemUUID"`
	BootID                  string `json:"bootID" protobuf:"bytes,3,opt,name=bootID"`
	KernelVersion           string `json:"kernelVersion" protobuf:"bytes,4,opt,name=kernelVersion"`
	OSImage                 string `json:"osImage" protobuf:"bytes,5,opt,name=osImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion" protobuf:"bytes,6,opt,name=containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion" protobuf:"bytes,7,opt,name=kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion" protobuf:"bytes,8,opt,name=kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem" protobuf:"bytes,9,opt,name=operatingSystem"`
	Architecture            string `json:"architecture" protobuf:"bytes,10,opt,name=architecture"`
}

type NodeConfigStatus struct {
	Assigned      *NodeConfigSource `json:"assigned,omitempty" protobuf:"bytes,1,opt,name=assigned"`
	Active        *NodeConfigSource `json:"active,omitempty" protobuf:"bytes,2,opt,name=active"`
	LastKnownGood *NodeConfigSource `json:"lastKnownGood,omitempty" protobuf:"bytes,3,opt,name=lastKnownGood"`
	Error         string            `json:"error,omitempty" protobuf:"bytes,4,opt,name=error"`
}

type NodeStatus struct {
	Capacity        ResourceList        `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	Allocatable     ResourceList        `json:"allocatable,omitempty" protobuf:"bytes,2,rep,name=allocatable,casttype=ResourceList,castkey=ResourceName"`
	Phase           NodePhase           `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=NodePhase"`
	Conditions      []NodeCondition     `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,4,rep,name=conditions"`
	Addresses       []NodeAddress       `json:"addresses,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,5,rep,name=addresses"`
	DaemonEndpoints NodeDaemonEndpoints `json:"daemonEndpoints,omitempty" protobuf:"bytes,6,opt,name=daemonEndpoints"`
	NodeInfo        NodeSystemInfo      `json:"nodeInfo,omitempty" protobuf:"bytes,7,opt,name=nodeInfo"`
	Images          []ContainerImage    `json:"images,omitempty" protobuf:"bytes,8,rep,name=images"`
	VolumesInUse    []UniqueVolumeName  `json:"volumesInUse,omitempty" protobuf:"bytes,9,rep,name=volumesInUse"`
	VolumesAttached []AttachedVolume    `json:"volumesAttached,omitempty" protobuf:"bytes,10,rep,name=volumesAttached"`
	Config          *NodeConfigStatus   `json:"config,omitempty" protobuf:"bytes,11,opt,name=config"`
}

type UniqueVolumeName string
type AttachedVolume struct {
	Name       UniqueVolumeName `json:"name" protobuf:"bytes,1,rep,name=name"`
	DevicePath string           `json:"devicePath" protobuf:"bytes,2,rep,name=devicePath"`
}

type AvoidPods struct {
	PreferAvoidPods []PreferAvoidPodsEntry `json:"preferAvoidPods,omitempty" protobuf:"bytes,1,rep,name=preferAvoidPods"`
}

type PreferAvoidPodsEntry struct {
	PodSignature PodSignature `json:"podSignature" protobuf:"bytes,1,opt,name=podSignature"`
	EvictionTime Time         `json:"evictionTime,omitempty" protobuf:"bytes,2,opt,name=evictionTime"`
	Reason       string       `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Message      string       `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
}

type PodSignature struct {
	PodController *OwnerReference `json:"podController,omitempty" protobuf:"bytes,1,opt,name=podController"`
}

type ContainerImage struct {
	Names     []string `json:"names" protobuf:"bytes,1,rep,name=names"`
	SizeBytes int64    `json:"sizeBytes,omitempty" protobuf:"varint,2,opt,name=sizeBytes"`
}

type NodePhase string

const (
	NodePending    NodePhase = "Pending"
	NodeRunning    NodePhase = "Running"
	NodeTerminated NodePhase = "Terminated"
)

type NodeConditionType string

const (
	NodeReady              NodeConditionType = "Ready"
	NodeOutOfDisk          NodeConditionType = "OutOfDisk"
	NodeMemoryPressure     NodeConditionType = "MemoryPressure"
	NodeDiskPressure       NodeConditionType = "DiskPressure"
	NodePIDPressure        NodeConditionType = "PIDPressure"
	NodeNetworkUnavailable NodeConditionType = "NetworkUnavailable"
)

type NodeCondition struct {
	Type               NodeConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeConditionType"`
	Status             ConditionStatus   `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	LastHeartbeatTime  Time              `json:"lastHeartbeatTime,omitempty" protobuf:"bytes,3,opt,name=lastHeartbeatTime"`
	LastTransitionTime Time              `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	Reason             string            `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	Message            string            `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type NodeAddressType string

const (
	NodeHostName    NodeAddressType = "Hostname"
	NodeExternalIP  NodeAddressType = "ExternalIP"
	NodeInternalIP  NodeAddressType = "InternalIP"
	NodeExternalDNS NodeAddressType = "ExternalDNS"
	NodeInternalDNS NodeAddressType = "InternalDNS"
)

type NodeAddress struct {
	Type    NodeAddressType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=NodeAddressType"`
	Address string          `json:"address" protobuf:"bytes,2,opt,name=address"`
}

type ResourceName string

const (
	ResourceCPU              ResourceName = "cpu"
	ResourceMemory           ResourceName = "memory"
	ResourceStorage          ResourceName = "storage"
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)
const (
	ResourceDefaultNamespacePrefix  = "kubernetes.io/"
	ResourceHugePagesPrefix         = "hugepages-"
	ResourceAttachableVolumesPrefix = "attachable-volumes-"
)

type ResourceList map[ResourceName]string
type Node struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NodeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Node `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type FinalizerName string

const (
	FinalizerKubernetes FinalizerName = "kubernetes"
)

type NamespaceSpec struct {
	Finalizers []FinalizerName `json:"finalizers,omitempty" protobuf:"bytes,1,rep,name=finalizers,casttype=FinalizerName"`
}

type NamespaceStatus struct {
	Phase NamespacePhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=NamespacePhase"`
}

type NamespacePhase string

const (
	NamespaceActive      NamespacePhase = "Active"
	NamespaceTerminating NamespacePhase = "Terminating"
)

type Namespace struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       NamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     NamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NamespaceList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Namespace `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Binding struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Target     ObjectReference `json:"target" protobuf:"bytes,2,opt,name=target"`
}

type Preconditions struct {
	UID *string `json:"uid,omitempty" protobuf:"bytes,1,opt,name=uid,casttype=k8s.io/apimachinery/pkg/string"`
}

type PodLogOptions struct {
	TypeMeta     `json:",inline"`
	Container    string `json:"container,omitempty" protobuf:"bytes,1,opt,name=container"`
	Follow       bool   `json:"follow,omitempty" protobuf:"varint,2,opt,name=follow"`
	Previous     bool   `json:"previous,omitempty" protobuf:"varint,3,opt,name=previous"`
	SinceSeconds *int64 `json:"sinceSeconds,omitempty" protobuf:"varint,4,opt,name=sinceSeconds"`
	SinceTime    *Time  `json:"sinceTime,omitempty" protobuf:"bytes,5,opt,name=sinceTime"`
	Timestamps   bool   `json:"timestamps,omitempty" protobuf:"varint,6,opt,name=timestamps"`
	TailLines    *int64 `json:"tailLines,omitempty" protobuf:"varint,7,opt,name=tailLines"`
	LimitBytes   *int64 `json:"limitBytes,omitempty" protobuf:"varint,8,opt,name=limitBytes"`
}

type PodAttachOptions struct {
	TypeMeta  `json:",inline"`
	Stdin     bool   `json:"stdin,omitempty" protobuf:"varint,1,opt,name=stdin"`
	Stdout    bool   `json:"stdout,omitempty" protobuf:"varint,2,opt,name=stdout"`
	Stderr    bool   `json:"stderr,omitempty" protobuf:"varint,3,opt,name=stderr"`
	TTY       bool   `json:"tty,omitempty" protobuf:"varint,4,opt,name=tty"`
	Container string `json:"container,omitempty" protobuf:"bytes,5,opt,name=container"`
}

type PodExecOptions struct {
	TypeMeta  `json:",inline"`
	Stdin     bool     `json:"stdin,omitempty" protobuf:"varint,1,opt,name=stdin"`
	Stdout    bool     `json:"stdout,omitempty" protobuf:"varint,2,opt,name=stdout"`
	Stderr    bool     `json:"stderr,omitempty" protobuf:"varint,3,opt,name=stderr"`
	TTY       bool     `json:"tty,omitempty" protobuf:"varint,4,opt,name=tty"`
	Container string   `json:"container,omitempty" protobuf:"bytes,5,opt,name=container"`
	Command   []string `json:"command" protobuf:"bytes,6,rep,name=command"`
}

type PodPortForwardOptions struct {
	TypeMeta `json:",inline"`
	Ports    []int32 `json:"ports,omitempty" protobuf:"varint,1,rep,name=ports"`
}

type PodProxyOptions struct {
	TypeMeta `json:",inline"`
	Path     string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
}

type NodeProxyOptions struct {
	TypeMeta `json:",inline"`
	Path     string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
}

type ServiceProxyOptions struct {
	TypeMeta `json:",inline"`
	Path     string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
}

type ObjectReference struct {
	Kind            string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	Namespace       string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
	Name            string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	UID             string `json:"uid,omitempty" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/string"`
	APIVersion      string `json:"apiVersion,omitempty" protobuf:"bytes,5,opt,name=apiVersion"`
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion"`
	FieldPath       string `json:"fieldPath,omitempty" protobuf:"bytes,7,opt,name=fieldPath"`
}

type LocalObjectReference struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

type TypedLocalObjectReference struct {
	APIGroup *string `json:"apiGroup" protobuf:"bytes,1,opt,name=apiGroup"`
	Kind     string  `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	Name     string  `json:"name" protobuf:"bytes,3,opt,name=name"`
}

type SerializedReference struct {
	TypeMeta  `json:",inline"`
	Reference ObjectReference `json:"reference,omitempty" protobuf:"bytes,1,opt,name=reference"`
}

type EventSource struct {
	Component string `json:"component,omitempty" protobuf:"bytes,1,opt,name=component"`
	Host      string `json:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
}

const (
	EventTypeNormal  string = "Normal"
	EventTypeWarning string = "Warning"
)

type Event struct {
	TypeMeta            `json:",inline"`
	ObjectMeta          `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	InvolvedObject      ObjectReference  `json:"involvedObject" protobuf:"bytes,2,opt,name=involvedObject"`
	Reason              string           `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Message             string           `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
	Source              EventSource      `json:"source,omitempty" protobuf:"bytes,5,opt,name=source"`
	FirstTimestamp      Time             `json:"firstTimestamp,omitempty" protobuf:"bytes,6,opt,name=firstTimestamp"`
	LastTimestamp       Time             `json:"lastTimestamp,omitempty" protobuf:"bytes,7,opt,name=lastTimestamp"`
	Count               int32            `json:"count,omitempty" protobuf:"varint,8,opt,name=count"`
	Type                string           `json:"type,omitempty" protobuf:"bytes,9,opt,name=type"`
	EventTime           MicroTime        `json:"eventTime,omitempty" protobuf:"bytes,10,opt,name=eventTime"`
	Series              *EventSeries     `json:"series,omitempty" protobuf:"bytes,11,opt,name=series"`
	Action              string           `json:"action,omitempty" protobuf:"bytes,12,opt,name=action"`
	Related             *ObjectReference `json:"related,omitempty" protobuf:"bytes,13,opt,name=related"`
	ReportingController string           `json:"reportingComponent" protobuf:"bytes,14,opt,name=reportingComponent"`
	ReportingInstance   string           `json:"reportingInstance" protobuf:"bytes,15,opt,name=reportingInstance"`
}

type EventSeries struct {
	Count            int32            `json:"count,omitempty" protobuf:"varint,1,name=count"`
	LastObservedTime MicroTime        `json:"lastObservedTime,omitempty" protobuf:"bytes,2,name=lastObservedTime"`
	State            EventSeriesState `json:"state,omitempty" protobuf:"bytes,3,name=state"`
}

type EventSeriesState string

const (
	EventSeriesStateOngoing  EventSeriesState = "Ongoing"
	EventSeriesStateFinished EventSeriesState = "Finished"
	EventSeriesStateUnknown  EventSeriesState = "Unknown"
)

type EventList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Event `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type LimitType string

const (
	LimitTypePod                   LimitType = "Pod"
	LimitTypeContainer             LimitType = "Container"
	LimitTypePersistentVolumeClaim LimitType = "PersistentVolumeClaim"
)

type LimitRangeItem struct {
	Type                 LimitType    `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=LimitType"`
	Max                  ResourceList `json:"max,omitempty" protobuf:"bytes,2,rep,name=max,casttype=ResourceList,castkey=ResourceName"`
	Min                  ResourceList `json:"min,omitempty" protobuf:"bytes,3,rep,name=min,casttype=ResourceList,castkey=ResourceName"`
	Default              ResourceList `json:"default,omitempty" protobuf:"bytes,4,rep,name=default,casttype=ResourceList,castkey=ResourceName"`
	DefaultRequest       ResourceList `json:"defaultRequest,omitempty" protobuf:"bytes,5,rep,name=defaultRequest,casttype=ResourceList,castkey=ResourceName"`
	MaxLimitRequestRatio ResourceList `json:"maxLimitRequestRatio,omitempty" protobuf:"bytes,6,rep,name=maxLimitRequestRatio,casttype=ResourceList,castkey=ResourceName"`
}

type LimitRangeSpec struct {
	Limits []LimitRangeItem `json:"limits" protobuf:"bytes,1,rep,name=limits"`
}

type LimitRange struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       LimitRangeSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type LimitRangeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []LimitRange `json:"items" protobuf:"bytes,2,rep,name=items"`
}

const (
	ResourcePods                     ResourceName = "pods"
	ResourceServices                 ResourceName = "services"
	ResourceReplicationControllers   ResourceName = "replicationcontrollers"
	ResourceQuotas                   ResourceName = "resourcequotas"
	ResourceSecrets                  ResourceName = "secrets"
	ResourceConfigMaps               ResourceName = "configmaps"
	ResourcePersistentVolumeClaims   ResourceName = "persistentvolumeclaims"
	ResourceServicesNodePorts        ResourceName = "services.nodeports"
	ResourceServicesLoadBalancers    ResourceName = "services.loadbalancers"
	ResourceRequestsCPU              ResourceName = "requests.cpu"
	ResourceRequestsMemory           ResourceName = "requests.memory"
	ResourceRequestsStorage          ResourceName = "requests.storage"
	ResourceRequestsEphemeralStorage ResourceName = "requests.ephemeral-storage"
	ResourceLimitsCPU                ResourceName = "limits.cpu"
	ResourceLimitsMemory             ResourceName = "limits.memory"
	ResourceLimitsEphemeralStorage   ResourceName = "limits.ephemeral-storage"
)
const (
	ResourceRequestsHugePagesPrefix = "requests.hugepages-"
	DefaultResourceRequestsPrefix   = "requests."
)

type ResourceQuotaScope string

const (
	ResourceQuotaScopeTerminating    ResourceQuotaScope = "Terminating"
	ResourceQuotaScopeNotTerminating ResourceQuotaScope = "NotTerminating"
	ResourceQuotaScopeBestEffort     ResourceQuotaScope = "BestEffort"
	ResourceQuotaScopeNotBestEffort  ResourceQuotaScope = "NotBestEffort"
	ResourceQuotaScopePriorityClass  ResourceQuotaScope = "PriorityClass"
)

type ResourceQuotaSpec struct {
	Hard          ResourceList         `json:"hard,omitempty" protobuf:"bytes,1,rep,name=hard,casttype=ResourceList,castkey=ResourceName"`
	Scopes        []ResourceQuotaScope `json:"scopes,omitempty" protobuf:"bytes,2,rep,name=scopes,casttype=ResourceQuotaScope"`
	ScopeSelector *ScopeSelector       `json:"scopeSelector,omitempty" protobuf:"bytes,3,opt,name=scopeSelector"`
}

type ScopeSelector struct {
	MatchExpressions []ScopedResourceSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,1,rep,name=matchExpressions"`
}

type ScopedResourceSelectorRequirement struct {
	ScopeName ResourceQuotaScope    `json:"scopeName" protobuf:"bytes,1,opt,name=scopeName"`
	Operator  ScopeSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=ScopedResourceSelectorOperator"`
	Values    []string              `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type ScopeSelectorOperator string

const (
	ScopeSelectorOpIn           ScopeSelectorOperator = "In"
	ScopeSelectorOpNotIn        ScopeSelectorOperator = "NotIn"
	ScopeSelectorOpExists       ScopeSelectorOperator = "Exists"
	ScopeSelectorOpDoesNotExist ScopeSelectorOperator = "DoesNotExist"
)

type ResourceQuotaStatus struct {
	Hard ResourceList `json:"hard,omitempty" protobuf:"bytes,1,rep,name=hard,casttype=ResourceList,castkey=ResourceName"`
	Used ResourceList `json:"used,omitempty" protobuf:"bytes,2,rep,name=used,casttype=ResourceList,castkey=ResourceName"`
}

type ResourceQuota struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       ResourceQuotaSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     ResourceQuotaStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ResourceQuotaList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ResourceQuota `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Secret struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Data       map[string][]byte `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
	StringData map[string]string `json:"stringData,omitempty" protobuf:"bytes,4,rep,name=stringData"`
	Type       SecretType        `json:"type,omitempty" protobuf:"bytes,3,opt,name=type,casttype=SecretType"`
}

const MaxSecretSize = 1 * 1024 * 1024

type SecretType string

const (
	SecretTypeOpaque              SecretType = "Opaque"
	SecretTypeServiceAccountToken SecretType = "kubernetes.io/service-account-token"
	ServiceAccountNameKey                    = "kubernetes.io/service-account.name"
	ServiceAccountUIDKey                     = "kubernetes.io/service-account.uid"
	ServiceAccountTokenKey                   = "token"
	ServiceAccountKubeconfigKey              = "kubernetes.kubeconfig"
	ServiceAccountRootCAKey                  = "ca.crt"
	ServiceAccountNamespaceKey               = "namespace"
	SecretTypeDockercfg           SecretType = "kubernetes.io/dockercfg"
	DockerConfigKey                          = ".dockercfg"
	SecretTypeDockerConfigJson    SecretType = "kubernetes.io/dockerconfigjson"
	DockerConfigJsonKey                      = ".dockerconfigjson"
	SecretTypeBasicAuth           SecretType = "kubernetes.io/basic-auth"
	BasicAuthUsernameKey                     = "username"
	BasicAuthPasswordKey                     = "password"
	SecretTypeSSHAuth             SecretType = "kubernetes.io/ssh-auth"
	SSHAuthPrivateKey                        = "ssh-privatekey"
	SecretTypeTLS                 SecretType = "kubernetes.io/tls"
	TLSCertKey                               = "tls.crt"
	TLSPrivateKeyKey                         = "tls.key"
	SecretTypeBootstrapToken      SecretType = "bootstrap.kubernetes.io/token"
)

type SecretList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Secret `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ConfigMap struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Data       map[string]string `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
	BinaryData map[string][]byte `json:"binaryData,omitempty" protobuf:"bytes,3,rep,name=binaryData"`
}

type ConfigMapList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ConfigMap `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ComponentConditionType string

const (
	ComponentHealthy ComponentConditionType = "Healthy"
)

type ComponentCondition struct {
	Type    ComponentConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ComponentConditionType"`
	Status  ConditionStatus        `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	Message string                 `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	Error   string                 `json:"error,omitempty" protobuf:"bytes,4,opt,name=error"`
}

type ComponentStatus struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Conditions []ComponentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
}

type ComponentStatusList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ComponentStatus `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type DownwardAPIVolumeSource struct {
	Items       []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"`
	DefaultMode *int32                  `json:"defaultMode,omitempty" protobuf:"varint,2,opt,name=defaultMode"`
}

const (
	DownwardAPIVolumeSourceDefaultMode int32 = 0644
)

type DownwardAPIVolumeFile struct {
	Path             string                 `json:"path" protobuf:"bytes,1,opt,name=path"`
	FieldRef         *ObjectFieldSelector   `json:"fieldRef,omitempty" protobuf:"bytes,2,opt,name=fieldRef"`
	ResourceFieldRef *ResourceFieldSelector `json:"resourceFieldRef,omitempty" protobuf:"bytes,3,opt,name=resourceFieldRef"`
	Mode             *int32                 `json:"mode,omitempty" protobuf:"varint,4,opt,name=mode"`
}

type DownwardAPIProjection struct {
	Items []DownwardAPIVolumeFile `json:"items,omitempty" protobuf:"bytes,1,rep,name=items"`
}

type SecurityContext struct {
	Capabilities             *Capabilities   `json:"capabilities,omitempty" protobuf:"bytes,1,opt,name=capabilities"`
	Privileged               *bool           `json:"privileged,omitempty" protobuf:"varint,2,opt,name=privileged"`
	SELinuxOptions           *SELinuxOptions `json:"seLinuxOptions,omitempty" protobuf:"bytes,3,opt,name=seLinuxOptions"`
	RunAsUser                *int64          `json:"runAsUser,omitempty" protobuf:"varint,4,opt,name=runAsUser"`
	RunAsGroup               *int64          `json:"runAsGroup,omitempty" protobuf:"varint,8,opt,name=runAsGroup"`
	RunAsNonRoot             *bool           `json:"runAsNonRoot,omitempty" protobuf:"varint,5,opt,name=runAsNonRoot"`
	ReadOnlyRootFilesystem   *bool           `json:"readOnlyRootFilesystem,omitempty" protobuf:"varint,6,opt,name=readOnlyRootFilesystem"`
	AllowPrivilegeEscalation *bool           `json:"allowPrivilegeEscalation,omitempty" protobuf:"varint,7,opt,name=allowPrivilegeEscalation"`
	ProcMount                *ProcMountType  `json:"procMount,omitempty" protobuf:"bytes,9,opt,name=procMount"`
}

type ProcMountType string

const (
	DefaultProcMount  ProcMountType = "Default"
	UnmaskedProcMount ProcMountType = "Unmasked"
)

type SELinuxOptions struct {
	User  string `json:"user,omitempty" protobuf:"bytes,1,opt,name=user"`
	Role  string `json:"role,omitempty" protobuf:"bytes,2,opt,name=role"`
	Type  string `json:"type,omitempty" protobuf:"bytes,3,opt,name=type"`
	Level string `json:"level,omitempty" protobuf:"bytes,4,opt,name=level"`
}

type RangeAllocation struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Range      string `json:"range" protobuf:"bytes,2,opt,name=range"`
	Data       []byte `json:"data" protobuf:"bytes,3,opt,name=data"`
}

const (
	DefaultSchedulerName                        = "default-scheduler"
	DefaultHardPodAffinitySymmetricWeight int32 = 1
)

type Sysctl struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

type NodeResources struct {
	Capacity ResourceList `protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
}

const (
	ExecStdinParam             = "input"
	ExecStdoutParam            = "output"
	ExecStderrParam            = "error"
	ExecTTYParam               = "tty"
	ExecCommandParam           = "command"
	StreamType                 = "streamType"
	StreamTypeStdin            = "stdin"
	StreamTypeStdout           = "stdout"
	StreamTypeStderr           = "stderr"
	StreamTypeData             = "data"
	StreamTypeError            = "error"
	StreamTypeResize           = "resize"
	PortHeader                 = "port"
	PortForwardRequestIDHeader = "requestID"
)

const (
	ControllerRevisionHashLabelKey = "controller-revision-hash"
	StatefulSetRevisionLabel       = ControllerRevisionHashLabelKey
	DeprecatedRollbackTo           = "deprecated.deployment.rollback.to"
	DeprecatedTemplateGeneration   = "deprecated.daemonset.template.generation"
	StatefulSetPodNameLabel        = "statefulset.kubernetes.io/pod-name"
)

type Time struct {
	time.Time `protobuf:"-"`
}

type IntOrString struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,casttype=Type"`
	IntVal int32  `protobuf:"varint,2,opt,name=intVal"`
	StrVal string `protobuf:"bytes,3,opt,name=strVal"`
}

type Type int
type RawExtension struct {
	Raw    []byte `protobuf:"bytes,1,opt,name=raw"`
	Object Object `json:"-"`
}

type Object interface {
	GetObjectKind() ObjectKind
	DeepCopyObject() Object
}

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

type ObjectKind interface {
	SetGroupVersionKind(kind GroupVersionKind)
	GroupVersionKind() GroupVersionKind
}

var EmptyObjectKind = emptyObjectKind{}

type emptyObjectKind struct{}

func (emptyObjectKind) SetGroupVersionKind(gvk GroupVersionKind) {}

func (emptyObjectKind) GroupVersionKind() GroupVersionKind {
	return GroupVersionKind{}
}

type StatefulSet struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       StatefulSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     StatefulSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PodManagementPolicyType string

const (
	OrderedReadyPodManagement PodManagementPolicyType = "OrderedReady"
	ParallelPodManagement                             = "Parallel"
)

type StatefulSetUpdateStrategy struct {
	Type          StatefulSetUpdateStrategyType     `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=StatefulSetStrategyType"`
	RollingUpdate *RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate"`
}

type StatefulSetUpdateStrategyType string

const (
	RollingUpdateStatefulSetStrategyType = "RollingUpdate"
	OnDeleteStatefulSetStrategyType      = "OnDelete"
)

type RollingUpdateStatefulSetStrategy struct {
	Partition *int32 `json:"partition,omitempty" protobuf:"varint,1,opt,name=partition"`
}

type StatefulSetSpec struct {
	Replicas             *int32                    `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	Selector             *LabelSelector            `json:"selector" protobuf:"bytes,2,opt,name=selector"`
	Template             PodTemplateSpec           `json:"template" protobuf:"bytes,3,opt,name=template"`
	VolumeClaimTemplates []PersistentVolumeClaim   `json:"volumeClaimTemplates,omitempty" protobuf:"bytes,4,rep,name=volumeClaimTemplates"`
	ServiceName          string                    `json:"serviceName" protobuf:"bytes,5,opt,name=serviceName"`
	PodManagementPolicy  PodManagementPolicyType   `json:"podManagementPolicy,omitempty" protobuf:"bytes,6,opt,name=podManagementPolicy,casttype=PodManagementPolicyType"`
	UpdateStrategy       StatefulSetUpdateStrategy `json:"updateStrategy,omitempty" protobuf:"bytes,7,opt,name=updateStrategy"`
	RevisionHistoryLimit *int32                    `json:"revisionHistoryLimit,omitempty" protobuf:"varint,8,opt,name=revisionHistoryLimit"`
}

type StatefulSetStatus struct {
	ObservedGeneration int64                  `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	Replicas           int32                  `json:"replicas" protobuf:"varint,2,opt,name=replicas"`
	ReadyReplicas      int32                  `json:"readyReplicas,omitempty" protobuf:"varint,3,opt,name=readyReplicas"`
	CurrentReplicas    int32                  `json:"currentReplicas,omitempty" protobuf:"varint,4,opt,name=currentReplicas"`
	UpdatedReplicas    int32                  `json:"updatedReplicas,omitempty" protobuf:"varint,5,opt,name=updatedReplicas"`
	CurrentRevision    string                 `json:"currentRevision,omitempty" protobuf:"bytes,6,opt,name=currentRevision"`
	UpdateRevision     string                 `json:"updateRevision,omitempty" protobuf:"bytes,7,opt,name=updateRevision"`
	CollisionCount     *int32                 `json:"collisionCount,omitempty" protobuf:"varint,9,opt,name=collisionCount"`
	Conditions         []StatefulSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,10,rep,name=conditions"`
}

type StatefulSetConditionType string
type StatefulSetCondition struct {
	Type               StatefulSetConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=StatefulSetConditionType"`
	Status             ConditionStatus          `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/ConditionStatus"`
	LastTransitionTime Time                     `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	Reason             string                   `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Message            string                   `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

type StatefulSetList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []StatefulSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Deployment struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       DeploymentSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     DeploymentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DeploymentSpec struct {
	Replicas                *int32             `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	Selector                *LabelSelector     `json:"selector" protobuf:"bytes,2,opt,name=selector"`
	Template                PodTemplateSpec    `json:"template" protobuf:"bytes,3,opt,name=template"`
	Strategy                DeploymentStrategy `json:"strategy,omitempty" patchStrategy:"retainKeys" protobuf:"bytes,4,opt,name=strategy"`
	MinReadySeconds         int32              `json:"minReadySeconds,omitempty" protobuf:"varint,5,opt,name=minReadySeconds"`
	RevisionHistoryLimit    *int32             `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit"`
	Paused                  bool               `json:"paused,omitempty" protobuf:"varint,7,opt,name=paused"`
	ProgressDeadlineSeconds *int32             `json:"progressDeadlineSeconds,omitempty" protobuf:"varint,9,opt,name=progressDeadlineSeconds"`
}

const (
	DefaultDeploymentUniqueLabelKey string = "pod-template-hash"
)

type DeploymentStrategy struct {
	Type          DeploymentStrategyType   `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=DeploymentStrategyType"`
	RollingUpdate *RollingUpdateDeployment `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate"`
}

type DeploymentStrategyType string

const (
	RecreateDeploymentStrategyType      DeploymentStrategyType = "Recreate"
	RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

type RollingUpdateDeployment struct {
	MaxUnavailable json.Number `json:"maxUnavailable,Number,omitempty" protobuf:"bytes,1,opt,name=maxUnavailable"`
	MaxSurge       json.Number `json:"maxSurge,Number,omitempty" protobuf:"bytes,2,opt,name=maxSurge"`
}

type DeploymentStatus struct {
	ObservedGeneration  int64                 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	Replicas            int32                 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas"`
	UpdatedReplicas     int32                 `json:"updatedReplicas,omitempty" protobuf:"varint,3,opt,name=updatedReplicas"`
	ReadyReplicas       int32                 `json:"readyReplicas,omitempty" protobuf:"varint,7,opt,name=readyReplicas"`
	AvailableReplicas   int32                 `json:"availableReplicas,omitempty" protobuf:"varint,4,opt,name=availableReplicas"`
	UnavailableReplicas int32                 `json:"unavailableReplicas,omitempty" protobuf:"varint,5,opt,name=unavailableReplicas"`
	Conditions          []DeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
	CollisionCount      *int32                `json:"collisionCount,omitempty" protobuf:"varint,8,opt,name=collisionCount"`
}

type DeploymentConditionType string

const (
	DeploymentAvailable      DeploymentConditionType = "Available"
	DeploymentProgressing    DeploymentConditionType = "Progressing"
	DeploymentReplicaFailure DeploymentConditionType = "ReplicaFailure"
)

type DeploymentCondition struct {
	Type               DeploymentConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=DeploymentConditionType"`
	Status             ConditionStatus         `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/ConditionStatus"`
	LastUpdateTime     Time                    `json:"lastUpdateTime,omitempty" protobuf:"bytes,6,opt,name=lastUpdateTime"`
	LastTransitionTime Time                    `json:"lastTransitionTime,omitempty" protobuf:"bytes,7,opt,name=lastTransitionTime"`
	Reason             string                  `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Message            string                  `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

type DeploymentList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Deployment `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type DaemonSetUpdateStrategy struct {
	Type          DaemonSetUpdateStrategyType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	RollingUpdate *RollingUpdateDaemonSet     `json:"rollingUpdate,omitempty" protobuf:"bytes,2,opt,name=rollingUpdate"`
}

type DaemonSetUpdateStrategyType string

const (
	RollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "RollingUpdate"
	OnDeleteDaemonSetStrategyType      DaemonSetUpdateStrategyType = "OnDelete"
)

type RollingUpdateDaemonSet struct {
	MaxUnavailable json.Number `json:"maxUnavailable,Number,omitempty" protobuf:"bytes,1,opt,name=maxUnavailable"`
}

type DaemonSetSpec struct {
	Selector             *LabelSelector          `json:"selector" protobuf:"bytes,1,opt,name=selector"`
	Template             PodTemplateSpec         `json:"template" protobuf:"bytes,2,opt,name=template"`
	UpdateStrategy       DaemonSetUpdateStrategy `json:"updateStrategy,omitempty" protobuf:"bytes,3,opt,name=updateStrategy"`
	MinReadySeconds      int32                   `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds"`
	RevisionHistoryLimit *int32                  `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit"`
}

type DaemonSetStatus struct {
	CurrentNumberScheduled int32                `json:"currentNumberScheduled" protobuf:"varint,1,opt,name=currentNumberScheduled"`
	NumberMisscheduled     int32                `json:"numberMisscheduled" protobuf:"varint,2,opt,name=numberMisscheduled"`
	DesiredNumberScheduled int32                `json:"desiredNumberScheduled" protobuf:"varint,3,opt,name=desiredNumberScheduled"`
	NumberReady            int32                `json:"numberReady" protobuf:"varint,4,opt,name=numberReady"`
	ObservedGeneration     int64                `json:"observedGeneration,omitempty" protobuf:"varint,5,opt,name=observedGeneration"`
	UpdatedNumberScheduled int32                `json:"updatedNumberScheduled,omitempty" protobuf:"varint,6,opt,name=updatedNumberScheduled"`
	NumberAvailable        int32                `json:"numberAvailable,omitempty" protobuf:"varint,7,opt,name=numberAvailable"`
	NumberUnavailable      int32                `json:"numberUnavailable,omitempty" protobuf:"varint,8,opt,name=numberUnavailable"`
	CollisionCount         *int32               `json:"collisionCount,omitempty" protobuf:"varint,9,opt,name=collisionCount"`
	Conditions             []DaemonSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,10,rep,name=conditions"`
}

type DaemonSetConditionType string

type DaemonSetCondition struct {
	Type               DaemonSetConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=DaemonSetConditionType"`
	Status             ConditionStatus        `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/ConditionStatus"`
	LastTransitionTime Time                   `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	Reason             string                 `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Message            string                 `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

type DaemonSet struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       DaemonSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     DaemonSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

const (
	DefaultDaemonSetUniqueLabelKey = ControllerRevisionHashLabelKey
)

type DaemonSetList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []DaemonSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ReplicaSet struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec       ReplicaSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     ReplicaSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ReplicaSetList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ReplicaSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ReplicaSetSpec struct {
	Replicas        *int32          `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	MinReadySeconds int32           `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds"`
	Selector        *LabelSelector  `json:"selector" protobuf:"bytes,2,opt,name=selector"`
	Template        PodTemplateSpec `json:"template,omitempty" protobuf:"bytes,3,opt,name=template"`
}

type ReplicaSetStatus struct {
	Replicas             int32                 `json:"replicas" protobuf:"varint,1,opt,name=replicas"`
	FullyLabeledReplicas int32                 `json:"fullyLabeledReplicas,omitempty" protobuf:"varint,2,opt,name=fullyLabeledReplicas"`
	ReadyReplicas        int32                 `json:"readyReplicas,omitempty" protobuf:"varint,4,opt,name=readyReplicas"`
	AvailableReplicas    int32                 `json:"availableReplicas,omitempty" protobuf:"varint,5,opt,name=availableReplicas"`
	ObservedGeneration   int64                 `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration"`
	Conditions           []ReplicaSetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
}

type ReplicaSetConditionType string

const (
	ReplicaSetReplicaFailure ReplicaSetConditionType = "ReplicaFailure"
)

type ReplicaSetCondition struct {
	Type               ReplicaSetConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ReplicaSetConditionType"`
	Status             ConditionStatus         `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/ConditionStatus"`
	LastTransitionTime Time                    `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	Reason             string                  `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	Message            string                  `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

type ControllerRevision struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Data       RawExtension `json:"data,omitempty" protobuf:"bytes,2,opt,name=data"`
	Revision   int64        `json:"revision" protobuf:"varint,3,opt,name=revision"`
}

type ControllerRevisionList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []ControllerRevision `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type TypeMeta struct {
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

type ListMeta struct {
	SelfLink        string `json:"selfLink,omitempty" protobuf:"bytes,1,opt,name=selfLink"`
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,2,opt,name=resourceVersion"`
	Continue        string `json:"continue,omitempty" protobuf:"bytes,3,opt,name=continue"`
}

const (
	FinalizerOrphanDependents string = "orphan"
	FinalizerDeleteDependents string = "foregroundDeletion"
)

type ObjectMeta struct {
	Name                       string            `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	GenerateName               string            `json:"generateName,omitempty" protobuf:"bytes,2,opt,name=generateName"`
	Namespace                  string            `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	SelfLink                   string            `json:"selfLink,omitempty" protobuf:"bytes,4,opt,name=selfLink"`
	UID                        string            `json:"uid,omitempty" protobuf:"bytes,5,opt,name=uid,casttype=k8s.io/kubernetes/pkg/types.UID"`
	ResourceVersion            string            `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion"`
	Generation                 int64             `json:"generation,omitempty" protobuf:"varint,7,opt,name=generation"`
	CreationTimestamp          Time              `json:"creationTimestamp,omitempty" protobuf:"bytes,8,opt,name=creationTimestamp"`
	DeletionTimestamp          *Time             `json:"deletionTimestamp,omitempty" protobuf:"bytes,9,opt,name=deletionTimestamp"`
	DeletionGracePeriodSeconds *int64            `json:"deletionGracePeriodSeconds,omitempty" protobuf:"varint,10,opt,name=deletionGracePeriodSeconds"`
	Labels                     map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`
	Annotations                map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`
	OwnerReferences            []OwnerReference  `json:"ownerReferences,omitempty" patchStrategy:"merge" patchMergeKey:"uid" protobuf:"bytes,13,rep,name=ownerReferences"`
	Initializers               *Initializers     `json:"initializers,omitempty" protobuf:"bytes,16,opt,name=initializers"`
	Finalizers                 []string          `json:"finalizers,omitempty" patchStrategy:"merge" protobuf:"bytes,14,rep,name=finalizers"`
	ClusterName                string            `json:"clusterName,omitempty" protobuf:"bytes,15,opt,name=clusterName"`
}

type Initializers struct {
	Pending []Initializer `json:"pending" protobuf:"bytes,1,rep,name=pending" patchStrategy:"merge" patchMergeKey:"name"`
	Result  *Status       `json:"result,omitempty" protobuf:"bytes,2,opt,name=result"`
}

type Initializer struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
}

const (
	NamespaceDefault string = "default"
	NamespaceAll     string = ""
	NamespaceNone    string = ""
	NamespaceSystem  string = "kube-system"
	NamespacePublic  string = "kube-public"
)

type OwnerReference struct {
	APIVersion         string `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion"`
	Kind               string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	Name               string `json:"name" protobuf:"bytes,3,opt,name=name"`
	UID                string `json:"uid" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/string"`
	Controller         *bool  `json:"controller,omitempty" protobuf:"varint,6,opt,name=controller"`
	BlockOwnerDeletion *bool  `json:"blockOwnerDeletion,omitempty" protobuf:"varint,7,opt,name=blockOwnerDeletion"`
}

type ListOptions struct {
	TypeMeta             `json:",inline"`
	LabelSelector        string `json:"labelSelector,omitempty" protobuf:"bytes,1,opt,name=labelSelector"`
	FieldSelector        string `json:"fieldSelector,omitempty" protobuf:"bytes,2,opt,name=fieldSelector"`
	IncludeUninitialized bool   `json:"includeUninitialized,omitempty" protobuf:"varint,6,opt,name=includeUninitialized"`
	Watch                bool   `json:"watch,omitempty" protobuf:"varint,3,opt,name=watch"`
	ResourceVersion      string `json:"resourceVersion,omitempty" protobuf:"bytes,4,opt,name=resourceVersion"`
	TimeoutSeconds       *int64 `json:"timeoutSeconds,omitempty" protobuf:"varint,5,opt,name=timeoutSeconds"`
	Limit                int64  `json:"limit,omitempty" protobuf:"varint,7,opt,name=limit"`
	Continue             string `json:"continue,omitempty" protobuf:"bytes,8,opt,name=continue"`
}

type ExportOptions struct {
	TypeMeta `json:",inline"`
	Export   bool `json:"export" protobuf:"varint,1,opt,name=export"`
	Exact    bool `json:"exact" protobuf:"varint,2,opt,name=exact"`
}

type GetOptions struct {
	TypeMeta             `json:",inline"`
	ResourceVersion      string `json:"resourceVersion,omitempty" protobuf:"bytes,1,opt,name=resourceVersion"`
	IncludeUninitialized bool   `json:"includeUninitialized,omitempty" protobuf:"varint,2,opt,name=includeUninitialized"`
}

type DeletionPropagation string

const (
	DeletePropagationOrphan     DeletionPropagation = "Orphan"
	DeletePropagationBackground DeletionPropagation = "Background"
	DeletePropagationForeground DeletionPropagation = "Foreground"
)
const (
	DryRunAll = "All"
)

type DeleteOptions struct {
	TypeMeta           `json:",inline"`
	GracePeriodSeconds *int64               `json:"gracePeriodSeconds,omitempty" protobuf:"varint,1,opt,name=gracePeriodSeconds"`
	Preconditions      *Preconditions       `json:"preconditions,omitempty" protobuf:"bytes,2,opt,name=preconditions"`
	OrphanDependents   *bool                `json:"orphanDependents,omitempty" protobuf:"varint,3,opt,name=orphanDependents"`
	PropagationPolicy  *DeletionPropagation `json:"propagationPolicy,omitempty" protobuf:"varint,4,opt,name=propagationPolicy"`
	DryRun             []string             `json:"dryRun,omitempty" protobuf:"bytes,5,rep,name=dryRun"`
}

type CreateOptions struct {
	TypeMeta             `json:",inline"`
	DryRun               []string `json:"dryRun,omitempty" protobuf:"bytes,1,rep,name=dryRun"`
	IncludeUninitialized bool     `json:"includeUninitialized,omitempty" protobuf:"varint,2,opt,name=includeUninitialized"`
}

type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty" protobuf:"bytes,1,rep,name=dryRun"`
}

type Status struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Status   string         `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
	Message  string         `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	Reason   StatusReason   `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason,casttype=StatusReason"`
	Details  *StatusDetails `json:"details,omitempty" protobuf:"bytes,5,opt,name=details"`
	Code     int32          `json:"code,omitempty" protobuf:"varint,6,opt,name=code"`
}

type StatusDetails struct {
	Name              string        `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Group             string        `json:"group,omitempty" protobuf:"bytes,2,opt,name=group"`
	Kind              string        `json:"kind,omitempty" protobuf:"bytes,3,opt,name=kind"`
	UID               string        `json:"uid,omitempty" protobuf:"bytes,6,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`
	Causes            []StatusCause `json:"causes,omitempty" protobuf:"bytes,4,rep,name=causes"`
	RetryAfterSeconds int32         `json:"retryAfterSeconds,omitempty" protobuf:"varint,5,opt,name=retryAfterSeconds"`
}

const (
	StatusSuccess = "Success"
	StatusFailure = "Failure"
)

type StatusReason string

const (
	StatusReasonUnknown              StatusReason = ""
	StatusReasonUnauthorized         StatusReason = "Unauthorized"
	StatusReasonForbidden            StatusReason = "Forbidden"
	StatusReasonNotFound             StatusReason = "NotFound"
	StatusReasonAlreadyExists        StatusReason = "AlreadyExists"
	StatusReasonConflict             StatusReason = "Conflict"
	StatusReasonGone                 StatusReason = "Gone"
	StatusReasonInvalid              StatusReason = "Invalid"
	StatusReasonServerTimeout        StatusReason = "ServerTimeout"
	StatusReasonTimeout              StatusReason = "Timeout"
	StatusReasonTooManyRequests      StatusReason = "TooManyRequests"
	StatusReasonBadRequest           StatusReason = "BadRequest"
	StatusReasonMethodNotAllowed     StatusReason = "MethodNotAllowed"
	StatusReasonNotAcceptable        StatusReason = "NotAcceptable"
	StatusReasonUnsupportedMediaType StatusReason = "UnsupportedMediaType"
	StatusReasonInternalError        StatusReason = "InternalError"
	StatusReasonExpired              StatusReason = "Expired"
	StatusReasonServiceUnavailable   StatusReason = "ServiceUnavailable"
)

type StatusCause struct {
	Type    CauseType `json:"reason,omitempty" protobuf:"bytes,1,opt,name=reason,casttype=CauseType"`
	Message string    `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	Field   string    `json:"field,omitempty" protobuf:"bytes,3,opt,name=field"`
}

type CauseType string

const (
	CauseTypeFieldValueNotFound       CauseType = "FieldValueNotFound"
	CauseTypeFieldValueRequired       CauseType = "FieldValueRequired"
	CauseTypeFieldValueDuplicate      CauseType = "FieldValueDuplicate"
	CauseTypeFieldValueInvalid        CauseType = "FieldValueInvalid"
	CauseTypeFieldValueNotSupported   CauseType = "FieldValueNotSupported"
	CauseTypeUnexpectedServerResponse CauseType = "UnexpectedServerResponse"
)

type List struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []RawExtension `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type APIVersions struct {
	TypeMeta                   `json:",inline"`
	Versions                   []string                    `json:"versions" protobuf:"bytes,1,rep,name=versions"`
	ServerAddressByClientCIDRs []ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs" protobuf:"bytes,2,rep,name=serverAddressByClientCIDRs"`
}

type APIGroupList struct {
	TypeMeta `json:",inline"`
	Groups   []APIGroup `json:"groups" protobuf:"bytes,1,rep,name=groups"`
}

type APIGroup struct {
	TypeMeta                   `json:",inline"`
	Name                       string                      `json:"name" protobuf:"bytes,1,opt,name=name"`
	Versions                   []GroupVersionForDiscovery  `json:"versions" protobuf:"bytes,2,rep,name=versions"`
	PreferredVersion           GroupVersionForDiscovery    `json:"preferredVersion,omitempty" protobuf:"bytes,3,opt,name=preferredVersion"`
	ServerAddressByClientCIDRs []ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs,omitempty" protobuf:"bytes,4,rep,name=serverAddressByClientCIDRs"`
}

type ServerAddressByClientCIDR struct {
	ClientCIDR    string `json:"clientCIDR" protobuf:"bytes,1,opt,name=clientCIDR"`
	ServerAddress string `json:"serverAddress" protobuf:"bytes,2,opt,name=serverAddress"`
}

type GroupVersionForDiscovery struct {
	GroupVersion string `json:"groupVersion" protobuf:"bytes,1,opt,name=groupVersion"`
	Version      string `json:"version" protobuf:"bytes,2,opt,name=version"`
}

type APIResource struct {
	Name         string   `json:"name" protobuf:"bytes,1,opt,name=name"`
	SingularName string   `json:"singularName" protobuf:"bytes,6,opt,name=singularName"`
	Namespaced   bool     `json:"namespaced" protobuf:"varint,2,opt,name=namespaced"`
	Group        string   `json:"group,omitempty" protobuf:"bytes,8,opt,name=group"`
	Version      string   `json:"version,omitempty" protobuf:"bytes,9,opt,name=version"`
	Kind         string   `json:"kind" protobuf:"bytes,3,opt,name=kind"`
	Verbs        Verbs    `json:"verbs" protobuf:"bytes,4,opt,name=verbs"`
	ShortNames   []string `json:"shortNames,omitempty" protobuf:"bytes,5,rep,name=shortNames"`
	Categories   []string `json:"categories,omitempty" protobuf:"bytes,7,rep,name=categories"`
}

type Verbs []string

func (vs Verbs) String() string {
	return fmt.Sprintf("%v", []string(vs))
}

type APIResourceList struct {
	TypeMeta     `json:",inline"`
	GroupVersion string        `json:"groupVersion" protobuf:"bytes,1,opt,name=groupVersion"`
	APIResources []APIResource `json:"resources" protobuf:"bytes,2,rep,name=resources"`
}

type RootPaths struct {
	Paths []string `json:"paths" protobuf:"bytes,1,rep,name=paths"`
}

func LabelSelectorQueryParam(version string) string {
	return "labelSelector"
}

func FieldSelectorQueryParam(version string) string {
	return "fieldSelector"
}

func (apiVersions APIVersions) String() string {
	return strings.Join(apiVersions.Versions, ",")
}

func (apiVersions APIVersions) GoString() string {
	return apiVersions.String()
}

type Patch struct{}

type LabelSelector struct {
	MatchLabels      map[string]string          `json:"matchLabels,omitempty" protobuf:"bytes,1,rep,name=matchLabels"`
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty" protobuf:"bytes,2,rep,name=matchExpressions"`
}

type LabelSelectorRequirement struct {
	Key      string                `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	Operator LabelSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=LabelSelectorOperator"`
	Values   []string              `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type LabelSelectorOperator string

const (
	LabelSelectorOpIn           LabelSelectorOperator = "In"
	LabelSelectorOpNotIn        LabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
)
