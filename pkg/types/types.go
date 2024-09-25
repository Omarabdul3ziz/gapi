package types

// api server types

type NodeFarmInfo struct {
	FarmId   uint   `json:"farm_id"`
	FarmName string `json:"farm_name"`
}

type Location struct {
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Capacity struct { // all in bytes right? GB vs GiB
	CRU uint `json:"cru"`
	SRU uint `json:"sru"`
	HRU uint `json:"hru"`
	MRU uint `json:"mru"`
}

type CapacityReport struct {
	Total Capacity `json:"total"` // is zos included?
	Used  Capacity `json:"used"`  // is zos included?
	Free  Capacity `json:"free"`  // is zos excluded?
}

type PowerReport struct {
	OverallUptime uint   `json:"overall_uptime"`
	LatestUptime  uint   `json:"latest_uptime"`
	Status        string `json:"status"`
	Healthy       bool   `json:"healthy"`
	State         string `json:"state"`
	Target        string `json:"target"`
}

type RentingReport struct {
	DedicatedFarm  bool `json:"dedicated_farm"`
	DedicatedNode  bool `json:"dedicated_node"`
	Rentable       bool `json:"rentable"`
	Rented         bool `json:"rented"`
	Renter         uint `json:"renter"`
	RentContractId uint `json:"rent_contract_id"`
}

type BIOS struct {
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}

type Baseboard struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
}

type Processor struct {
	Version     string `json:"version"`
	ThreadCount string `json:"thread_count"`
}

type Memory struct {
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
}

type Dmi struct {
	NodeTwinId uint32      `json:"node_twin_id,omitempty"`
	BIOS       BIOS        `json:"bios"`
	Baseboard  Baseboard   `json:"baseboard"`
	Processor  []Processor `json:"processor"`
	Memory     []Memory    `json:"memory"`
}

type Speed struct {
	Upload   float64 `json:"upload"`   // in bit/sec
	Download float64 `json:"download"` // in bit/sec
}

type GPUCard struct {
	ID       string `json:"id"`
	Vendor   string `json:"vendor"`
	Device   string `json:"device"`
	Contract int    `json:"contract"`
}

type PriceReport struct {
	ExtraFee   uint `json:"extra_fee"`
	TotalPrice uint `json:"total_price"`
}

type PublicConfig struct {
	Domain string `json:"domain"`
	Gw4    string `json:"gw4"`
	Gw6    string `json:"gw6"`
	Ipv4   string `json:"ipv4"`
	Ipv6   string `json:"ipv6"`
}

type Metadata struct {
	GridVersion       uint   `json:"grid_version"`
	SerialNumber      string `json:"serial_number"`
	CertificationType string `json:"certification_type"`
	FarmingPolicyID   int    `json:"farming_policy_id"`
	CreationTime      uint   `json:"creation_time"`
}

type Node struct {
	NodeId       uint           `json:"node_id"`
	TwinId       uint           `json:"twin_id"`
	Farm         NodeFarmInfo   `json:"farm"`
	Location     Location       `json:"location"`
	Capacity     CapacityReport `json:"capacity"`
	Gpus         []GPUCard      `json:"gpus"`
	Power        PowerReport    `json:"power"`
	Renting      RentingReport  `json:"renting"`
	Dmi          Dmi            `json:"dmi"`
	Speed        Speed          `json:"speed"`
	PriceUsd     PriceReport    `json:"price"`
	PublicConfig PublicConfig   `json:"public_config"`
	Metadata     Metadata       `json:"metadata"`
}

type Farm struct{}
type Twin struct{}
type Contract struct{}
