[Unit] 
Description=Consul agent bootstrap
After=docker.service
Requires=docker.service

[Service]
EnvironmentFile=/etc/environment


ExecStartPre=-/usr/bin/mkdir -p /opt/bin
ExecStartPre=-/usr/bin/mkdir -p /opt/consul
ExecStartPre=-/usr/bin/mkdir -p /etc/consul.d
ExecStartPre=-/usr/bin/wget --retry-connrefused -t 5 -N -P /opt/consul/ https://releases.hashicorp.com/consul/1.0.7/consul_1.0.7_linux_amd64.zip
ExecStartPre=-/usr/bin/unzip /opt/consul/consul_1.0.7_linux_amd64.zip -d /opt/bin

ExecStartPre=-/usr/bin/chmod +x /etc/bootstrap-consulw

ExecStart=/etc/bootstrap-consulw

[Install]
WantedBy=multi-user.target

[X-Fleet]
Global=true

