# Git issues

When I get a new computer, some things just have to be re-done.

This link has succinct info:
https://stackoverflow.com/questions/38556096/github-permission-denied-publickey-fatal-could-not-read-from-remote-reposit

1. Generate SSH locally by:

```sh
ssh-keygen -t rsa -b 4096 -C "cecil.new@gmail.com"
```

2. Copy key to clipboard

```
$ cat ~/.ssh/id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCwYgZdK59J18pplzzPnIGbnwf5s4ojM4n+jYi+so2zyZO52Y5o7WsQuPuru+tVgHNux0/zE9MheVClsy6mlZLz6ksLq0+YeZPKwye4g65KpaakbaeTidyDFKXtswn8Wt2ZvPBZRYwmBCukyAcIpdEGjvo+WaUKtlhag6dG06IPnqM4wDbVnWJZcCgvVn4v8l0VCsvn8QybQIUK4ONNTzM8Vhff9e766AqfYbPdiPvG6bJW6019UVR507Kgp9FVnmqtoc021aH+OxVzk4JhkCF/aUokPPfTMjBdqeE0xvChyC4TfL+xGwsxatcCoT3po3yCF/qhaAeoMa/ILgZ0pdrKl4OQB7qlkxr9L7HtXD1ACmh1TTPs9bEoRcKC/S7fZoVRIUCn3eDb1OHPv4H/u8Qk7CnKJ+jgDLQI7U3FP+xh3ID8ZnVhNCvoWXrhobuxxKCWWDhpkR6tE+t5mlkKy1IKCpWq+dBOnPbJXMQTYGwaHYq+ZGQ7W6bOJWbx+2QDNBE= mando@LAPTOP-GVLOCNGD
```

3. Paste the above-copied output to the form at [https://github.com/settings/ssh/new](https://github.com/settings/ssh/new)