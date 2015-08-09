# openstack-ssh
[![Build Status](https://travis-ci.org/pyama86/openstack-ssh.svg?branch=master)](https://travis-ci.org/pyama86/openstack-ssh)

ssh authentication key by openstack key-pair wrapper

## Installation

### Download
https://github.com/pyama86/openstack-ssh/releases

### Configure
 /etc/ssh/openstack-ssh.conf
 ```
 auth_url = "http://your-api-host:35357/v2.0"
 user = "your name"
 password = "your password"
 tenant = "your tenant"
 region = "your region"(default regionOne)
 ```
 /etc/ssh/sshd_config
 ```
 AuthorizedKeysCommand  /path/to/openstack-ssh
 AuthorizedKeysCommandUser nobody
 ```

## Usage
 create user `pyama` to server
```
[root@test-server.com] useradd pyama
```

create key-pair

```
nova keypair-add --pub-key ~/.ssh/github_rsa.pub pyama
```
login
```
ssh pyama@test-server.com
```


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D


## License

MIT
