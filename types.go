package megos

// State represents the JSON from the state.json of a mesos node
type State struct {
	ActivatedSlaves        int         `json:"activated_slaves"`
	BuildDate              string      `json:"build_date"`
	BuildTime              int         `json:"build_time"`
	BuildUser              string      `json:"build_user"`
	Cluster                string      `json:"cluster"`
	CompletedFrameworks    []Framework `json:"completed_frameworks"`
	DeactivatedSlaves      int         `json:"deactivated_slaves"`
	ElectedTime            float32     `json:"elected_time"`
	Flags                  Flags       `json:"flags"`
	Frameworks             []Framework `json:"frameworks"`
	GitSHA                 string      `json:"git_sha"`
	GitBranch              string      `json:"git_branch"`
	GitTag                 string      `json:"git_tag"`
	Hostname               string      `json:"hostname"`
	ID                     string      `json:"id"`
	Leader                 string      `json:"leader"`
	LogDir                 string      `json:"log_dir"`
	ExternalLogFile        string      `json:"external_log_file"`
	OrphanTasks            []Task      `json:"orphan_tasks"`
	PID                    string      `json:"pid"`
	Slaves                 []Slave     `json:"slaves"`
	StartTime              float32     `json:"start_time"`
	UnregisteredFrameworks []string    `json:"unregistered_frameworks"`
	Version                string      `json:"version"`
}

// Flags represents the flags of a mesos state
type Flags struct {
	AllocationInterval          string `json:"allocation_interval"`
	Allocator                   string `json:"allocator"`
	Authenticate                string `json:"authenticate"`
	Authenticatee               string `json:"authenticatee"`
	AuthenticateSlaves          string `json:"authenticate_slaves"`
	Authenticators              string `json:"authenticators"`
	Authorizers                 string `json:"authorizers"`
	CgroupsEnableCfs            string `json:"cgroups_enable_cfs"`
	CgroupsHierarchy            string `json:"cgroups_hierarchy"`
	CgroupsLimitSwap            string `json:"cgroups_limit_swap"`
	CgroupsRoot                 string `json:"cgroups_root"`
	Cluster                     string `json:"cluster"`
	ContainerDiskWatchInterval  string `json:"container_disk_watch_interval"`
	Containerizers              string `json:"containerizers"`
	DefaultRole                 string `json:"default_role"`
	DiskWatchInterval           string `json:"disk_watch_interval"`
	Docker                      string `json:"docker"`
	DockerRemoveDelay           string `json:"docker_remove_delay"`
	DockerSandboxDirectory      string `json:"docker_sandbox_directory"`
	DockerStopTimeout           string `json:"docker_stop_timeout"`
	EnforceContainerDiskQuota   string `json:"enforce_container_disk_quota"`
	ExecutorRegistrationTimeout string `json:"executor_registration_timeout"`
	ExecutorShutdownGracePeriod string `json:"executor_shutdown_grace_period"`
	FrameworksHome              string `json:"frameworks_home"`
	FrameworkSorter             string `json:"framework_sorter"`
	GCDelay                     string `json:"gc_delay"`
	GCDiskHeadroom              string `json:"gc_disk_headroom"`
	HadoopHome                  string `json:"hadoop_home"`
	Help                        string `json:"help"`
	Hostname                    string `json:"hostname"`
	InitializeDriverLogging     string `json:"initialize_driver_logging"`
	IP                          string `json:"ip"`
	Isolation                   string `json:"isolation"`
	LauncherDir                 string `json:"launcher_dir"`
	LogAutoInitialize           string `json:"log_auto_initialize"`
	LogDir                      string `json:"log_dir"`
	Logbufsecs                  string `json:"logbufsecs"`
	LoggingLevel                string `json:"logging_level"`
	MaxSlavePingTimeouts        string `json:"max_slave_ping_timeouts"`
	Master                      string `json:"master"`
	PerfDuration                string `json:"perf_duration"`
	PerfInterval                string `json:"perf_interval"`
	Port                        string `json:"port"`
	Quiet                       string `json:"quiet"`
	Quorum                      string `json:"quorum"`
	Recover                     string `json:"recover"`
	RecoverySlaveRemovalLimit   string `json:"recovery_slave_removal_limit"`
	RecoveryTimeout             string `json:"recovery_timeout"`
	RegistrationBackoffFactor   string `json:"registration_backoff_factor"`
	Registry                    string `json:"registry"`
	RegistryFetchTimeout        string `json:"registry_fetch_timeout"`
	RegistryStoreTimeout        string `json:"registry_store_timeout"`
	RegistryStrict              string `json:"registry_strict"`
	ResourceMonitoringInterval  string `json:"resource_monitoring_interval"`
	RootSubmissions             string `json:"root_submissions"`
	SlavePingTimeout            string `json:"slave_ping_timeout"`
	SlaveReregisterTimeout      string `json:"slave_reregister_timeout"`
	Strict                      string `json:"strict"`
	SwitchUser                  string `json:"switch_user"`
	UserSorter                  string `json:"user_sorter"`
	Version                     string `json:"version"`
	WebuiDir                    string `json:"webui_dir"`
	WorkDir                     string `json:"work_dir"`
	ZK                          string `json:"zk"`
	ZKSessionTimeout            string `json:"zk_session_timeout"`
}

