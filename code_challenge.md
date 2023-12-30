# DevOps Code Challenge

You have made a great impression on us in our first steps and we would very much like to jump straight to the next part of our assessment.

## Task
Setup an infrastructure capable of running one of our web apps as described below. You are given a Docker image, a Docker Compose file, and two Vagrant configuration files (x86 and ARM64 architecture) that we will use to test your solution. You should only use the Vagrant configuration file relevant to your machine architecture. You are supposed to do as many requirements as you think is necessary for MVP, but please don't spend more than 2-3 hours on the task. We are interested to see what you come up with.

### Requirements
* Single user `Xorg` should have SSH access to the servers using SSH key (please generate one yourself).
* In case one server becomes unavailable, the observed change in behaviour of the system should be minimal.
* The web app should support 5 concurrent users on various HTTP endpoints at any moment.
* The web app will at peak times be accessed by 100 concurrent users.
* The web app logs that begin with `[ERROR]` should be saved into a file at `/var/logs/app1/errors.log` on the host VM in format that includes docker service name, docker container id, timestamp, and content of the logline.
* The web app's HTTP response body should be completely removed from the HTTP response whenever its status code is >= 500. The response code and the headers should stay the same.

### What you'll get from us
The web app image we want you to run can be found at [Xorghearing/challenge-devops-sample-app](https://hub.docker.com/repository/docker/Xorghearing/challenge-devops-sample-app) on Docker Hub. The app runs on port 3000. The app will respond with 500 status code and write error logs to standard output whenever it receives an HTTP request containing letter `s` in the URL path.

The Docker Compose file describes the three services that will be run. The `sample-app` and `mongo` services do not need any additional configuration to run. You will need to add an nginx configuration file to the `nginx` service.

With this code challenge, you should also receive a `Vagrantfile`. You can use it to configure Vagrant for provisioning VMs. We will use the same `Vagrantfile` when testing your solution.

### Tech constraints
Configure the solution taking into account any security configuration you deem necessary for the requirements.

We use Ansible to configure and manage our servers. We greatly appreciate solutions using the same or similar tech stack.

## Logistics
* You are **not** allowed to publish this document or mention our name whatsoever. If you want to reference it, edit to use wrappers. Don’t write "Xorg" but write "a company" :)
* If you don’t want to publish your code, it is 100% acceptable to just zip/tar it and send it to the person you got this document from.

_Good luck and have fun!_
