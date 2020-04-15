# Helm + Codefresh with Lightstep

![Helm plus Codefresh](codefresh-helm.jpg)
![With Lightstep](lslogo.png)

This is an example Go application packaged with Docker and Helm.
It is compiled using Codefresh. 
This application has been instrumented with OpenTelemetry and Lightstep in order to monitor its performance.

## Create a multi-stage docker image

To compile and package using Docker multi-stage builds

```bash
docker build . -t my-app
```

## To run the docker image

```bash
docker run -p 8080:8080 my-app
```

And then visit http://localhost:8080 in your browser.

## Editing the chart

The chart was created using [Draft](draft.sh). Make sure to edit the templates and values
with your own settings (e.g. docker image deployed).

## To use this project in Codefresh

There is also a [codefresh.yml](codefresh.yml) for easy usage with the [Codefresh](codefresh.io) CI/CD platform.

More details can be found in [Codefresh documentation](https://codefresh.io/docs/docs/yaml-examples/examples/helm)

## Using this project with Lightstep

Sign up for a [Lightstep account](https://go.lightstep.com/trial) and set your project access token as the `LS_KEY` value in Helm to enable reporting. You can set this value in the `codefresh.yml` file, by passing it as a environment value to the `deploy` step in Codefresh.

