## Solution Write up

For this challenge, I had to break down the problem into manageable tasks. Around 90% of the subtask are complete. Even though everything else worked. The provided docker image, for web, couldn't run. I had to create a simple http-based web server in golang. In my solution, you will find Ansible playbook to provision the vms. The playbook completes the following:
- Creates a user (Xorg).
- Creates a home directory for Xorg.
- Generate and copys ssh public key to the vms for ssh access.
- Installs docker and docker compose.
- Copies nginx configuration file, and docker-compose files.

With more time, I could have done the following:
- Figure out the Bad gateway error with my nginx configurations, and parse the logs generated.
- Secure the ssh access by a passphrase or password for the user Xorg.
- Encrypt the HTTP traffic by implementing TLS/SSL.
- Implementing loadbalancing using nginx to distribute traffic across all the three vms (nodes). 

### Blocker:
- The provided sample image couldn't run, and had to replicate a web server in go. 


### To test the solution:
- Install the following dependencies: Virtualbox (for x86 Architecture), Vagrant, and Ansible.
- Change into the challenge root directory. 
- Run `vagrant up --provision`. This command will spin up and provision all the VMs. This should complete successfully.
- To ssh into the VMs, run `ssh -i ./keys/Xorg_key Xorg@192.168.56.3`. This command should give you access to `node01` user Xorg. 

