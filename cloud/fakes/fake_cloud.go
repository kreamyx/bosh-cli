package fakes

type FakeCloud struct {
	CreateStemcellInputs []CreateStemcellInput
	CreateStemcellCID    string
	CreateStemcellErr    error

	HasVMInput HasVMInput
	HasVMFound bool
	HasVMErr   error

	CreateVMInput CreateVMInput
	CreateVMCID   string
	CreateVMErr   error

	CreateDiskInput CreateDiskInput
	CreateDiskCID   string
	CreateDiskErr   error

	AttachDiskInput AttachDiskInput
	AttachDiskErr   error

	DetachDiskInput DetachDiskInput
	DetachDiskErr   error

	DeleteVMInput DeleteVMInput
	DeleteVMErr   error

	DeleteDiskInputs []DeleteDiskInput
	DeleteDiskErr    error

	DeleteStemcellInputs []DeleteStemcellInput
	DeleteStemcellErr    error
}

type CreateStemcellInput struct {
	CloudProperties map[string]interface{}
	ImagePath       string
}

type HasVMInput struct {
	VMCID string
}

type CreateVMInput struct {
	StemcellCID     string
	CloudProperties map[string]interface{}
	NetworksSpec    map[string]interface{}
	Env             map[string]interface{}
}

type CreateDiskInput struct {
	Size            int
	CloudProperties map[string]interface{}
	InstanceID      string
}

type AttachDiskInput struct {
	VMCID   string
	DiskCID string
}

type DetachDiskInput struct {
	VMCID   string
	DiskCID string
}

type DeleteVMInput struct {
	VMCID string
}

type DeleteDiskInput struct {
	DiskCID string
}

type DeleteStemcellInput struct {
	StemcellCID string
}

func NewFakeCloud() *FakeCloud {
	return &FakeCloud{
		CreateStemcellInputs: []CreateStemcellInput{},
		DeleteDiskInputs:     []DeleteDiskInput{},
	}
}

func (c *FakeCloud) CreateStemcell(cloudProperties map[string]interface{}, imagePath string) (string, error) {
	c.CreateStemcellInputs = append(c.CreateStemcellInputs, CreateStemcellInput{
		CloudProperties: cloudProperties,
		ImagePath:       imagePath,
	})

	return c.CreateStemcellCID, c.CreateStemcellErr
}

func (c *FakeCloud) DeleteStemcell(stemcellCID string) error {
	c.DeleteStemcellInputs = append(c.DeleteStemcellInputs, DeleteStemcellInput{
		StemcellCID: stemcellCID,
	})
	return c.DeleteStemcellErr
}

func (c *FakeCloud) HasVM(vmCID string) (bool, error) {
	c.HasVMInput = HasVMInput{
		VMCID: vmCID,
	}
	return c.HasVMFound, c.HasVMErr
}

func (c *FakeCloud) CreateVM(
	stemcellCID string,
	cloudProperties map[string]interface{},
	networksSpec map[string]interface{},
	env map[string]interface{},
) (string, error) {
	c.CreateVMInput = CreateVMInput{
		StemcellCID:     stemcellCID,
		CloudProperties: cloudProperties,
		NetworksSpec:    networksSpec,
		Env:             env,
	}

	return c.CreateVMCID, c.CreateVMErr
}

func (c *FakeCloud) CreateDisk(
	size int,
	cloudProperties map[string]interface{},
	instanceID string,
) (string, error) {
	c.CreateDiskInput = CreateDiskInput{
		Size:            size,
		CloudProperties: cloudProperties,
		InstanceID:      instanceID,
	}

	return c.CreateDiskCID, c.CreateDiskErr
}

func (c *FakeCloud) AttachDisk(vmCID, diskCID string) error {
	c.AttachDiskInput = AttachDiskInput{
		VMCID:   vmCID,
		DiskCID: diskCID,
	}
	return c.AttachDiskErr
}

func (c *FakeCloud) DetachDisk(vmCID, diskCID string) error {
	c.DetachDiskInput = DetachDiskInput{
		VMCID:   vmCID,
		DiskCID: diskCID,
	}
	return c.DetachDiskErr
}

func (c *FakeCloud) DeleteVM(vmCID string) error {
	c.DeleteVMInput = DeleteVMInput{
		VMCID: vmCID,
	}
	return c.DeleteVMErr
}

func (c *FakeCloud) DeleteDisk(diskCID string) error {
	c.DeleteDiskInputs = append(c.DeleteDiskInputs, DeleteDiskInput{
		DiskCID: diskCID,
	})
	return c.DeleteDiskErr
}

func (c *FakeCloud) String() string {
	return "FakeCloud{}"
}
