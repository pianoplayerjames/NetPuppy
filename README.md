# Requirements
- Golang v.1.22.1

# Installation
cd ~
curl -OL https://golang.org/dl/go1.22.1.linux-amd64.tar.gz
sha256sum go1.22.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.22.1.linux-amd64.tar.gz
sudo nano ~/.profile
export PATH=$PATH:/usr/local/go/bin
source ~/.profile
git clone https://github.com/pianoplayerjames/netsquirrel
cd netsquirrel
chmod +x net
./net


# How to start server
```./net```

# How to use
- You can connect to the server by installing a tcp client such as https://github.com/trshpuppy/netpuppy or telnet.

# Plugins
- There is a plugin store at https://github.com/pianoplayerjames/netsquirrel_plugins or you can use your own store by changing the url in plugins/install.go

# commands
- all commands are based on what plugins you have. so if you have a go file in your plugins directory called "encrypt.go", the command would be "encrypt". You can also type "help" to see a list of the plugins directory with the description of how to use it.

# How to make a plugin
- In the plugins directory there is a template.go file. You can use this boilerplate code to start a new plugin. Just replace all instances of Template with your plugin name.