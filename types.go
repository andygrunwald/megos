package megos

// State represents the JSON from the state.json of a mesos node
type State struct {
	// Missing fields...
	// TODO "attributes": {},
	// TODO Better type as string for "2015-05-05 06:16:58" ?
	// TODO: completed_frameworks "completed_frameworks": [],
	// TODO "orphan_tasks": [],
	ActivatedSlaves        int         `json:"activated_slaves"`
	BuildDate              string      `json:"build_date"`
	BuildTime              int         `json:"build_time"`
	BuildUser              string      `json:"build_user"`
	Cluster                string      `json:"cluster"`
	DeactivatedSlaves      int         `json:"deactivated_slaves"`
	ElectedTime            float32     `json:"elected_time"`
	FailedTasks            int         `json:"failed_tasks"`
	FinishedTasks          int         `json:"finished_tasks"`
	Flags                  Flags       `json:"flags"`
	Frameworks             []Framework `json:"frameworks"`
	GitSHA                 string      `json:"git_sha"`
	GitTag                 string      `json:"git_tag"`
	Hostname               string      `json:"hostname"`
	ID                     string      `json:"id"`
	KilledTasks            int         `json:"killed_tasks"`
	Leader                 string      `json:"leader"`
	LogDir                 string      `json:"log_dir"`
	LostTasks              int         `json:"lost_tasks"`
	MasterHostname         string      `json:"master_hostname"`
	PID                    string      `json:"pid"`
	Resources              Resources   `json:"resources"`
	Slaves                 []Slave     `json:"slaves"`
	StagedTasks            int         `json:"staged_tasks"`
	StartTime              float32     `json:"start_time"`
	StartedTasks           int         `json:"started_tasks"`
	UnregisteredFrameworks []Framework `json:"unregistered_frameworks"`
	Version                string      `json:"version"`
}

// Flags represents the flags of a mesos state
type Flags struct {
	AllocationInterval          string `json:"allocation_interval"`
	Authenticate                string `json:"authenticate"`
	Authenticatee               string `json:"authenticatee"`
	AuthenticateSlaves          string `json:"authenticate_slaves"`
	Authenticators              string `json:"authenticators"`
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
	// Missing fields
	// TODO: "offers": [],
	Active             bool       `json:"active"`
	Checkpoint         bool       `json:"checkpoint"`
	CompletedExecutors []Executor `json:"completed_executors"`
	CompletedTasks     []Task     `json:"completed_tasks"`
	FailoverTimeout    int        `json:"failover_timeout"`
	Hostname           string     `json:"hostname"`
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	OfferedResources   Resources  `json:"offered_resources"`
	RegisteredTime     float32    `json:"registered_time"`
	ReregisteredTime   float32    `json:"reregistered_time"`
	Resources          Resources  `json:"resources"`
	Role               string     `json:"role"`
	Tasks              []Task     `json:"tasks"`
	UnregisteredTime   float32    `json:"unregistered_time"`
	UsedResources      Resources  `json:"used_resources"`
	User               string     `json:"user"`
	WebuiURL           string     `json:"webui_url"`
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
