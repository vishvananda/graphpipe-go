# graphpipe-tf - Serve TensorFlow Models via Graphpipe

The headlines are true! You can serve your Tensorflow models via graphpipe
easily using this server.

If all you want to do is deploy models using GraphPipe, we recommend you read 
our [project documentation](https://oracle.github.io/graphpipe/).  If you are
interested in hacking on `graphpipe-tf`, read on.

## Development Quickstart
Because of the relative complexity of system configuration when dealing with machine
learning libraries, our dev and build system for graphpipe-tf is 100% docker-driven.

Our build system can output images in three configurations:

* *cpu* (default) - create an Ubuntu-based build for cpu inference.  In this configuration, the BLAS backend is MKL.
* *gpu* - create an Ubuntu-based build for gpu inference.  If no physical gpu is present, inference falls back to 
  MKL cpu inference

You can switch between these configurations by setting the *RUN_TYPE* environment variable.

```
export RUN_TYPE=gpu
```

In order to support streamlined development and deployment, each build configuration
has 2 containers: one for development, and one for deployment.
```
make dev-container  # compiles the server inside the dev-container
make # compiles the server inside the dev-container
make runtime-container # compiles the runtime-container and injects build artifacts
```

Additionally, you can build all three of these steps at the same time:
```
make all
```

During development, it is usually sufficient to run the server from the development image.
An example invocation of a development server can be invoked like this:
```
make devserver  # observe the docker command that is output, and tweak it for your own testing
```

Similarly, you can invoke a test instance of the deployment
```
make runserver  # observe the docker command that is output, and tweak it for your own testing
```

If things seem broken, try dropping into a shell in your dev-container to figure things out:

```
make devshell
```

## Proxies
If you are behind a proxy, set the *http_proxy* and *https_proxy* environment variables so our build system
can forward this configuration to docker.

## Running the Server Manually
If you want to run the server manually, you will need to install some additional dependencies:

  - golang: version 1.8 or higher
  - libtensorflow: https://www.tensorflow.org/install/install_go or `make install-tensorflow`
  - govendor: https://github.com/kardianos/govendor or `make install-govendor`
  - CUDA9+: required for building gpu-capable binaries

If you manage to get these and their recursive deps installed, you should now be able to call build
without docker:

```
make deps
make local
```


## Running the server
Before running the server, you first need a tensorflow model to serve.
The model should be in protobuf (.pb) format, or in the tensorflow-serving
directory format.  We provide a tool to easily convert keras to the correct
model format [here](https://github.com/oracle/graphpipe-tf-py/blob/master/examples/convert.py).

To see the server's command line options:

```
./graphpipe-tf --help
graphpipe-tf - serving up ml models

Usage:
  graphpipe-tf [flags]

Flags:
  -n, --cache            do not cache results
  -d, --dir string       dir for local state (default "~/.graphpipe-tf")
  -h, --help             help for graphpipe-tf
  -i, --inputs string    comma seprated default inputs
  -l, --listen string    listen string (default "127.0.0.1:9000")
  -m, --model string     tensorflow model to load (accepts local files and unauthenticated http/https urls)
  -o, --outputs string   comma separated default outputs
  -v, --verbose          verbose output
  -V, --version          show version
```

The only required parameter is --model (see above).  If not specified, inputs and outputs
will be automatically determined by the graphdef as the first input and last output; if your
model has multiple inputs and/or outputs, you must specify these parameters manually.

Once you have your model, start the server .

```
./graphpipe-tf --model=mymodel.pb
```

## Environment Variables
For convenience, the key parameters of the service can be configured with environment variables,
 GP_MODEL, GP_INPUTS, GP_OUTPUTS, and GP_CACHE.


## Troubleshooting

### govendor can't fetch private libs
The in-docker setup _should_ forward `ssh-agent`s correctly if you
have it set up on your systems. Don't forget to `ssh-add` your key!

This link might be helpful: https://developer.github.com/v3/guides/using-ssh-agent-forwarding/

### proxies :(
Proxying should be forwarded for all our commands, but you may need
to configure your docker runtime to use them as well. Probably lives
(or needs to be created) at:

  `/etc/systemd/system/docker.service.d/http-proxy.conf`
