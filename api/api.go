package api

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/kolo/xmlrpc"
	"github.com/megamsys/libgo/cmd"
)

const (
	ENDPOINT        = "endpoint"
	USERID          = "userid"
	PASSWORD        = "password"
	TEMPLATE        = "template"
	IMAGE           = "image"
	ONEZONE         = "region"
	VCPU_PERCENTAGE = "vcpu_percentage"
	CLUSTER         = "cluster"

	VMPOOL_ACCOUNTING      = "one.vmpool.accounting"
	VMPOOL_INFO            = "one.vmpool.info"
	TEMPLATEPOOL_INFO      = "one.templatepool.info"
	TEMPLATE_UPDATE        = "one.template.update"
	VM_INFO                = "one.vm.info"
	DISK_ATTACH            = "one.vm.attach"
	DISK_DETACH            = "one.vm.detach"
	VNET_CREATE            = "one.vn.allocate"
	VNET_ADDIP             = "one.vn.add_ar"
	VNET_SHOW              = "one.vn.info"
	VNET_LIST              = "one.vnpool.info"
	ONE_HOST_INFO          = "one.host.info"
	ONE_DATASTORE_INFO     = "one.datastore.info"
	ONE_DATASTOREPOOL_INFO = "one.datastorepool.info"
	ONE_HOST_ALLOCATE      = "one.host.allocate"
	ONE_HOST_DELETE        = "one.host.delete"
	ONE_DATASTORE_ALLOCATE = "one.datastore.allocate"
	ONE_TEMPLATE_ALLOCATE  = "one.template.allocate"
	ONE_HOST_POOL          = "one.hostpool.info"
)

var (
	ErrArgsNotSatisfied = errors.New("[" + ENDPOINT + "," + USERID + "," + PASSWORD + "] one (or) more args missing!")
)

type Rpc struct {
	Client xmlrpc.Client
	Key    string
}

func NewClient(config map[string]string) (*Rpc, error) {
	log.Debugf(cmd.Colorfy("  > [one-go] connecting", "blue", "", "bold"))

	if !satisfied(config) {
		return nil, ErrArgsNotSatisfied
	}

	client, err := xmlrpc.NewClient(config[ENDPOINT], nil)
	if err != nil {
		return nil, err
	}
	//log.Debugf(cmd.Colorfy("  > [one-go] connection response", "blue", "", "bold")+"%#v",client)
	log.Debugf(cmd.Colorfy("  > [one-go] connected", "blue", "", "bold")+" %s", config[ENDPOINT])

	return &Rpc{
		Client: *client,
		Key:    config[USERID] + ":" + config[PASSWORD]}, nil
}

func (c *Rpc) Call(command string, args []interface{}) ([]interface{}, error) {
	log.Debugf(cmd.Colorfy("  > [one-go] ", "blue", "", "bold")+"%s", command)
	log.Debugf(cmd.Colorfy("\n> args   ", "cyan", "", "bold")+" %v\n", args)

	result := []interface{}{}

	if err := c.Client.Call(command, args, &result); err != nil {
		return nil, err
	}
	//log.Debugf(cmd.Colorfy("\n> response ", "cyan", "", "bold")+" %v", result)
	log.Debugf(cmd.Colorfy("  > [one-go] ( ´ ▽ ` ) SUCCESS", "blue", "", "bold"))
	return result, nil
}

func satisfied(c map[string]string) bool {
	return len(strings.TrimSpace(c[ENDPOINT])) > 0 &&
		len(strings.TrimSpace(c[USERID])) > 0 &&
		len(strings.TrimSpace(c[PASSWORD])) > 0
}
