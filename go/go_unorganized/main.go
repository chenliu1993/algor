package main

import "fmt"

var (
	text = `
	time="2020-09-14T13:26:22.445493000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.448124000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.448161000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.448188000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:22.448501000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:22.448932000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:22.448965000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.448990000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:22.449013000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:22.449036000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.449062000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:22.449090000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:22.451155000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:22.452907000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:22.452973000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.453026000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.453074000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.453123000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.453318000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:22.453907000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.451579000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:26.451633000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:26.451662000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.451689000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.451713000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.451742000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.451767000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.451790000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:26.454429000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:26.456549000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:26.456596000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.456645000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.456747000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.456858000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.457070000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.457247000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.457679000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:26.457712000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:26.457736000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.457757000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.457784000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.457804000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.457824000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.457844000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:26.458956000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:26.461018000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:26.461097000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.461169000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.461243000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.461331000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.461500000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.461673000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.462065000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:26.462097000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:26.462135000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.462171000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.462198000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.462221000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.462242000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:26.462262000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:26.463670000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:26.465526000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:26.465570000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.465626000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.465697000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.465796000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.465958000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:26.466156000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.452454000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:27.452502000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:27.452527000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.452561000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.452583000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.452606000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.452627000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.452647000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:27.453908000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:27.455879000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:27.455954000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.456007000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.456123000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.456221000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.456445000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.456677000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.457337000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:27.457406000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:27.457486000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.457528000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.457566000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.457612000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.457697000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:27.457762000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:27.459206000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:27.461867000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:27.461940000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.462001000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.462045000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.462217000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.462557000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:27.462857000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.469464000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:32.469538000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:32.469571000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.469592000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.469628000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.469654000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.469679000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.469697000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:32.471747000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:32.475631000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:32.475701000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.475747000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.475870000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.475983000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.476269000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.476494000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.477340000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:32.477424000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:32.477482000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.477559000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.477601000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.477633000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.477671000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.477708000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:32.479177000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:32.481587000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:32.481643000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.481672000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.481750000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.481877000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.482016000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.482215000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.482686000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:32.482730000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:32.482762000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.482793000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.482812000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.482833000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.482858000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:32.482876000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:32.485246000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:32.487010000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:32.487066000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.487104000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.487176000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.487275000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.487455000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:32.487589000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:38.496716000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:38.496775000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:38.496796000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.496817000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:38.496835000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:38.496852000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.496870000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:38.496887000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:38.499025000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:38.500851000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:38.500913000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.500940000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.501012000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.501107000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.502383000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:38.503317000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:40.501121000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:40.501193000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:40.501220000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.501246000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:40.501273000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:40.501303000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.501327000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:40.501348000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:40.502988000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:40.505419000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:40.505518000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.505579000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.505657000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.505788000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.506041000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:40.506386000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:43.511044000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:43.511097000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:43.511150000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.511195000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:43.511231000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:43.511264000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.511296000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:43.511326000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:43.512797000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:43.514776000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:43.514830000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.514864000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.515830000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.515911000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.516073000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:43.516255000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:48.524516000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:48.524655000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:48.524705000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.524758000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:48.524823000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:48.524870000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.524944000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:48.524984000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:48.527311000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:48.529995000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:48.530063000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.530111000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.531133000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.531258000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.531582000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:48.531905000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:49.526277000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:49.526330000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:49.526358000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.526386000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:49.526412000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:49.526437000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.526476000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:49.526507000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:49.528043000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:49.530418000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:49.530490000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.530551000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.530649000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.530768000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.531064000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:49.531355000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.527096000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:50.527164000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:50.527210000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.527252000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.527311000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.527356000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.527405000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.527480000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:50.529870000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:50.532595000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:50.532668000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.532771000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.532886000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.533038000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.533327000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.533574000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.534307000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:50.534622000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:50.534755000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.534891000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.534962000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.535023000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.535117000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.535177000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:50.536358000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:50.539633000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:50.539707000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.539820000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.539873000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.539996000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.540315000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.541397000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.541922000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:50.541999000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:50.542033000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.542066000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.542102000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.542134000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.542168000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.542237000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:50.543712000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:50.546055000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:50.546118000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.546160000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.546270000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.546429000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.546782000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.546925000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.547435000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:50.547476000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:50.547513000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.547545000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.547577000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.547604000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.547638000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:50.547667000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:50.549211000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:50.551240000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:50.551297000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.551342000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.551449000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.551561000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.551755000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:50.551937000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.530533000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.530631000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.530710000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.530764000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.530820000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.530878000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.530941000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.530966000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.533859000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.536869000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.536946000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.536998000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.537140000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.537271000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.537562000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.537846000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.538373000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.538412000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.538442000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.538465000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.538486000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.538507000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.538529000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.538550000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.540251000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.543491000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.543575000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.543622000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.543663000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.543791000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.544009000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.544192000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.544672000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.544711000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.544738000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.544764000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.544804000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.544826000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.544848000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.544867000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.546214000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.548198000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.548239000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.548276000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.548338000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.548433000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.548663000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.548847000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.549179000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.549212000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.549235000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.549259000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.549280000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.549301000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.549322000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.549342000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.550679000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.552477000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.552518000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.552549000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.552645000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.552717000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.552871000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.553029000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.553392000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.553437000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.553463000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.553486000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.553536000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.553574000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.553611000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.553645000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.554888000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.556922000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.556980000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.557016000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.557086000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.557184000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.557370000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.557561000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.558010000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.558043000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.558069000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.558092000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.558113000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.558134000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.558155000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.558175000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.559352000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.561377000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.561429000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.561462000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.561495000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.561615000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.561836000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.562080000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.562468000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.562519000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.562608000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.562645000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.562674000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.562695000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.562718000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.562738000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.563915000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.565792000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.565838000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.565865000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.565960000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.566043000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.566207000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.566417000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.566797000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:51.566828000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:51.566851000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.566871000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.566889000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.566906000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.566926000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:51.566946000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:51.568035000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:51.570090000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:51.570152000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.570180000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.570211000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.570326000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.570481000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:51.570637000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:53.535961000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:53.536014000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:53.536047000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.536071000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:53.536092000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:53.536113000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.536145000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:53.536180000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:53.538866000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:53.541012000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:53.541063000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.541095000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.541178000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.541297000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.541500000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:53.541682000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.540477000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:55.540549000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:55.540592000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.540631000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.540667000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.540702000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.540737000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.540772000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:55.541977000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:55.544066000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:55.544113000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.544167000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.544347000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.544448000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.544636000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.544817000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.545285000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:55.545331000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:55.545371000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.545426000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.545463000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.545492000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.545516000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:55.545537000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:55.546566000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:55.548553000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:55.548607000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.548659000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.548747000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.548836000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.548998000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:55.549223000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.543190000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:56.543237000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:56.543262000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.543302000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.543324000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.543344000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.543366000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.543393000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:56.544485000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:56.546349000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:56.546391000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.546423000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.546504000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.546592000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.546792000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.546951000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.547348000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:56.547376000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:56.547401000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.547422000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.547443000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.547463000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.547484000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.547504000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:56.548633000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:56.550726000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:56.550780000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.550829000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.550886000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.550982000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.551164000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.551333000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.551737000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:56.551767000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:56.551793000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.551823000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.551844000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.551864000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.551890000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:56.551927000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:56.552993000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:56.554810000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:56.554850000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.554879000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.554978000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.555056000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.555221000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:56.555386000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:58.547676000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:58.547748000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:58.547807000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.547866000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:58.547896000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:58.547926000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.547954000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:58.547975000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:58.549129000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:58.551424000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:58.551480000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.551512000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.551578000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.551667000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.551909000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:58.552106000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:59.552188000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:26:59.552258000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:26:59.552290000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.552346000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:59.552388000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:59.552428000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.552456000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:26:59.552482000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:26:59.553959000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:26:59.556684000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:26:59.556770000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.556837000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.556870000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.557017000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.557214000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:26:59.557450000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:01.558589000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:01.558680000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:01.558712000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.558758000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:01.558807000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:01.558852000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.558891000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:01.558922000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:01.560171000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:01.562062000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:01.562121000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.562173000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.562281000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.562404000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.562639000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:01.562878000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.567396000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:05.567924000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:05.568056000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.568118000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.568183000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.568289000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.568348000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.568391000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:05.570877000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:05.573127000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:05.573215000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.573287000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.573340000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.573454000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.573680000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.573942000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.574427000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:05.574470000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:05.574536000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.574573000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.574606000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.574670000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.574712000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:05.574742000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:05.577434000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:05.579271000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:05.579332000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.579371000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.579421000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.579521000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.579682000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:05.579865000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.570310000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:06.570393000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:06.570431000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.570463000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.570492000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.570530000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.570563000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.570595000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:06.572096000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:06.574257000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:06.574327000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.574365000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.574441000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.574529000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.574742000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.574960000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.575414000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:06.575456000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:06.575487000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.575513000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.575537000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.575561000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.575597000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.575625000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:06.576667000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:06.578686000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:06.578741000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.578779000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.578866000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.578959000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.579188000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.579378000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.579741000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:06.579796000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:06.579827000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.579852000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.579876000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.579915000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.579953000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:06.579979000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:06.580972000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:06.582805000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:06.582859000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.582912000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.582969000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.583028000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.583228000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:06.583391000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:07.573537000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:07.573637000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:07.573685000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.573716000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:07.573762000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:07.573813000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.573851000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:07.573885000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:07.575244000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:07.577455000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:07.577518000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.577562000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.577645000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.577745000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.577965000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:07.578099000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.574656000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:09.574709000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:09.574741000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.574764000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.574784000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.574806000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.574828000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.574855000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:09.576003000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:09.577521000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:09.577558000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.577586000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.577701000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.577775000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.577970000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.578157000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.578524000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:09.578556000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:09.578580000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.578601000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.578619000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.578637000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.578657000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:09.578678000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:09.581072000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:09.582533000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:09.582576000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.582603000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.582660000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.582737000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.582869000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:09.583048000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.574933000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:10.575031000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:10.575060000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.575087000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.575107000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.575433000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.575635000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.575704000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:10.577295000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:10.578957000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:10.579010000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.579041000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.579091000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.579193000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.579368000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.579541000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.579968000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:10.580010000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:10.580046000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.580080000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.580113000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.580144000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.580178000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.580217000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:10.581210000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:10.583100000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:10.583149000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.583184000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.583233000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.583291000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.583438000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.583564000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.583909000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:10.583937000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:10.583958000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.583985000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.584015000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.584056000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.584102000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:10.584143000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:10.585044000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:10.586584000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:10.586617000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.586642000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.586731000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.586821000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.586958000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:10.587869000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.782127000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:11.782189000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:11.782221000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.782248000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.782273000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.782300000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.782326000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.782349000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:11.783686000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:11.785403000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:11.785450000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.785486000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.785576000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.785673000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.785851000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.786036000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.786483000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:11.786519000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:11.786545000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.786569000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.786593000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.786616000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.786642000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:11.786665000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:11.788014000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:11.789664000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:11.789705000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.789741000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.789821000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.789918000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.790140000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:11.790350000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:12.575558000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:12.575645000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:12.575686000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.575719000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:12.575749000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:12.575776000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.575802000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:12.575826000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:12.578195000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:12.580455000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:12.580517000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.580559000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.580692000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.580790000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.581015000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:12.581308000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:15.580099000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:15.580168000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:15.580224000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.580274000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:15.580322000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:15.580369000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.580414000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:15.580453000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:15.582206000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:15.584179000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:15.584256000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.584298000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.584813000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.584943000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.585148000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:15.585362000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.593223000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:20.593273000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:20.593302000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.593323000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.593343000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.593361000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.593380000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.593402000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:20.594466000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:20.596061000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:20.596100000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.596129000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.596210000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.596292000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.596448000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.596580000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.596919000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:20.596950000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:20.596972000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.596992000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.597012000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.597033000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.597052000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:20.597070000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:20.598099000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:20.599681000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:20.599727000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.599757000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.599815000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.599879000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.600080000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:20.600273000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:23.604759000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:23.604815000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:23.604844000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.604874000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:23.604900000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:23.604924000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.604949000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:23.604972000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:23.606042000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:23.608269000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:23.608342000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.608400000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.608450000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.608526000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.608742000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:23.608941000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.605023000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:24.605080000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:24.605112000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.605143000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.605182000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.605241000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.605270000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.605296000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:24.606548000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:24.608340000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:24.608399000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.608440000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.608488000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.608590000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.608807000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.608940000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.609411000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:24.609442000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:24.609466000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.609488000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.609514000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.609536000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.609558000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:24.609580000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:24.610508000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:24.612272000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:24.612324000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.612361000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.612394000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.612468000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.612653000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:24.612823000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:27.617878000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:27.617926000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:27.617955000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.617980000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:27.617999000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:27.618017000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.618036000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:27.618053000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:27.619122000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:27.620674000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:27.620732000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.621123000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.621681000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.621753000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.621892000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:27.622025000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:28.621983000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:28.622025000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:28.622047000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.622068000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:28.622091000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:28.622111000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.622131000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:28.622150000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:28.623238000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:28.624987000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:28.625038000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.625083000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.625122000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.625206000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.625377000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:28.625570000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.648240000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:35.648333000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:35.648394000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.648435000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.648493000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.648524000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.648551000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.648578000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:35.649836000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:35.651901000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:35.651949000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.651986000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.652067000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.652165000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.652950000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.653476000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.653849000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:35.653886000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:35.653912000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.653937000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.653960000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.653985000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.654019000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:35.654043000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:35.655018000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:35.656993000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:35.657040000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.657077000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.657129000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.657227000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.657399000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:35.657885000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.656424000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:38.656724000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:38.656806000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.656853000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.656898000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.656991000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.657030000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.657060000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:38.658676000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:38.660568000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:38.660620000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.660655000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.660755000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.660832000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.661024000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.661191000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.661684000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:38.661721000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:38.661756000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.661788000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.661813000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.661836000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.661860000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:38.661883000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:38.663108000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:38.664918000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:38.664976000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.665013000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.665043000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.665143000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.665318000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:38.665478000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:39.659046000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:39.659123000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:39.659176000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.659208000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:39.659244000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:39.659270000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.659296000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:39.659325000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:39.660680000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:39.662823000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:39.662895000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.662927000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.663008000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.663134000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.663365000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:39.663653000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:41.664678000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:41.664795000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:41.664843000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.664871000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:41.664894000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:41.664914000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.664937000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:41.664959000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:41.666637000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:41.668584000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:41.668632000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.668667000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.668757000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.668859000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.669057000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:41.669255000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:42.666908000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:42.666956000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:42.666983000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.667005000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:42.667026000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:42.667059000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.667084000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:42.667108000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:42.668398000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:42.670460000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:42.670506000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.670538000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.670673000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.670774000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.670970000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:42.671182000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.682207000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:46.682259000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:46.682282000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.682300000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.682317000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.682333000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.682350000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.682367000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:46.684812000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:46.686328000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:46.686371000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.686396000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.687278000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.687346000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.687529000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.687707000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.688097000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:46.688126000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:46.688147000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.688165000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.688183000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.688199000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.688216000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:46.688232000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:46.689284000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:46.690901000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:46.690947000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.690974000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.691044000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.691118000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.691273000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:46.691411000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:48.689309000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:48.689357000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:48.689401000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.689433000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:48.689532000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:48.689570000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.689596000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:48.689617000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:48.690739000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:48.692988000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:48.693061000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.693107000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.693221000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.693318000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.693513000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:48.693725000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:49.693158000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:49.693207000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:49.693239000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.693266000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:49.693292000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:49.693320000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.693344000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:49.693364000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:49.694734000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:49.696594000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:49.696639000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.696673000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.696765000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.696898000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.697094000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:49.697293000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:50.698202000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:50.698249000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:50.698276000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.698299000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:50.698323000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:50.698344000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.698370000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:50.698394000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:50.699466000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:50.701636000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:50.701687000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.701721000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.701800000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.701934000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.702079000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:50.702243000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:52.703502000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:52.703558000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:52.703587000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.703618000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:52.703641000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:52.703661000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.703687000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:52.703730000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:52.704928000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:52.706945000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:52.707009000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.707065000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.707108000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.707183000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.707370000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:52.707537000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:53.707512000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:53.707553000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:53.707580000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.707602000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:53.707630000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:53.707650000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.707674000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:53.707702000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:53.708908000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:53.710574000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:53.710613000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.710664000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.710722000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.710846000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.710991000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:53.711149000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:55.716581000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:55.716644000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:55.716674000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.716707000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:55.716731000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:55.716751000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.716774000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:55.716809000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:55.718577000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:55.720661000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:55.720732000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.720798000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.720906000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.721083000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.721303000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:55.721531000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:56.721328000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:56.721426000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:56.721468000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.721523000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:56.721588000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:56.721644000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.721677000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:56.721700000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:56.722892000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:56.724907000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:56.724965000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.725047000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.725099000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.725209000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.725459000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:56.725598000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.725374000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:57.725424000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:57.725740000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.725840000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.725980000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.726088000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.726157000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.726210000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:57.728072000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:57.730240000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:57.730318000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.730372000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.730491000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.730576000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.730739000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.730978000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.731654000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:57.731694000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:57.731720000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.731749000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.731772000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.731793000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.731817000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.731850000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:57.734358000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:57.736331000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:57.736376000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.736409000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.736443000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.736555000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.736738000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.736902000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.737345000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:27:57.737377000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:27:57.737403000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.737425000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.737448000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.737468000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.737489000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:27:57.737509000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:27:57.738678000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:27:57.740898000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:27:57.740960000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.741015000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.741126000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.741246000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.741433000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:27:57.741611000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:00.736474000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:00.736535000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:00.736564000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.736592000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:00.736615000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:00.736636000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.736660000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:00.736684000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:00.738005000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:00.739930000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:00.739969000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.739999000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.741405000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.741495000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.741655000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:00.741850000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:02.741562000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:02.741607000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:02.741635000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.741660000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:02.741685000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:02.741707000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.741730000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:02.741753000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:02.744235000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:02.746028000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:02.746069000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.746102000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.746254000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.746361000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.746579000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:02.746742000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.745880000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:03.745925000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:03.745951000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.745979000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.746003000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.746023000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.746045000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.746064000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:03.747872000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:03.749941000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:03.750006000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.750066000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.750184000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.750289000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.750504000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.750781000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.751480000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:03.751517000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:03.751544000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.751574000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.751609000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.751633000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.751655000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:03.751717000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:03.753437000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:03.756155000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:03.756225000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.756277000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.756326000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.756518000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.756755000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:03.756958000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:09.762032000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:09.762991000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:09.763044000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.763080000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:09.763106000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:09.763130000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.763155000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:09.763186000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:09.763651000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:09.765912000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:09.765966000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.766022000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.766729000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.766798000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.767617000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:09.768305000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:11.765900000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:11.765992000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:11.766016000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.766054000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:11.766073000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:11.766091000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.766111000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:11.766152000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:11.767309000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:11.769310000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:11.769389000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.769426000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.769499000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.769597000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.769777000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:11.769923000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:14.775050000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:14.775145000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:14.775207000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.775258000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:14.775299000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:14.775358000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.775406000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:14.775440000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:14.776989000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:14.779447000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:14.779521000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.779601000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.779705000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.779795000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.780071000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:14.780335000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:16.780801000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:16.780856000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:16.780890000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.780916000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:16.780941000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:16.780964000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.780989000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:16.781012000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:16.784071000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:16.786397000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:16.786471000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.786523000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.787351000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.787472000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.787670000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:16.787949000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.783659000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:17.783718000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:17.783748000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.783781000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.783807000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.783831000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.783856000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.783880000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:17.785430000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:17.787786000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:17.787858000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.787914000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.787965000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.788045000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.788272000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.788444000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.788952000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:17.788997000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:17.789027000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.789077000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.789110000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.789134000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.789162000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:17.789186000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:17.790807000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:17.793140000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:17.793203000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.793318000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.793426000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.793544000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.793718000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:17.793913000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.787049000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:19.787104000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:19.787136000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.787162000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.787187000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.787215000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.787240000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.787262000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:19.788389000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:19.790492000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:19.790567000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.790616000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.790741000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.790853000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.791040000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.791233000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.791708000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:19.791748000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:19.791776000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.791808000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.791834000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.791857000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.791883000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:19.791911000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:19.792951000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:19.794773000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:19.794814000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.794849000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.794915000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.795016000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.795180000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:19.795366000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.791547000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:21.791600000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:21.791629000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.791657000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.791682000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.791708000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.791734000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.791767000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:21.792988000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:21.794810000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:21.794854000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.794890000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.794991000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.795080000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.795259000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.795455000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.795820000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:21.795858000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:21.795885000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.795919000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.795945000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.795969000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.795993000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:21.796019000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:21.797013000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:21.799084000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:21.799142000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.799177000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.799256000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.799349000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.799548000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:21.799743000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:22.795072000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:22.795132000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:22.795161000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.795187000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:22.795213000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:22.795237000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.795263000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:22.795287000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:22.796590000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:22.798837000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:22.798897000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.798933000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.799081000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.799209000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.799440000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:22.799731000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:23.798885000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:23.798939000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:23.798978000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.799020000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:23.799056000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:23.799084000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.799107000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:23.799126000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:23.800668000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:23.803034000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:23.803090000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.803125000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.803268000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.803362000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.803575000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:23.803784000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.803589000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:24.803634000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:24.803657000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.803675000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.803693000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.803710000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.803727000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.803743000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:24.804872000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:24.806443000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:24.806490000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.806518000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.806606000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.806674000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.806849000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.807006000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.807357000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:24.807388000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:24.807414000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.807433000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.807451000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.807468000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.807486000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.807502000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:24.808739000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:24.810372000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:24.810414000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.810438000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.810519000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.810641000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.810867000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.811060000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.811531000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:24.811574000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:24.811597000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.811620000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.811637000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.811653000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.811670000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:24.811686000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:24.812668000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:24.814348000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:24.814400000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.814452000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.814488000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.814582000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.814765000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:24.814897000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.806221000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:26.806268000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:26.806297000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.806318000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.806338000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.806359000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.806378000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.806396000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:26.808673000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:26.810170000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:26.810210000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.810235000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.810383000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.810562000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.810747000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.810917000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.811229000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:26.811256000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:26.811276000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.811296000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.811314000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.811331000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.811358000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:26.811383000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:26.813579000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:26.814982000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:26.815022000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.815047000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.815125000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.815209000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.815362000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:26.815502000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:27.809013000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:27.809056000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:27.809079000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.809097000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:27.809121000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:27.809139000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.809162000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:27.809178000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:27.811423000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:27.812960000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:27.812995000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.813029000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.813084000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.813149000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.813318000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:27.813505000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.815526000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:29.815918000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:29.815966000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.815995000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.816020000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.816047000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.816074000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.816097000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:29.818497000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:29.820326000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:29.820374000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.820409000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.820525000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.820659000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.820875000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.821095000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.821554000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:29.821598000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:29.821640000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.821673000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.821702000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.821725000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.821752000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:29.821774000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:29.823175000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:29.825300000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:29.825394000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.825455000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.825534000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.825732000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.825970000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:29.826291000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.816581000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:30.816896000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:30.817047000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.817116000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.817176000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.817226000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.817329000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.817410000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:30.820128000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:30.822495000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:30.822551000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.822589000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.824090000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.824149000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.824359000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.824533000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.824966000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:30.825006000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:30.825037000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.825063000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.825088000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.825112000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.825136000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.825159000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:30.826684000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:30.828726000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:30.828777000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.828817000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.828934000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.829037000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.829237000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.829451000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.829952000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:30.829998000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:30.830046000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.830088000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.830115000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.830140000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.830165000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:30.830190000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:30.831677000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:30.833985000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:30.834141000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.834212000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.834431000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.834633000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.834959000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:30.835195000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:32.822055000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:32.822113000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:32.822142000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.822171000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:32.822196000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:32.822219000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.822244000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:32.822266000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:32.824791000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:32.826584000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:32.826642000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.826696000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.826790000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.826883000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.827080000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:32.827259000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:38.826102000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:38.826154000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:38.826180000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.826206000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:38.826229000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:38.826249000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.826272000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:38.826293000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:38.827465000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:38.829360000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:38.829409000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.829444000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.829526000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.829617000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.829792000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:38.829963000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.834101000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:41.834169000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:41.834206000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.834234000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.834258000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.834280000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.834309000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.834333000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:41.835615000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:41.837744000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:41.837843000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.837901000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.838517000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.838595000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.839244000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.839629000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.840033000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:41.840071000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:41.840101000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.840126000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.840151000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.840173000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.840196000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:41.840218000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:41.841374000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:41.843262000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:41.843358000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.843416000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.843462000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.843498000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.843953000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:41.844124000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.840205000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:43.840592000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:43.840630000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.840663000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.840689000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.840716000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.840744000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.840766000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:43.843054000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:43.845368000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:43.845447000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.845535000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.845606000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.845766000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.846022000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.846264000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.846838000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:43.846892000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:43.846924000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.846947000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.846969000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.846993000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.847021000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.847044000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:43.849442000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:43.851350000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:43.851396000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.851432000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.851497000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.851598000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.851748000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.851924000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.852435000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:43.852474000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:43.852503000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.852530000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.852554000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.852576000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.852608000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:43.852645000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:43.854355000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:43.856543000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:43.856641000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.856721000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.856865000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.856992000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.857232000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:43.857398000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.841987000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:44.842046000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:44.842077000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.842109000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.842138000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.842163000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.842192000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.842219000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:44.843527000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:44.845662000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:44.845714000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.845755000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.845825000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.845922000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.846142000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.846313000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.846742000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:44.846778000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:44.846808000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.846833000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.846858000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.846882000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.846910000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:44.846940000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:44.849493000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:44.851600000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:44.851646000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.851689000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.851734000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.851822000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.852006000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:44.852183000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:46.844476000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:46.844551000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:46.844627000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.844677000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:46.844708000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:46.844733000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.844765000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:46.844790000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:46.846559000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:46.850094000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:46.850178000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.850238000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.850289000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.850444000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.850862000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:46.851263000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.849609000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:47.849728000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:47.849798000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.849853000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.849902000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.849936000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.849995000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.850079000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:47.852978000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:47.855229000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:47.855380000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.855427000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.856349000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.856542000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.857167000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.857757000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.858493000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:47.858546000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:47.858583000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.858628000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.858672000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.858707000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.858747000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:47.858788000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:47.859842000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:47.862950000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:47.863055000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.863170000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.863351000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.863694000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.864080000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:47.864469000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.782596000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.782682000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.782738000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.782784000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.782842000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.782931000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.783011000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.783102000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.784724000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.787127000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.787231000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.787303000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.787429000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.787546000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.787953000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.788097000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.788695000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.788771000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.788845000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.788900000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.788968000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.789025000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.789060000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.789096000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.790447000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.793051000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.793130000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.793211000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.793315000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.793384000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.793562000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.793861000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.794375000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.794423000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.794460000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.794496000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.794531000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.794567000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.794608000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.794644000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.796065000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.798838000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.798910000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.799045000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.799137000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.799317000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.799558000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.799903000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.800357000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.800394000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.800425000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.800450000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.800482000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.800507000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.800532000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.800561000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.801706000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.804151000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.804217000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.804265000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.804374000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.804508000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.804782000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.805070000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.805635000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.805698000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.805747000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.805811000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.805861000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.805929000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.805993000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.806055000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.807512000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.810129000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.810234000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.810283000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.810335000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.810492000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.810746000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.811046000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.811582000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.811639000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.811687000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.811716000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.811742000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.811765000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.811789000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.811812000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.813548000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.815811000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.815894000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.815962000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.815996000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.816035000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.816278000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.816470000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.816938000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.817003000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.817053000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.817093000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.817126000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.817160000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.817207000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.817241000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.818399000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.820582000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.820629000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.820659000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.820788000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.820916000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.821134000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.821421000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.821963000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.822019000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.822060000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.822099000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.822124000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.822147000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.822169000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.822197000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.823608000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.826069000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.826133000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.826168000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.826320000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.826464000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.826689000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.826904000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.859177000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.859376000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.859416000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.859448000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.859478000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.859502000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.859530000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.859555000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.861829000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.864226000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.864302000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.864351000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.864446000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.864588000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.864850000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.865163000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.865877000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:51.865945000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:51.865988000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.866019000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.866054000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.866092000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.866134000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:51.866186000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:51.867559000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:51.869592000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:51.869669000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.869732000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.869822000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.869942000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.870195000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:51.870427000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.861398000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:52.861492000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:52.861535000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.861569000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.861604000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.861630000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.861655000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.861680000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:52.862963000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:52.864584000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:52.864627000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.864661000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.864769000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.864924000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.865115000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.865315000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.865662000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:52.865693000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:52.865722000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.865755000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.865783000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.865806000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.865831000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:52.865854000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:52.867000000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:52.868988000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:52.869043000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.869078000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.869148000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.869239000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.869455000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:52.869634000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.863801000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:53.863883000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:53.863942000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.863986000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.864031000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.864071000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.864113000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.864151000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:53.866502000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:53.868566000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:53.868630000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.868668000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.868756000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.868848000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.869095000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.869414000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.869929000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:53.869989000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:53.870031000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.870062000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.870089000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.870113000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.870148000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:53.870173000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:53.871697000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:53.873392000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:53.873443000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.873484000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.873557000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.873643000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.873822000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:53.874021000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.873530000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:56.873605000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:56.873655000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.873687000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.873713000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.873739000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.873764000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.873787000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:56.874909000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:56.877424000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:56.877484000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.877523000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.877567000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.877683000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.877858000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.878031000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.878467000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:56.878503000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:56.878533000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.878559000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.878582000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.878605000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.878638000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:56.878662000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:56.879876000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:56.881610000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:56.881658000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.881695000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.881775000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.881881000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.882085000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:56.882301000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:58.878374000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:28:58.878462000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:28:58.878529000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.878569000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:58.878598000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:58.878626000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.878675000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:28:58.878708000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:28:58.880336000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:28:58.882592000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:28:58.882692000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.882778000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.882845000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.882955000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.883223000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:28:58.883598000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.779036000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:01.779544000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:01.779598000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.779626000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.779659000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.779684000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.779709000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.779733000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:01.781678000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:01.783537000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:01.783581000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.783622000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.784588000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.784679000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.784844000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.785070000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.786189000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:01.786252000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:01.786281000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.786308000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.786333000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.786356000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.786380000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.786403000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:01.787253000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:01.789033000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:01.789083000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.789447000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.789557000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.789632000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.789787000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.789958000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.790327000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:01.790359000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:01.790388000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.790414000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.790438000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.790461000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.790486000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.790508000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:01.791568000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:01.793259000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:01.793300000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.793335000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.793392000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.793484000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.793615000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.793776000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.794134000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:01.794172000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:01.794206000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.794234000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.794260000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.794284000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.794310000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.794333000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:01.796566000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:01.798561000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:01.798607000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.798652000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.798715000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.798838000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.798968000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.799157000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.887052000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:01.887108000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:01.887139000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.887167000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.887196000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.887221000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.887249000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.887280000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:01.888519000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:01.890451000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:01.890493000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.890529000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.890621000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.890711000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.890877000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.891043000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.891446000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:01.891486000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:01.891517000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.891543000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.891568000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.891591000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.891618000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:01.891645000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:01.892652000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:01.894274000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:01.894321000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.894359000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.894415000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.894496000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.894674000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:01.894846000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:03.891894000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:03.891939000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:03.891963000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.891987000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:03.892006000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:03.892025000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.892044000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:03.892062000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:03.892930000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:03.894674000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:03.894716000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.894752000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.894824000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.894884000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.895040000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:03.895180000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.895928000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:06.895983000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:06.896010000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.896036000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.896060000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.896090000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.896114000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.896141000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:06.897111000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:06.899034000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:06.899108000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.899165000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.899204000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.899296000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.899468000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.899661000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.900030000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:06.900067000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:06.900095000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.900119000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.900142000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.900164000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.900188000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:06.900209000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:06.903227000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:06.905416000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:06.905505000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.905566000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.905657000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.905778000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.905990000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:06.906173000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:07.900492000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:07.900545000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:07.900589000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.900626000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:07.900653000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:07.900676000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.900704000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:07.900733000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:07.902932000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:07.904612000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:07.904654000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.904689000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.904769000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.904858000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.905052000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:07.905212000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:08.905205000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:08.905258000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:08.905288000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.905336000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:08.905367000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:08.905394000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.905422000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:08.905446000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:08.908072000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:08.910270000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:08.910337000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.910396000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.910512000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.910614000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.910851000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:08.911080000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:11.913493000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:11.913553000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:11.913585000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.913615000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:11.913647000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:11.913675000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.913709000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:11.913740000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:11.915053000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:11.916761000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:11.916807000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.916842000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.916941000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.917023000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.917196000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:11.917358000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.914522000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:12.914647000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:12.914733000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.914796000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.914907000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.914971000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.915036000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.915133000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:12.916861000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:12.919407000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:12.919466000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.919511000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.919622000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.919753000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.919976000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.920182000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.920705000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:12.920747000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:12.920779000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.920806000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.920832000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.920860000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.920887000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:12.920911000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:12.921974000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:12.923915000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:12.923966000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.924005000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.924062000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.924167000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.924331000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:12.924507000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:14.923627000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:14.923760000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:14.923809000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.923855000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:14.923901000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:14.923941000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.923983000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:14.924021000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:14.926040000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:14.928591000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:14.928651000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.928701000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.928763000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.928889000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.929122000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:14.929346000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.927623000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:15.927741000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:15.927810000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.927884000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.927942000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.927996000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.928053000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.928105000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:15.930549000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:15.933111000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:15.933172000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.933222000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.933354000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.933413000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.933685000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.933862000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.934395000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:15.934437000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:15.934479000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.934512000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.934543000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.934572000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.934604000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.934633000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:15.935980000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:15.937801000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:15.937854000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.937896000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.938005000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.938126000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.938327000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.938503000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.938925000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:15.938968000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:15.938998000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.939028000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.939054000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.939079000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.939110000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:15.939135000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:15.940149000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:15.942002000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:15.942051000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.942085000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.942144000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.942238000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.942417000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:15.942618000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:20.944559000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:20.944644000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:20.944740000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.944787000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:20.944819000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:20.944850000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.944901000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:20.944988000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:20.946153000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:20.948272000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:20.948343000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.948376000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.948459000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.948570000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.948744000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:20.949248000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:21.948667000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:21.948743000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:21.948790000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.948833000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:21.948868000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:21.948902000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.948938000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:21.948971000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:21.950536000+08:00" level=debug msg="success:true " module=libcrx package=agent
time="2020-09-14T13:29:21.952712000+08:00" level=debug msg="container state in agent after the kill signal:running" module=libcrx package=sandbox
time="2020-09-14T13:29:21.952760000+08:00" level=debug msg="update the state of container:nginx, new state:running" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.952802000+08:00" level=debug msg="save the config of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.953012000+08:00" level=debug msg="port mapping of the container:nginx before saving:[]" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.953177000+08:00" level=debug msg="tries to update the indexed container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.953451000+08:00" level=debug msg="save the status of the container into:/Users/cliu2/.vctl/.r/containers/nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:21.953730000+08:00" level=debug msg="stop monitoring state for container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.478092000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.478136000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:23.478174000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.478218000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.478247000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.478279000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.478304000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.478327000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:23.479006000+08:00" level=debug msg="[shim] ===>> &CloseIORequest{ID:nginx,ExecID:,Stdin:true,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.497986000+08:00" level=debug msg="agent client stderr done:&{0xc000010078 0xc0000aa700 0xc000151080}" module=libcrx package=agent
time="2020-09-14T13:29:23.498060000+08:00" level=debug msg="agent client stderr err:rpc error: code = Unavailable desc = transport is closing" module=libcrx package=agent
time="2020-09-14T13:29:23.498105000+08:00" level=debug msg="failed to copy container stderr" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.498147000+08:00" level=debug msg="[shim] iostream.Wait done"
time="2020-09-14T13:29:23.498187000+08:00" level=debug msg="[shim] SandboxWaitProcess..."
time="2020-09-14T13:29:23.498222000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.498286000+08:00" level=debug msg="failed to kill a container using the agent client, err:rpc error: code = Unavailable desc = transport is closing" module=libcrx package=agent
time="2020-09-14T13:29:23.498330000+08:00" level=debug msg="failed to kill a container:nginx in sandbox:nginx when performing kill request" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.498359000+08:00" level=debug msg="wait container process response:<nil>" module=libcrx package=agent
time="2020-09-14T13:29:23.498389000+08:00" level=debug msg="[shim] SandboxWaitProcess done: false -1 rpc error: code = Unavailable desc = connection closed"
time="2020-09-14T13:29:23.498417000+08:00" level=debug msg="[shim] <<=== &WaitResponse{ExitStatus:255,ExitedAt:2020-09-14 13:29:23.498157 +0800 CST m=+1238.512711592,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.498494000+08:00" level=debug msg="event forwarded" ns=vctl topic=/tasks/exit type=containerd.events.TaskExit
time="2020-09-14T13:29:23.498648000+08:00" level=debug msg="[shim] <<=== send event=/tasks/exit"
time="2020-09-14T13:29:23.499164000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.499214000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:23.499253000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.499280000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.499308000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.499330000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.499353000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.499373000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:23.499421000+08:00" level=debug msg="failed to kill a container using the agent client, err:rpc error: code = Unavailable desc = connection error: desc = \\transport: Error while dialing dial unix /Users/cliu2/.vctl/.r/vms/nginx/agent.sock: connect: connection refused\\" module=libcrx package=agent
time="2020-09-14T13:29:23.499456000+08:00" level=debug msg="failed to kill a container:nginx in sandbox:nginx when performing kill request" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.499695000+08:00" level=debug msg="[shim] ===>> &StateRequest{ID:nginx,ExecID:,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.499968000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:1,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.500009000+08:00" level=debug msg="[shim] kill, signal " name=SIGHUP number=1
time="2020-09-14T13:29:23.500040000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.500063000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.500122000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.500174000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.500225000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.500255000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:23.500333000+08:00" level=debug msg="failed to kill a container using the agent client, err:rpc error: code = Unavailable desc = connection error: desc = \\transport: Error while dialing dial unix /Users/cliu2/.vctl/.r/vms/nginx/agent.sock: connect: connection refused\\" module=libcrx package=agent
time="2020-09-14T13:29:23.500385000+08:00" level=debug msg="failed to kill a container:nginx in sandbox:nginx when performing kill request" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.500455000+08:00" level=debug msg="[shim] <<=== &StateResponse{ID:nginx,Bundle:/Users/cliu2/.vctl/storage/containerd/state/io.containerd.runtime.v2.task/vctl/nginx,Pid:27086,Status:STOPPED,Stdin:/var/folders/t2/5zp40rkx5t5_8xzhg76cj6_40000gn/T/vctl/fifo/755773382/nginx-stdin,Stdout:/var/folders/t2/5zp40rkx5t5_8xzhg76cj6_40000gn/T/vctl/fifo/755773382/nginx-stdout,Stderr:/var/folders/t2/5zp40rkx5t5_8xzhg76cj6_40000gn/T/vctl/fifo/755773382/nginx-stderr,Terminal:false,ExitStatus:255,ExitedAt:0001-01-01 00:00:00 +0000 UTC,ExecID:,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.500935000+08:00" level=debug msg="[shim] ===>> &KillRequest{ID:nginx,ExecID:,Signal:16,All:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.500971000+08:00" level=debug msg="[shim] kill, signal " name=SIGURG number=16
time="2020-09-14T13:29:23.501003000+08:00" level=debug msg="[shim] ===>> &DeleteRequest{ID:nginx,ExecID:,XXX_unrecognized:[],}"
time="2020-09-14T13:29:23.501031000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.501052000+08:00" level=debug msg="[shim] deleteContainerSandbox start..."
time="2020-09-14T13:29:23.501074000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.501096000+08:00" level=debug msg="tries to send signal to container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.501117000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.501137000+08:00" level=debug msg="[shim] StopContainer"
time="2020-09-14T13:29:23.501157000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.501178000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.501203000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.501227000+08:00" level=debug msg="port mapping:{[]} after fetching the container:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.501247000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:23.501298000+08:00" level=debug msg="failed to kill a container using the agent client, err:rpc error: code = Unavailable desc = connection error: desc = \\transport: Error while dialing dial unix /Users/cliu2/.vctl/.r/vms/nginx/agent.sock: connect: connection refused\\" module=libcrx package=agent
time="2020-09-14T13:29:23.501337000+08:00" level=debug msg="failed to kill a container:nginx in sandbox:nginx when performing kill request" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.501362000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.501384000+08:00" level=debug msg="stop container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:23.501405000+08:00" level=debug msg="kill container with ID:nginx" module=libcrx package=agent
time="2020-09-14T13:29:23.501447000+08:00" level=debug msg="failed to kill a container using the agent client, err:rpc error: code = Unavailable desc = connection error: desc = \\transport: Error while dialing dial unix /Users/cliu2/.vctl/.r/vms/nginx/agent.sock: connect: connection refused\\" module=libcrx package=agent
time="2020-09-14T13:29:23.501486000+08:00" level=debug msg="[shim] call libcrx: StopContainer, err=rpc error: code = Unavailable desc = connection error: desc = \\transport: Error while dialing dial unix /Users/cliu2/.vctl/.r/vms/nginx/agent.sock: connect: connection re" err=rpc
time="2020-09-14T13:29:23.501512000+08:00" level=debug msg="[shim] DeleteContainer"
time="2020-09-14T13:29:23.501535000+08:00" level=debug msg="DeleteContainer, podSandboxID:nginx containerID:nginx" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.501556000+08:00" level=debug msg="tries to get a container: nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.501581000+08:00" level=debug msg="tries to delete the container ID:nginx" module=libcrx package=ocicontainer
time="2020-09-14T13:29:23.502585000+08:00" level=debug msg="[shim] StopPodSandbox"
time="2020-09-14T13:29:23.502622000+08:00" level=debug msg="StopPodSandbox, podSandboxID:nginx force:true" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.502653000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.502682000+08:00" level=debug msg="vmrun location:/Applications/VMware Fusion Tech Preview.app/Contents/Public/vmrun" module=libcrx package=vmrun
time="2020-09-14T13:29:23.502711000+08:00" level=debug msg="command line arguments for vmrun:[stop /Users/cliu2/.vctl/.r/vms/nginx/nginx.vmx hard]" module=libcrx package=vmrun
time="2020-09-14T13:29:23.502743000+08:00" level=debug msg="cmd to run:/Applications/VMware Fusion Tech Preview.app/Contents/Public/vmrun -T fusion stop /Users/cliu2/.vctl/.r/vms/nginx/nginx.vmx hard" module=libcrx package=vmrun
time="2020-09-14T13:29:23.730225000+08:00" level=debug msg="result of command execution:Error: The virtual machine is not powered on: /Users/cliu2/.vctl/.r/vms/nginx/nginx.vmx\\n" module=libcrx package=vmrun
time="2020-09-14T13:29:23.730273000+08:00" level=debug msg="failed to run command:/Applications/VMware Fusion Tech Preview.app/Contents/Public/vmrun -T fusion stop /Users/cliu2/.vctl/.r/vms/nginx/nginx.vmx hard, err:exit status 255" module=libcrx package=vmrun
time="2020-09-14T13:29:23.730297000+08:00" level=debug msg="health check routine stoppped successfully" module=libcrx package=crxvm
time="2020-09-14T13:29:23.730322000+08:00" level=error msg="failed to stop a sandbox, err:error code:exit status 255, message:Error: The virtual machine is not powered on: /Users/cliu2/.vctl/.r/vms/nginx/nginx.vmx\\n" module=libcrx package=sandbox
time="2020-09-14T13:29:23.730345000+08:00" level=debug msg="failed to stop a sandbox" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.730361000+08:00" level=debug msg="stop health routine" module=libcrx package=crxvm
time="2020-09-14T13:29:23.730385000+08:00" level=debug msg="[shim] call libcrx: StopPodSandbox, err=error code:exit status 255, message:Error: The virtual machine is not powered on: /Users/cliu2/.vctl/.r/vms/nginx/ng" err=error
time="2020-09-14T13:29:23.730401000+08:00" level=debug msg="[shim] DeletePodSandbox"
time="2020-09-14T13:29:23.730417000+08:00" level=debug msg="DeletePodSandbox, podSandboxID:nginx" module=libcrx package=crxruntime
time="2020-09-14T13:29:23.730436000+08:00" level=debug msg="get the sandbox from store with ID:nginx" module=libcrx package=sandbox
time="2020-09-14T13:29:23.756046000+08:00" level=debug msg="Detach with arguments ([detach /dev/disk5])."
time="2020-09-14T13:29:24.110182000+08:00" level=debug msg="event forwarded" ns=vctl topic=/tasks/delete type=containerd.events.TaskDelete
time="2020-09-14T13:29:24.110417000+08:00" level=debug msg="[shim] ===>> &ShutdownRequest{ID:nginx,Now:false,XXX_unrecognized:[],}"
time="2020-09-14T13:29:24.110470000+08:00" level=debug msg="[shim] <<=== send event=/tasks/delete"
time="2020-09-14T13:29:24.111706000+08:00" level=debug msg="[shim] shim process is going to exit..."
time="2020-09-14T13:29:24.113267000+08:00" level=info msg="shim disconnected" id=nginx
time="2020-09-14T13:29:35.218289000+08:00" level=debug msg="received signal" signal=terminated
`
)

func main() {
	fmt.Println(text)
}
