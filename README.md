# curl-a-joke-faas
This is a variation on the [curl-a-joke](https://github.com/dvdmuckle/curl-a-joke) project designed for use with the OpenFaaS system. It is able to give you a random joke from a selection of twenty (very bad) jokes. Unlike the original curl-a-joke, this variation will not allow for POSTing new jokes.

### Build and Deploy
```bash
faas-cli build -f ./curl-a-joke.yml && faas-cli deploy -f ./curl-a-joke.yml
```

Make sure to run `docker-machine env` before this if you're deploying on a non-local instance. 

If you're deploying to Kubernetes, please use the `curl-a-joke-k8s.yml` variation.