// Framework represent a single framework of a mesos node
type Framework struct {
	Active           bool       `json:"active"`
	Checkpoint       bool       `json:"checkpoint"`
	CompletedTasks   []Task     `json:"completed_tasks"`
	Executors        []Executor `json:"executors"`
	FailoverTimeout  int        `json:"failover_timeout"`
	Hostname         string     `json:"hostname"`
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	OfferedResources Resources  `json:"offered_resources"`
	Offers           []Offer    `json:"offers"`
	RegisteredTime   float32    `json:"registered_time"`
	ReregisteredTime float32    `json:"reregistered_time"`
	Resources        Resources  `json:"resources"`
	Role             string     `json:"role"`
	Tasks            []Task     `json:"tasks"`
	UnregisteredTime float32    `json:"unregistered_time"`
	UsedResources    Resources  `json:"used_resources"`
	User             string     `json:"user"`
	WebuiURL         string     `json:"webui_url"`
	Labels           []Label    `json:"label"`
}

type Offer struct {
	ID          string            `json:"id"`
	FrameworkID string            `json:"framework_id"`
	SlaveID     string            `json:"slave_id"`
	Hostname    string            `json:"hostname"`
	URL         URL               `json:"url"`
	Resources   Resources         `json:"resources"`
	Attributes  map[string]string `json:"attributes"`
}

type URL struct {
	Scheme     string      `json:"scheme"`
	Address    Address     `json:"address"`
	Path       string      `json:"path"`
	Parameters []Parameter `json:"parameters"`
}

type Address struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Label struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Task represent a single Mesos task
type Task struct {
	// Missing fields
	// TODO: "labels": [],
	ExecutorID  string       `json:"executor_id"`
	FrameworkID string       `json:"framework_id"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Resources   Resources    `json:"resources"`
	SlaveID     string       `json:"slave_id"`
	State       string       `json:"state"`
	Statuses    []TaskStatus `json:"statuses"`
}

// Resources represents a resource type for a task
type Resources struct {
	CPUs  float32 `json:"cpus"`
	Disk  int     `json:"disk"`
	Mem   float64 `json:"mem"`
	Ports string  `json:"ports"`
}

// TaskStatus represents the status of a single task
type TaskStatus struct {
	State     string  `json:"state"`
	Timestamp float32 `json:"timestamp"`
}

// Slave represents a single mesos slave node
type Slave struct {
	Active         bool              `json:"active"`
	Hostname       string            `json:"hostname"`
	ID             string            `json:"id"`
	PID            string            `json:"pid"`
	RegisteredTime float32           `json:"registered_time"`
	Resources      Resources         `json:"resources"`
	Attributes     map[string]string `json:"attributes"`
}

// Executor represents a single executor of a framework
type Executor struct {
	// Missing fields
	// TODO "queued_tasks": [],
	// TODO "tasks": []
	CompletedTasks []Task    `json:"completed_tasks"`
	Container      string    `json:"container"`
	Directory      string    `json:"directory"`
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Resources      Resources `json:"resources"`
	Source         string    `json:"source"`
}
