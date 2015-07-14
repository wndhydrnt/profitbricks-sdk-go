package main
	
	import (
		"fmt"
		"profitbricks-sdk-go"
	)
	
	func main() {

		//Sets username and password 
		profitbricks.SetAuth("username", "password")
		//Sets depth.
		profitbricks.SetDepth("5")

		dcrequest := profitbricks.CreateDatacenterRequest{
			DCProperties: profitbricks.DCProperties{
				Name:        "test",
				Description: "description",
				Location:    "us/lasdev",
			},
		}
	
		datacenter := profitbricks.CreateDatacenter(dcrequest)
	
		serverrequest := profitbricks.CreateServerRequest{
			ServerProperties: profitbricks.ServerProperties{
				Name:  "go01",
				Ram:   1024,
				Cores: 2,
			},
		}
		server := profitbricks.CreateServer(datacenter.Id, serverrequest)
	
		images := profitbricks.ListImages()
	
		fmt.Println(images.Items)
	
		volumerequest := profitbricks.CreateVolumeRequest{
			VolumeProperties: profitbricks.VolumeProperties{
				Size:        1,
				Name:        "Volume Test",
				LicenceType: "LINUX",
			},
		}
	
		storage := profitbricks.CreateVolume(datacenter.Id, volumerequest)
	
		serverupdaterequest := profitbricks.ServerProperties{
			Name:  "go01renamed",
			Cores: 1,
			Ram:   256,
		}
		resp := profitbricks.PatchServer(datacenter.Id, server.Id, serverupdaterequest)
	
		volumes := profitbricks.ListVolumes(datacenter.Id)
		servers := profitbricks.ListServers(datacenter.Id)
		datacenters := profitbricks.ListDatacenters()
		
		profitbricks.DeleteServer(datacenter.Id, server.Id)
		profitbricks.DeleteDatacenter(datacenter.Id)
	}